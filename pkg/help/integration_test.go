// SPDX-Licence-Identifier: EUPL-1.2
package help

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntegration_Good_FullPipeline(t *testing.T) {
	// 1. Create content directory
	contentDir := t.TempDir()
	require.NoError(t, os.MkdirAll(filepath.Join(contentDir, "cli"), 0o755))
	require.NoError(t, os.MkdirAll(filepath.Join(contentDir, "go"), 0o755))

	require.NoError(t, os.WriteFile(filepath.Join(contentDir, "cli", "dev-work.md"), []byte(`---
title: Dev Work
tags: [cli]
order: 1
---

## Usage

core dev work syncs your workspace.

## Flags

--status  Show status only
`), 0o644))

	require.NoError(t, os.WriteFile(filepath.Join(contentDir, "go", "go-scm.md"), []byte(`---
title: Go SCM
tags: [go, library]
order: 2
---

## Overview

Registry and git operations for the workspace.
`), 0o644))

	// 2. Load into catalog
	catalog, err := LoadContentDir(contentDir)
	require.NoError(t, err)
	assert.Len(t, catalog.List(), 2)

	// 3. Generate static site
	outputDir := t.TempDir()
	err = Generate(catalog, outputDir)
	require.NoError(t, err)

	// 4. Verify file structure
	assert.FileExists(t, filepath.Join(outputDir, "index.html"))
	assert.FileExists(t, filepath.Join(outputDir, "search.html"))
	assert.FileExists(t, filepath.Join(outputDir, "search-index.json"))
	assert.FileExists(t, filepath.Join(outputDir, "404.html"))
	assert.FileExists(t, filepath.Join(outputDir, "topics", "dev-work.html"))
	assert.FileExists(t, filepath.Join(outputDir, "topics", "go-scm.html"))

	// 5. Verify search index
	indexData, err := os.ReadFile(filepath.Join(outputDir, "search-index.json"))
	require.NoError(t, err)
	var entries []searchIndexEntry
	require.NoError(t, json.Unmarshal(indexData, &entries))
	assert.Len(t, entries, 2)

	// 6. Verify HTML content has HLCRF structure
	indexHTML, err := os.ReadFile(filepath.Join(outputDir, "index.html"))
	require.NoError(t, err)
	assert.Contains(t, string(indexHTML), "Dev Work")
	assert.Contains(t, string(indexHTML), "Go SCM")
	assert.Contains(t, string(indexHTML), `role="banner"`)
	assert.Contains(t, string(indexHTML), `role="main"`)

	// 7. Verify topic page has section anchors
	topicHTML, err := os.ReadFile(filepath.Join(outputDir, "topics", "dev-work.html"))
	require.NoError(t, err)
	assert.Contains(t, string(topicHTML), `href="#usage"`)
	assert.Contains(t, string(topicHTML), `href="#flags"`)
	assert.Contains(t, string(topicHTML), `role="complementary"`)

	// 8. Search works
	results := catalog.Search("workspace")
	assert.NotEmpty(t, results)
	assert.Equal(t, "dev-work", results[0].Topic.ID)
}
