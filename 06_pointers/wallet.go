package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (w *Wallet) Deposit(amount Bitcoin) {
	//fmt.Println("address of balance in Deposit is", &w.balance)
	w.balance += amount
}

var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return InsufficientFundsError
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	//fmt.Println("address of balance in Balance is", &w.balance)
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
