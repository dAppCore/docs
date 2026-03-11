# go-i18n

Grammar engine for Go.

**Module**: `forge.lthn.ai/core/go-i18n`

Provides forward composition primitives (PastTense, Gerund, Pluralize, Article, composite progress and label functions), a `T()` translation entry point with namespace key handlers, and a reversal engine that recovers base forms and grammatical roles from inflected text. The reversal package produces `GrammarImprint` feature vectors for semantic similarity scoring, builds reference domain distributions, performs anomaly detection, and includes a 1B model pre-sort pipeline for training data classification.

## Quick Start

```go
import "forge.lthn.ai/core/go-i18n"

// Grammar primitives
fmt.Println(i18n.PastTense("delete"))   // "deleted"
fmt.Println(i18n.Gerund("build"))       // "building"
fmt.Println(i18n.Pluralize("file", 3))  // "files"

// Translation with auto-composed output
fmt.Println(i18n.T("i18n.progress.build"))  // "Building..."
fmt.Println(i18n.T("i18n.done.delete", "file"))  // "File deleted"

// Reversal: recover grammar from text
tokeniser := reversal.NewTokeniser()
tokens := tokeniser.Tokenise("deleted the files")
imprint := reversal.NewImprint(tokens)
```

## Components

| Component | Description |
|-----------|-------------|
| Forward | PastTense, Gerund, Pluralize, Article, progress/label composers |
| `T()` | Translation entry point with namespace key handlers |
| Reversal | Base form recovery, grammatical role detection |
| GrammarImprint | Feature vectors for semantic similarity scoring |
| 1B Pipeline | Training data pre-sort and classification |
