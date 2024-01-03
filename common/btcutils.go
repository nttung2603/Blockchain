package common

import (
	"Blockchain/transaction"
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
)

func NewMOfNRedeemScript(m int, n int, publicKeys [][]byte) ([]byte, error) {
	//Get OP Code for m and n.
	//81 is OP_1, 82 is OP_2 etc.
	//80 is not a valid OP_Code, so we floor at 81
	mOPCode := OP_1 + (m - 1)
	nOPCode := OP_1 + (n - 1)
	//Multisig redeemScript format:
	//<OP_m> <A pubkey> <B pubkey> <C pubkey>... <OP_n> OP_CHECKMULTISIG
	var redeemScript bytes.Buffer
	redeemScript.WriteByte(byte(mOPCode)) //m
	for _, publicKey := range publicKeys {
		redeemScript.WriteByte(byte(len(publicKey))) //PUSH
		redeemScript.Write(publicKey)                //<pubkey>
	}
	redeemScript.WriteByte(byte(nOPCode)) //n
	redeemScript.WriteByte(byte(OP_CHECKMULTISIG))
	return redeemScript.Bytes(), nil
}

// NewP2SHScriptPubKey creates a scriptPubKey for a P2SH transaction given the redeemScript hash
func NewP2SHScriptPubKey(redeemScriptHash []byte) ([]byte, error) {
	if redeemScriptHash == nil {
		return nil, errors.New("redeemScriptHash can't be empty.")
	}
	//P2SH scriptSig format:
	//<OP_HASH160> <Hash160(redeemScript)> <OP_EQUAL>
	var scriptPubKey bytes.Buffer
	scriptPubKey.WriteByte(byte(OP_HASH160))
	scriptPubKey.WriteByte(byte(len(redeemScriptHash))) //PUSH
	scriptPubKey.Write(redeemScriptHash)
	scriptPubKey.WriteByte(byte(OP_EQUAL))
	return scriptPubKey.Bytes(), nil
}

// NewP2PKHScriptPubKey creates a scriptPubKey for a P2PKH transaction given the destination public key hash
func NewP2PKHScriptPubKey(publicKeyHash []byte) ([]byte, error) {
	if publicKeyHash == nil {
		return nil, errors.New("publicKeyHash can't be empty.")
	}
	//P2PKH scriptSig format:
	//<OP_DUP> <OP_HASH160> <pubKeyHash> <OP_EQUALVERIFY> <OP_CHECKSIG>
	var scriptPubKey bytes.Buffer
	scriptPubKey.WriteByte(byte(OP_DUP))
	scriptPubKey.WriteByte(byte(OP_HASH160))
	scriptPubKey.WriteByte(byte(len(publicKeyHash))) //PUSH
	scriptPubKey.Write(publicKeyHash)
	scriptPubKey.WriteByte(byte(OP_EQUALVERIFY))
	scriptPubKey.WriteByte(byte(OP_CHECKSIG))
	return scriptPubKey.Bytes(), nil
}

func NewAddressP2PKH(publicKey []byte) ([]byte, error) {
	pubHash := PublicKeyHash(publicKey)

	versionedHash := append([]byte{version}, pubHash...)
	checksum := Checksum(versionedHash)

	fullHash := append(versionedHash, checksum...)
	address := Base58Encode(fullHash)

	return address, nil
}

func NewAddressP2SH(redeemScript []byte) ([]byte, string) {
	redeemScriptHash := Hash160(redeemScript)

	versionedHash := append([]byte{version}, redeemScriptHash...)
	checksum := Checksum(versionedHash)

	fullHash := append(versionedHash, checksum...)
	address := Base58Encode(fullHash)

	//Get redeemScript in Hex
	redeemScriptHex := hex.EncodeToString(redeemScript)

	return address, redeemScriptHex
}

func NewSignature(tx transaction.Transaction, privKey ecdsa.PrivateKey) ([]byte, error) {
	txCopy := tx.TrimmedCopy()
	dataToSign := fmt.Sprintf("%x\n", txCopy)
	r, s, err := ecdsa.Sign(rand.Reader, &privKey, []byte(dataToSign))
	if err != nil {
		log.Panic(err)
	}
	signature := append(r.Bytes(), s.Bytes()...)

	return signature, nil
}

func signP2PKHTransaction(tx transaction.Transaction, privKey ecdsa.PrivateKey) {

}

func signP2SHTransaction(tx transaction.Transaction, privKey ecdsa.PrivateKey) {

}
