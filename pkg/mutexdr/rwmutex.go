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

func (mutex *Mutex[T]) AWRun(f func(old T) (new T)) chan<- T {
	c := make(chan T)
	go func(c chan T) {
		mutex.mutex.Lock()
		defer mutex.mutex.Unlock()

		mutex.value = f(mutex.value)
		c <- mutex.value
	}(c)

	return c
}

func (mutex *Mutex[T]) ARRun(f func(old T) (new T)) chan<- struct{} {
	c := make(chan struct{})
	go func(c chan struct{}) {
		mutex.mutex.Lock()
		defer mutex.mutex.Unlock()

		mutex.value = f(mutex.value)
		close(c)
	}(c)

	return c
}
