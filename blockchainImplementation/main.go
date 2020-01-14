package main

import (
  "flag"
  "fmt"
  "github.com/grace0925/blockchain-research/blockchain"
  "os"
  "runtime"
  "strconv"
)

type CommandLine struct {
  blockchain *blockchain.BlockChain
}

func (cli *CommandLine) printUsage() {
  fmt.Println("Usage: add-block BLOCK DATA; print")
}

func (cli *CommandLine) validateArgs() {
  if len(os.Args) < 2 {
    cli.printUsage()
    runtime.Goexit()
  }
}

func (cli *CommandLine) addBlock(data string) {
  cli.blockchain.AddBlock(data)
  fmt.Println("Added block!")
}

func (cli *CommandLine) printBlock() {
  iterator := cli.blockchain.Iterator()
  for {
    block := iterator.Next()
    fmt.Printf("Previous Hash: %x\n", block.PrevHash)
    fmt.Printf("Data: %s\n", block.Data)
    fmt.Printf("Hash: %x\n", block.Hash)

    pow := blockchain.NewProof(block)
    fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
    fmt.Println()

    // if previous hash is 0, the reach the end of the block chain
    if len(block.PrevHash) == 0 {
      break
    }
  }
}

func (cli *CommandLine) run() {
  cli.validateArgs()

  addCmd := flag.NewFlagSet("add", flag.ExitOnError)
  printCmd := flag.NewFlagSet("print", flag.ExitOnError)
  // if user types in block after typing add, then they can type in the data
  addBlockData := addCmd.String("block", "", "Block data")

  switch os.Args[1] {
  case "add":
    // parse all argument after first argument
    err := addCmd.Parse(os.Args[2:])
    blockchain.Handle(err)

  case "print":
    err := printCmd.Parse(os.Args[2:])
    blockchain.Handle(err)

  default:
    cli.printUsage()
    runtime.Goexit()
  }

  if addCmd.Parsed() {
    if *addBlockData == "" {
      addCmd.Usage()
      runtime.Goexit()
    }
    // if add block data is not empty, add a new block with the new data
    cli.addBlock(*addBlockData)
  }

  if printCmd.Parsed() {
    cli.printBlock()
  }
}

func main() {
  fmt.Println("Hi from blockchain")
  defer os.Exit(0)
  bchain := blockchain.InitBlockChain()
  // call defer since InitBlockChain() opens a database, need to close the database
  defer bchain.Database.Close()

  cli := CommandLine{bchain}
  cli.run()

}
