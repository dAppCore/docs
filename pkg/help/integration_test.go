// SPDX-Licence-Identifier: EUPL-1.2
package help

import (
	. "dappco.re/go"
	"encoding/json"
	"os"
	"path/filepath"
)

func TestIntegration_Good_FullPipeline(t *T) {
	// 1. Create content directory
	contentDir := t.TempDir()
	RequireNoError(t, os.MkdirAll(filepath.Join(contentDir, "cli"), 0o755))
	RequireNoError(t, os.MkdirAll(filepath.Join(contentDir, "go"), 0o755))

	RequireNoError(t, os.WriteFile(filepath.Join(contentDir, "cli", "dev-work.md"), []byte(`---
title: Dev Work
tags: [cli]
order: 1
---

## Usage

core dev work syncs your workspace.

## Flags

--status  Show status only
`), 0o644))

	RequireNoError(t, os.WriteFile(filepath.Join(contentDir, "go", "go-scm.md"), []byte(`---
title: Go SCM
tags: [go, library]
order: 2
---

## Overview

Registry and git operations for the workspace.
`), 0o644))

	// 2. Load into catalog
	catalog, err := LoadContentDir(contentDir)
	RequireNoError(t, err)
	AssertLen(t, catalog.List(), 2)

	// 3. Generate static site
	outputDir := t.TempDir()
	err = Generate(catalog, outputDir)
	RequireNoError(t, err)

	// 4. Verify file structure
	_, err = os.Stat(filepath.Join(outputDir, "index.html"))
	AssertNoError(t, err)
	_, err = os.Stat(filepath.Join(outputDir, "search.html"))
	AssertNoError(t, err)
	_, err = os.Stat(filepath.Join(outputDir, "search-index.json"))
	AssertNoError(t, err)
	_, err = os.Stat(filepath.Join(outputDir, "404.html"))
	AssertNoError(t, err)
	_, err = os.Stat(filepath.Join(outputDir, "topics", "dev-work.html"))
	AssertNoError(t, err)
	_, err = os.Stat(filepath.Join(outputDir, "topics", "go-scm.html"))
	AssertNoError(t, err)

	// 5. Verify search index
	indexData, err := os.ReadFile(filepath.Join(outputDir, "search-index.json"))
	RequireNoError(t, err)
	var entries []searchIndexEntry
	RequireNoError(t, json.Unmarshal(indexData, &entries))
	AssertLen(t, entries, 2)

	// 6. Verify HTML content has HLCRF structure
	indexHTML, err := os.ReadFile(filepath.Join(outputDir, "index.html"))
	RequireNoError(t, err)
	AssertContains(t, string(indexHTML), "Dev Work")
	AssertContains(t, string(indexHTML), "Go SCM")
	AssertContains(t, string(indexHTML), `role="banner"`)
	AssertContains(t, string(indexHTML), `role="main"`)

	// 7. Verify topic page has section anchors
	topicHTML, err := os.ReadFile(filepath.Join(outputDir, "topics", "dev-work.html"))
	RequireNoError(t, err)
	AssertContains(t, string(topicHTML), `href="#usage"`)
	AssertContains(t, string(topicHTML), `href="#flags"`)
	AssertContains(t, string(topicHTML), `role="complementary"`)

	// 8. Search works
	results := catalog.Search("workspace")
	AssertNotEmpty(t, results)
	AssertEqual(t, "dev-work", results[0].Topic.ID)
}
