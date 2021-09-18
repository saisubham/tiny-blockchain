package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	"strconv"
)

const targetBits = 16

type POW struct {
	block  *Block
	target *big.Int
}

func MakePOW(b *Block) *POW {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	return &POW{block: b, target: target}
}

func (pow *POW) MakeData(nonce int) []byte {
	return bytes.Join([][]byte{
		pow.block.PrevHash,
		pow.block.Data,
		[]byte(strconv.FormatInt(pow.block.Timestamp, 16)),
		[]byte(strconv.FormatInt(int64(targetBits), 16)),
		[]byte(strconv.FormatInt(int64(nonce), 16)),
	}, []byte{})
}

func (pow *POW) Run() (int, []byte) {
	nonce := 0
	var hashInt big.Int
	var hash [32]byte

	fmt.Printf("Mining block with data \"%s\" ...\n", string(pow.block.Data))
	for nonce < math.MaxInt64 {
		data := pow.MakeData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)

		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Printf("\n\n")

	return nonce, hash[:]
}

func (pow *POW) Validate() bool {
	var hashInt big.Int

	data := pow.MakeData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	return hashInt.Cmp(pow.target) == -1
}
