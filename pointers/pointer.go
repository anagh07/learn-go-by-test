package pointers

import "fmt"

type Bitcoin float64

type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%.2f BTC", b)
}

func (wallet *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %p\n", &wallet.balance)
	wallet.balance += amount
}

func (wallet *Wallet) Balance() Bitcoin {
	return wallet.balance
}
