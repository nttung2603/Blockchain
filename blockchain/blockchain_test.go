package blockchain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlockchain_AddBlock(t *testing.T) {
	// Create a new blockchain with the genesis block
	blockchain := InitBlockchain("test")

	transactions := []*Transaction{
		{Data: []byte("Transaction 1")},
		{Data: []byte("Transaction 2")},
	}
	blockchain.AddBlock(transactions)
	numBlocks := len(blockchain.blocks)
	assert.Equal(t, 2, numBlocks, "Blockchain have 2 blocks")
}

func TestBlockchain_Blockchain(t *testing.T) {
	// Create a new blockchain with the genesis block
	blockchain := InitBlockchain("test")
	// Create Block 1
	transactions1 := []*Transaction{
		{Data: []byte("Transaction 3")},
		{Data: []byte("Transaction 4")},
	}
	blockchain.AddBlock(transactions1)

	// Create Block 2
	transactions2 := []*Transaction{
		{Data: []byte("Transaction 5")},
		{Data: []byte("Transaction 6")},
	}
	blockchain.AddBlock(transactions2)

	// Check Hash
	hash1 := blockchain.blocks[0].Hash
	prevHash2 := blockchain.blocks[1].PrevHash

	assert.Equal(t, hash1, prevHash2, "Hash of the first block is equal to the hash of the second block's previous block")

	hash2 := blockchain.blocks[1].Hash
	prevHash3 := blockchain.blocks[2].PrevHash
	assert.Equal(t, hash2, prevHash3, "Hash of the second block is equal to the hash of the third block's previous block")
}
