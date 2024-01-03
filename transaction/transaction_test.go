package transaction

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/btcsuite/btcd/txscript"
	"testing"
)

func TestP2PKH(t *testing.T) {
	PrivateKey1, PublicKey1 := NewKeyPair() //Send
	_, PublicKey2 := NewKeyPair()           //Receive

	txInput := TxInput{
		PreviousTx: []byte("f5d8e39a430901c91a5917b9f2dc19d6d1a0e9cea205b009ca73đ04470b9a6"),
		Index:      0,
		ScriptSig:  nil,
	}

	txOutput := TxOutput{
		Value:        5000000000,
		ScriptPubKey: nil,
	}

	scriptPubKey, _ := NewP2PKHScriptPubKey(PublicKeyHash(PublicKey2))
	txOutput.ScriptPubKey = scriptPubKey

	newTransaction := Transaction{
		ID:      nil, // Placeholder for the transaction ID
		Inputs:  txInput,
		Outputs: txOutput,
	}

	newTransaction.ID = newTransaction.Hash()
	newTransaction, _ = SignP2PKHTransaction(newTransaction, PrivateKey1, PublicKey1)

	fmt.Println("Transaction:")
	fmt.Println("ID: ", newTransaction.ID)
	fmt.Println("Inputs:")
	input := newTransaction.Inputs
	fmt.Println("- PreviousTx: ", input.PreviousTx)
	fmt.Println("  Index: ", input.Index)
	fmt.Println("  ScriptSig: ", input.ScriptSig)
	tmpScriptSig, _ := txscript.DisasmString(input.ScriptSig)
	fmt.Println("  ScriptSig String: ", tmpScriptSig)
	//
	fmt.Println("Outputs:")
	output := newTransaction.Outputs
	fmt.Println("- Value: ", output.Value)
	fmt.Println("  ScriptPubKey: ", output.ScriptPubKey)
	tmpScriptPubKey, _ := txscript.DisasmString(output.ScriptPubKey)
	fmt.Println("  ScriptPubKey String: ", tmpScriptPubKey)
}

func TestP2SH(t *testing.T) {
	PrivateKey1, PublicKey1 := NewKeyPair() //Send
	PrivateKey2, PublicKey2 := NewKeyPair() //Send

	txInput := TxInput{
		PreviousTx: []byte("f5d8e39a430901c91a5917b9f2dc19d6d1a0e9cea205b009ca73đ04470b9a6"),
		Index:      0,
		ScriptSig:  nil,
	}

	txOutput := TxOutput{
		Value:        5000000000,
		ScriptPubKey: nil,
	}
	PubKeys := [][]byte{PublicKey1, PublicKey2}
	redeemScript, _ := NewMOfNRedeemScript(2, 2, PubKeys)
	redeemScriptHash := Hash160(redeemScript)
	scriptPubKey, _ := NewP2SHScriptPubKey(Hash160(redeemScriptHash))
	txOutput.ScriptPubKey = scriptPubKey

	newTransaction := Transaction{
		ID:      nil, // Placeholder for the transaction ID
		Inputs:  txInput,
		Outputs: txOutput,
	}

	newTransaction.ID = newTransaction.Hash()
	PrivateKeys := []ecdsa.PrivateKey{PrivateKey1, PrivateKey2}
	newTransaction, _ = SignP2SHTransaction(newTransaction, PrivateKeys)

	fmt.Println("Transaction:")
	fmt.Println("ID: ", newTransaction.ID)
	fmt.Println("Inputs:")
	input := newTransaction.Inputs
	fmt.Println("- PreviousTx: ", input.PreviousTx)
	fmt.Println("  Index: ", input.Index)
	fmt.Println("  ScriptSig: ", input.ScriptSig)
	tmpScriptSig, _ := txscript.DisasmString(input.ScriptSig)
	fmt.Println("  ScriptSig String: ", tmpScriptSig)
	//
	fmt.Println("Outputs:")
	output := newTransaction.Outputs
	fmt.Println("- Value: ", output.Value)
	fmt.Println("  ScriptPubKey: ", output.ScriptPubKey)
	tmpScriptPubKey, _ := txscript.DisasmString(output.ScriptPubKey)
	fmt.Println("  ScriptPubKey String: ", tmpScriptPubKey)
}
