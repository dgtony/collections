package concurrency

import (
	"context"
	"fmt"
)

// Future runs until completion and stays unchanged ever after.
func NewFuture(cons func() (interface{}, error)) *Future {
	var f = &Future{complete: make(chan struct{})}

	go func() {
		f.value, f.err = cons()
		close(f.complete)
	}()

	return f
}

type Future struct {
	value    interface{}
	err      error
	complete chan struct{}
}

func (f *Future) Wait(ctx context.Context) (interface{}, error) {
	select {
	case <-f.complete:
		return f.value, f.err
	case <-ctx.Done():
		return nil, fmt.Errorf("connection wait interrupted: %v", ctx.Err())
	}
}

func (f *Future) Get() (value interface{}, err error, complete bool) {
	select {
	case <-f.complete:
		value, err, complete = f.value, f.err, true
	default:
	}
	return
}

func (f *Future) Complete() bool {
	select {
	case <-f.complete:
		return true
	default:
		return false
	}
}
