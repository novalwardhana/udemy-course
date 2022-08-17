package belajar_golang_context

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func ContextDeadlineCounter(ctx context.Context) <-chan int {
	result := make(chan int)
	go func() {
		defer close(result)
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
					result <- counter
					time.Sleep(1 * time.Second)
				}
			}
		}
	}()
	return result
}

func TestContextDeadline(t *testing.T) {
	fmt.Println("Number goroutine:", runtime.NumGoroutine())
	parent := context.Background()
	ctx, deadline := context.WithDeadline(parent, time.Now().Add(20*time.Second))
	defer deadline()

	counter := ContextDeadlineCounter(ctx)
	fmt.Println("Number goroutine:", runtime.NumGoroutine())
	for data := range counter {
		fmt.Println(data)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Total goroutine:", runtime.NumGoroutine())
}

func ContextDeadlineCounterFloat(ctx context.Context) <-chan float64 {
	result := make(chan float64)
	go func() {
		defer close(result)
		var counter float64
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

func TestContextMultiDeadline(t *testing.T) {
	fmt.Println("Total goroutine: ", runtime.NumGoroutine())
	parent := context.Background()

	ctx1, deadline1 := context.WithDeadline(parent, time.Now().Add(10*time.Second))
	ctx2, deadline2 := context.WithDeadline(parent, time.Now().Add(30*time.Second))
	ctx3, deadline3 := context.WithDeadline(parent, time.Now().Add(20*time.Second))
	ctx4, deadline4 := context.WithDeadline(parent, time.Now().Add(12*time.Second))
	ctx5, deadline5 := context.WithDeadline(parent, time.Now().Add(18*time.Second))
	defer deadline1()
	defer deadline2()
	defer deadline3()
	defer deadline4()
	defer deadline5()

	counter1 := ContextDeadlineCounterFloat(ctx1)
	counter2 := ContextDeadlineCounterFloat(ctx2)
	counter3 := ContextDeadlineCounterFloat(ctx3)
	counter4 := ContextDeadlineCounterFloat(ctx4)
	counter5 := ContextDeadlineCounterFloat(ctx5)
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
