package wallet

//
//import (
//	"Blockchain/common"
//	"bytes"
//	"crypto/ecdsa"
//	"crypto/sha256"
//)
//
//type Wallet struct {
//	PrivateKey ecdsa.PrivateKey
//	PublicKey  []byte
//}
//
//func (w Wallet) Address() []byte {
//	pubHash := PublicKeyHash(w.PublicKey)
//
//	versionedHash := append([]byte{version}, pubHash...)
//	checksum := Checksum(versionedHash)
//
//	fullHash := append(versionedHash, checksum...)
//	address := common.Base58Encode(fullHash)
//
//	return address
//}
//
//func MakeWallet() *Wallet {
//	private, public := NewKeyPair()
//	wallet := Wallet{private, public}
//
//	return &wallet
//}
//
//func Checksum(payload []byte) []byte {
//	firstHash := sha256.Sum256(payload)
//	secondHash := sha256.Sum256(firstHash[:])
//
//	return secondHash[:checksumLength]
//}
//
//func ValidateAddress(address string) bool {
//	pubKeyHash := common.Base58Decode([]byte(address))
//	actualChecksum := pubKeyHash[len(pubKeyHash)-checksumLength:]
//	version := pubKeyHash[0]
//	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-checksumLength]
//	targetChecksum := Checksum(append([]byte{version}, pubKeyHash...))
//
//	return bytes.Compare(actualChecksum, targetChecksum) == 0
//}
