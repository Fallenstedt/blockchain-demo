package main

import (
	"fmt"
	"github.com/fallenstedt/blockchain-exp/blockchain"
	"strconv"
)


func main() {
	chain := blockchain.NewBlockChain()
	chain.AddBlock("First thing added")
	chain.AddBlock("Second thing added!")
	chain.AddBlock("Third thing added")

	for _, b := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", string(b.PrevHash))
		fmt.Printf("Data In Block: %s\n", string(b.Data))
		fmt.Printf("Hash: %x\n", string(b.Hash))

		pow := blockchain.NewProofOfWork(b)

		fmt.Printf("Proof of work: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
