package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
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
}

// Method to calculate the hash of the block by simply concating the prev block's hash, timestamp and the data
// and calculating a SHA256 hash on it
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

// Function to create a new block
// requires data and the hash of the prev block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}
