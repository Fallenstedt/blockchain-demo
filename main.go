package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)
type BlockChain struct {
	// placeholder
	blocks []*Block
}

func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, newBlock)
}

type Block struct {
	Hash []byte
	Data []byte
	PrevHash []byte
}

func (b *Block) DeriveHash() {
	// placeholder
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func NewBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}


func Genesis() *Block {
	return NewBlock("Genesis", []byte{})
}

func main() {
	chain := NewBlockChain()
	chain.AddBlock("First thing added")
	chain.AddBlock("Second thing added!")
	chain.AddBlock("Third thing added")
	for _, b := range chain.blocks {
		fmt.Printf("%x", string(b.Hash))
		fmt.Println()
	}
}
