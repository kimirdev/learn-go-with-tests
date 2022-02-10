package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(sum Bitcoin) {
	w.balance += sum
}

func (w Wallet) Balance() Bitcoin {
	return w.balance
}

var ErrInsufficientFunds = errors.New("Not enough BTC on balance")

func (w *Wallet) Withdraw(sum Bitcoin) error {
	if w.balance < sum {
		return ErrInsufficientFunds
	}

	w.balance -= sum

	return nil
}
