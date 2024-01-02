package wallet

import (
	"fmt"
	"testing"
)

func TestP2PKH(t *testing.T) {
	wallet := MakeWallet()
	fmt.Println(wallet.Address())
	fmt.Println(wallet.PrivateKey)
	fmt.Println(wallet.PublicKey)
}
