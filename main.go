package main

import (
  "bytes"
  "crypto/sha256"
  "fmt"
)

type BlockChain struct {
  blocks []*Block
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
  prevBlock := chain.blocks[len(chain.blocks)-1] // get last block in the blockchain
  new := CreateBlock(data, prevBlock.Hash) // create new block
  chain.blocks = append(chain.blocks, new) // append new block to blockchain
}

// create the first block in the chain
func Genesis() *Block {
  return CreateBlock("Genesis", []byte{})
}

// create new blockchain
func InitBlockChain() *BlockChain {
  return &BlockChain{[]*Block{Genesis()}}
}

func main() {
  bchain := InitBlockChain()
  bchain.AddBlock("First block")
  bchain.AddBlock("Second block")
  bchain.AddBlock("Third block")

  for _, block := range bchain.blocks {
    fmt.Printf("Data: %s\n", block.Data)
    fmt.Printf("Hash: %x\n", block.Hash)
  }
}
