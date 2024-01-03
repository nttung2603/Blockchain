package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
	"testing"
)

func Test(t *testing.T) {
	// Generate two random private keys
	privateKey1, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		fmt.Println("Error generating private key 1:", err)
		return
	}

	privateKey2, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		fmt.Println("Error generating private key 2:", err)
		return
	}

	// Derive the public keys
	// Derive the public keys
	publicKey1 := elliptic.MarshalCompressed(privateKey1.PublicKey.Curve, privateKey1.PublicKey.X, privateKey1.PublicKey.Y)
	publicKey2 := elliptic.MarshalCompressed(privateKey2.PublicKey.Curve, privateKey2.PublicKey.X, privateKey2.PublicKey.Y)

	// Create instances of btcutil.AddressPubKey
	//addrPubKey1, _ := btcutil.NewAddressPubKey(publicKey1, &chaincfg.MainNetParams)
	//addrPubKey2, _ := btcutil.NewAddressPubKey(publicKey2, &chaincfg.MainNetParams)

	// Create a 2-of-2 multisig redeem script
	//redeemScript, err := txscript.MultiSigScript([]*btcutil.AddressPubKey{addrPubKey1, addrPubKey2}, 2)
	//if err != nil {
	//	fmt.Println("Error creating redeem script:", err)
	//	return
	//}

	// Create a P2SH address from the redeem script
	//scriptHash := btcutil.Hash160(redeemScript)
	//address, err := btcutil.NewAddressScriptHashFromHash(scriptHash, &chaincfg.MainNetParams)
	//if err != nil {
	//	fmt.Println("Error creating P2SH address:", err)
	//	return
	//}

	builder := txscript.NewScriptBuilder()
	builder.AddOp(txscript.OP_2)
	// add the 3 public key
	builder.AddData(publicKey1).AddData(publicKey2)
	// add the total number of public keys in the multi-sig screipt
	builder.AddOp(txscript.OP_2)
	// add the check-multi-sig op-code
	builder.AddOp(txscript.OP_CHECKMULTISIG)

	redeemScript, err := builder.Script()
	if err != nil {
		return
	}

	// Create a P2SH address from the redeem script
	scriptHash := btcutil.Hash160(redeemScript)
	address, err := btcutil.NewAddressScriptHash(scriptHash, &chaincfg.MainNetParams)
	if err != nil {
		fmt.Println("Error creating P2SH address:", err)
		return
	}

	// dis asemble the script program, so can see its structure
	redeemStr, err := txscript.DisasmString(redeemScript)
	if err != nil {
		return
	}
	fmt.Println(redeemStr)
	fmt.Printf("Multisig Address: %s\n", address.EncodeAddress())
	txscript.SignatureScript(redeemScript, 0, redeemScript, txscript.SigHashAll, privateKey1, nil)
	engine, err := txscript.NewEngine(redeemScript, redeemScript, txscript.StandardVerifyFlags, nil, nil, nil)
	if err != nil {
		return
	}
	engine.Execute(nil, nil)
}
