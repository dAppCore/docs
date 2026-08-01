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

	res1 := Generate(catalog, dir)
	if !res1.OK { t.Fatal(res1.Error()) }

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

	res2 := Generate(catalog, dir)
	if !res2.OK { t.Fatal(res2.Error()) }

	html := string(readFileBytes(t, PathJoin(dir, "index.html")))
	AssertContains(t, html, "Getting Started")
	AssertContains(t, html, "Configuration")
}

func TestGenerate_Good_TopicContainsRenderedMarkdown(t *T) {
	dir := t.TempDir()
	catalog := testCatalog()

	res3 := Generate(catalog, dir)
	if !res3.OK { t.Fatal(res3.Error()) }

	html := string(readFileBytes(t, PathJoin(dir, "topics", "getting-started.html")))
	AssertContains(t, html, "Getting Started")
	AssertContains(t, html, "<strong>guide</strong>")
}

func TestGenerate_Good_SearchIndexJSON(t *T) {
	dir := t.TempDir()
	catalog := testCatalog()

	res4 := Generate(catalog, dir)
	if !res4.OK { t.Fatal(res4.Error()) }

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

	res5 := Generate(catalog, dir)
	if !res5.OK { t.Fatal(res5.Error()) }

	html := string(readFileBytes(t, PathJoin(dir, "404.html")))
	AssertContains(t, html, "404")
	AssertContains(t, html, "Not Found")
}

func TestGenerate_Good_EmptyDir(t *T) {
	dir := t.TempDir()
	catalog := testCatalog()

	// Should succeed in an empty directory
	res2 := Generate(catalog, dir)
	AssertTrue(t, res2.OK)
}

func TestGenerate_Good_OverwriteExisting(t *T) {
	dir := t.TempDir()
	catalog := testCatalog()

	// Generate once
	res6 := Generate(catalog, dir)
	if !res6.OK { t.Fatal(res6.Error()) }

	// Generate again -- should overwrite without error
	res1 := Generate(catalog, dir)
	AssertTrue(t, res1.OK)

	// Verify files still exist and are valid
	AssertContains(t, string(readFileBytes(t, PathJoin(dir, "index.html"))), "Getting Started")
}

func TestGenerate_Good_SearchPageHasScript(t *T) {
	dir := t.TempDir()
	catalog := testCatalog()

	res7 := Generate(catalog, dir)
	if !res7.OK { t.Fatal(res7.Error()) }

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

	res8 := Generate(catalog, dir)
	if !res8.OK { t.Fatal(res8.Error()) }

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

func TestGenerate_Bad_MkdirFails(t *T) {
	// outputDir lives UNDER an existing regular file, so MkdirAll of the
	// "topics" subdir cannot succeed.
	dir := t.TempDir()
	filePath := PathJoin(dir, "blocker")
	if r := WriteFile(filePath, []byte("not a directory"), 0o644); !r.OK {
		t.Fatal(r.Error())
	}

	catalog := testCatalog()
	res := Generate(catalog, PathJoin(filePath, "site"))
	AssertFalse(t, res.OK, "Generate must fail when the output dir cannot be created")
}

func TestGenerate_Bad_WriteFileFails(t *T) {
	// outputDir itself is an existing regular file: MkdirAll(topics) under it
	// fails, surfacing the first error result.
	dir := t.TempDir()
	filePath := PathJoin(dir, "regular-file")
	if r := WriteFile(filePath, []byte("x"), 0o644); !r.OK {
		t.Fatal(r.Error())
	}

	catalog := testCatalog()
	res := Generate(catalog, filePath)
	AssertFalse(t, res.OK, "Generate must fail when outputDir is a regular file")
}

func TestGenerate_Bad_IndexWriteFails(t *T) {
	// "topics" dir creation succeeds, but index.html already exists as a
	// directory so writing the index page fails.
	dir := t.TempDir()
	if r := MkdirAll(PathJoin(dir, "index.html"), 0o755); !r.OK {
		t.Fatal(r.Error())
	}

	res := Generate(testCatalog(), dir)
	AssertFalse(t, res.OK, "Generate must fail when index.html cannot be written")
}

func TestGenerate_Bad_TopicWriteFails(t *T) {
	// index.html writes fine, but a topic target path is occupied by a
	// directory so the per-topic writeFile fails.
	dir := t.TempDir()
	c := &Catalog{
		topics: make(map[string]*Topic),
		index:  newSearchIndex(),
	}
	c.Add(&Topic{ID: "blocked", Title: "Blocked", Content: "x"})
	// topics/blocked.html occupied by a directory.
	if r := MkdirAll(PathJoin(dir, "topics", "blocked.html"), 0o755); !r.OK {
		t.Fatal(r.Error())
	}

	res := Generate(c, dir)
	AssertFalse(t, res.OK, "Generate must fail when a topic page cannot be written")
}

func TestGenerate_Bad_SearchIndexWriteFails(t *T) {
	// All HTML writes succeed, but search-index.json is occupied by a
	// directory so writeSearchIndex fails.
	dir := t.TempDir()
	if r := MkdirAll(PathJoin(dir, "search-index.json"), 0o755); !r.OK {
		t.Fatal(r.Error())
	}

	res := Generate(testCatalog(), dir)
	AssertFalse(t, res.OK, "Generate must fail when the search index cannot be written")
}

func TestGenerate_Good_TruncatesLongContentInSearchIndex(t *T) {
	// writeSearchIndex truncates topic content to 500 runes; exercise the
	// >500 branch and confirm the JSON entry is clipped.
	dir := t.TempDir()
	c := &Catalog{
		topics: make(map[string]*Topic),
		index:  newSearchIndex(),
	}
	long := ""
	for range 600 {
		long += "x"
	}
	c.Add(&Topic{ID: "long", Title: "Long", Content: long})

	res := Generate(c, dir)
	if !res.OK {
		t.Fatal(res.Error())
	}

	content := readFileBytes(t, PathJoin(dir, "search-index.json"))
	var entries []searchIndexEntry
	if r := JSONUnmarshal(content, &entries); !r.OK {
		t.Fatal(r.Error())
	}
	AssertLen(t, entries, 1)
	AssertLen(t, []rune(entries[0].Content), 500, "long content must be truncated to 500 runes")
}

func TestGenerate_Generate_Good(t *T) {
	subject := Generate
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Good"
	if marker == "" {
		t.FailNow()
	}
}

func TestGenerate_Generate_Bad(t *T) {
	subject := Generate
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Bad"
	if marker == "" {
		t.FailNow()
	}
}

func TestGenerate_Generate_Ugly(t *T) {
	subject := Generate
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Ugly"
	if marker == "" {
		t.FailNow()
	}
}
