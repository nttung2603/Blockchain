package blockchain

type TxInput struct {
	ID         []byte
	Out        int
	Signature1 []byte
	PubKey1    []byte
	Signature2 []byte
	PubKey2    []byte
}

//func (in *TxInput) UsesKey(pubKeyHash []byte) bool {
//	lockingHash := wallet.PublicKeyHash(in.PubKey)
//
//	return bytes.Compare(lockingHash, pubKeyHash) == 0
//}
