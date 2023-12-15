package main

import "fmt"

type IBankAccount interface {
	getBalance() int // 100 = 1 dollar
	Deposit(amount int)
	Withdraw(amount int) error
}

func main() {

	myAccounts := []IBankAccount {
		NewWellsFargo(),
		NewBitCoinAccount(),
	}

	for _, val := range myAccounts {
		val.Deposit(500)
		if err := val.Withdraw(70); err != nil {
			fmt.Printf("ERR: %d\n", err)
		}
		balance := val.getBalance()
		fmt.Printf("balance = %d\n", balance)
	}


	// wf := NewWellsFargo()
	// wf.Deposit(200)
	// wf.Deposit(100)
	// wf.Deposit(370)
	// if err := wf.Withdraw(117); err != nil {
	// 	panic(err)
	// }
	// currentBalance := wf.getBalance()
	// fmt.Printf("WF balance: %d\n", currentBalance)
}