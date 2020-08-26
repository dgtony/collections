package concurrency

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSemaphore(t *testing.T) {
	s := NewSemaphore(2)

	s.Acquire()
	assert.True(t, s.TryAcquire())
	assert.False(t, s.TryAcquire(), "full semaphore acquired")

	s.Release()
	assert.True(t, s.TryAcquire(), "cannot acquire vacant semaphore")
	s.Release()
	s.Release()

	// never block on release
	s.Release()
}
