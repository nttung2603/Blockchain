package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"log"
	"time"
)

type Block struct {
	//Header
	Timestamp  int64  //present the time the block was generated.
	PrevHash   []byte //present the hash of the previous block
	Hash       []byte //present the hash of the block
	MerkleRoot []byte //present the merkle root tree hash of the block transactions
	Nonce      int    //present the nonce of the block
	// Body
	Transactions []*Transaction //present the list of transactions in the block
}

// SetHash calculates and sets the hash of the block.
func (b *Block) SetHash() {
	// Concatenate PrevBlockHash, Transactions, and Timestamp
	data := append(append(b.PrevHash, b.HashTransactions()...), IntToHex(b.Timestamp)...)

	// Calculate SHA-256 hash
	hash := sha256.Sum256(data)
	// pow := NewProofOfWork(block)
	// nonce, hash := pow.Run()

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

	b.MerkleRoot = hash.RootNode.Hash
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
	// block.SetHash()
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
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

func FakeTransactionData() []*Transaction {
	transactions := []*Transaction{
		{Data: []byte("Transaction 1")},
		{Data: []byte("Transaction 2")},
	}
	return transactions
}

func (b *Block) Serialize() ([]byte, error) {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)

	return res.Bytes(), err
}

func Deserialize(data []byte) (*Block, error) {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)

	return &block, err
}

func (b *Block) PrintBlock() {
	fmt.Println("Block Information:")
	fmt.Println("* Timestamp: ", time.Unix(b.Timestamp, 0).Format(time.RFC3339))
	fmt.Println("* PrevHash: ", b.PrevHash)
	fmt.Println("* Hash: ", b.Hash)
	fmt.Println("* Nonce: ", b.Nonce)
	fmt.Println("* Transactions: ")
	for _, tx := range b.Transactions {
		fmt.Println("- Transaction detail: ", string(tx.Data))
		fmt.Println("Data Bytes: ", tx.Data)
	}
}
