package mutexdr

import "sync"

// NewRW creates a reader and writer mutex (RW) with the provided type [T any].
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

// NewRWWith does the same thing as NewRW.
// But you can set a default value of [T any] for the new RW.
func NewRWWith[T any](value T) RW[T] {
	rw := NewRW[T]()
	rw.value = value

	return rw
}

// W is an extended version of sync.RWMutex with run functions and generic-type.
// Use NewRW or NewRWWith to create a new RW.
type RW[T any] struct {
	standardMutex *sync.RWMutex

	W[T]
}

// RRun locks the reader mutex and runs the provded function.
// Variable value is the curernt value of the mutex.
func (mu *RW[T]) RRun(f func(value T)) {
	mu.standardMutex.RLock()
	defer mu.standardMutex.RUnlock()

	f(mu.value)
}

// ARRun (await-RRun) does the same thing as RRun.
// But returns a channel which will be closed after the reader mutex re-locks.
func (mu *RW[T]) ARRun(f func(value T)) chan<- struct{} {
	type cT = chan struct{}

	c := make(cT)
	go func(mu *RW[T], c cT, f func(value T)) {
		mu.RRun(f)

		close(c)
	}(mu, c, f)

	return c
}
