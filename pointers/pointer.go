package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin float64

type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%.2f BTC", b)
}

func (wallet *Wallet) Deposit(amount Bitcoin) {
	wallet.balance += amount
}

var ErrInsufficiantFunds = errors.New("cannot withdraw, insufficient funds")

func (wallet *Wallet) Withdraw(amount Bitcoin) error {
	if amount > wallet.balance {
		return ErrInsufficiantFunds
	}
	wallet.balance -= amount
	return nil
}

func (wallet *Wallet) Balance() Bitcoin {
	return wallet.balance
}
