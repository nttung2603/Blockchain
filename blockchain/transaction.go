package blockchain

import (
	"Blockchain/wallet"
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"math/big"
)

type Transaction struct {
	ID      []byte
	Inputs  TxInput
	Outputs TxOutput
}

func (tx *Transaction) Hash() []byte {
	var hash [32]byte

	txCopy := *tx
	txCopy.ID = []byte{}

	hash = sha256.Sum256(txCopy.SerializeTransaction())

	return hash[:]
}

func (tx Transaction) SerializeTransaction() []byte {
	var encoded bytes.Buffer

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}

	return encoded.Bytes()
}

func DeserializeTransaction(data []byte) Transaction {
	var transaction Transaction

	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&transaction)
	if err != nil {
		log.Panic(err)
	}

	return transaction
}

//func (tx *Transaction) IsCoinbase() bool {
//	return len(tx.Inputs) == 1 && len(tx.Inputs[0].ID) == 0 && tx.Inputs[0].Out == -1
//}

func (tx *Transaction) Sign(privKey ecdsa.PrivateKey) {
	txCopy := tx.TrimmedCopy()
	dataToSign := fmt.Sprintf("%x\n", txCopy)
	r, s, err := ecdsa.Sign(rand.Reader, &privKey, []byte(dataToSign))
	if err != nil {
		log.Panic(err)
	}
	signature := append(r.Bytes(), s.Bytes()...)

	tx.Inputs.Signature = signature
}

func (tx *Transaction) Verify(PubKeyHash []byte) bool {
	if !bytes.Equal(wallet.PublicKeyHash(tx.Inputs.PubKey), PubKeyHash) {
		return false
	}

	txCopy := tx.TrimmedCopy()
	curve := elliptic.P256()

	txCopy.Inputs.Signature = nil

	r1, s1 := GenerateXY(tx.Inputs.Signature)
	x1, y1 := GenerateXY(tx.Inputs.PubKey)

	dataToVerify := fmt.Sprintf("%x\n", txCopy)

	rawPubKey1 := ecdsa.PublicKey{Curve: curve, X: &x1, Y: &y1}
	if ecdsa.Verify(&rawPubKey1, []byte(dataToVerify), &r1, &s1) == false {
		return false
	}

	return true
}

func GenerateXY(data []byte) (big.Int, big.Int) {
	x := big.Int{}
	y := big.Int{}
	keyLen := len(data)
	x.SetBytes(data[:(keyLen / 2)])
	y.SetBytes(data[(keyLen / 2):])
	return x, y
}

func (tx *Transaction) TrimmedCopy() Transaction {

	txCopy := Transaction{tx.ID, tx.Inputs, tx.Outputs}

	return txCopy
}
