package queue

import (
	"sync"
	"time"

	"github.com/dgtony/collections/hashset"
)

// Generic queue G/D/1 with periodic flushing and value deduplication.
// It is safe to run queue more than once from the different
// goroutines, e.g for load balancing. All the partial sinks
// will be closed on the stop.
type BulkQueue struct {
	input   chan interface{}
	stop    chan struct{}
	stopper sync.Once
	period  time.Duration
	sinkBuf int
}

func NewBulkQueue(
	period time.Duration,
	inBufSize, outBufSize int,
) *BulkQueue {
	return &BulkQueue{
		period:  period,
		sinkBuf: outBufSize,
		input:   make(chan interface{}, inBufSize),
		stop:    make(chan struct{}),
	}
}

// Get results one-by-one.
func (c *BulkQueue) Run() <-chan interface{} {
	var (
		sink = make(chan interface{}, c.sinkBuf)
		ws   = hashset.New()
	)

	go c.run(
		ws,
		func() {
			for {
				v, ok := ws.Pop()
				if !ok {
					break
				}
				sink <- v
			}
		},
		func() { close(sink) },
	)

	return sink
}

// Get packs of results.
func (c *BulkQueue) RunBulk() <-chan []interface{} {
	var (
		sink = make(chan []interface{}, c.sinkBuf)
		ws   = hashset.New()
	)

	go c.run(
		ws,
		func() {
			if workingSetSize := ws.Len(); workingSetSize > 0 {
				var bulk = make([]interface{}, 0, workingSetSize)
				for {
					v, ok := ws.Pop()
					if !ok {
						break
					}
					bulk = append(bulk, v)
				}
				sink <- bulk
			}
		},
		func() { close(sink) },
	)

	return sink
}

func (c *BulkQueue) run(ws hashset.HashSet, onTick, onStop func()) {
	var tick = time.NewTicker(c.period)
	defer tick.Stop()

	for {
		select {
		case <-c.stop:
			onStop()
			return

		case <-tick.C:
			onTick()

		case v := <-c.input:
			ws.Add(v)
		}
	}
}

func (c *BulkQueue) Stop() {
	c.stopper.Do(func() { close(c.stop) })
}

func (c *BulkQueue) Push(value interface{}) {
	c.input <- value
}

func (c *BulkQueue) PushTimeout(value interface{}, timeout time.Duration) bool {
	var t = time.NewTimer(timeout)
	defer t.Stop()

	select {
	case c.input <- value:
		return true
	case <-t.C:
		return false
	}
}
