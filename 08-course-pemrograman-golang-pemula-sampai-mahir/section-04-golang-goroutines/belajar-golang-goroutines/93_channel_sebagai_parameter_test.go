package belajar_golang_goroutines

import (
	"fmt"
	"testing"
)

func channelParameter(channel chan string) {
	channel <- "Novalita Kusuma Wardhana"
}

func TestChannelSebagaiParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go channelParameter(channel)

	data := <-channel
	fmt.Println(data)
}

func TestChannelSebagaiParameterAdvance(t *testing.T) {

	type person struct {
		Name    string
		Age     int
		Address string
	}
	channel := make(chan person)

	go func() {
		channel <- person{"Noval", 29, "Bantul"}
	}()

	data := <-channel
	fmt.Println(data)

}
