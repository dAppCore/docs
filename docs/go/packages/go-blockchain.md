# go-blockchain

Pure Go implementation of the Lethean blockchain protocol.

**Module**: `forge.lthn.ai/core/go-blockchain`

Provides chain configuration, core cryptographic data types, CryptoNote wire serialisation, and LWMA difficulty adjustment for the Lethean CryptoNote/Zano-fork chain. Lineage: CryptoNote to IntenseCoin (2017) to Lethean to Zano rebase.

## Quick Start

```go
import (
    "forge.lthn.ai/core/go-blockchain/config"
    "forge.lthn.ai/core/go-blockchain/types"
    "forge.lthn.ai/core/go-blockchain/wire"
    "forge.lthn.ai/core/go-blockchain/difficulty"
)

// Query the active hardfork version at a given block height
version := config.VersionAtHeight(config.MainnetForks, 10081) // returns HF2

// Encode and decode a Lethean address
addr := &types.Address{SpendPublicKey: spendKey, ViewPublicKey: viewKey}
encoded := addr.Encode(config.AddressPrefix)
decoded, prefix, err := types.DecodeAddress(encoded)

// Varint encoding for the wire protocol
buf := wire.EncodeVarint(0x1eaf7)
val, n, err := wire.DecodeVarint(buf)

// Calculate next block difficulty
nextDiff := difficulty.NextDifficulty(timestamps, cumulativeDiffs, 120)
```

## Packages

| Package | Description |
|---------|-------------|
| `config` | Hardfork schedule, network parameters, address prefixes |
| `types` | Addresses, keys, hashes, amounts |
| `wire` | CryptoNote binary serialisation, varint codec |
| `difficulty` | LWMA difficulty adjustment algorithm |
