package belajar_golang_context

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func ContextTimeoutCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 0
		for {
			select {
			case <-ctx.Done():
				{
					return
				}
			default:
				{
					counter++
					destination <- counter
					time.Sleep(1 * time.Second)
				}
			}
		}
	}()

	return destination
}

func TestContextTimeout(t *testing.T) {
	fmt.Println("Total goroutine: ", runtime.NumGoroutine())
	parent := context.Background()
	ctx, timeout := context.WithTimeout(parent, 20*time.Second)
	defer timeout()

	counter := ContextTimeoutCounter(ctx)
	fmt.Println("Total goroutine: ", runtime.NumGoroutine())
	for data := range counter {
		fmt.Println("Value: ", data)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Total goroutine: ", runtime.NumGoroutine())
}

func ContextTimeoutCounterFLoat(ctx context.Context) <-chan float64 {
	result := make(chan float64)
	go func() {
		defer close(result)
		counter := float64(0)
		for {
			select {
			case <-ctx.Done():
				{
					return
				}
			default:
				{
					counter++
					result <- counter
					time.Sleep(1 * time.Second)
				}
			}
		}
	}()
	return result
}

func TestContextTimeoutMultiCounter(t *testing.T) {
	fmt.Println("Total goroutine: ", runtime.NumGoroutine())

	parent := context.Background()
	ctx1, timeout1 := context.WithTimeout(parent, 5*time.Second)
	ctx2, timeout2 := context.WithTimeout(parent, 20*time.Second)
	ctx3, timeout3 := context.WithTimeout(parent, 6*time.Second)
	ctx4, timeout4 := context.WithTimeout(parent, 10*time.Second)
	ctx5, timeout5 := context.WithTimeout(parent, 25*time.Second)
	defer timeout1()
	defer timeout2()
	defer timeout3()
	defer timeout4()
	defer timeout5()

	counter1 := ContextTimeoutCounterFLoat(ctx1)
	counter2 := ContextTimeoutCounterFLoat(ctx2)
	counter3 := ContextTimeoutCounterFLoat(ctx3)
	counter4 := ContextTimeoutCounterFLoat(ctx4)
	counter5 := ContextTimeoutCounterFLoat(ctx5)
	fmt.Println("Total goroutine: ", runtime.NumGoroutine())

	wg := &sync.WaitGroup{}
	wg.Add(5)

	go func() {
		defer wg.Done()
		for data := range counter1 {
			message := fmt.Sprintf("Data from counter1: %v", data)
			fmt.Println(message)
		}
	}()

	go func() {
		defer wg.Done()
		for data := range counter2 {
			message := fmt.Sprintf("Data from counter2: %v", data)
			fmt.Println(message)
		}
	}()

	go func() {
		defer wg.Done()
		for data := range counter3 {
			message := fmt.Sprintf("Data from counter3: %v", data)
			fmt.Println(message)
		}
	}()

	go func() {
		defer wg.Done()
		for data := range counter4 {
			message := fmt.Sprintf("Data from counter4: %v", data)
			fmt.Println(message)
		}
	}()

	go func() {
		defer wg.Done()
		for data := range counter5 {
			message := fmt.Sprintf("Data from counter5: %v", data)
			fmt.Println(message)
		}
	}()

	wg.Wait()

	time.Sleep(5 * time.Second)
	fmt.Println("Total goroutine: ", runtime.NumGoroutine())
}
