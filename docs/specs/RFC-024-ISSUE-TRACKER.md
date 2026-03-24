# RFC-024: Issue Tracker & Sprint System

**Date:** 2026-03-16
**Status:** Draft
**Author:** Cladius Maximus
**Scope:** CorePHP (php-agentic module)

## Summary

A lightweight issue tracker built into the CorePHP agentic module. Issues flow from discovery (scans, webhooks, user reports) through triage, sprint planning, agent dispatch, PR, merge, and auto-release.

## Motivation

Manually tagging 40+ Go repos bottom-up is painful. No system connects "work found" to "work done" to "version released". The agentic dispatch system executes work, but nothing orchestrates what to work on or when to release.

## Design Principles

1. Issues are a spectrum — from "fix this typo" to "add GoDot to the IDE"
2. Sprints are dumb — just "start" and "complete"
3. Milestones are buckets — next-patch, next-minor, next-major, ideas, backlog
4. Changelogs are derived — auto-generated from merged PRs per milestone
5. Agents are assignees — Cladius, Charon, Gemini, local model, or human
6. Projects are repos — tracked by php-uptelligence
7. Triage is automated — local model reports size/scope

## Data Model

### Issue

| Field | Type | Purpose |
|-------|------|---------|
| repo | string | Project = repo name |
| title, body | string, text | What needs doing |
| status | enum | open, assigned, in_progress, review, done, closed |
| priority | enum | critical, high, normal, low |
| milestone | enum | next-patch, next-minor, next-major, ideas, backlog |
| size | enum | trivial, small, medium, large, epic |
| source | string | scan, user, forge, github, discovery, cve |
| source_ref | string | URL or workspace ID |
| assignee | string | Agent name or user handle |
| labels | JSON | e.g. ["security", "conventions"] |
| pr_url | string | Linked PR |
| plan_id | FK nullable | Escalated to AgentPlan |
| parent_id | FK nullable | Epic child relationship |
| metadata | JSON | Flexible (scan results, triage notes) |

### Sprint

| Field | Type | Purpose |
|-------|------|---------|
| name | string | "Sprint 2026-W12" or custom |
| status | enum | planning, active, completed |
| started_at | timestamp | When sprint was started |
| completed_at | timestamp | When sprint was completed |
| notes | text | Retrospective / release notes |
| metadata | JSON | Repos touched, tags created |

### IssueComment

| Field | Type | Purpose |
|-------|------|---------|
| issue_id | FK | Parent issue |
| author | string | Agent name or user |
| body | text | Comment content |
| type | enum | comment, triage, scan_result, status_change |

## Issue Sizing

| Size | Scope | Plans | Example |
|------|-------|-------|---------|
| trivial | One line | 0 | Fix typo |
| small | One file | 0 | Alias import |
| medium | Multiple files, one repo | 1 | Refactor to use go-io |
| large | Multiple plans, one repo | 2+ | Security audit fixes |
| epic | Multiple repos | N children | Add GoDot to IDE |

## Sprint Flow

Start sprint: takes everything in next-* milestones, marks active.
Complete sprint: triggers core dev tag, generates changelogs, closes done issues.

## Auto-Versioning Pipeline

PR opened: v0.3.2-alpha.{pr}
PR updated: v0.3.2-alpha.{pr}+build.{n}
PR merged: v0.3.2-beta.{n}
Sprint complete: v0.3.2 (stable release)

core dev tag handles bottom-up dependency chain.
Downstream repos get webhook, go get -u, test, auto-PR.

## Discovery Engine

Scheduled action finding repos needing attention. Repos tracked by php-uptelligence with no open issues get scanned automatically. Scan types: conventions, dependency, security, CVE, test gaps, doc sync, dead code. Results create issues in the inbox.

## API Endpoints

Issues: GET/POST /v1/issues, GET/PATCH/DEL /v1/issues/{id}
Comments: POST/GET /v1/issues/{id}/comments
Sprints: GET/POST /v1/sprints, POST /v1/sprints/{id}/start, POST /v1/sprints/{id}/complete
Projects: GET /v1/projects, GET /v1/projects/{repo}/changelog
Milestones: GET /v1/milestones

## MCP Tools

agentic_issue_create, agentic_issue_list, agentic_issue_update, agentic_issue_triage, agentic_sprint_start, agentic_sprint_complete

## Integration Points

Forge/GitHub (bidirectional issue sync), php-uptelligence (discovery), agentic dispatch (issue to PR), core dev tag (sprint to release), CodeRabbit (PR review), OpenBrain (context), Sentry (error reports)

## UI (Flux UI Pro)

Board (Kanban), List (filterable table with project dropdown), Sprint (progress + repo breakdown)

## Implementation Order

1. Models + migrations
2. API endpoints
3. MCP tools
4. Discovery cron
5. Forge sync
6. UI
7. Auto-changelog
8. Auto-tag
