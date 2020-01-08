# Verifiable Credentials Data Model 1.0
Make credentials on the web cryptographically secure, privacy respecting and machine-verifiable.

## Ecosystem
**Holder:** An entity who possesses one or more verifiable credentials and generates verifiable presentations from them.\
**Issuer:** An entity who asserts claims about one or more subjects, creates a verifiable credential from these claims, and transmits the verifiable credential to a holder.\
**Subject::** An entity about which claims are made. (Human, animals and things). Like the pet.\
**Verifier:** An entity who receives one or more verifiable credentials. (Employers, security personnel).\
![Ecosystem Roles](/Images/Ecosystem1.png)

## Core Data Model
### Claims
**Claim:** Statement about a subject.\
**Subject:** A thing about which claims can be made.\
### Credentials
**Credentials:** A set of one or more claims amde by the same entity. Might include identifiers or metadata to describe properties of the credential.\
**Verifiable Credentials:** A set of tamper-evident claims and metadata that cryptographically prove who issued it.\
![Verifiable Credential Comp](/Images/VerifiableCred.png)\
First graph expresses the verifiable credential itself (contains credential metadata and claims).\
Second graph expresses the digital proof.\
![Complete Verifiable Credential Comp](/Images/CompleteVeriCred.png)\
### Presentations
Expresses data from one or more verifiable credentials. Data in a presentation is often about the same subject, might have been issued by multiple issuers.
![Presentation](/Images/Presentations.png)\
![Presentation Graph](/Images/PresentationGraph.png)\
