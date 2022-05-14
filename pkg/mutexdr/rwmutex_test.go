package mutexdr

import (
	"sync"
	"testing"
)

func testRW(
	rw *RW[int],
	t *testing.T,
	expect int,
) {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)

		go func() {
			v := rw.Load()
			rw.Run(func(int) int {
				return v + 1
			})

			wg.Done()
		}()
	}
	wg.Wait()

	v := rw.Load()
	if v != expect {
		t.Fatalf("Expect RW's value to be %v, got %v", expect, v)
	}
}

func TestRW(t *testing.T) {
	rw := NewRW[int]()
	testRW(&rw, t, 3)
}
