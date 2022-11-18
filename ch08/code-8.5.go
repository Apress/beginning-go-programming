package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var balanceAmount int
var mutex = &sync.Mutex{}

func creditToBalance() {
	for i := 0; i < 5; i++ {
		mutex.Lock()
		balanceAmount += 1000
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		fmt.Println("Balance After Credit: ", balanceAmount)
		mutex.Unlock()
	}
}
func debitFromBalance() {
	for i := 0; i < 5; i++ {
		mutex.Lock()
		balanceAmount -= 1000
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		fmt.Println("Balance After Debit: ", balanceAmount)
		mutex.Unlock()
	}
}
func main() {
	balanceAmount = 3000
	fmt.Println("Initial balance: ", balanceAmount)
	go creditToBalance()
	go debitFromBalance()
	fmt.Scanln()
}

/*
OUTPUT
Initial balance:  3000
Balance After Credit:  4000
Balance After Credit:  5000
Balance After Debit:  4000
Balance After Debit:  3000
Balance After Credit:  4000
Balance After Credit:  5000
Balance After Debit:  4000
Balance After Debit:  3000
Balance After Credit:  4000
Balance After Debit:  3000
*/
