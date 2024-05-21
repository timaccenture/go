package main

import "testing"

func TestBlockChain(t *testing.T) {
	bc := NewBlockChain()

	bc.AddBlock("Send 1 BTC to Max")

	block0 := bc.blocks[0]
	expectedBlock0 := NewBlock("Genesis Block", []byte{})

	if block0.Data[0] != expectedBlock0.Data[0] {
		t.Errorf("expected the blocks to be the same, but are not")
	}

	block1 := bc.blocks[1]
	block1Expected := NewBlock("Send 1 BTC to Max", block0.Hash)

	if block1.PrevBlockHash[0] != block0.Hash[0] {
		t.Errorf("something went wrong with the calclulation of the block hash")
	}

	if block1.Data[0] != block1Expected.Data[0] {
		t.Errorf("something went wrong with the data in the block")
	}
}
