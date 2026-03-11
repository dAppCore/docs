# go-crypt

Cryptographic primitives, authentication, and trust policy engine.

**Module**: `forge.lthn.ai/core/go-crypt`

Provides symmetric encryption (ChaCha20-Poly1305 and AES-256-GCM with Argon2id KDF), OpenPGP challenge-response authentication with online and air-gapped courier modes, Argon2id password hashing, RSA-OAEP key generation, RFC-0004 deterministic content hashing, and a three-tier agent trust policy engine with an audit log and approval queue.

## Quick Start

```go
import (
    "forge.lthn.ai/core/go-crypt/crypt"
    "forge.lthn.ai/core/go-crypt/auth"
    "forge.lthn.ai/core/go-crypt/trust"
)

// Encrypt with ChaCha20-Poly1305 + Argon2id KDF
ciphertext, err := crypt.Encrypt(plaintext, passphrase)

// OpenPGP authentication
a := auth.New(medium, auth.WithSessionStore(auth.NewSQLiteSessionStore(dbPath)))
session, err := a.Login(userID, password)

// Trust policy evaluation
engine := trust.NewPolicyEngine(registry)
decision := engine.Evaluate("Charon", "repo.push", "core/go-crypt")
```

## Packages

| Package | Description |
|---------|-------------|
| `crypt` | ChaCha20-Poly1305, AES-256-GCM, Argon2id KDF |
| `auth` | OpenPGP challenge-response, session management |
| `trust` | Three-tier policy engine, audit log, approval queue |
| `hash` | RFC-0004 deterministic content hashing |
| `keys` | RSA-OAEP key generation |
