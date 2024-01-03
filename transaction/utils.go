package transaction

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
	"log"
	"math/big"

	"github.com/mr-tron/base58"
)

func Base58Encode(input []byte) string {
	encode := base58.Encode(input)

	return encode
}

func Base58Decode(input string) []byte {
	decode, err := base58.Decode(input)
	if err != nil {
		log.Panic(err)
	}

	return decode
}

func GenerateXY(data []byte) (big.Int, big.Int) {
	x := big.Int{}
	y := big.Int{}
	keyLen := len(data)
	x.SetBytes(data[:(keyLen / 2)])
	y.SetBytes(data[(keyLen / 2):])
	return x, y
}

func Hash160(data []byte) []byte {
	hasher := ripemd160.New()
	_, err := hasher.Write(data[:])
	if err != nil {
		log.Panic(err)
	}

	dataRipMD := hasher.Sum(nil)

	return dataRipMD
}
func PublicKeyHash(pubKey []byte) []byte {
	pubHash := sha256.Sum256(pubKey)

	return Hash160(pubHash[:])
}

func NewKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()

	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}

	pub := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)
	return *private, pub
}

func Checksum(payload []byte) []byte {
	firstHash := sha256.Sum256(payload)
	secondHash := sha256.Sum256(firstHash[:])

	return secondHash[:checksumLength]
}
