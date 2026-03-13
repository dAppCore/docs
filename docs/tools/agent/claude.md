---
title: Claude Code Plugins
description: Claude Code plugin marketplace, installation, and command reference for the Core platform.
---

# Claude Code Plugins

Core Agent provides a **plugin marketplace** for Claude Code with development workflows, code review, verification, and language-specific tooling.


## Marketplace

The marketplace is hosted in `core/agent` and distributed via npm as `@lthn/core-agent`.

### Install the Marketplace

```bash
# From npm (recommended for most users)
claude plugin marketplace add @lthn/core-agent

# From git (Forge access)
claude plugin marketplace add https://forge.lthn.ai/core/agent.git
```

### Install All Plugins

Once the marketplace is added, install plugins individually:

```bash
claude plugin install code@core-agent
claude plugin install review@core-agent
claude plugin install verify@core-agent
claude plugin install core-php@core-agent
claude plugin install go-build@core-agent
claude plugin install devops@core-agent
```


## Plugins

### code

Development hooks, auto-approve workflow, and research data collection.

| Type | Name | Description |
|------|------|-------------|
| Command | `/remember` | Save context to memory |
| Command | `/yes` | Auto-approve workflow |
| Skill | `bitcointalk` | BitcoinTalk forum research |
| Skill | `block-explorer` | Block explorer data collection |
| Skill | `coinmarketcap` | CoinMarketCap data collection |
| Skill | `community-chat` | Community chat research |
| Skill | `cryptonote-discovery` | CryptoNote project discovery |
| Skill | `github-history` | GitHub history research |
| Skill | `job-collector` | Job posting collection |
| Skill | `ledger-papers` | Academic paper archive (crypto/blockchain) |
| Skill | `mining-pools` | Mining pool research |
| Skill | `project-archaeology` | Dead project salvage reports |
| Skill | `wallet-releases` | Wallet release tracking |
| Skill | `whitepaper-archive` | Whitepaper collection |

**Source:** `core/agent` &middot; **npm:** `@lthn/core-agent`

---

### review

Code review automation with multi-agent pipeline.

| Type | Name | Description |
|------|------|-------------|
| Command | `/review` | Perform code review on staged changes or PRs |
| Command | `/security` | Security-focused code review |
| Command | `/pr` | Review a pull request |
| Command | `/pipeline` | Run the 5-agent review pipeline |
| Skill | `architecture-review` | Architecture-level review |
| Skill | `reality-check` | Sanity check implementation against intent |
| Skill | `security-review` | Security vulnerability analysis |
| Skill | `senior-dev-fix` | Senior developer fix suggestions |
| Skill | `test-analysis` | Test coverage and quality analysis |

**Source:** `core/agent` &middot; **npm:** `@lthn/core-agent`

---

### verify

Work verification before committing.

| Type | Name | Description |
|------|------|-------------|
| Command | `/verify` | Verify work is complete before stopping |
| Command | `/ready` | Quick check if work is ready to commit |
| Command | `/tests` | Verify tests pass for changed files |

**Source:** `core/agent` &middot; **npm:** `@lthn/core-agent`

---

### core-php

PHP/Laravel development skills and API generation.

| Type | Name | Description |
|------|------|-------------|
| Command | `/api-generate` | Generate TypeScript/JavaScript API client from Laravel routes |
| Skill | `php` | PHP development conventions and patterns |
| Skill | `laravel` | Laravel framework patterns |
| Skill | `php-agent` | PHP agent development guidance |

**Source:** `core/php` &middot; **npm:** `@lthn/core-claude-php`

---

### go-build

Go QA pipeline and build tooling.

| Type | Name | Description |
|------|------|-------------|
| Command | `/qa` | Run full QA pipeline and fix all issues iteratively |
| Command | `/check` | Run QA checks without fixing (report only) |
| Command | `/fix` | Fix a specific QA issue |
| Command | `/lint` | Run linter and fix issues |

**Source:** `core/go-build` &middot; **npm:** `@lthn/core-claude-build`

---

### devops

CI/CD, deployment, and issue tracking.

| Type | Name | Description |
|------|------|-------------|
| Command | `/ci` | Check CI status and manage workflows |
| Command | `/ci-status` | Show CI status for current branch |
| Command | `/ci-run` | Trigger a CI workflow run |
| Command | `/ci-fix` | Analyse and fix failing CI |
| Command | `/ci-workflow` | Create or update CI workflow |
| Command | `/coolify-deploy` | Deploy a service to Coolify |
| Command | `/coolify-status` | Check Coolify deployment status |
| Command | `/issue-list` | List open issues |
| Command | `/issue-view` | View issue details |
| Command | `/issue-start` | Start working on an issue |
| Command | `/issue-close` | Close an issue with a commit |

**Source:** `core/go-devops` &middot; **npm:** `@lthn/core-claude-devops`


## npm Packages

All plugins are published to npm under the `@lthn` scope for public distribution:

| Package | Plugin | Version |
|---------|--------|---------|
| `@lthn/core-agent` | Marketplace (code, review, verify) | 0.2.0 |
| `@lthn/core-claude-php` | core-php | 0.1.0 |
| `@lthn/core-claude-build` | go-build | 0.1.0 |
| `@lthn/core-claude-devops` | devops | 0.1.0 |


## Plugin Development

Plugins live in `.claude-plugin/` directories within their source repositories. Each plugin has:

- `plugin.json` -- manifest with name, description, version
- `commands/` -- slash command markdown files (auto-discovered)
- `skills/` -- skill directories with `SKILL.md` files (auto-discovered)
- `hooks.json` -- optional event hooks

See the [Claude Code plugin documentation](https://code.claude.com/docs/en/plugins) for the full specification.


## Licence

EUPL-1.2
