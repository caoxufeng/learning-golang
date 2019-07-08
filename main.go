package main

import "fmt"

func main() {
	bc := NewBlockchain()
	bc.AddBlock("Send 1 BTC to Jeffrey.")
	bc.AddBlock("Send more BTC to Jeffrey")

	for _, block := range bc.blcoks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockChain)
		fmt.Printf("Data %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
