package belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func InChannel(channel chan<- string) {
	channel <- "Novalita Kusuma Wardhana"
}

func OutChannel(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go InChannel(channel)
	go OutChannel(channel)

	time.Sleep(5 * time.Second)
}

func OutChannelReturn(param string) <-chan string {
	result := make(chan string)
	go func() {
		defer close(result)

		result <- param
	}()
	return result
}

func TestOutChannelReturn(t *testing.T) {
	result := <-OutChannelReturn("Halo Noval")
	fmt.Println(result)
}

func InChannelAdvance(channel chan<- string, iteration int) {
	for i := 0; i < iteration; i++ {
		channel <- fmt.Sprintf("Index: %d", i)
	}
}

func OutChannelAdvance(channel <-chan string, iteration int) {
	for i := 0; i < iteration; i++ {
		data := <-channel
		fmt.Println(data)
	}
}

func TestInOutChannelAdvance(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	iteration := 5

	go InChannelAdvance(channel, iteration)
	go OutChannelAdvance(channel, iteration)

	time.Sleep(5 * time.Second)
}
