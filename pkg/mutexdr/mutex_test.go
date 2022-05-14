package mutexdr

import (
	"testing"
)

func testW(
	w *W[int],
	t *testing.T,
	expect int,
) {
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
	w.Run(func(v int) int {
		if v != expect {
			t.Fatalf("Expect W's value to be %v, got %v", expect, v)
		}

		return v
	})
}

func TestW(t *testing.T) {
	w := NewW[int]()
	testW(&w, t, 1)
}

func TestWWith(t *testing.T) {
	w := NewWWith(1)
	testW(&w, t, 2)
}
