# blockchain-research
## Decentralized Identifiers (DIDs)
New type of identifier to provide verifiable, decentralized digital identity.\
**DIDs:** URLs that relate a DID subject to a DID document allowing trustable interactions with that subject. Consist of 3 parts:\
  - URL scheme Identifier
  - Idnetifier for the DID method
  - DID method-specific Identifier
Eg. did:example:123456789abcdefghi

**DID document:** simple documents describing how to use the specific DID.

- Conventional identity management systems are based on centralized authorities. Each centralized authority serves as its own root of trust.
- In DI system, entities are free to use any shared root of trust. Manage a root of trust without introducing a centralized authority or a single point of failure.
  - Entities are identified by DIDs and can authenticate using proofs.
  - DIDs point to DID documents (contain a set of service endpoints for interacting with what the DID identifies).

### Design Goal
- Eliminate the need for centralized authorities or single point failure in identifier management.
- Fuve both human and non-human entities the power to directly control their digital identifiers without the need to rely on external authorities. Each entity controls own privacy.
- Proof based.

### Data Model
