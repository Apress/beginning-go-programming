package main

import (
	"fmt"
	"math/rand"
	"time"
)

var balanceAmount int

func creditToBalance() {
	for i := 0; i < 5; i++ {
		balanceAmount += 1000
		time.Sleep(time.Duration(rand.Intn(100)) *
			time.Millisecond)
		fmt.Println("Balance After Credit: ", balanceAmount)
	}
}
func debitFromBalance() {
	for i := 0; i < 5; i++ {
		balanceAmount -= 1000
		time.Sleep(time.Duration(rand.Intn(100)) *
			time.Millisecond)
		fmt.Println("Balance After Debit: ", balanceAmount)
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
Balance After Credit:  3000
Balance After Debit:  3000
Balance After Credit:  3000
Balance After Debit:  4000
Balance After Debit:  3000
Balance After Debit:  2000
Balance After Credit:  1000
Balance After Debit:  2000
Balance After Credit:  2000
Balance After Credit:  3000
*/
