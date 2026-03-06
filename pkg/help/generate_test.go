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

// testCatalog builds a small catalog for generator tests.
func testCatalog() *Catalog {
	c := &Catalog{
		topics: make(map[string]*Topic),
		index:  newSearchIndex(),
	}
	c.Add(&Topic{
		ID:      "getting-started",
		Title:   "Getting Started",
		Content: "# Getting Started\n\nWelcome to the **guide**.\n",
		Tags:    []string{"intro"},
		Sections: []Section{
			{ID: "getting-started", Title: "Getting Started", Level: 1},
		},
	})
	c.Add(&Topic{
		ID:      "config",
		Title:   "Configuration",
		Content: "# Configuration\n\nSet up your environment.\n",
		Tags:    []string{"setup"},
		Related: []string{"getting-started"},
	})
	return c
}

func TestGenerate_Good_FileStructure(t *testing.T) {
	dir := t.TempDir()
	catalog := testCatalog()

	err := Generate(catalog, dir)
	require.NoError(t, err)

	// Verify expected file structure
	expectedFiles := []string{
		"index.html",
		"search.html",
		"search-index.json",
		"404.html",
		"topics/getting-started.html",
		"topics/config.html",
	}

	for _, f := range expectedFiles {
		path := filepath.Join(dir, f)
		_, err := os.Stat(path)
		assert.NoError(t, err, "expected file %s to exist", f)
	}
}

func TestGenerate_Good_IndexContainsTopics(t *testing.T) {
	dir := t.TempDir()
	catalog := testCatalog()

	err := Generate(catalog, dir)
	require.NoError(t, err)

	content, err := os.ReadFile(filepath.Join(dir, "index.html"))
	require.NoError(t, err)

	html := string(content)
	assert.Contains(t, html, "Getting Started")
	assert.Contains(t, html, "Configuration")
}

func TestGenerate_Good_TopicContainsRenderedMarkdown(t *testing.T) {
	dir := t.TempDir()
	catalog := testCatalog()

	err := Generate(catalog, dir)
	require.NoError(t, err)

	content, err := os.ReadFile(filepath.Join(dir, "topics", "getting-started.html"))
	require.NoError(t, err)

	html := string(content)
	assert.Contains(t, html, "Getting Started")
	assert.Contains(t, html, "<strong>guide</strong>")
}

func TestGenerate_Good_SearchIndexJSON(t *testing.T) {
	dir := t.TempDir()
	catalog := testCatalog()

	err := Generate(catalog, dir)
	require.NoError(t, err)

	content, err := os.ReadFile(filepath.Join(dir, "search-index.json"))
	require.NoError(t, err)

	var entries []searchIndexEntry
	require.NoError(t, json.Unmarshal(content, &entries))
	assert.Len(t, entries, 2, "search index should contain all topics")

	// Verify fields are populated
	ids := make(map[string]bool)
	for _, e := range entries {
		ids[e.ID] = true
		assert.NotEmpty(t, e.Title)
		assert.NotEmpty(t, e.Content)
	}
	assert.True(t, ids["getting-started"])
	assert.True(t, ids["config"])
}

func TestGenerate_Good_404Exists(t *testing.T) {
	dir := t.TempDir()
	catalog := testCatalog()

	err := Generate(catalog, dir)
	require.NoError(t, err)

	content, err := os.ReadFile(filepath.Join(dir, "404.html"))
	require.NoError(t, err)

	html := string(content)
	assert.Contains(t, html, "404")
	assert.Contains(t, html, "Not Found")
}

func TestGenerate_Good_EmptyDir(t *testing.T) {
	dir := t.TempDir()
	catalog := testCatalog()

	// Should succeed in an empty directory
	err := Generate(catalog, dir)
	assert.NoError(t, err)
}

func TestGenerate_Good_OverwriteExisting(t *testing.T) {
	dir := t.TempDir()
	catalog := testCatalog()

	// Generate once
	err := Generate(catalog, dir)
	require.NoError(t, err)

	// Generate again -- should overwrite without error
	err = Generate(catalog, dir)
	assert.NoError(t, err)

	// Verify files still exist and are valid
	content, err := os.ReadFile(filepath.Join(dir, "index.html"))
	require.NoError(t, err)
	assert.Contains(t, string(content), "Getting Started")
}

func TestGenerate_Good_SearchPageHasScript(t *testing.T) {
	dir := t.TempDir()
	catalog := testCatalog()

	err := Generate(catalog, dir)
	require.NoError(t, err)

	content, err := os.ReadFile(filepath.Join(dir, "search.html"))
	require.NoError(t, err)

	html := string(content)
	assert.Contains(t, html, "<script>")
	assert.Contains(t, html, "search-index.json")
}

func TestGenerate_Good_EmptyCatalog(t *testing.T) {
	dir := t.TempDir()
	catalog := &Catalog{
		topics: make(map[string]*Topic),
		index:  newSearchIndex(),
	}

	err := Generate(catalog, dir)
	require.NoError(t, err)

	// index.html should still exist
	_, err = os.Stat(filepath.Join(dir, "index.html"))
	assert.NoError(t, err)

	// search-index.json should be valid empty array
	content, err := os.ReadFile(filepath.Join(dir, "search-index.json"))
	require.NoError(t, err)
	var entries []searchIndexEntry
	require.NoError(t, json.Unmarshal(content, &entries))
	assert.Empty(t, entries)
}
