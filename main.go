package main

import (
  "fmt"
  "github.com/grace0925/blockchain-research/blockchain"
)

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
