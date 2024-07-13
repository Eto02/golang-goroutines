package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <-timer.C
	fmt.Println(time)
}

func TestAfter(t *testing.T) {
	timer := time.After(5 * time.Second)
	fmt.Println(time.Now())

	time := <-timer
	fmt.Println(time)
}

func TestAfterFunc(t *testing.T) {
	group = sync.WaitGroup{}
	group.Add(1)
	time.AfterFunc(5*time.Second, func() {
		defer group.Done()
		fmt.Println(time.Now())
	})

	fmt.Println(time.Now())
	group.Wait()
}
