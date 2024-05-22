package main

import (
	"time"
)

// Block is defined by four essential parts: the timestamp of its creation,
// the data it should contain, the hash of the previous block (they should be linked together hence Blockchain ;))
// and finally the hash of the block
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

// Function to create a new block
// requires data and the hash of the prev block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash
	block.Nonce = nonce

	return block
}
