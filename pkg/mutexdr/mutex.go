package mutexdr

import "sync"

// NewW creates a writer mutex (W).
// T is the value's type.
func NewW[T any]() W[T] {
	return W[T]{
		standardMutex: &sync.Mutex{},
	}
}

// NewWWith does the same thing as NewW.
// But W's value can be predefined.
func NewWWith[T any](value T) W[T] {
	w := NewW[T]()
	w.value = value

	return w
}

// W is wrapper of sync.Mutex with Run, ARun and holds a value.
// Use NewW or NewWWith to create a new W.
type W[T any] struct {
	value         T
	standardMutex interface {
		Lock()
		Unlock()
	}
}

// Run locks the writer mutex and runs the provded function.
// Variable old is the W's curernt value.
// It can be updated by returning a modified value.
func (mu *W[T]) Run(f func(old T) (new T)) {
	mu.standardMutex.Lock()
	defer mu.standardMutex.Unlock()

	mu.value = f(mu.value)
}

// ARun (await-Run) does the same thing as Run.
// But returns a channel that the updated value will be sent to after the writer mutex re-locks.
// A channel should only has one receiver.
func (mu *W[T]) ARun(f func(old T) (new T)) chan<- T {
	c := make(chan T, 1)
	go func(mu *W[T], c chan T) {
		mu.standardMutex.Lock()
		mu.value = f(mu.value)
		mu.standardMutex.Unlock()

		select {
		case c <- mu.value:
		default:
		}
		close(c)
	}(mu, c)

	return c
}
