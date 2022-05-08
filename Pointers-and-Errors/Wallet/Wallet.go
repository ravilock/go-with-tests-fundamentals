package Wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}

func (bitcoin Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", bitcoin)
}

type Wallet struct {
	balance Bitcoin
}

func (wallet *Wallet) Deposit(amount Bitcoin) {
	wallet.balance += amount
}

var ErrorInsufficientFunds = errors.New("Cannot withdraw: Insufficient funds")

func (wallet *Wallet) Withdraw(amount Bitcoin) error {
	if amount > wallet.balance {
		return ErrorInsufficientFunds
	}
	wallet.balance -= amount
	return nil
}

func (wallet *Wallet) Balance() Bitcoin {
	return wallet.balance
}
