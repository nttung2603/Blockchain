package blockchain

import (
	wallet "Blockchain/wallet"
	"fmt"
	"testing"
)

func TestP2PKH(t *testing.T) {
	// Create a sample transaction input
	walletA := wallet.MakeWallet()
	walletB := wallet.MakeWallet()
	walletC := wallet.MakeWallet()

	txInput := TxInput{
		ID:         []byte("sampleTransactionID"),
		Out:        0,
		Signature1: nil, // Placeholder for the signature
		PubKey1:    walletA.PublicKey,
		Signature2: nil, // Placeholder for the signature
		PubKey2:    walletB.PublicKey,
	}

	txOutput := TxOutput{
		Value:      30,                                      // Sample value in satoshis
		PubKeyHash: wallet.PublicKeyHash(walletC.PublicKey), //Destination address
	}

	newTransaction := Transaction{
		ID:      nil, // Placeholder for the transaction ID
		Inputs:  txInput,
		Outputs: txOutput,
	}

	// Set the transaction ID by hashing its contents
	newTransaction.ID = newTransaction.Hash()

	newTransaction.Sign(walletA.PrivateKey)
	newTransaction.Sign(walletB.PrivateKey)

	fmt.Println("New Transaction:")
	fmt.Printf("ID: %x\n", newTransaction.ID)
	fmt.Println("Inputs:")
	input := newTransaction.Inputs
	fmt.Printf("- ID: %x\n", input.ID)
	fmt.Printf("  Out: %d\n", input.Out)
	fmt.Printf("  Signature 1: %x\n", input.Signature1)
	fmt.Printf("  PubKey 1: %x\n", input.PubKey1)
	fmt.Printf("  Signature 2: %x\n", input.Signature2)
	fmt.Printf("  PubKey 2: %x\n", input.PubKey2)
	//
	output := newTransaction.Outputs
	fmt.Printf("- Value: %d\n", output.Value)
	fmt.Printf("  PubKeyHash: %x\n", output.PubKeyHash)
	//
	fmt.Println("- Verifying signature: ", newTransaction.Verify())
}
