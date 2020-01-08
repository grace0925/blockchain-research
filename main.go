package main

import (
  "fmt"
  "github.com/grace0925/blockchain-research/blockchain"
  "strconv"
)

func main() {
  bchain := blockchain.InitBlockChain()
  bchain.AddBlock("First block")
  bchain.AddBlock("Second block")
  bchain.AddBlock("Third block")

  for _, block := range bchain.Blocks {
    fmt.Printf("Previous Hash: %x\n", block.PrevHash)
    fmt.Printf("Data: %s\n", block.Data)
    fmt.Printf("Hash: %x\n", block.Hash)

    pow := blockchain.NewProof(block)
    fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
    fmt.Println()
  }
}
