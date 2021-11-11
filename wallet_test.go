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
	type args struct {
		b Bitcoin
	}
	tests := []struct {
		name    string
		w       *Wallet
		args    args
		wantErr bool
	}{
		{
			name:    "1",
			w:       &Wallet{balance: 1},
			args:    args{b: 1},
			wantErr: false,
		},
		{
			name:    "2",
			w:       &Wallet{balance: 1},
			args:    args{b: -1},
			wantErr: true,
		},
		{
			name:    "3",
			w:       &Wallet{balance: -1},
			args:    args{b: 1},
			wantErr: true,
		},
		{
			name:    "4",
			w:       &Wallet{balance: 1},
			args:    args{b: 2},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.w.Withdraw(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("Wallet.Withdraw() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWallet_Deposit(t *testing.T) {
	type args struct {
		b Bitcoin
	}
	tests := []struct {
		name    string
		w       *Wallet
		args    args
		wantErr bool
	}{
		{
			name:    "1",
			w:       &Wallet{balance: 0},
			args:    args{b: 1},
			wantErr: false,
		},
		{
			name:    "2",
			w:       &Wallet{balance: -1},
			args:    args{b: 1},
			wantErr: false,
		},
		{
			name:    "1",
			w:       &Wallet{balance: 1},
			args:    args{b: -1},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.w.Deposit(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("Wallet.Deposit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
