---
title: go-forge
description: Forgejo API client covering actions, admin, repos, branches, and more
---

# go-forge

`forge.lthn.ai/core/go-forge`

Comprehensive API client for Forgejo (Gitea-compatible) instances. Covers the full Forgejo REST API surface including repository management, CI/CD actions, admin operations, branch protection, and user management. Handles authentication, rate limiting, and pagination.

## Key Types

- `ActionsService` — CI/CD actions and workflow management
- `AdminService` — server administration operations
- `BranchService` — branch creation, protection, and deletion
- `APIError` — structured error from the Forgejo API
- `RateLimit` — rate limit state tracking for API calls
