# go-mlx

Native Apple Metal GPU inference via mlx-c CGO bindings.

**Module**: `forge.lthn.ai/core/go-mlx`

Implements the `inference.Backend` and `inference.TextModel` interfaces from go-inference for Apple Silicon (M1-M4). Supports Gemma 3, Qwen 2/3, and Llama 3 architectures from HuggingFace safetensors format, with fused Metal kernels for RMSNorm, RoPE, and scaled dot-product attention, KV cache management, LoRA fine-tuning with AdamW, and batch inference. A Python subprocess backend (`mlxlm`) is provided as a CGO-free alternative.

!!! warning "Platform restriction"
    `darwin/arm64` only. A no-op stub compiles on all other platforms.

## Quick Start

```go
import (
    "forge.lthn.ai/core/go-inference"
    _ "forge.lthn.ai/core/go-mlx"  // registers "metal" backend via init()
)

model, err := inference.LoadModel("/Volumes/Data/lem/safetensors/gemma-3-1b/")
defer model.Close()

for tok := range model.Generate(ctx, "Hello", inference.WithMaxTokens(256)) {
    fmt.Print(tok.Text)
}
```

## Supported Architectures

| Architecture | Models |
|-------------|--------|
| Gemma 3 | gemma-3-1b, gemma-3-4b |
| Qwen 2/3 | qwen3-8b, qwen2.5-* |
| Llama 3 | llama-3.2-*, llama-3.1-* |
