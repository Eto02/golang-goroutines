package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
		done <- true
	}()

	for {
		select {
		case t := <-ticker.C:
			fmt.Println(t)
		case <-done:
			fmt.Println("Ticker stopped")
			return
		}
	}
}

func TestTick(t *testing.T) {
	ticker := time.Tick(1 * time.Second)
	stop := time.After(5 * time.Second)

	for {
		select {
		case t := <-ticker:
			fmt.Println(t)
		case <-stop:
			fmt.Println("Ticker stopped")
			return
		}
	}
}
