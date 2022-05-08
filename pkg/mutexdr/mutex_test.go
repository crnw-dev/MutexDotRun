package mutexdr

import (
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	type mT = []string

	m := NewW[mT]()
	m.Run(func() {
		time.Sleep()
	})

	go func() {
		m.Run(func(v mT) {

		})
	}()
}
