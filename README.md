# blockchain-research
## Decentralized Identifiers (DIDs)
New type of identifier to provide verifiable, decentralized digital identity.\
**DIDs:** URLs that relate a DID subject to a DID document allowing trustable interactions with that subject.\
**DID document:** simple documents describing how to use the specific DID.\

- Conventional identity management systems are based on centralized authorities. Each centralized authority serves as its own root of trust.
- In DI system, entities are free to use any shared root of trust. Manage a root of trust without introducing a centralized authority or a single point of failure.
  - Entities are identified by DIDs and can authenticate using proofs.
  - DIDs point to DID documents (contain a set of service endpoints for interacting with what the DID identifies).
