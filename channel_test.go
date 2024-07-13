package golang_goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Tahta Failah Mubarak"
	fmt.Println("Selesai Mengirim Data ke Channel")

}
func TestCreateCahnnel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Tahta Failah Mubarak"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(2 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Tahta"
		channel <- "Failah"
		channel <- "Mubarak"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)

	}()
	time.Sleep(2 * time.Second)

	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string, 3)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perlunagan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for v := range channel {
		fmt.Println("Menerima data", v)
	}

	fmt.Println("Selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)
	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari Channel 1:", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari Channel 2:", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)
	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari Channel 1:", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari Channel 2:", data)
			counter++

		default:
			fmt.Println("Menunggu data")
		}
		if counter == 2 {
			break
		}
	}
}

func TestRaceCondition(t *testing.T) {
	x := 0
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				x = x + 1
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Counter", x)
}
