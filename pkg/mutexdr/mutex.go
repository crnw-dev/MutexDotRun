package mutexdr

import "sync"

type Mutex[T any] struct {
	value T
	mutex sync.Mutex
}

func (mutex *Mutex[T]) Run(f func(old T) (new T)) {
	mutex.mutex.Lock()
	defer mutex.mutex.Unlock()

	mutex.value = f(mutex.value)
}

func (mutex *Mutex[T]) ARun(f func(old T) (new T)) chan<- T {
	c := make(chan T)
	go func(c chan T) {
		mutex.mutex.Lock()
		defer mutex.mutex.Unlock()

		mutex.value = f(mutex.value)
		c <- mutex.value
	}(c)

	return c
}
