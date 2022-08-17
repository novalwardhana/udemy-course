package belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestMembuatChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(5 * time.Second)
		channel <- fmt.Sprintf("Novalwardhana Receive at: %s", time.Now().Format("2006-01-02 15:04:05"))
	}()

	data := <-channel
	fmt.Println(data)
}

func TestMembuatChannelTanpaPenerima(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		channel <- fmt.Sprintf("Novalwardhana receive at: %s", time.Now().Format("2006-01-02 15:04:05"))
	}()

	time.Sleep(5 * time.Second)
}

func TestMembuatChannelTanpaPengirim(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {

	}()

	data := <-channel
	fmt.Println(data)
}
