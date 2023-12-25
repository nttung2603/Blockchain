package blockchain

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Blockchain struct {
	blocks   []*Block
	dataPath string
}

func (chain *Blockchain) AppendBlock(b *Block) {
	chain.blocks = append(chain.blocks, b)
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

func (chain *Blockchain) GetPrevHash() []byte {
	return chain.blocks[len(chain.blocks)-1].Hash
}

func GetChain(address string) *Blockchain {
	dataPath := address + ".json"
	content, _ := ioutil.ReadFile(dataPath)
	chain := new(Blockchain)
	json.Unmarshal(content, &chain.blocks)
	return chain
	// chain.AddBlock(transData)
	// blockToJson, _ := json.Marshal(chain.blocks)
	// ioutil.WriteFile(dataPath, blockToJson, os.ModePerm)
}

func SetChain(chain *Blockchain, address string) {
	dataPath := address + ".json"
	blockToJson, _ := json.Marshal(chain.blocks)
	ioutil.WriteFile(dataPath, blockToJson, os.ModePerm)
}

func MineBlock(address string, transData []*Transaction) {
	dataPath := address + ".json"
	content, _ := ioutil.ReadFile(dataPath)
	chain := new(Blockchain)
	json.Unmarshal(content, &chain.blocks)
	chain.AddBlock(transData)
	blockToJson, _ := json.Marshal(chain.blocks)
	ioutil.WriteFile(dataPath, blockToJson, os.ModePerm)
}
