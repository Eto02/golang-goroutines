package golang_goroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			for j := 0; j < 100; j++ {
				atomic.AddInt64(&x, 1)
			}

		}()
	}
	group.Wait()
	fmt.Println("Counter", x)
}
