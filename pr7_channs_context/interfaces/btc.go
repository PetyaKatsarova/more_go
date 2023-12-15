package main

import (
	"errors"
)

type BitcoinAccount struct {
	balance int
	fee int
}

// constructor for the struct:
func NewBitCoinAccount() *BitcoinAccount {
	return &BitcoinAccount{
		balance: 0,
		fee: 300, // 300 is 3 dollars
	}
}

func (b *BitcoinAccount) getBalance() int {
	return b.balance
}

func (b *BitcoinAccount) Deposit(amount int) {
	b.balance += amount
}

func (b *BitcoinAccount) Withdraw(amount int) error {
	newBalance := b.balance - amount - b.fee
	if newBalance < 0 {
		return errors.New("insufficient funds")
	}
	b.balance = newBalance
	return nil
}