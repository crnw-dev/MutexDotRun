package mutexdr

import (
	"testing"
)

func testW(w *W[int]) {
	c1 := make(chan struct{})
	c2 := make(chan struct{})
	c3 := make(chan struct{})

	go func() {
		<-c1
		w.Run(func(v int) int {
			close(c2)

			return v - 1
		})
	}()
	go func() {
		close(c1)

		<-c2
		w.Run(func(v int) int {
			return v + 2
		})

		close(c3)
	}()

	<-c3
}

func TestW(t *testing.T) {
	w := NewW[int]()
	testW(&w)

	w.Run(func(v int) int {
		const expect = 1
		if v != expect {
			t.Fatalf("Expect W's value to be %v, got %v", expect, v)
		}

		return v
	})
}
