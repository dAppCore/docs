---
title: go-webview
description: Chrome DevTools Protocol client for browser automation and testing
---

# go-webview

`forge.lthn.ai/core/go-webview`

Chrome DevTools Protocol (CDP) client for browser automation, testing, and scraping. Supports navigation, DOM queries, click/type interactions, console capture, JavaScript evaluation, screenshots, multi-tab management, and viewport emulation. Includes an ActionSequence builder for chaining operations and Angular-specific helpers.

## Key Types

- `ClickAction` — clicks an element by selector
- `TypeAction` — types text into an input element
- `NavigateAction` — navigates to a URL and waits for load
- `WaitAction` — waits for a selector, timeout, or network idle
