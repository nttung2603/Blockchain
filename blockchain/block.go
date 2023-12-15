package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
	"time"
)

type Block struct {
	//Header
	Timestamp  int64       //present the time the block was generated.
	PrevHash   []byte      //present the hash of the previous block
	Hash       []byte      //present the hash of the block
	MerkleTree *MerkleTree //present the merkle tree of the block
	// Body
	Nonce        int            //present the nonce of the block
	Transactions []*Transaction //present the list of transactions in the block
}

// SetHash calculates and sets the hash of the block.
func (b *Block) SetHash() {
	// Concatenate PrevBlockHash, Transactions, and Timestamp
	data := append(append(b.PrevHash, b.HashTransactions()...), IntToHex(b.Timestamp)...)

	// Calculate SHA-256 hash
	hash := sha256.Sum256(data)

	// Set the hash of the block
	b.Hash = hash[:]
}

func (b *Block) HashTransactions() []byte {
	var transactionsData [][]byte

	// Concatenate hashes of each transaction
	for _, tx := range b.Transactions {
		//transactionsData = append(transactionsData, tx.Hash())
		transactionsData = append(transactionsData, tx.Data)
	}

	// Calculate SHA-256 hash of the concatenated transaction data
	//hash := sha256.Sum256(transactionsData)
	hash := NewMerkleTree(transactionsData)

	return hash.RootNode.Hash
	//return hash[:]
}

func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)

	}
	return buff.Bytes()
}

func CreateBlock(txs []*Transaction, prevHash []byte) *Block {
	block := &Block{Timestamp: time.Now().Unix(), Transactions: txs, PrevHash: prevHash}
	block.SetHash()
	return block
}

// AddBlock adds a new block to the blockchain.
func (bc *Blockchain) AddBlock(transactions []*Transaction) {
	// Get the previous block
	prevBlock := bc.blocks[len(bc.blocks)-1]

	// Create a new block
	newBlock := CreateBlock(transactions, prevBlock.Hash)

	// Append the new block to the blockchain
	bc.blocks = append(bc.blocks, newBlock)
}

func GenesisBlock(coinbase *Transaction) *Block {
	return CreateBlock([]*Transaction{coinbase}, []byte{})
}
