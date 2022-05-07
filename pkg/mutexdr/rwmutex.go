package mutexdr

import "sync"

type RWMutex[T any] struct {
	value T
	mutex sync.RWMutex
}

func (mutex *RWMutex[T]) WRun(f func(old T) (new T)) {
	mutex.mutex.Lock()
	defer mutex.mutex.Unlock()

	mutex.value = f(mutex.value)
}

func (mutex *RWMutex[T]) RRun(f func(old T)) {
	mutex.mutex.RLock()
	defer mutex.mutex.RUnlock()

	f(mutex.value)
}
