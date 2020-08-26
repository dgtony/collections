package concurrency

type WorkerPoolBlocking interface {
	Handle(func())
}

type WorkerPoolNonBlocking interface {
	TryHandle(func()) bool
}

/* Implementations */

var _ WorkerPoolBlocking = &dynamicPool{}
var _ WorkerPoolNonBlocking = &dynamicPool{}

type dynamicPool struct {
	sema *Semaphore
}

func NewDynamicPool(maxWorkers int) *dynamicPool {
	return &dynamicPool{sema: NewSemaphore(maxWorkers)}
}

func (d *dynamicPool) Handle(action func()) {
	d.sema.Acquire()
	go func() {
		action()
		d.sema.Release()
	}()
}

func (d *dynamicPool) TryHandle(action func()) bool {
	if d.sema.TryAcquire() {
		go func() {
			action()
			d.sema.Release()
		}()
		return true
	}

	return false
}
