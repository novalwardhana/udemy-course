package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

var OnlyOnceCounter int

func OnlyOnce() {
	OnlyOnceCounter++
}

func TestOnceCounter(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			group.Add(1)
			go func() {
				defer group.Done()

				once.Do(OnlyOnce)
			}()
		}
	}

	group.Wait()
	fmt.Println("OnlyOnceCounter: ", OnlyOnceCounter)
}

func OnceFinishRandom() {
	fmt.Println("Finish process")
}

func TestOnceRandom(t *testing.T) {
	once := &sync.Once{}
	wg := &sync.WaitGroup{}
	var indexI int
	var indexJ int

	for i := 1; i <= 100; i++ {
		for j := 0; j <= 100; j++ {
			wg.Add(1)
			i := i
			j := j
			go func() {
				defer wg.Done()

				indexI = i
				indexJ = j
				once.Do(OnceFinishRandom)
			}()
		}
	}

	wg.Wait()
	fmt.Println("Random index I: ", indexI)
	fmt.Println("Random index J: ", indexJ)
}
