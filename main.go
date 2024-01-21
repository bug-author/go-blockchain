package main

import (
	"fmt"
)

func main() {
	blockchain := BlockChain{}

	block, _ := Genesis()

	blockchain.Blocks = append(blockchain.Blocks, block)

	blockchain.AddBlock("Second Block!")
	blockchain.AddBlock("Third Block!")

	for _, block := range blockchain.Blocks {
		fmt.Printf("PrevHash: %s\nTimestamp: %d\nData: %s\nHash: %s\n\n",
			block.PrevHash, block.Timestamp, block.Data, block.Hash)

		fmt.Println("-------------------------------------------")
	}
}
