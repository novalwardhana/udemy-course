package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimeTimer(t *testing.T) {
	timer := time.NewTimer(15 * time.Second)
	fmt.Println(time.Now())

	fmt.Println("Test async 1")
	fmt.Println("Test async 2")
	fmt.Println("Test async 3")

	time := <-timer.C
	fmt.Println(time)
}

func TestTimeAfter(t *testing.T) {
	after := time.After(10 * time.Second)
	fmt.Println(time.Now())

	fmt.Println("Test async 1")
	fmt.Println("Test async 2")
	fmt.Println("Test async 3")

	time := <-after
	fmt.Println(time)
}

func TestTimeAfterFunc(t *testing.T) {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	time.AfterFunc(15*time.Second, func() {
		defer wg.Done()
		fmt.Println("Test after func")
	})
	wg.Wait()

	fmt.Println("Finish")

}

func TestLoopTimeTimer(t *testing.T) {
	fmt.Println("Start at:", time.Now().Format("2006-01-02 15:04:05"))
	timer1 := time.NewTimer(20 * time.Second)
	timer2 := time.NewTimer(10 * time.Second)
	timer3 := time.NewTimer(30 * time.Second)

	fmt.Println("...Run main goroutine...")
	fmt.Println("...Run main goroutine...")
	fmt.Println("...Run main goroutine...")

	counter := 0
	for {
		select {
		case data := <-timer1.C:
			{
				message := fmt.Sprintf("Timer 1, value: %s", data.Format("2006-01-02 15:04:05"))
				fmt.Println(message)
				counter++
			}
		case data := <-timer2.C:
			{
				message := fmt.Sprintf("Timer 2, value: %s", data.Format("2006-01-02 15:04:05"))
				fmt.Println(message)
				counter++
			}
		case data := <-timer3.C:
			{
				message := fmt.Sprintf("Timer 3, value: %s", data.Format("2006-01-02 15:04:05"))
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

	fmt.Println("Finish at:", time.Now().Format("2006-01-02 15:04:05"))
}

func TestLoopTimeAfter(t *testing.T) {
	fmt.Println("Start at:", time.Now().Format("2006-01-02 15:04:05"))
	timeAfter1 := time.After(30 * time.Second)
	timeAfter2 := time.After(3 * time.Second)
	timeAfter3 := time.After(10 * time.Second)

	fmt.Println("...Run main goroutine...")
	fmt.Println("...Run main goroutine...")
	fmt.Println("...Run main goroutine...")

	counter := 0
	for {
		select {
		case data := <-timeAfter1:
			{
				message := fmt.Sprintf("Timer 1, value: %s", data.Format("2006-01-02 15:04:05"))
				fmt.Println(message)
				counter++
			}
		case data := <-timeAfter2:
			{
				message := fmt.Sprintf("Timer 2, value: %s", data.Format("2006-01-02 15:04:05"))
				fmt.Println(message)
				counter++
			}
		case data := <-timeAfter3:
			{
				message := fmt.Sprintf("Timer 3, value: %s", data.Format("2006-01-02 15:04:05"))
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

	fmt.Println("Finish at:", time.Now().Format("2006-01-02 15:04:05"))
}

var timeTimerWG = &sync.WaitGroup{}

func TimeTimerFunc() {
	defer timeTimerWG.Done()

	message := fmt.Sprintf("Run callback at: %s", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(message)
}

func TestMultipleTimeAfterFunc(t *testing.T) {
	fmt.Println("Start at: ", time.Now().Format("2006-01-02 15:04:05"))

	timeTimerWG.Add(5)
	time.AfterFunc(5*time.Second, TimeTimerFunc)
	time.AfterFunc(3*time.Second, TimeTimerFunc)
	time.AfterFunc(2*time.Second, TimeTimerFunc)
	time.AfterFunc(20*time.Second, TimeTimerFunc)
	time.AfterFunc(15*time.Second, TimeTimerFunc)

	timeTimerWG.Wait()
	fmt.Println("Finish at:", time.Now().Format("2006-01-02 15:04:05"))
}
