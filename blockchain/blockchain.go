package blockchain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Blockchain struct {
	blocks   []*Block
	dataPath string
}

func InitBlockchain(address string) *Blockchain {
	coinbaseTx := &Transaction{Data: []byte("Coinbase Transaction")}
	genesisBlock := GenesisBlock(coinbaseTx)
	blockchain := &Blockchain{
		blocks: []*Block{genesisBlock},
	}
	blockchain.dataPath = address + ".json"
	blockToJson, _ := json.Marshal(blockchain.blocks)
	ioutil.WriteFile(blockchain.dataPath, blockToJson, os.ModePerm)
	return blockchain
}

// func (chain *Blockchain) MineBlock(block *Block) {
// 	file, _ := os.OpenFile(chain.dataPath, os.O_CREATE, os.ModePerm)
// 	defer file.Close()
// 	encoder := json.NewEncoder(file)
// 	fmt.Print(encoder)
// }

func MineBlock(address string) {
	dataPath := address + ".json"
	content, _ := ioutil.ReadFile(dataPath)
	chain := new(Blockchain)
	json.Unmarshal(content, &chain.blocks)
	fmt.Println(chain.blocks[0].Timestamp)
}
