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
	buffer  int
}

func NewBulkQueue(
	period time.Duration,
	buffer int,
) *BulkQueue {
	return &BulkQueue{
		period: period,
		buffer: buffer,
		input:  make(chan interface{}, 1),
		stop:   make(chan struct{}),
	}
}

func (c *BulkQueue) Run() <-chan interface{} {
	var (
		tick = time.NewTicker(c.period)
		sink = make(chan interface{}, c.buffer)
		ws   = hashset.New()
	)

	go func() {
		for {
			select {
			case <-c.stop:
				close(sink)
				tick.Stop()
				return

			case <-tick.C:
				for {
					v, ok := ws.Pop()
					if !ok {
						break
					}
					sink <- v
				}

			case v := <-c.input:
				ws.Add(v)
			}
		}
	}()

	return sink
}

func (c *BulkQueue) Stop() {
	c.stopper.Do(func() { close(c.stop) })
}

func (c *BulkQueue) Push(value interface{}) {
	c.input <- value
}
