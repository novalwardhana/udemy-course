package belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestBufferedChannelWithoutGoroutine(t *testing.T) {
	channel := make(chan int, 5)
	defer close(channel)

	for i := 1; i <= 5; i++ {
		channel <- i
	}

	for i := 1; i <= 5; i++ {
		fmt.Println(<-channel)
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Finish")
}

func TestBufferedChannelWithGoroutine(t *testing.T) {
	channel := make(chan float64, 6)
	defer close(channel)

	go func() {
		for i := 1; i <= 12; i++ {
			channel <- float64(i)
		}
	}()

	go func() {
		for i := 1; i <= 6; i++ {
			data := fmt.Sprintf("Batch 1,  channel value: %f, capacity: %d, length: %d", <-channel, cap(channel), len(channel))
			fmt.Println(data)
		}
	}()

	go func() {
		for i := 1; i <= 6; i++ {
			data := fmt.Sprintf("Batch 2, channel value: %f, capacity: %d, length: %d", <-channel, cap(channel), len(channel))
			fmt.Println(data)
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("Finish")
}

func SendBufferedChannelAdvance(channel chan string, capacity int) {
	for i := 0; i < capacity; i++ {
		channel <- fmt.Sprintf("Index: %d", i)
	}
}

func ReceiveBufferedChannelAdvance(channel chan string, capacity int, batch int) {
	for i := 0; i < capacity; i++ {
		data := <-channel
		fmt.Println("Batch: ", batch, " : ", data)
	}
}

func TestBufferedChannelAdvance(t *testing.T) {
	capacity := 5
	channel := make(chan string, capacity)
	defer close(channel)

	go SendBufferedChannelAdvance(channel, capacity*2)
	go ReceiveBufferedChannelAdvance(channel, capacity, 1)
	go ReceiveBufferedChannelAdvance(channel, capacity, 2)

	time.Sleep(5 * time.Second)
}
