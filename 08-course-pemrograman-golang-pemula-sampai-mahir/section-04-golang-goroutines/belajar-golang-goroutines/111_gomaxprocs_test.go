package belajar_golang_goroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGomaxprocs(t *testing.T) {
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(5 * time.Second)
		}()
	}

	totalCPU := runtime.NumCPU()
	fmt.Println("Total CPU: ", totalCPU)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total thread: ", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total goroutine: ", totalGoroutine)

	wg.Wait()
}

func TestGomaxprocsSetThread(t *testing.T) {
	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU: ", totalCpu)

	runtime.GOMAXPROCS(30)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread: ", totalThread)

	wg := &sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(5 * time.Second)
		}()
	}

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total goroutine: ", totalGoroutine)

	wg.Wait()

	totalGoroutineNow := runtime.NumGoroutine()
	fmt.Println("Total goroutine: ", totalGoroutineNow)
}

func TestGomaxprocsSetMillionGoroutine(t *testing.T) {
	numCPU := runtime.NumCPU()
	fmt.Println("Number of CPU: ", numCPU)

	runtime.GOMAXPROCS(20)
	numThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Number of Thread: ", numThread)

	wg := &sync.WaitGroup{}
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(15 * time.Second)
		}()
	}

	numGoroutine := runtime.NumGoroutine()
	fmt.Println("Number of Goroutine: ", numGoroutine)
	wg.Wait()
}
