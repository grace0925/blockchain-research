# blockchain
A chain of blocks that contains information. Used to timestamp digital documents to avoid backdating or tempering.\

## Architecture
### Block
Data stored inside a block depends on the type of blockchain. Eg: Bitcoin Book contains information about the Sender, Receiver and number of bitcoins transferred.\
Each new block is linked to the previous block.
### Hash
Unique to each blocks, identifies a block and all of its content.\
Any change inside the block will cause the hash to change.\
Each block has:
- Data
- Hash
- Hash of the previous block
Prevent tempering.
### Proof of Work
Mechanism that slows down the creation of the new blocks. Makes blockchain secure.
### Distributed P2P Network
Distributed peer-peer network that everyone is allowed to join.\
Each computer is called a node and everyone gets the full copy of the blockchain.\
**Consensus:** all the nodes in the P2P network agree on what blocks are valid and reject blocks that are tampered with.\
To tamper a blockchain:
- Need to tamper all blocks
- Need to redo proof-of-work for each block
- Take control of greater than 50% of the peer-peer network

## Blockchain Transaction
