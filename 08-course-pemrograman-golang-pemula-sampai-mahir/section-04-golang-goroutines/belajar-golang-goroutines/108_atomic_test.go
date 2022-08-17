package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestWithoutAtomic(t *testing.T) {
	counter := 0
	wg := &sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				counter++
			}
		}()
	}
	wg.Wait()

	fmt.Println(counter)
}

func TestWithAtomic(t *testing.T) {
	counter := int64(0)
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for j := 0; j < 100; j++ {
				atomic.AddInt64(&counter, int64(1))
			}
		}()
	}
	wg.Wait()

	fmt.Println(counter)
}

func TestAtomicInt32(t *testing.T) {
	var counter int32
	var counterExpected int32 = 100000
	var wg = &sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				atomic.AddInt32(&counter, int32(1))
			}
		}()
	}

	wg.Wait()
	if counter != counterExpected {
		t.Fatal(fmt.Sprintf("Detected race condition. value: %d, expected value: %d", counter, counterExpected))
	}
	fmt.Println(fmt.Sprintf("Success. Value: %d, expected value: %d", counter, counterExpected))
}
