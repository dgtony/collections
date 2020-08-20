package time

import (
	"sync"
	"time"
)

func NewRightAwayTicker(d time.Duration) *rightAwayTicker {
	c := make(chan time.Time, 1)
	s := make(chan struct{})

	ticker := time.NewTicker(d)
	c <- time.Now()

	go func() {
		for {
			select {
			case t := <-ticker.C:
				c <- t
			case <-s:
				ticker.Stop()
				return
			}
		}
	}()

	return &rightAwayTicker{C: c, s: s}
}

type rightAwayTicker struct {
	C    chan time.Time
	s    chan struct{}
	stop sync.Once
}

func (t *rightAwayTicker) Stop() {
	t.stop.Do(func() { close(t.s) })
}
