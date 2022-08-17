package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type RWMutexBankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (r *RWMutexBankAccount) AddBalance(amount int) {
	r.RWMutex.Lock()
	r.Balance += amount
	r.RWMutex.Unlock()
}

func (r *RWMutexBankAccount) GetBalance() int {
	r.RWMutex.RLock()
	balance := r.Balance
	r.RWMutex.RUnlock()
	return balance
}

func TestRWMutexBankAccount(t *testing.T) {
	account := RWMutexBankAccount{}
	iteration1 := 1000
	iteration2 := 100
	expectedBalance := iteration1 * iteration2

	for i := 0; i < iteration1; i++ {
		go func() {
			for j := 0; j < iteration2; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(30 * time.Second)
	if account.GetBalance() != expectedBalance {
		t.Fatal(fmt.Sprintf("Detected race condition. Value: %d, expected value: %d", account.GetBalance(), expectedBalance))
	}
	fmt.Println(fmt.Sprintf("Success, no race conditon. Value: %d, expected value: %d", account.GetBalance(), expectedBalance))
}

type RWMutexTransaction struct {
	RWMutex     sync.RWMutex
	TotalAmount int
}

func (r *RWMutexTransaction) AddAmount(amount int) {
	r.RWMutex.Lock()
	r.TotalAmount += amount
	r.RWMutex.Unlock()
}

func (r *RWMutexTransaction) GetAmount() int {
	r.RWMutex.RLock()
	result := r.TotalAmount
	r.RWMutex.RUnlock()
	return result
}

func TestRWMutexTransaction(t *testing.T) {
	trx1 := RWMutexTransaction{}
	trx2 := RWMutexTransaction{}
	trx3 := RWMutexTransaction{}
	trx1.AddAmount(100)
	trx2.AddAmount(200)
	trx3.AddAmount(300)

	iterationI := 100
	iterationJ := 100
	iterationK := 100
	expectedAmountTrx1 := 100 + iterationI*iterationJ*iterationK
	expectedAmountTrx2 := 200 + iterationI*iterationJ*iterationK
	expextedAmountTrx3 := 300 + iterationI*iterationJ*iterationK

	for i := 0; i < iterationI; i++ {
		for j := 0; j < iterationJ; j++ {
			for k := 0; k < iterationK; k++ {
				go func() {
					trx1.AddAmount(1)
					trx2.AddAmount(1)
					trx3.AddAmount(1)
				}()
			}
		}
	}

	time.Sleep(30 * time.Second)
	if trx1.GetAmount() != expectedAmountTrx1 {
		t.Fatal(fmt.Sprintf("Detected race condition in transaction 1. Value: %d, expected value: %d", trx1.GetAmount(), expectedAmountTrx1))
	}
	if trx2.GetAmount() != expectedAmountTrx2 {
		t.Fatal(fmt.Sprintf("Detected race condition in transaction 2. Value: %d, expected value: %d", trx2.GetAmount(), expectedAmountTrx2))
	}
	if trx3.GetAmount() != expextedAmountTrx3 {
		t.Fatal(fmt.Sprintf("Detected race condition in transaction 3. Value: %d, expected value: %d", trx3.GetAmount(), expextedAmountTrx3))
	}
	message := "Success nor race condition in all transactions \n"
	message += fmt.Sprintf("Transaction 1 | value: %d, expected value: %d \n", trx1.GetAmount(), expectedAmountTrx1)
	message += fmt.Sprintf("Transaction 2 | value: %d, expected value: %d \n", trx2.GetAmount(), expectedAmountTrx2)
	message += fmt.Sprintf("Transaction 3 | value: %d, expected value: %d \n", trx3.GetAmount(), expextedAmountTrx3)
	fmt.Println(message)
}
