package mutexdr

import "sync"

type Mutex[T any] struct {
	value T
	mutex sync.Mutex
}

func (mutex *Mutex) Run(f func(old T) (new T)) {
	mutex.mutex.Lock()
	defer mutex.mutex.UnLock()

	mutex.value = f(mutex.value)
}
