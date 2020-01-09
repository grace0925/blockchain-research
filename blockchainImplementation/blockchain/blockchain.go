package blockchain

import (
	"fmt"
	"github.com/dgraph-io/badger"
)

const (
	dbPath = "./tmp/blocks" // path to database
)

type BlockChain struct {
	LastHash []byte // hash of last block in the chain
	Database *badger.DB // pointer to the database
}

type BlockChainIterator struct {
	CurrentHash []byte
	Database *badger.DB
}

// create new blockchain
func InitBlockChain() *BlockChain{
	var lastHash []byte

	opts := badger.DefaultOptions(dbPath)
	opts.Dir = dbPath // where keys and metadata are stored
	opts.ValueDir = dbPath // where values are stored
	db, err := badger.Open(opts) // open database
	Handle(err)

	var lastH []byte

	err = db.Update(func(txn *badger.Txn)error{
		// if "lh" can't be found in the db
		if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {
			fmt.Println("No blockchain found")
			genesis := Genesis()
			fmt.Println("Genesis generated and proved")

			err = txn.Set(genesis.Hash, genesis.Serialize()) // use hash of genesis as key for genesis []byte
			Handle(err)
			err = txn.Set([]byte("lh"), genesis.Hash) // set the genesis hash as the last hash of the database
			lastHash = genesis.Hash
		} else {
			lh, er := txn.Get([]byte("lh"))
			Handle(er)
			er = lh.Value(func(lastHash []byte)error {
				lastH = lastHash
				return nil
			})
		}
		return nil
	})
	Handle(err)

	blockchain := BlockChain{lastH, db}
	return &blockchain
}

// add blocks to blockchain
func (chain *BlockChain) AddBlock(data string) {
	var lastHash []byte
	chain.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		Handle(err)
		err = item.Value(func(lh []byte)error {
			lastHash = lh
			return nil
		})

		newBlock := CreateBlock(data, lastHash)

		err = chain.Database.Update(func(txn *badger.Txn) error{
			err := txn.Set(newBlock.Hash, newBlock.Serialize())
			Handle(err)
			err = txn.Set([]byte("lh"), newBlock.Hash)
			chain.LastHash = newBlock.Hash
			return err
		})
		return err
	})
}

// iterator continues calling next until it reaches the end
func (chain *BlockChain) Iterator() *BlockChainIterator {
	iterator := &BlockChainIterator{chain.LastHash, chain.Database}
	return iterator
}

func (iterator *BlockChainIterator) Next() *Block {
	var block *Block
	var encodedBlock []byte
	err := iterator.Database.View(func(txn *badger.Txn) error {
		// get value of serialized block using value of hash
		item, err := txn.Get(iterator.CurrentHash)
		Handle(err)
		err = item.Value(func(val []byte)error {
			encodedBlock = val
			block = Deserialize(encodedBlock)
			return err
		})

		return err
	})
	Handle(err)
	// iterate to the previous block in the blockchain
	iterator.CurrentHash = block.PrevHash
	return block
}