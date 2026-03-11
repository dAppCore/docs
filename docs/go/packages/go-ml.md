# go-ml

ML inference backends, multi-suite scoring engine, and agent orchestrator.

**Module**: `forge.lthn.ai/core/go-ml`

Provides pluggable backends (Apple Metal via go-mlx, managed llama-server subprocesses, and OpenAI-compatible HTTP APIs), a concurrent scoring engine that evaluates model responses across heuristic, semantic, content, and standard benchmark suites, 23 capability probes, GGUF model management, and an SSH-based agent orchestrator that streams checkpoint evaluation results to InfluxDB and DuckDB.

## Quick Start

```go
import "forge.lthn.ai/core/go-ml"

// HTTP backend (Ollama, LM Studio, any OpenAI-compatible endpoint)
backend := ml.NewHTTPBackend("http://localhost:11434", "qwen3:8b")
resp, err := backend.Generate(ctx, "Hello", ml.GenOpts{MaxTokens: 256})

// Scoring engine
engine := ml.NewEngine(backend, ml.Options{Suites: "heuristic,semantic", Concurrency: 4})
scores := engine.ScoreAll(responses)
```

## Components

| Component | Description |
|-----------|-------------|
| Backends | Metal (go-mlx), llama-server subprocess, HTTP (OpenAI-compatible) |
| Scoring | Heuristic, semantic, content, and benchmark suites |
| Probes | 23 capability probes for model evaluation |
| GGUF | Model management and metadata parsing |
| Orchestrator | SSH-based agent dispatch with InfluxDB/DuckDB streaming |
