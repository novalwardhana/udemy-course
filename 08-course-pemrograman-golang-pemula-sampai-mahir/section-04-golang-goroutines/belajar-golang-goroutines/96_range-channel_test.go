package belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func GenerateRangeChannel(channel chan<- float64, iteration int) {
	for i := 1; i <= iteration; i++ {
		channel <- float64(i) / float64(i+1)
	}
	close(channel)
}

func TestRangeChannelOne(t *testing.T) {
	channel := make(chan float64)
	iteration := 7

	go GenerateRangeChannel(channel, iteration)

	for data := range channel {
		fmt.Println("Data: ", data)
	}
	time.Sleep(5 * time.Second)
}

func TestRangeChannelTwo(t *testing.T) {
	channel := make(chan string)
	iteration := 8

	go func() {
		for i := 1; i <= iteration; i++ {
			channel <- fmt.Sprintf("Iteration: %d, time: %s", i, time.Now().Format("2006-01-02 15:04:05"))
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}
}

func GenerateRangeChannelStringLoop(channel chan<- string, iteration int) {
	for i := 1; i < iteration; i++ {
		channel <- time.Now().Format("20060102-150405")
	}
}

func TestRangeChannelStringLoop(t *testing.T) {
	channel := make(chan string)
	iteration := 7
	defer close(channel)

	go GenerateRangeChannelStringLoop(channel, iteration)
	count := 1
	for data := range channel {
		fmt.Println("Data from channel: ", data)
		count++
		if count >= iteration {
			break
		}
	}
}

func GenerateRangeChannelFLoatLoop(channel chan<- float64, iteration int) {
	for i := 1; i < iteration; i++ {
		channel <- float64(i) / float64(i+1)
	}
}

func TestRangeChannelFLoatLoop(t *testing.T) {
	channel := make(chan float64)
	iteration := 10
	defer close(channel)

	go GenerateRangeChannelFLoatLoop(channel, iteration)
	count := 1
	for data := range channel {
		fmt.Println("Data from channel: ", data)
		count++
		if count >= iteration {
			break
		}
	}

}
