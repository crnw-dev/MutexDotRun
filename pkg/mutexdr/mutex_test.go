package mutexdr

import (
	"testing"
)

func TestMutex(t *testing.T) {
	w := NewW[int]()
	c1 := make(chan struct{})
	c2 := make(chan struct{})

	go func() {
		w.Run(func(int) int {
			close(c1)

			return 1
		})
	}()
	go func() {
		<-c1
		w.Run(func(int) int {
			close(c2)

			return 2
		})
	}()

	<-c2
	w.Run(func(v int) int {
		const expect = 2
		if v != expect {
			t.Fatalf("Expect W's value to be %v, got %v", expect, v)
		}

		return v
	})
}
