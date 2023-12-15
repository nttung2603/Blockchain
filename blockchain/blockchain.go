package blockchain

type Blockchain struct {
	blocks []*Block
}

func InitBlockchain() *Blockchain {
	coinbaseTx := &Transaction{Data: []byte("Coinbase Transaction")}
	genesisBlock := GenesisBlock(coinbaseTx)
	blockchain := &Blockchain{
		blocks: []*Block{genesisBlock},
	}
	return blockchain
}
