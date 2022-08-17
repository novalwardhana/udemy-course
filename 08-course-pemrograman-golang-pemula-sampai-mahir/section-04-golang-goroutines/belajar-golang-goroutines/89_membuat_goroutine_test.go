package belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Haloo Noval")
}

func TestMembuatGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Run")
	time.Sleep(5 * time.Second)
}

func TestMembuatGoroutineAdvance(t *testing.T) {
	for i := 0; i < 5; i++ {
		go RunHelloWorld()
	}
	time.Sleep(5 * time.Second)
}
