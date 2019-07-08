package main

type Blockchain struct {
	blcoks []*Block
}

func (bc *Blockchain) AddBlock(data string) {
	preBlock := bc.blcoks[len(bc.blcoks)-1]
	newBlock := NewBlock(data, preBlock.Hash)
	bc.blcoks = append(bc.blcoks, newBlock)
}

//NewGenesisBlock create first block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// NewBlockchain create first blockchain
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
