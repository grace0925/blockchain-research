# Hyperledger-Fabric 
**Consortium:** The group of organizations that are members of a Fabric network.\
**Peers:**  The fundamental components of any Fabric network.
    - Store the blockchain ledger and validate transactions before they are committed to the ledger. 
    - Run smart contracts.
    - Every peer in the network needs to belong to a member of the consortium. \
**Ordering Service:** Decide on the order of transactions.
    - Allows peers to focus on validating transactions and committing them to the ledger.\
**Channel:** Can be used only by organizations that are invited to the channel. Each channel has a separate blockchain ledger.\
**Smart Contracts:** Contain the business logic that governs assets on the blockchain ledger.
    - Transactions created using smart contracts typically need to be signed by multiple organizations to be committedto the channel ledger.