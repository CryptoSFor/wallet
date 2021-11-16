package wallet_test

import (
	"fmt"

	"github.com/CryptoSFor/wallet"
)

func ExampleWallet() {
	w := wallet.Wallet{}
	w.Deposit(0.75)
	w.Withdraw(0.1)
	fmt.Println(w.Balance())
	//Output: 0.65
}
