package main

import "time"

type Block struct {
	Timestamp int64
	Data      []byte
	Hash      []byte
	PrevHash  []byte
	Nonce     int
}

func MakeBlock(data string, prevHash []byte) *Block {
	b := &Block{
		Timestamp: time.Now().Unix(),
		Data:      []byte(data),
		PrevHash:  prevHash,
		Hash:      []byte{},
		Nonce:     0,
	}
	pow := MakePOW(b)
	nonce, hash := pow.Run()
	b.Hash = hash
	b.Nonce = nonce
	return b
}

func MakeGenBlock() *Block {
	return MakeBlock("", []byte{})
}
