package mutexdr

import "sync"

func NewW[T any]() W[T] {
	return W[T]{
		standardMutex: &sync.Mutex{},
	}
}

type W[T any] struct {
	value         T
	standardMutex interface {
		Lock()
		Unlock()
	}
}

func (mu *W[T]) WRun(f func(old T) (new T)) {
	mu.standardMutex.Lock()
	defer mu.standardMutex.Unlock()

	mu.value = f(mu.value)
}

func (mu *W[T]) AWRun(f func(old T) (new T)) chan<- T {
	c := make(chan T)
	go func(c chan T) {
		mu.standardMutex.Lock()
		defer mu.standardMutex.Unlock()

		mu.value = f(mu.value)
		c <- mu.value
	}(c)

	return c
}
