package main

import (
	"log"

	"github.com/boltdb/bolt"
)

const (
	dbFile       = "bc.db"
	blocksBucket = "blocks"
)

type Blockchain struct {
	tip []byte
	db  *bolt.DB
}

func (bc *Blockchain) Append(data string) {
	var prevHash []byte

	err := bc.db.View(func(t *bolt.Tx) error {
		b := t.Bucket([]byte(blocksBucket))
		prevHash = b.Get([]byte("l"))
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	newBlock := MakeBlock(data, prevHash)
	err = bc.db.Update(func(t *bolt.Tx) error {
		b := t.Bucket([]byte(blocksBucket))

		err = b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}

		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}

		bc.tip = newBlock.Hash
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

func MakeBlockchain() *Blockchain {
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	var tip []byte

	db.Update(func(t *bolt.Tx) error {
		b := t.Bucket([]byte(blocksBucket))

		if b == nil {
			gen := MakeGenBlock()
			b, err = t.CreateBucket([]byte(blocksBucket))

			err = b.Put(gen.Hash, gen.Serialize())
			if err != nil {
				log.Panic(err)
			}

			err = b.Put([]byte("l"), gen.Hash)
			if err != nil {
				log.Panic(err)
			}

			tip = gen.Hash
		} else {
			tip = b.Get([]byte("l"))
		}
		return nil
	})
	return &Blockchain{tip, db}
}

type BlockchainIterator struct {
	curHash []byte
	db      *bolt.DB
}

func (bc *Blockchain) Iterator() *BlockchainIterator {
	return &BlockchainIterator{bc.tip, bc.db}
}

func (it *BlockchainIterator) Next() *Block {
	var block *Block

	err := it.db.View(func(t *bolt.Tx) error {
		b := t.Bucket([]byte(blocksBucket))
		block = Deserialize(b.Get(it.curHash))
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	it.curHash = block.PrevHash
	return block
}
