---
title: AI
description: LEM — Lethean Evaluation Model, training pipeline, scoring, and inference
---

# AI

`forge.lthn.ai/lthn/lem`

LEM (Lethean Evaluation Model) is an AI training and evaluation platform. A 1-billion-parameter model trained with 5 axioms consistently outperforms untrained models 27 times its size. 29 models tested, 3,000+ individual runs, two independent probe sets. Fully reproducible on Apple Silicon.

## Benchmark Highlights

| Model | Params | v2 Score | Notes |
|-------|--------|----------|-------|
| Gemma3 12B + LEK kernel | 12B | **23.66** | Best kernel-boosted |
| Gemma3 27B + LEK kernel | 27B | 23.26 | |
| **LEK-Gemma3 1B baseline** | **1B** | **21.74** | **Axioms in weights** |
| Base Gemma3 4B | 4B | 21.12 | Untrained |
| Base Gemma3 12B | 12B | 20.47 | Untrained |

## Packages

### pkg/lem

The core engine — 75+ files covering the full pipeline:

- **Training**: distillation, conversation generation, attention analysis, grammar integration
- **Scoring**: heuristic probes, tiered scoring, coverage analysis, judge evaluation
- **Inference**: Metal and mlx-lm backends, worker pool, client API
- **Data**: InfluxDB time-series storage, Parquet export, zstd compression, ingestion
- **Publishing**: Forgejo, Hugging Face, Docker registry integration
- **Analytics**: cluster analysis, feature extraction, metrics, comparison tools

### pkg/lab

LEM Lab — model store, configuration, local experimentation environment.

### pkg/heuristic

Standalone heuristic scoring engine for probe evaluation.

## Binaries

| Command | Purpose |
|---------|---------|
| `lemcmd` | Main CLI — training, scoring, publishing |
| `scorer` | Standalone scoring binary |
| `lem-desktop` | Desktop app (LEM Lab UI) |
| `composure-convert` | Training data format conversion |
| `dedup-check` | Dataset deduplication checker |

## Repository

- **Source**: [forge.lthn.ai/lthn/lem](https://forge.lthn.ai/lthn/lem)
- **Go module**: `forge.lthn.ai/lthn/lem`
- **Models**: [huggingface.co/lthn](https://huggingface.co/lthn)
