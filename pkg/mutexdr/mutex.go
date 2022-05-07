package mutexdr

import "sync"

type W[T any] struct {
	value         T
	standardMutex sync.Mutex
}

func (mu *W[T]) Run(f func(old T) (new T)) {
	mu.standardMutex.Lock()
	defer mu.standardMutex.Unlock()

	mu.value = f(mu.value)
}

func (mu *W[T]) ARun(f func(old T) (new T)) chan<- T {
	c := make(chan T)
	go func(c chan T) {
		mu.standardMutex.Lock()
		defer mu.standardMutex.Unlock()

		mu.value = f(mu.value)
		c <- mu.value
	}(c)

	return c
}
