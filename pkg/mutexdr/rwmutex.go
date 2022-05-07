package mutexdr

import "sync"

type RWMutex[T any] struct {
	value T
	mutex sync.Mutex
}

func (mutex *RWMutex[T]) WRun(f func(old T) (new T)) {
	mutex.mutex.Lock()
	defer mutex.mutex.Unlock()

	mutex.value = f(mutex.value)
}

func (mutex *RWMutex[T]) RRun(f func(old T)) {
	mutex.mutex.Lock()
	defer mutex.mutex.Unlock()

	f(mutex.value)
}
