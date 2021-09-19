package main

import (
	"flag"
	"fmt"
	"os"
)

type CLI struct {
	bc *Blockchain
}

func (cli *CLI) Usage() {
	fmt.Printf("Usage: ./bc-go [option]\n")
	fmt.Printf("Options:\n")
	fmt.Printf("  a -data <string>\tmine and add a block with specified data in blockchain\n")
	fmt.Printf("  h\t\t\tshow this help\n")
	fmt.Printf("  p\t\t\tdisplay blockchain\n")
	fmt.Printf("\n")
}

func (cli *CLI) addBlock(data string) {
	cli.bc.Append(data)
}

func (cli *CLI) printBlockchain() {
	it := cli.bc.Iterator()

	for {
		block := it.Next()

		fmt.Printf("data: %s\n", block.Data)
		fmt.Printf("hash: %x\n", block.Hash)
		fmt.Printf("prev_hash: %x\n", block.PrevHash)
		pow := MakePOW(block)
		fmt.Printf("proof of work: %v\n", pow.Validate())
		fmt.Printf("\n")

		if len(block.PrevHash) == 0 {
			break
		}
	}
}

func (cli *CLI) Run() {
	addCmd := flag.NewFlagSet("a", flag.ExitOnError)
	printCmd := flag.NewFlagSet("p", flag.ExitOnError)
	addBlockData := addCmd.String("data", "", "block data")

	switch os.Args[1] {
	case "a":
		addCmd.Parse(os.Args[2:])
	case "p":
		printCmd.Parse(os.Args[2:])
	case "h":
		fallthrough
	default:
		cli.Usage()
		os.Exit(1)
	}

	if addCmd.Parsed() {
		if *addBlockData == "" {
			addCmd.Usage()
			os.Exit(1)
		}
		cli.addBlock(*addBlockData)
	}

	if printCmd.Parsed() {
		cli.printBlockchain()
	}
}
