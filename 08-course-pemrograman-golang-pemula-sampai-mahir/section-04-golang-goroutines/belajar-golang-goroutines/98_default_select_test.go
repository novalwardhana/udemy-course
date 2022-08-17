package belajar_golang_goroutines

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func GenerateDefaultSelectString(channel chan<- string) {
	channel <- time.Now().Format("150405")
}

func GenerateDefaultSelectFLoat(channel chan<- float64) {
	rand.Seed(time.Now().Unix())
	channel <- rand.Float64()
}

func TestDefaultSelectWithoutLoop(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan float64)
	defer close(channel1)
	defer close(channel2)

	go GenerateDefaultSelectString(channel1)
	go GenerateDefaultSelectFLoat(channel2)

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
		default:
			{
				fmt.Println("Waiting channel")
			}
		}
		if counter == 2 {
			break
		}
	}
}

func GenerateDefaultSelectStringLoop(channnel chan<- string, iteration int) {
	for i := 1; i <= iteration; i++ {
		channnel <- time.Now().Format("150405")
	}
}

func GenerateDefaultSelectFLoatLoop(channel chan<- float64, iteration int) {
	rand.Seed(time.Now().Unix())
	for i := 1; i <= iteration; i++ {
		channel <- rand.Float64()
	}
}

func TestDefaultSelectWithLoop(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan float64)
	defer close(channel1)
	defer close(channel2)
	iteration := 8

	go GenerateDefaultSelectStringLoop(channel1, iteration)
	go GenerateDefaultSelectFLoatLoop(channel2, iteration)
	count := 0
	for {
		select {
		case data := <-channel1:
			{
				fmt.Println("Data from channel 1: ", data)
				count++
			}
		case data := <-channel2:
			{
				fmt.Println("Data from channel 2: ", data)
				count++
			}
		default:
			{
				fmt.Println("Waiting channel")
			}
		}
		if count >= iteration*2 {
			break
		}
	}
}
