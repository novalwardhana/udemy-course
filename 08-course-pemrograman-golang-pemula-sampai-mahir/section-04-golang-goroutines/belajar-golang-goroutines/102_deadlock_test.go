package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type DeadlockBankAccount struct {
	sync.Mutex
	Name    string
	Balance float64
}

func (d *DeadlockBankAccount) Lock() {
	d.Mutex.Lock()
}

func (d *DeadlockBankAccount) Unlock() {
	d.Mutex.Unlock()
}

func (d *DeadlockBankAccount) Change(amount float64) {
	d.Balance = d.Balance + amount
}

func DeadlockTransferBankAccount(senderAccount *DeadlockBankAccount, receiverAccount *DeadlockBankAccount, amount float64) {
	senderAccount.Lock()
	fmt.Println("Lock sender: ", senderAccount.Name)
	senderAccount.Change(-1 * amount)
	time.Sleep(1 * time.Second)

	receiverAccount.Lock()
	fmt.Println("Lock receiver: ", receiverAccount.Name)
	receiverAccount.Change(amount)
	time.Sleep(1 * time.Second)

	senderAccount.Unlock()
	receiverAccount.Unlock()
}

func TestDeadlockBankAccount(t *testing.T) {
	account1 := &DeadlockBankAccount{
		Name:    "Account 1",
		Balance: 1000000,
	}
	account2 := &DeadlockBankAccount{
		Name:    "Account 2",
		Balance: 1000000,
	}

	go DeadlockTransferBankAccount(account1, account2, 100000)
	go DeadlockTransferBankAccount(account2, account1, 200000)

	time.Sleep(10 * time.Second)
	fmt.Println(account1.Name, ": ", account1.Balance)
	fmt.Println(account2.Name, ": ", account2.Balance)
}

type Ewallet struct {
	Mutex   sync.Mutex
	Name    string
	Balance float64
}

func (e *Ewallet) Transfer(amount float64) {
	e.Balance = e.Balance + amount
}

func DeadlockTransferEwallet(sender *Ewallet, receiver *Ewallet, amount float64) {
	sender.Mutex.Lock()
	fmt.Println("Ewallet lock sender: ", sender.Name)
	sender.Transfer(-amount)

	receiver.Mutex.Lock()
	fmt.Println("Ewallet lock receiver: ", receiver.Name)
	receiver.Transfer(amount)

	sender.Mutex.Unlock()
	receiver.Mutex.Unlock()
}

func TestDeadlockEwallet(t *testing.T) {
	ewallet1 := &Ewallet{
		Name:    "OVO",
		Balance: 150000,
	}
	ewallet2 := &Ewallet{
		Name:    "Gopay",
		Balance: 250000,
	}

	go DeadlockTransferEwallet(ewallet1, ewallet2, 20000)
	go DeadlockTransferEwallet(ewallet2, ewallet1, 50000)
	var ewallet1ExpectedBalance float64 = 180000
	var ewallet2ExpectedBalance float64 = 220000

	time.Sleep(5 * time.Second)
	message := fmt.Sprintf("Ewallet 1: %s, Balance: %f, Expected balance: %f\n", ewallet1.Name, ewallet1.Balance, ewallet1ExpectedBalance)
	message += fmt.Sprintf("Ewallet 2: %s, Balance: %f, Expected balance: %f\n", ewallet2.Name, ewallet2.Balance, ewallet2ExpectedBalance)
	if ewallet1.Balance != ewallet1ExpectedBalance || ewallet2.Balance != ewallet2ExpectedBalance {
		t.Fatal(fmt.Sprintf("Detected deadlock !!!.\n%s", message))
	}
	fmt.Println(fmt.Sprintf("Success, no deadlock.\n%s", message))
}

func SolveDeadlockTransferEwallet(sender *Ewallet, receiver *Ewallet, amount float64) {
	sender.Mutex.Lock()
	fmt.Println("Ewallet lock sender: ", sender.Name)
	sender.Transfer(-amount)
	sender.Mutex.Unlock()

	receiver.Mutex.Lock()
	fmt.Println("Ewallet lock receiver: ", receiver.Name)
	receiver.Transfer(amount)
	receiver.Mutex.Unlock()
}

func TestSolveDeadlockEwallet(t *testing.T) {
	ewallet1 := &Ewallet{
		Name:    "OVO",
		Balance: 150000,
	}
	ewallet2 := &Ewallet{
		Name:    "Gopay",
		Balance: 250000,
	}

	go SolveDeadlockTransferEwallet(ewallet1, ewallet2, 20000)
	go SolveDeadlockTransferEwallet(ewallet2, ewallet1, 50000)
	var ewallet1ExpectedBalance float64 = 180000
	var ewallet2ExpectedBalance float64 = 220000

	time.Sleep(5 * time.Second)
	message := fmt.Sprintf("Ewallet 1: %s, Balance: %f, Expected balance: %f\n", ewallet1.Name, ewallet1.Balance, ewallet1ExpectedBalance)
	message += fmt.Sprintf("Ewallet 1: %s, Balance: %f, Expected balance: %f\n", ewallet2.Name, ewallet2.Balance, ewallet2ExpectedBalance)
	if ewallet1.Balance != ewallet1ExpectedBalance || ewallet2.Balance != ewallet2ExpectedBalance {
		t.Fatal(fmt.Sprintf("Detected deadlock !!!.\n%s", message))
	}
	fmt.Println(fmt.Sprintf("Success, no deadlock.\n%s", message))
}
