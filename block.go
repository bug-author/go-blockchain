package main

import (
	"crypto/sha256"
	"errors"
	"time"
)

type Block struct {
	PrevHash  []byte
	Data      string
	Timestamp int64
	Hash      []byte
}

type BlockChain struct {
	Blocks []*Block
}

func (block *Block) calculateHash() [32]byte {
	dataToHash := append(block.PrevHash, byte(block.Timestamp))
	hash := sha256.Sum256(dataToHash)
	return hash
}

func NewBlock(prevHash []byte, data string) (*Block, error) {
	if data == "" {
		return &Block{}, errors.New("invalid data")
	}

	block := &Block{
		PrevHash:  prevHash,
		Data:      data,
		Timestamp: time.Now().Unix(),
	}

	hash := block.calculateHash()
	block.Hash = hash[:]

	return block, nil
}

func (blockChain *BlockChain) AddBlock(data string) {
	prevBlock := blockChain.Blocks[len(blockChain.Blocks)-1]
	newBlock, _ := NewBlock(prevBlock.PrevHash, data)

	blockChain.Blocks = append(blockChain.Blocks, newBlock)
}
