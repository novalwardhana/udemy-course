package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestSyncMutexHighIteration(t *testing.T) {
	var counter int
	var counterMutex sync.Mutex
	var iteration1 int = 1000
	var iteration2 int = 100
	var counterExpected = iteration1 * iteration2

	for i := 0; i < iteration1; i++ {
		go func() {
			for j := 0; j < iteration2; j++ {
				counterMutex.Lock()
				counter++
				counterMutex.Unlock()
			}
		}()
	}

	time.Sleep(10 * time.Second)
	if counter != counterExpected {
		t.Fatal(fmt.Sprintf("Detected race condition. Value: %d, expected value: %d", counter, counterExpected))
	}
	fmt.Println(fmt.Sprintf("Success, no race condition. Value: %d, expected value: %d", counter, counterExpected))
}

func TestSyncMutexLowIteration(t *testing.T) {
	var counter int
	var counterMutex sync.Mutex
	var iteration int = 1000
	var counterExpected int = iteration

	for i := 0; i < iteration; i++ {
		go func() {
			counterMutex.Lock()
			counter++
			counterMutex.Unlock()
		}()
	}

	time.Sleep(10 * time.Second)
	if counter != counterExpected {
		t.Fatal(fmt.Sprintf("Detected race condition. Value: %d, expected value: %d", counter, counterExpected))
	}
	fmt.Println(fmt.Sprintf("Success, no race condition. Value: %d, expected value: %d", counter, counterExpected))
}

func TestSyncMutexManyGoroutine(t *testing.T) {
	var counter int
	var counterMutex sync.Mutex
	var iteration1 int = 1000
	var iteration2 int = 100
	var iteration3 int = 100
	var counterExpected = iteration1 * iteration2 * iteration3

	for i := 0; i < iteration1; i++ {
		for j := 0; j < iteration2; j++ {
			for k := 0; k < iteration3; k++ {
				go func() {
					counterMutex.Lock()
					counter++
					counterMutex.Unlock()
				}()
			}
		}
	}

	time.Sleep(60 * time.Second)
	if counter != counterExpected {
		t.Fatal(fmt.Sprintf("Detected race condition. Value: %d, expected value: %d", counter, counterExpected))
	}
	fmt.Println(fmt.Sprintf("Success, no race condition. Value: %d, expected value: %d", counter, counterExpected))
}
