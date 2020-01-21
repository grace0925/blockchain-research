# CHAPI
## Handling Credential Request Flow
1. An origin requests permission from user to handle credential request and storage
2. Credential handlers are defined in service worker code, use CredentialManager to set:
    - a list of enabled WebCredential types
    - Info used in the display of hints supported by the credential handler
3. Two entry points for interaction with a credential handler
    - Relying party (verifier) calls credential-mgmt-1 get(); (when user pushes button that requires identity attributes with WebCredentialRequestOptions)\
    - When issuer calls the credential-mgmt-1 store(); (when user pushes button to receive a credential with a Web Credential)
4. User agent displays a set of choices to the user (the registered hints of the candidate credential handlers)
5. User selects a hint, user agent fires a CredentialRequestEvent or a CredentialStoreEvent in the service worker whose CredentialManager the hint was registered with.