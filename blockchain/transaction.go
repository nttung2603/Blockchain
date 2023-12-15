package blockchain

import "crypto/sha256"

type Transaction struct {
	Data []byte
}

func (tx *Transaction) Hash() []byte {
	// Calculate SHA-256 hash of the transaction data
	hash := sha256.Sum256(tx.Data)
	return hash[:]
}
