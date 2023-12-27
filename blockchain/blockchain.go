package blockchain

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Blockchain struct {
	blocks []*Block
	// dataPath string
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
	dataPath := address + ".json"
	blockToJson, _ := json.Marshal(blockchain.blocks)
	ioutil.WriteFile(dataPath, blockToJson, os.ModePerm)
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

func (chain *Blockchain) PrintBlock(index int) {
	chain.blocks[index].PrintBlock()
}

func (chain *Blockchain) PrintChain() {
	fmt.Println("Blockchain Information")
	fmt.Println("* Num blocks: ", len(chain.blocks))
	fmt.Println("* Last block")
	chain.blocks[len(chain.blocks)-1].PrintBlock()
}

func SerializeBase(address string) []byte {
	dataPath := address + ".json"
	content, _ := ioutil.ReadFile(dataPath)
	// var by []byte
	// err := json.Unmarshal(content, &by)
	return content
}

func DeserializeBase(data []byte, address string) {
	ioutil.WriteFile(address+".json", data, os.ModePerm)
}

func (chain *Blockchain) SerializeChain() ([]byte, error) {
	chain.PrintChain()
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)
	err := encoder.Encode(chain)
	return res.Bytes(), err
}

func DeserializeChain(data []byte) (*Blockchain, error) {
	var chain Blockchain

	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&chain)
	return &chain, err
}
