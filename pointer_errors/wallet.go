package wallet

import (
	"errors"
	"fmt"
)

// Bitcoin type
type Bitcoin int

// Wallet holds our balance
type Wallet struct {
	balance Bitcoin
}

// Stringer interface
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
    return fmt.Sprintf("%d BTC", b)
}

// Deposit a certain amount to my wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

//ErrInsufficientFunds error message
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Withdraw a certain amount from my wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {

    if amount > w.balance {
        return ErrInsufficientFunds
    }

    w.balance -= amount
    return nil
}

// Balance preview
func (w *Wallet) Balance() Bitcoin {
    return w.balance
}