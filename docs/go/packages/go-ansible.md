---
title: go-ansible
description: Ansible executor for programmatic playbook and ad-hoc command execution
---

# go-ansible

`forge.lthn.ai/core/go-ansible`

Go wrapper around Ansible for executing playbooks and ad-hoc commands programmatically. Provides a clean API for running Ansible operations with structured output parsing and SSH client abstraction for testing.

## Key Types

- `Executor` — runs Ansible playbooks and ad-hoc commands, captures structured output
- `MockSSHClient` — test double for SSH connections, enables unit testing without live hosts
