package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	x := 0

	var mutex sync.Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Counter", x)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (acc *BankAccount) AddBalance(amount int) {
	acc.RWMutex.Lock()
	acc.Balance = acc.Balance + amount
	acc.RWMutex.Unlock()
}

func (acc *BankAccount) GetBalance() int {
	acc.RWMutex.Lock()
	balance := acc.Balance
	acc.RWMutex.Unlock()
	return balance
}

func TestReadWriteMutex(t *testing.T) {
	acc := BankAccount{}
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				acc.AddBalance(1)
				fmt.Println(acc.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Final Balance :", acc.GetBalance())
}

type UserBalance struct {
	sync.Mutex
	Balance int
	Name    string
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock user1", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user2", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadLock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Tahta",
		Balance: 10000,
	}
	user2 := UserBalance{
		Name:    "Mubarak",
		Balance: 10000,
	}
	go Transfer(&user1, &user2, 1000)
	go Transfer(&user2, &user1, 2000)
	time.Sleep(10 * time.Second)

	fmt.Println("User ", user1.Name, " Balance ", user1.Balance)
	fmt.Println("User ", user2.Name, " Balance ", user2.Balance)
}
