package wallet_test

import (
	"fmt"

	"github.com/CryptoSFor/wallet/wallet"
)

func ExampleWallet() {
	w := wallet.Wallet{}
	w.Deposit(0.75)
	w.Withdraw(0.1)
	fmt.Println(w.GetBalance())
	//Output: 0.65
}
