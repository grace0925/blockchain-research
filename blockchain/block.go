package blockchain

import (
)

type BlockChain struct {
	Blocks []*Block
}

type Block struct {
	Hash      []byte // hash of current block
	Data      []byte // data inside block
	PrevHash  []byte // previous block hash
	Nonce	  int    // counter
}

// create new blocks
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0} // new block

	// run proof of work on every block created
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// add blocks to blockchain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1] // get last block in the blockchain
	new := CreateBlock(data, prevBlock.Hash) // create new block
	chain.Blocks = append(chain.Blocks, new) // append new block to blockchain
}

// create the first block in the chain
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// create new blockchain
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}