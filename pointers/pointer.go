package pointers

import "fmt"

type Wallet struct {
	balance float64
}

func (wallet *Wallet) Deposit(amount float64) {
	fmt.Printf("address of balance in Deposit is %p\n", &wallet.balance)
	wallet.balance += amount
}

func (wallet *Wallet) Balance() float64 {
	return wallet.balance
}
