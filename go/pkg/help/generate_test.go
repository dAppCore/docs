// SPDX-Licence-Identifier: EUPL-1.2
package help

import (
	. "dappco.re/go"
)

// statExists is a small test helper that reports OK iff the path exists.
func statExists(path string) error {
	r := Stat(path)
	if !r.OK {
		return r.Value.(error)
	}
	return nil
}

// readFileBytes returns the bytes at path or fails the test.
func readFileBytes(t *T, path string) []byte {
	t.Helper()
	r := ReadFile(path)
	if !r.OK {
		t.Fatal(r.Error())
	}
	return r.Value.([]byte)
}

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

func TestGenerate_Good_FileStructure(t *T) {
	dir := t.TempDir()
	catalog := testCatalog()

	err := Generate(catalog, dir)
	RequireNoError(t, err)

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
		path := PathJoin(dir, f)
		AssertNoError(t, statExists(path), "expected generated file to exist: "+f)
	}
}

func TestGenerate_Good_IndexContainsTopics(t *T) {
	dir := t.TempDir()
	catalog := testCatalog()

	err := Generate(catalog, dir)
	RequireNoError(t, err)

	html := string(readFileBytes(t, PathJoin(dir, "index.html")))
	AssertContains(t, html, "Getting Started")
	AssertContains(t, html, "Configuration")
}

func TestGenerate_Good_TopicContainsRenderedMarkdown(t *T) {
	dir := t.TempDir()
	catalog := testCatalog()

	err := Generate(catalog, dir)
	RequireNoError(t, err)

	html := string(readFileBytes(t, PathJoin(dir, "topics", "getting-started.html")))
	AssertContains(t, html, "Getting Started")
	AssertContains(t, html, "<strong>guide</strong>")
}

func TestGenerate_Good_SearchIndexJSON(t *T) {
	dir := t.TempDir()
	catalog := testCatalog()

	err := Generate(catalog, dir)
	RequireNoError(t, err)

	content := readFileBytes(t, PathJoin(dir, "search-index.json"))

	var entries []searchIndexEntry
	if r := JSONUnmarshal(content, &entries); !r.OK {
		t.Fatal(r.Error())
	}
	AssertLen(t, entries, 2, "search index should contain all topics")

	// Verify fields are populated
	ids := make(map[string]bool)
	for _, e := range entries {
		ids[e.ID] = true
		AssertNotEmpty(t, e.Title)
		AssertNotEmpty(t, e.Content)
	}
	AssertTrue(t, ids["getting-started"])
	AssertTrue(t, ids["config"])
}

func TestGenerate_Good_404Exists(t *T) {
	dir := t.TempDir()
	catalog := testCatalog()

	err := Generate(catalog, dir)
	RequireNoError(t, err)

	html := string(readFileBytes(t, PathJoin(dir, "404.html")))
	AssertContains(t, html, "404")
	AssertContains(t, html, "Not Found")
}

func TestGenerate_Good_EmptyDir(t *T) {
	dir := t.TempDir()
	catalog := testCatalog()

	// Should succeed in an empty directory
	err := Generate(catalog, dir)
	AssertNoError(t, err)
}

func TestGenerate_Good_OverwriteExisting(t *T) {
	dir := t.TempDir()
	catalog := testCatalog()

	// Generate once
	err := Generate(catalog, dir)
	RequireNoError(t, err)

	// Generate again -- should overwrite without error
	err = Generate(catalog, dir)
	AssertNoError(t, err)

	// Verify files still exist and are valid
	AssertContains(t, string(readFileBytes(t, PathJoin(dir, "index.html"))), "Getting Started")
}

func TestGenerate_Good_SearchPageHasScript(t *T) {
	dir := t.TempDir()
	catalog := testCatalog()

	err := Generate(catalog, dir)
	RequireNoError(t, err)

	html := string(readFileBytes(t, PathJoin(dir, "search.html")))
	AssertContains(t, html, "<script>")
	AssertContains(t, html, "search-index.json")
}

func TestGenerate_Good_EmptyCatalog(t *T) {
	dir := t.TempDir()
	catalog := &Catalog{
		topics: make(map[string]*Topic),
		index:  newSearchIndex(),
	}

	err := Generate(catalog, dir)
	RequireNoError(t, err)

	// index.html should still exist
	AssertNoError(t, statExists(PathJoin(dir, "index.html")))

	// search-index.json should be valid empty array
	content := readFileBytes(t, PathJoin(dir, "search-index.json"))
	var entries []searchIndexEntry
	if r := JSONUnmarshal(content, &entries); !r.OK {
		t.Fatal(r.Error())
	}
	AssertEmpty(t, entries)
}
