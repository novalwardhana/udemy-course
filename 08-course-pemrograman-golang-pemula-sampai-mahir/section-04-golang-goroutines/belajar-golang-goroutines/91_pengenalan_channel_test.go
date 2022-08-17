package belajar_golang_goroutines

import (
	"fmt"
	"testing"
)

func TestPengenalanChannel(t *testing.T) {
	channel := make(chan string)

	channel <- "Noval wardhana"

	fmt.Println(<-channel)
}
