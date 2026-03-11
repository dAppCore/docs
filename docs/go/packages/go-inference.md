# go-inference

Shared interface contract for text generation backends.

**Module**: `forge.lthn.ai/core/go-inference`

Defines `TextModel`, `Backend`, `Token`, `Message`, and associated configuration types that GPU-specific backends implement and consumers depend on. Zero external dependencies — stdlib only — and compiles on all platforms regardless of GPU availability. The backend registry supports automatic selection (Metal preferred on macOS, ROCm on Linux) and explicit pinning.

## Quick Start

```go
import (
    "forge.lthn.ai/core/go-inference"
    _ "forge.lthn.ai/core/go-mlx"   // registers "metal" backend on darwin/arm64
)

model, err := inference.LoadModel("/path/to/safetensors/model/")
defer model.Close()

for tok := range model.Generate(ctx, "Hello", inference.WithMaxTokens(256)) {
    fmt.Print(tok.Text)
}
```

## Key Interfaces

```go
type TextModel interface {
    Generate(ctx context.Context, prompt string, opts ...Option) iter.Seq[Token]
    Close() error
}

type Backend interface {
    LoadModel(path string) (TextModel, error)
    Name() string
}
```

## Backend Registry

| Backend | Platform | Registration |
|---------|----------|-------------|
| Metal (go-mlx) | darwin/arm64 | Auto via `init()` |
| ROCm (go-rocm) | linux/amd64 | Auto via `init()` |
| HTTP | All | Explicit via `NewHTTPBackend()` |
