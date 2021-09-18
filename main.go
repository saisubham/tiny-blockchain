package main

import (
	"fmt"
)

func main() {
	bc := MakeBlockchain()
	bc.Append("one")
	bc.Append("two")

	for _, block := range bc.blocks {
		fmt.Printf("Data: \"%s\"\n", string(block.Data))
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("PrevHash: %x\n", block.PrevHash)
		isValid := MakePOW(block).Validate()
		fmt.Printf("pow: %v\n\n", isValid)
	}
}
