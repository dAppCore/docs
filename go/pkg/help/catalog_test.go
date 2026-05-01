package help

import (
	. "dappco.re/go"
)

func TestDefaultCatalog_Good(t *T) {
	c := DefaultCatalog()

	AssertNotNil(t, c)
	AssertNotNil(t, c.topics)
	AssertNotNil(t, c.index)

	t.Run("contains built-in topics", func(t *T) {
		topics := c.List()
		AssertGreaterOrEqual(t, len(topics), 2, "should have at least 2 default topics")
	})

	t.Run("getting-started topic exists", func(t *T) {
		topic, err := c.Get("getting-started")
		RequireNoError(t, err)
		AssertEqual(t, "Getting Started", topic.Title)
		AssertContains(t, topic.Content, "Common Commands")
	})

	t.Run("config topic exists", func(t *T) {
		topic, err := c.Get("config")
		RequireNoError(t, err)
		AssertEqual(t, "Configuration", topic.Title)
		AssertContains(t, topic.Content, "Environment Variables")
	})
}

func TestCatalog_Add_Good(t *T) {
	c := &Catalog{
		topics: make(map[string]*Topic),
		index:  newSearchIndex(),
	}

	topic := &Topic{
		ID:      "test-topic",
		Title:   "Test Topic",
		Content: "This is a test topic for unit testing.",
		Tags:    []string{"test", "unit"},
	}

	c.Add(topic)

	t.Run("topic is retrievable after add", func(t *T) {
		got, err := c.Get("test-topic")
		RequireNoError(t, err)
		AssertEqual(t, topic, got)
	})

	t.Run("topic is searchable after add", func(t *T) {
		results := c.Search("test")
		AssertNotEmpty(t, results)
	})

	t.Run("overwrite existing topic", func(t *T) {
		replacement := &Topic{
			ID:      "test-topic",
			Title:   "Replaced Topic",
			Content: "Replacement content.",
		}
		c.Add(replacement)

		got, err := c.Get("test-topic")
		RequireNoError(t, err)
		AssertEqual(t, "Replaced Topic", got.Title)
	})
}

func TestCatalog_List_Good(t *T) {
	c := &Catalog{
		topics: make(map[string]*Topic),
		index:  newSearchIndex(),
	}

	t.Run("empty catalog returns empty list", func(t *T) {
		list := c.List()
		AssertEmpty(t, list)
	})

	t.Run("returns all added topics", func(t *T) {
		c.Add(&Topic{ID: "alpha", Title: "Alpha"})
		c.Add(&Topic{ID: "beta", Title: "Beta"})
		c.Add(&Topic{ID: "gamma", Title: "Gamma"})

		list := c.List()
		AssertLen(t, list, 3)

		// Collect IDs (order is not guaranteed from map)
		ids := make(map[string]bool)
		for _, t := range list {
			ids[t.ID] = true
		}
		AssertTrue(t, ids["alpha"])
		AssertTrue(t, ids["beta"])
		AssertTrue(t, ids["gamma"])
	})
}

func TestCatalog_Search_Good(t *T) {
	c := DefaultCatalog()

	t.Run("finds default topics", func(t *T) {
		results := c.Search("configuration")
		AssertNotEmpty(t, results)
	})

	t.Run("empty query returns nil", func(t *T) {
		results := c.Search("")
		AssertNil(t, results)
	})

	t.Run("no match returns empty", func(t *T) {
		results := c.Search("zzzyyyxxx")
		AssertEmpty(t, results)
	})
}

func TestCatalog_Get_Good(t *T) {
	c := &Catalog{
		topics: make(map[string]*Topic),
		index:  newSearchIndex(),
	}

	c.Add(&Topic{ID: "exists", Title: "Existing Topic"})

	t.Run("existing topic", func(t *T) {
		topic, err := c.Get("exists")
		RequireNoError(t, err)
		AssertEqual(t, "Existing Topic", topic.Title)
	})

	t.Run("missing topic returns error", func(t *T) {
		topic, err := c.Get("does-not-exist")
		AssertNil(t, topic)
		AssertError(t, err)
		AssertContains(t, err.Error(), "topic not found")
		AssertContains(t, err.Error(), "does-not-exist")
	})
}

func TestCatalog_Search_Good_ScoreTiebreaking(t *T) {
	// Tests the alphabetical tie-breaking in search result sorting (search.go:165).
	c := &Catalog{
		topics: make(map[string]*Topic),
		index:  newSearchIndex(),
	}

	// Add topics with identical content so they receive the same score.
	c.Add(&Topic{
		ID:      "zebra-topic",
		Title:   "Zebra",
		Content: "Unique keyword zephyr.",
	})
	c.Add(&Topic{
		ID:      "alpha-topic",
		Title:   "Alpha",
		Content: "Unique keyword zephyr.",
	})

	results := c.Search("zephyr")
	AssertLen(t, results, 2)

	// With equal scores, results should be sorted alphabetically by title.
	AssertEqual(t, "Alpha", results[0].Topic.Title)
	AssertEqual(t, "Zebra", results[1].Topic.Title)
	AssertEqual(t, results[0].Score, results[1].Score,
		"scores should be equal for tie-breaking to apply")
}

func TestCatalog_LoadContentDir_Good(t *T) {
	dir := t.TempDir()
	MkdirAll(PathJoin(dir, "cli"), 0o755)
	WriteFile(PathJoin(dir, "cli", "dev-work.md"), []byte("---\ntitle: Dev Work\ntags: [cli, dev]\n---\n\n## Usage\n\ncore dev work syncs your workspace.\n"), 0o644)
	WriteFile(PathJoin(dir, "cli", "setup.md"), []byte("---\ntitle: Setup\ntags: [cli]\n---\n\n## Installation\n\nRun core setup to get started.\n"), 0o644)

	catalog, err := LoadContentDir(dir)
	RequireNoError(t, err)
	AssertLen(t, catalog.List(), 2)

	devWork, err := catalog.Get("dev-work")
	RequireNoError(t, err)
	AssertEqual(t, "Dev Work", devWork.Title)
	AssertContains(t, devWork.Tags, "cli")

	results := catalog.Search("workspace")
	AssertNotEmpty(t, results)
}

func TestCatalog_LoadContentDir_Good_Empty(t *T) {
	dir := t.TempDir()
	catalog, err := LoadContentDir(dir)
	RequireNoError(t, err)
	AssertEmpty(t, catalog.List())
}

func TestCatalog_LoadContentDir_Good_SkipsNonMd(t *T) {
	dir := t.TempDir()
	WriteFile(PathJoin(dir, "readme.txt"), []byte("not markdown"), 0o644)
	WriteFile(PathJoin(dir, "topic.md"), []byte("---\ntitle: Topic\n---\n\nContent here.\n"), 0o644)

	catalog, err := LoadContentDir(dir)
	RequireNoError(t, err)
	AssertLen(t, catalog.List(), 1)
}

func BenchmarkSearch(b *B) {
	// Build a catalog with 100+ topics for benchmarking.
	c := &Catalog{
		topics: make(map[string]*Topic),
		index:  newSearchIndex(),
	}

	for i := range 150 {
		c.Add(&Topic{
			ID:      Sprintf("topic-%d", i),
			Title:   Sprintf("Topic Number %d About Various Subjects", i),
			Content: Sprintf("This is the content of topic %d. It covers installation, configuration, deployment, and testing of the system.", i),
			Tags:    []string{"generated", Sprintf("tag%d", i%10)},
			Sections: []Section{
				{
					ID:      Sprintf("section-%d-a", i),
					Title:   "Overview",
					Content: "An overview of the topic and its purpose.",
				},
				{
					ID:      Sprintf("section-%d-b", i),
					Title:   "Details",
					Content: "Detailed information about the topic including examples and usage.",
				},
			},
		})
	}

	b.ResetTimer()
	b.ReportAllocs()
	for range b.N {
		c.Search("installation configuration")
	}
}
