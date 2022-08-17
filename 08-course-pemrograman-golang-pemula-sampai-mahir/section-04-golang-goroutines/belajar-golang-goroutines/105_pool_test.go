package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPoolString(t *testing.T) {
	var pool sync.Pool
	pool.Put("Noval")
	pool.Put("Kusuma")
	pool.Put("Wardhana")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println("Data: ", data)
			pool.Put(data)
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Finish")
}

func TestPoolWithoutPut(t *testing.T) {
	var pool sync.Pool
	pool.Put(1)
	pool.Put("Noval")
	pool.Put(12)
	pool.Put([]byte("Test byte"))

	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		index := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			data := pool.Get()
			fmt.Println(fmt.Sprintf("Iteration %d, value: %v", index, data))
		}()
	}
	wg.Wait()

	fmt.Println("Finish")
}

type PoolBankAccount struct {
	Name    string
	Balance float64
}

type PoolEwalletAccount struct {
	Name    string
	Balance float64
}

func TestPoolMultipleData(t *testing.T) {
	var pool = sync.Pool{
		New: func() interface{} {
			return "New data"
		},
	}
	pool.Put(1)
	pool.Put("Noval")
	pool.Put("Kusuma")
	pool.Put(12.5)
	pool.Put(22.9)
	pool.Put(true)
	pool.Put("Wardhana")
	pool.Put(PoolBankAccount{
		Name:    "Mandiri",
		Balance: 120000,
	})
	pool.Put(PoolEwalletAccount{
		Name:    "OVO",
		Balance: 15000,
	})
	pool.Put(false)
	pool.Put(true)
	pool.Put(30)
	pool.Put(32)
	pool.Put(&PoolEwalletAccount{
		Name:    "Dana",
		Balance: 300000,
	})
	pool.Put(&PoolBankAccount{
		Name:    "BPD",
		Balance: 32000,
	})

	wg := &sync.WaitGroup{}
	for i := 0; i < 30; i++ {
		index := i
		wg.Add(1)
		go func() {
			defer wg.Done()

			data := pool.Get()
			fmt.Println(fmt.Sprintf("Iteration: %d, value: %v", index, data))
			pool.Put(data)
		}()
	}
	wg.Wait()

	fmt.Println("Finish")
}

func TestPoolAutomated(t *testing.T) {
	var pool = sync.Pool{
		New: func() interface{} {
			return "New data"
		},
	}
	wgGenerated := &sync.WaitGroup{}

	for i := 1; i <= 10; i++ {
		index := i
		wgGenerated.Add(1)
		go func() {
			defer wgGenerated.Done()

			data := float64(index) / float64(index+1)
			pool.Put(data)
		}()
	}
	wgGenerated.Wait()
	fmt.Println("Finish generate pool data")

	wgIteration := &sync.WaitGroup{}
	for i := 0; i <= 50; i++ {
		index := i
		wgIteration.Add(1)
		go func() {
			defer wgIteration.Done()

			data := pool.Get()
			fmt.Println(fmt.Sprintf("Iteration: %d, value: %v", index, data))
			pool.Put(data)
		}()

	}
	wgIteration.Wait()
	fmt.Println("Finish iteration")

}
