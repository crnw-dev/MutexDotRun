package mutexdr

import "sync"

// NewW creates a writer mutex (W) with the provided type [T any].
func NewW[T any]() W[T] {
	return W[T]{
		standardMutex: &sync.Mutex{},
	}
}

// NewWWith does the same thing as NewW.
// But you can set a default value of [T any] for the new W.
func NewWWith[T any](value T) W[T] {
	w := NewW[T]()
	w.value = value

	return w
}

// W is an extended version of sync.Mutex with run functions and generic-type.
// Use NewW or NewWWith to create a new W.
type W[T any] struct {
	value         T
	standardMutex interface {
		Lock()
		Unlock()
	}
}

// Run locks the writer mutex and runs the provded function.
// Variable old is the curernt value of the mutex.
// It can be updated by returning a modified value.
func (mu *W[T]) Run(f func(old T) (new T)) {
	mu.standardMutex.Lock()
	defer mu.standardMutex.Unlock()

	mu.value = f(mu.value)
}

// ARun (await-Run) does the same thing as Run.
// But returns a channel that the updated value will be sent to after the writer mutex re-locks.
//
// It closes right after sending the updated value.
func (mu *W[T]) ARun(f func(old T) (new T)) chan<- T {
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
