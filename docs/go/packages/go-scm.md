# go-scm

SCM integration, AgentCI dispatch automation, and data collection.

**Module**: `forge.lthn.ai/core/go-scm`

Provides a Forgejo API client and a Gitea client for the public mirror, multi-repo git operations with parallel status checks, the Clotho Protocol orchestrator for dual-run agent verification, a PR automation pipeline (poll -> dispatch -> journal) driven by epic issue task lists, and pluggable data collectors for BitcoinTalk, GitHub, market data, and research papers.

## Quick Start

```go
import (
    "forge.lthn.ai/core/go-scm/forge"
    "forge.lthn.ai/core/go-scm/git"
    "forge.lthn.ai/core/go-scm/jobrunner"
)

// Forgejo client
client, err := forge.NewFromConfig("", "")

// Multi-repo status
statuses := git.Status(ctx, git.StatusOptions{Paths: repoPaths})

// AgentCI dispatch loop
poller := jobrunner.NewPoller(jobrunner.PollerConfig{
    Sources:      []jobrunner.JobSource{forgejoSrc},
    Handlers:     []jobrunner.JobHandler{dispatch, tickParent},
    PollInterval: 60 * time.Second,
})
poller.Run(ctx)
```

## Components

| Component | Description |
|-----------|-------------|
| `forge` | Forgejo API client |
| `git` | Multi-repo operations, parallel status |
| `jobrunner` | AgentCI dispatch pipeline |
| `clotho` | Dual-run verification protocol |
| `collectors` | BitcoinTalk, GitHub, market, research scrapers |
