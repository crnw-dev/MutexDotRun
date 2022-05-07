package mutexdr

import "sync"

func NewW[T any]() W[T] {
	return W[T]{
		standardMutex: &sync.Mutex{},
	}
}

func NewWWith[T any](value T) W[T] {
	w := NewW[T]()
	w.value = value

	return w
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
	type cT = chan T

	c := make(cT)
	go func(mu *W[T], c cT) {
		mu.standardMutex.Lock()
		mu.value = f(mu.value)
		mu.standardMutex.Unlock()

		c <- mu.value
		close(c)
	}(mu, c)

	return c
}
