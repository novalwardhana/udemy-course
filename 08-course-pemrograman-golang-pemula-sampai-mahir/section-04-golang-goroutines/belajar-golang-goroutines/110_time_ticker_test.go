package belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(5 * time.Second)
	breakTicker := make(chan bool)

	go func() {
		time.Sleep(16 * time.Second)
		ticker.Stop()
		breakTicker <- true
	}()

	counter := 0
	for {
		select {
		case data := <-ticker.C:
			{
				fmt.Println("Ticker: ", data.Format("2006-01-02 15:04:05"))
				counter++
			}
		default:
			{
				//
			}
		}
		if counter == 3 {
			break
		}
	}

	fmt.Println("Finish")
}

func TestTickOnly(t *testing.T) {
	tick := time.Tick(5 * time.Second)

	go func() {
		time.Sleep(16 * time.Second)

	}()

	counter := 0
	for {
		select {
		case data := <-tick:
			{
				message := fmt.Sprintf("Data from tick, value: %s", data.Format("2006-01-02 15:04:05"))
				fmt.Println(message)
				counter++
			}
		default:
			{
				//
			}
		}
		if counter == 3 {
			break
		}
	}

	fmt.Println("Finish")
}

func TestMultipleTicker(t *testing.T) {
	ticker1 := time.NewTicker(5 * time.Second)
	ticker2 := time.NewTicker(2 * time.Second)
	ticker3 := time.NewTicker(7 * time.Second)

	counter := 0
	maxCounter := 15
	for {
		select {
		case data := <-ticker1.C:
			{
				message := fmt.Sprintf("Data ticker 1, value: %s", data.Format("2006-01-02 15:04:05"))
				fmt.Println(message)
				counter++
			}
		case data := <-ticker2.C:
			{
				message := fmt.Sprintf("Data ticker 2, value: %s", data.Format("2006-01-02 15:04:05"))
				fmt.Println(message)
				counter++
			}
		case data := <-ticker3.C:
			{
				message := fmt.Sprintf("Data ticker 3, value: %s", data.Format("2006-01-02 15:04:05"))
				fmt.Println(message)
				counter++
			}
		default:
			{
				//
			}
		}
		if counter >= maxCounter {
			ticker1.Stop()
			ticker2.Stop()
			ticker3.Stop()
			break
		}
	}

	fmt.Println("Finish")
}

func TestMultipleTickOnly(t *testing.T) {
	tick1 := time.Tick(9 * time.Second)
	tick2 := time.Tick(2 * time.Second)
	tick3 := time.Tick(6 * time.Second)

	counter := 0
	maxCounter := 25
	for {
		select {
		case data := <-tick1:
			{
				message := fmt.Sprintf("Data tick 1, value: %s", data.Format("2006-01-02 15:04:05"))
				fmt.Println(message)
				counter++
			}
		case data := <-tick2:
			{
				message := fmt.Sprintf("Data tick 2, value: %s", data.Format("2006-01-02 15:04:05"))
				fmt.Println(message)
				counter++
			}
		case data := <-tick3:
			{
				message := fmt.Sprintf("Data tick 3, value: %s", data.Format("2006-01-02 15:04:05"))
				fmt.Println(message)
				counter++
			}
		default:
			{
				//
			}
		}
		if counter >= maxCounter {
			break
		}
	}

	fmt.Println("Finish")
}
