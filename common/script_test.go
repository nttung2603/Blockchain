package common

import (
	"fmt"
	"github.com/btcsuite/btcd/txscript"
	"testing"
)

func Test_NewMOfNRedeemScript(t *testing.T) {
	_, PublicKey1 := NewKeyPair()
	_, PublicKey2 := NewKeyPair()

	PubKeys := [][]byte{PublicKey1, PublicKey2}

	redeemScript, err := NewMOfNRedeemScript(2, 2, PubKeys)
	if err != nil {
		fmt.Println("Error creating redeem script:", err)
		return
	}
	redeemStr, err := txscript.DisasmString(redeemScript)
	if err != nil {
		return
	}

	fmt.Println("Redeem Script String: ", redeemStr)
	fmt.Println("Redeem Script: ", redeemScript)
}

func Test_NewP2PKHScriptPubKey(t *testing.T) {
	_, PublicKey1 := NewKeyPair()

	scriptPubKey, err := NewP2PKHScriptPubKey(PublicKeyHash(PublicKey1))
	if err != nil {
		fmt.Println("Error creating redeem script:", err)
		return
	}

	scriptPubKeyStr, err := txscript.DisasmString(scriptPubKey)
	if err != nil {
		return
	}

	fmt.Println("ScriptPubKey String: ", scriptPubKeyStr)
	fmt.Println("ScriptPubKey: ", scriptPubKey)
}

func Test_NewP2SHScriptPubKey(t *testing.T) {
	_, PublicKey1 := NewKeyPair()
	_, PublicKey2 := NewKeyPair()

	PubKeys := [][]byte{PublicKey1, PublicKey2}

	redeemScript, err := NewMOfNRedeemScript(2, 2, PubKeys)
	if err != nil {
		fmt.Println("Error creating redeem script:", err)
		return
	}

	redeemScriptHash := Hash160(redeemScript)
	scriptPubKey, err := NewP2SHScriptPubKey(Hash160(redeemScriptHash))
	if err != nil {
		fmt.Println("Error creating redeem script:", err)
		return
	}

	scriptPubKeyStr, err := txscript.DisasmString(scriptPubKey)
	if err != nil {
		return
	}

	fmt.Println("ScriptPubKey String: ", scriptPubKeyStr)
	fmt.Println("ScriptPubKey: ", scriptPubKey)
}

func Test_NewP2PKHAddress(t *testing.T) {
	_, PublicKey1 := NewKeyPair()
}
