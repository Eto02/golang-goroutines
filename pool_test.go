package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() any {
			return "New"
		},
	}

	pool.Put("Tahta")
	pool.Put("Failah")
	pool.Put("Mubarak")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}
	time.Sleep(11 * time.Second)
	fmt.Println("Selesai")
}
