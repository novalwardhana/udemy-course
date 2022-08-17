package belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestRaceConditionHighIteration(t *testing.T) {
	var counter int
	var iteration1 int = 1000
	var iteration2 int = 100
	var counterExpected int = iteration1 * iteration2

	for i := 0; i < iteration1; i++ {
		go func() {
			for j := 0; j < iteration2; j++ {
				counter++
			}
		}()
	}

	time.Sleep(10 * time.Second)
	if counter != counterExpected {
		t.Fatal(fmt.Sprintf("Detected race condition. Value: %d, expected value: %d", counter, counterExpected))
	}
	fmt.Println("Success, no race condition")
}

func TestRaceConditionLowIteration(t *testing.T) {
	var counter int
	var iteration int = 1000
	var counterExpected = iteration

	for i := 0; i < iteration; i++ {
		go func() {
			counter++
		}()
	}

	time.Sleep(10 * time.Second)
	if counter != counterExpected {
		t.Fatal(fmt.Sprintf("Detected race condition. Value: %d, expected value: %d", counter, counterExpected))
	}
	fmt.Println("Success, no race condition")
}
