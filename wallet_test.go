// Package implements a wallet, that holds bitcoin and operations: to deposit and withdraw money and get the wallet's balance

package wallet

import (
	"reflect"
	"testing"
)

func TestWallet_GetBalance(t *testing.T) {
	tests := []struct {
		name string
		w    *Wallet
		want Bitcoin
	}{
		{
			name: "zero balance",
			w:    &Wallet{balance: 0},
			want: 0,
		},
		{
			name: "positive balance",
			w:    &Wallet{balance: 1},
			want: 1,
		},
		{
			name: "negative balance",
			w:    &Wallet{balance: -1},
			want: -1,
		},
		{
			name: "small float balance",
			w:    &Wallet{balance: 0.00000001},
			want: 0.00000001,
		},
		{
			name: "huge float balance",
			w:    &Wallet{balance: 10000000.00000001},
			want: 10000000.00000001,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.w.GetBalance(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Wallet.GetBalance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWallet_Withdraw(t *testing.T) {
	tests := []struct {
		name           string
		w              *Wallet
		withdrawAmount Bitcoin
		want           Bitcoin
		wantErr        error
	}{
		{
			name:           "Withdrawing positive value from positive balance",
			w:              &Wallet{balance: 0.1},
			withdrawAmount: 0.05,
			want:           0.05,
			wantErr:        nil,
		},
		{
			name:           "Withdrawing negative value",
			w:              &Wallet{balance: 0.01},
			withdrawAmount: -0.0001,
			want:           0.01,
			wantErr:        NegativeInputError,
		},
		{
			name:           "Withdrawing from negative balance",
			w:              &Wallet{balance: -0.001},
			withdrawAmount: 0.1,
			want:           -0.001,
			wantErr:        InsufficientFundsError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.w.Withdraw(tt.withdrawAmount); err != tt.wantErr || tt.w.balance != tt.want {
				t.Errorf("Wallet.Withdraw() error = %v, wantErr = %v, want = %v, got = %v", err, tt.wantErr, tt.w.balance, tt.want)
			}
		})
	}
}

func TestWallet_Deposit(t *testing.T) {
	tests := []struct {
		name          string
		w             *Wallet
		depositAmount Bitcoin
		want          Bitcoin
		wantErr       error
	}{
		{
			name:          "Deposit positive value on positive balance",
			w:             &Wallet{balance: 0.1},
			depositAmount: 0.001,
			want:          0.101,
			wantErr:       nil,
		},
		{
			name:          "Deposit positive value on negative balance",
			w:             &Wallet{balance: -0.5},
			depositAmount: 0.05,
			want:          -0.45,
			wantErr:       nil,
		},
		{
			name:          "Deposit negative value",
			w:             &Wallet{balance: 0.01},
			depositAmount: -0.0001,
			want:          0.01,
			wantErr:       NegativeInputError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.w.Deposit(tt.depositAmount); err != tt.wantErr || tt.w.balance != tt.want {
				t.Errorf("Wallet.Deposit() error = %v, wantErr = %v, want = %v, got = %v", err, tt.wantErr, tt.w.balance, tt.want)
			}
		})
	}
}
