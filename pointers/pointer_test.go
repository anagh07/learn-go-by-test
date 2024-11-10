package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10)

	got := wallet.Balance()
	want := 10.0

	fmt.Printf("address of balance in test is %p\n", &wallet.balance)

	if got != want {
		t.Errorf("got %f, want %f", got, want)
	}

	wallet.Deposit(15)
	got = wallet.Balance()
	want += 15

	if got != want {
		t.Errorf("got %f, want %f", got, want)
	}
}
