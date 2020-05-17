package wallet

import (
	"errors"
	"fmt"
)

// ErrInsufficientFunds エラーメッセージ
var ErrInsufficientFunds = "cannot withdraw, insufficient funds"

// Bitcoin ビットコイン
type Bitcoin int

// Wallet 財布
type Wallet struct {
	balance Bitcoin
}

// Stringer 文字列
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Deposit 預ける
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Withdraw 引き出す
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errors.New(ErrInsufficientFunds)
	}

	w.balance -= amount
	return nil
}

// Balance 残高照会
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
