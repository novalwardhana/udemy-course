package belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func GenerateSelectChannelString(channel chan<- string, param string) {
	channel <- param
}

func GenerateSelectChannelFloat(channel chan<- float64, param float64) {
	channel <- param
}

func TestSelectChannelSameType(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GenerateSelectChannelString(channel1, "Real Madrid")
	go GenerateSelectChannelString(channel2, "Barcelona")

	counter := 0
	for {
		select {
		case data := <-channel1:
			{
				fmt.Println("Data from channel 1: ", data)
				counter++
			}
		case data := <-channel2:
			{
				fmt.Println("Data from channel 2: ", data)
				counter++
			}
		}
		if counter == 2 {
			break
		}
	}
}

func TestSelectChannelDifferentType(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan float64)
	defer close(channel1)
	defer close(channel2)

	go GenerateSelectChannelString(channel1, "Arsenal")
	go GenerateSelectChannelFloat(channel2, 75.59)

	counter := 0
	for {
		select {
		case data := <-channel1:
			{
				fmt.Println("Data from channel 1: ", data)
				counter++
			}
		case data := <-channel2:
			{
				fmt.Println("Data from channel 2: ", data)
				counter++
			}
		}
		if counter == 2 {
			break
		}
	}
}

func GenerateSelectChannelStringLoop(channel chan<- string, iteration int) {
	for i := 1; i <= iteration; i++ {
		channel <- time.Now().Format("20060102-150405")
	}
}

func GenerateSelectChannelFloatLoop(channel chan<- float64, iteration int) {
	for i := 1; i <= iteration; i++ {
		channel <- float64(i) / float64(i+1)
	}
}

func TestSelectChannelWithLoop(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan float64)
	defer close(channel1)
	defer close(channel2)
	iteration := 5

	go GenerateSelectChannelStringLoop(channel1, iteration)
	go GenerateSelectChannelFloatLoop(channel2, iteration)

	counter := 0
	for {
		select {
		case data := <-channel1:
			{
				fmt.Println("Data from channel 1: ", data)
				counter++
			}
		case data := <-channel2:
			{
				fmt.Println("Data from channel 2: ", data)
				counter++
			}
		}
		if counter == iteration*2 {
			break
		}
	}

}
