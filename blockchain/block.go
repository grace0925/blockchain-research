package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type BlockChain struct {
	Blocks []*Block
}

type Block struct {
	Hash      []byte // hash of current block
	Data      []byte // data inside block
	PrevHash  []byte // previous block hash
}

// create the hash for certain info in the block
func (b *Block) GetHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info) // create actual hash based on info
	b.Hash = hash[:]
}

// create new blocks
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash} // new block
	block.GetHash()
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