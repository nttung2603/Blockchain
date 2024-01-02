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

	txInput := TxInput{
		ID:        []byte("sampleTransactionID"),
		Out:       0,
		Signature: nil, // Placeholder for the signature
		PubKey:    walletA.PublicKey,
	}

	txOutput := TxOutput{
		Value:      30,                                      // Sample value in satoshis
		PubKeyHash: wallet.PublicKeyHash(walletB.PublicKey), //Destination address
	}

	newTransaction := Transaction{
		ID:      nil, // Placeholder for the transaction ID
		Inputs:  txInput,
		Outputs: txOutput,
	}

	// Set the transaction ID by hashing its contents
	newTransaction.ID = newTransaction.Hash()

	newTransaction.Sign(walletA.PrivateKey)

	fmt.Println("New Transaction:")
	fmt.Printf("ID: %x\n", newTransaction.ID)
	fmt.Println("Inputs:")
	input := newTransaction.Inputs
	fmt.Printf("- ID: %x\n", input.ID)
	fmt.Printf("  Out: %d\n", input.Out)
	fmt.Printf("  Signature: %x\n", input.Signature)
	fmt.Printf("  PubKey: %x\n", input.PubKey)
	//
	output := newTransaction.Outputs
	fmt.Printf("- Value: %d\n", output.Value)
	fmt.Printf("  PubKeyHash: %x\n", output.PubKeyHash)
	//
	fmt.Println("- Verifying signature: ", newTransaction.Verify(wallet.PublicKeyHash(walletA.PublicKey)))
}
