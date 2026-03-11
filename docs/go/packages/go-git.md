---
title: go-git
description: Git operations helper for status, push, and repository queries
---

# go-git

`forge.lthn.ai/core/go-git`

Helper library for common Git operations. Provides structured output for repository status, push results, and error handling around Git CLI commands.

## Key Types

- `RepoStatus` — parsed state of a Git repository (branch, dirty files, ahead/behind)
- `StatusOptions` — options controlling what status information to collect
- `PushResult` — structured result of a push operation
- `GitError` — typed error wrapping Git CLI failures with exit codes
- `QueryStatus` — batch query result for multiple repositories
