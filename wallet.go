// Package implements a wallet, that holds bitcoin and operations: to deposit and withdraw money and get the wallet's balance
package wallet

import (
	"errors"
	"sync"
)

// Bitcoin is a type based on float64
type Bitcoin float64

// Wallet is a type which contains balance and mutex(for concurrent use)
type Wallet struct {
	balance Bitcoin
	mutex   sync.Mutex
}

// Deposit allows depositing bitcoin to the  wallet
func (w *Wallet) Deposit(b Bitcoin) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	if float64(b) < 0 {
		return errors.New("negative input")
	}

	w.balance += b

	return nil
}
