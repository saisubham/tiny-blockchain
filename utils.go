package main

import (
	"bytes"
	"encoding/gob"
	"log"
)

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	enc := gob.NewEncoder(&result)
	if err := enc.Encode(b); err != nil {
		log.Fatal("encode error: ", err)
	}
	return result.Bytes()
}

func Deserialize(data []byte) *Block {
	var block Block
	dec := gob.NewDecoder(bytes.NewReader(data))
	if err := dec.Decode(&block); err != nil {
		log.Fatal("decode error: ", err)
	}
	return &block
}
