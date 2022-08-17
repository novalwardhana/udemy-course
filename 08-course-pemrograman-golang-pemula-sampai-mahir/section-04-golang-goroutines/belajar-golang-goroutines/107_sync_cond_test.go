package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func SyncCondWait(cond *sync.Cond, value int, wg *sync.WaitGroup) {
	cond.L.Lock()
	cond.Wait()
	fmt.Println("Iteration: ", value)
	cond.L.Unlock()
	wg.Done()
}

func TestSyncCondSignal(t *testing.T) {
	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}
	cond := sync.NewCond(mutex)

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go SyncCondWait(cond, i, wg)
	}

	go func() {
		for i := 0; i <= 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal()
		}
	}()

	wg.Wait()
}

func TestSyncCondBroadcast(t *testing.T) {
	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}
	cond := sync.NewCond(mutex)

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go SyncCondWait(cond, i, wg)
	}

	go func() {
		time.Sleep(5 * time.Second)
		cond.Broadcast()
	}()

	wg.Wait()
}

type SyncCondBankAccount struct {
	Name    string
	Balance float64
}

func SyncCondBankAccountAdd(cond *sync.Cond, account *SyncCondBankAccount, wg *sync.WaitGroup) {
	defer wg.Done()

	cond.L.Lock()
	cond.Wait()
	account.Balance += 1
	cond.L.Unlock()
}

func TestSyncCondBankAccountSignal(t *testing.T) {
	account := &SyncCondBankAccount{
		Name:    "My OVO",
		Balance: 1000000,
	}
	mutex := &sync.Mutex{}
	cond := sync.NewCond(mutex)
	wg := &sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			wg.Add(1)
			go SyncCondBankAccountAdd(cond, account, wg)
		}
	}

	go func() {
		for i := 0; i < 100000; i++ {
			cond.Signal()
		}
	}()

	wg.Wait()
	fmt.Println(fmt.Sprintf("Account balance: %.2f", account.Balance))
}

func SyncCondBroadcastMessage(cond *sync.Cond, indexI int, indexJ int, wg *sync.WaitGroup) {
	cond.L.Lock()
	cond.Wait()
	fmt.Println(fmt.Sprintf("Run index: %d, batch: %d", indexI, indexJ))
	cond.L.Unlock()
	wg.Done()
}

func TestSyncCondBroadcastMessage(t *testing.T) {

	mutex := &sync.Mutex{}
	cond := sync.NewCond(mutex)
	wg := &sync.WaitGroup{}

	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			wg.Add(1)
			go SyncCondBroadcastMessage(cond, i, j, wg)
		}
	}

	go func() {
		time.Sleep(5 * time.Second)
		cond.Broadcast()
	}()

	wg.Wait()
	fmt.Println("Finish")
}
