package mutexdr

import "sync"

func NewRW[T any]() RW[T] {
	m := &sync.RWMutex{}
	rw := RW[T]{
		standardMutex: m,
		W: W[T]{
			standardMutex: m,
		},
	}

	return rw
}

type RW[T any] struct {
	value         T
	standardMutex *sync.RWMutex

	W[T]
}

func (mu *RW[T]) RRun(f func(old T)) {
	mu.standardMutex.RLock()
	defer mu.standardMutex.RUnlock()

	f(mu.value)
}

func (mu *RW[T]) ARRun(f func(old T) (new T)) chan<- struct{} {
	c := make(chan struct{})
	go func(c chan struct{}) {
		mu.standardMutex.Lock()
		defer mu.standardMutex.Unlock()

		mu.value = f(mu.value)
		close(c)
	}(c)

	return c
}
