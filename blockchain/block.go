package blockchain

type BlockChain struct {
	// placeholder
	Blocks []*Block
}

func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)
}

type Block struct {
	Hash []byte
	Data []byte
	PrevHash []byte
	Nonce int
}

func NewBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}


func Genesis() *Block {
	return NewBlock("Genesis", []byte{})
}
