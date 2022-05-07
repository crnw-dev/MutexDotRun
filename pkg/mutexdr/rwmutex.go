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

func NewRWWith[T any](value T) RW[T] {
	rw := NewRW[T]()
	rw.value = value

	return rw
}

type RW[T any] struct {
	standardMutex *sync.RWMutex

	W[T]
}

func (mu *RW[T]) RRun(f func(value T)) {
	mu.standardMutex.RLock()
	defer mu.standardMutex.RUnlock()

	f(mu.value)
}

func (mu *RW[T]) ARRun(f func(value T)) chan<- struct{} {
	type cT = chan struct{}

	c := make(cT)
	go func(mu *RW[T], c cT, f func(value T)) {
		mu.RRun(f)

		close(c)
	}(mu, c, f)

	return c
}
