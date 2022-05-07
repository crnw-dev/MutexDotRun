package mutexdr

import "sync"

func NewRW[T any]() Rw[T] {
	return Rw[T]{}
}

type Rw[T any] struct {
	value         T
	standardMutex sync.RWMutex
}

func (mu *Rw[T]) WRun(f func(old T) (new T)) {
	mu.standardMutex.Lock()
	defer mu.standardMutex.Unlock()

	mu.value = f(mu.value)
}

func (mu *Rw[T]) RRun(f func(old T)) {
	mu.standardMutex.RLock()
	defer mu.standardMutex.RUnlock()

	f(mu.value)
}

func (mu *Rw[T]) AWRun(f func(old T) (new T)) chan<- T {
	c := make(chan T)
	go func(c chan T) {
		mu.standardMutex.Lock()
		defer mu.standardMutex.Unlock()

		mu.value = f(mu.value)
		c <- mu.value
	}(c)

	return c
}

func (mu *Rw[T]) ARRun(f func(old T) (new T)) chan<- struct{} {
	c := make(chan struct{})
	go func(c chan struct{}) {
		mu.standardMutex.Lock()
		defer mu.standardMutex.Unlock()

		mu.value = f(mu.value)
		close(c)
	}(c)

	return c
}
