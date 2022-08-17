package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunWaitGroup(group *sync.WaitGroup, index int) {
	group.Add(1)
	defer group.Done()

	fmt.Println("Run gorutine:", index)
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunWaitGroup(group, i)
	}
	group.Wait()

	fmt.Println("Finish")
}

type WaitGroupBankAccount struct {
	Mutex   sync.Mutex
	Name    string
	Balance float64
}

func (w *WaitGroupBankAccount) Add(amount float64) {
	w.Balance = w.Balance + amount
}

func WaitGroupTransfer(account *WaitGroupBankAccount, amount float64, group *sync.WaitGroup) {
	defer group.Done()

	account.Mutex.Lock()
	account.Add(amount)
	account.Mutex.Unlock()
}

func TestWaitGroupBankAccount(t *testing.T) {
	group := &sync.WaitGroup{}
	account := &WaitGroupBankAccount{
		Name:    "OVO",
		Balance: 100000,
	}
	iterationI := 1000
	iterationJ := 100
	expectedBalance := account.Balance + float64(iterationI)*float64(iterationJ)

	for i := 0; i < iterationI; i++ {
		for j := 0; j < iterationJ; j++ {
			group.Add(1) // Harus diletakkan di goroutine utama
			go WaitGroupTransfer(account, 1, group)
		}
	}

	group.Wait()
	if account.Balance != expectedBalance {
		message := fmt.Sprintf("Balance: %f, expected balance: %f", account.Balance, expectedBalance)
		t.Fatal("Race condition detected.", message)
	}
	message := fmt.Sprintf("Balance: %f, expected balance: %f", account.Balance, expectedBalance)
	fmt.Println("No race condition.", message)
}
