// SPDX-Licence-Identifier: EUPL-1.2
package help

import (
	. "dappco.re/go"
)

func TestIntegration_Good_FullPipeline(t *T) {
	// 1. Create content directory
	contentDir := t.TempDir()
	if r := MkdirAll(PathJoin(contentDir, "cli"), 0o755); !r.OK {
		t.Fatal(r.Error())
	}
	if r := MkdirAll(PathJoin(contentDir, "go"), 0o755); !r.OK {
		t.Fatal(r.Error())
	}

	if r := WriteFile(PathJoin(contentDir, "cli", "dev-work.md"), []byte(`---
title: Dev Work
tags: [cli]
order: 1
---

## Usage

core dev work syncs your workspace.

## Flags

--status  Show status only
`), 0o644); !r.OK {
		t.Fatal(r.Error())
	}

	if r := WriteFile(PathJoin(contentDir, "go", "go-scm.md"), []byte(`---
title: Go SCM
tags: [go, library]
order: 2
---

## Overview

Registry and git operations for the workspace.
`), 0o644); !r.OK {
		t.Fatal(r.Error())
	}

	// 2. Load into catalog
	res1 := LoadContentDir(contentDir)
	if !res1.OK { t.Fatal(res1.Error()) }
	catalog := res1.Value.(*Catalog)
	AssertLen(t, catalog.List(), 2)

	// 3. Generate static site
	outputDir := t.TempDir()
	res2 := Generate(catalog, outputDir)
	if !res2.OK { t.Fatal(res2.Error()) }

	// 4. Verify file structure
	AssertNoError(t, statExists(PathJoin(outputDir, "index.html")))
	AssertNoError(t, statExists(PathJoin(outputDir, "search.html")))
	AssertNoError(t, statExists(PathJoin(outputDir, "search-index.json")))
	AssertNoError(t, statExists(PathJoin(outputDir, "404.html")))
	AssertNoError(t, statExists(PathJoin(outputDir, "topics", "dev-work.html")))
	AssertNoError(t, statExists(PathJoin(outputDir, "topics", "go-scm.html")))

	// 5. Verify search index
	indexData := readFileBytes(t, PathJoin(outputDir, "search-index.json"))
	var entries []searchIndexEntry
	if r := JSONUnmarshal(indexData, &entries); !r.OK {
		t.Fatal(r.Error())
	}
	AssertLen(t, entries, 2)

	// 6. Verify HTML content has HLCRF structure
	indexHTML := readFileBytes(t, PathJoin(outputDir, "index.html"))
	AssertContains(t, string(indexHTML), "Dev Work")
	AssertContains(t, string(indexHTML), "Go SCM")
	AssertContains(t, string(indexHTML), `role="banner"`)
	AssertContains(t, string(indexHTML), `role="main"`)

	// 7. Verify topic page has section anchors
	topicHTML := readFileBytes(t, PathJoin(outputDir, "topics", "dev-work.html"))
	AssertContains(t, string(topicHTML), `href="#usage"`)
	AssertContains(t, string(topicHTML), `href="#flags"`)
	AssertContains(t, string(topicHTML), `role="complementary"`)

	// 8. Search works
	results := catalog.Search("workspace")
	AssertNotEmpty(t, results)
	AssertEqual(t, "dev-work", results[0].Topic.ID)
}
