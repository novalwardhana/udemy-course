package belajar_golang_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func ContextWithCancelCreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				{
					return
				}
			default:
				{
					destination <- counter
					counter++
					time.Sleep(1 * time.Second)
				}
			}
		}
	}()
	return destination
}

func TestCancelWithContextCreateCounter(t *testing.T) {

	fmt.Println("Num goroutine: ", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	counter := ContextWithCancelCreateCounter(ctx)
	for data := range counter {
		fmt.Println(data)
		if data == 10 {
			break
		}
	}

	cancel()
	time.Sleep(5 * time.Second)

	fmt.Println("Num goroutine: ", runtime.NumGoroutine())

}

func ContextWithCancelCreateFloat(ctx context.Context) chan float64 {
	result := make(chan float64)

	go func() {
		defer close(result)
		counter := float64(1)
		for {
			select {
			case <-ctx.Done():
				{
					return
				}
			default:
				{
					result <- counter
					counter++
					time.Sleep(1 * time.Second)
				}
			}
		}
	}()

	return result
}

func TestContextWithMultipleChannel(t *testing.T) {
	fmt.Println("Num goroutine: ", runtime.NumGoroutine())

	parent := context.Background()
	ctx1, cancel1 := context.WithCancel(parent)
	ctx2, cancel2 := context.WithCancel(parent)
	ctx3, cancel3 := context.WithCancel(parent)
	ctx4, cancel4 := context.WithCancel(parent)
	ctx5, cancel5 := context.WithCancel(parent)
	ctx6, cancel6 := context.WithCancel(parent)
	ctx7, cancel7 := context.WithCancel(parent)

	counter1 := ContextWithCancelCreateFloat(ctx1)
	counter2 := ContextWithCancelCreateFloat(ctx2)
	counter3 := ContextWithCancelCreateFloat(ctx3)
	counter4 := ContextWithCancelCreateFloat(ctx4)
	counter5 := ContextWithCancelCreateFloat(ctx5)
	counter6 := ContextWithCancelCreateFloat(ctx6)
	counter7 := ContextWithCancelCreateFloat(ctx7)

	traceCounter := 0
	maxCounter := 35
	for {
		select {
		case data := <-counter1:
			{
				message := fmt.Sprintf("Data from counter 1: %f", data)
				fmt.Println(message)
				traceCounter++
			}
		case data := <-counter2:
			{
				message := fmt.Sprintf("Data from counter 2: %f", data)
				fmt.Println(message)
				traceCounter++
			}
		case data := <-counter3:
			{
				message := fmt.Sprintf("Data from counter 3: %f", data)
				fmt.Println(message)
				traceCounter++
			}
		case data := <-counter4:
			{
				message := fmt.Sprintf("Data from counter 4: %f", data)
				fmt.Println(message)
				traceCounter++
			}
		case data := <-counter5:
			{
				message := fmt.Sprintf("Data from counter 5: %f", data)
				fmt.Println(message)
				traceCounter++
			}
		case data := <-counter6:
			{
				message := fmt.Sprintf("Data from counter 6: %f", data)
				fmt.Println(message)
				traceCounter++
			}
		case data := <-counter7:
			{
				message := fmt.Sprintf("Data from counter 7: %f", data)
				fmt.Println(message)
				traceCounter++
			}
		}

		if traceCounter >= maxCounter {
			break
		}
	}

	cancel1()
	cancel2()
	cancel3()
	cancel4()
	cancel5()
	cancel6()
	cancel7()

	time.Sleep(50 * time.Second)

	fmt.Println("Num goroutine: ", runtime.NumGoroutine())
}
