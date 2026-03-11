---
title: go-rag
description: Retrieval-Augmented Generation with document chunking, embeddings, and vector search
---

# go-rag

`forge.lthn.ai/core/go-rag`

Retrieval-Augmented Generation pipeline. Splits documents using three-level Markdown-aware chunking, generates embeddings via Ollama, and stores/searches vectors in Qdrant over gRPC. Supports keyword boosting and multiple result formats (plain text, XML, JSON).

## Key Types

- `ChunkConfig` — configuration for document splitting (sizes, overlap, strategy)
- `Chunk` — individual text chunk with metadata and position information
- `IngestConfig` — configuration for the document ingestion pipeline
- `IngestStats` — statistics from a completed ingestion run (chunks created, embeddings generated)
