package mutexdr

import "sync"

// NewRW creates a reader and writer mutex (RW).
// T is the value's type.
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
// But RW's value can be predefined.
func NewRWWith[T any](value T) RW[T] {
	rw := NewRW[T]()
	rw.value = value

	return rw
}

// RW is wrapper of sync.RWMutex with Load and holds a value.
// It embeds W.
// Use NewRW or NewRWWith to create a new RW.
type RW[T any] struct {
	standardMutex *sync.RWMutex

	W[T]
}

// Load locks the reader mutex and returns the RW's curernt value.
func (mu *RW[T]) Load() T {
	mu.standardMutex.RLock()
	defer mu.standardMutex.RUnlock()

	return mu.value
}
