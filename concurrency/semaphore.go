package concurrency

// Simple non-weighted semaphore.
type Semaphore struct {
	inner chan struct{}
}

func NewSemaphore(size int) *Semaphore {
	return &Semaphore{inner: make(chan struct{}, size)}
}

func (s *Semaphore) Acquire() {
	s.inner <- struct{}{}
}

func (s *Semaphore) TryAcquire() bool {
	select {
	case s.inner <- struct{}{}:
		return true
	default:
		return false
	}
}

func (s *Semaphore) Release() {
	select {
	case <-s.inner:
	default:
	}
}
