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

// Load locks the reader mutex and returns the curernt value of the mutex.
func (mu *RW[T]) Load() T {
	mu.standardMutex.RLock()
	defer mu.standardMutex.RUnlock()

	return mu.value
}
