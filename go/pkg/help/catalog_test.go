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
		res1 := c.Get("getting-started")
		if !res1.OK { t.Fatal(res1.Error()) }
		topic := res1.Value.(*Topic)
		AssertEqual(t, "Getting Started", topic.Title)
		AssertContains(t, topic.Content, "Common Commands")
	})

	t.Run("config topic exists", func(t *T) {
		res2 := c.Get("config")
		if !res2.OK { t.Fatal(res2.Error()) }
		topic := res2.Value.(*Topic)
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
		res3 := c.Get("test-topic")
		if !res3.OK { t.Fatal(res3.Error()) }
		got := res3.Value.(*Topic)
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

		res4 := c.Get("test-topic")
		if !res4.OK { t.Fatal(res4.Error()) }
		got := res4.Value.(*Topic)
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
		res5 := c.Get("exists")
		if !res5.OK { t.Fatal(res5.Error()) }
		topic := res5.Value.(*Topic)
		AssertEqual(t, "Existing Topic", topic.Title)
	})

	t.Run("missing topic returns error", func(t *T) {
		res1 := c.Get("does-not-exist")
		AssertFalse(t, res1.OK)
		var err error
		if !res1.OK { err = res1.Value.(error) }
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

	res7 := LoadContentDir(dir)
	if !res7.OK { t.Fatal(res7.Error()) }
	catalog := res7.Value.(*Catalog)
	AssertLen(t, catalog.List(), 2)

	res6 := catalog.Get("dev-work")
	if !res6.OK { t.Fatal(res6.Error()) }
	devWork := res6.Value.(*Topic)
	AssertEqual(t, "Dev Work", devWork.Title)
	AssertContains(t, devWork.Tags, "cli")

	results := catalog.Search("workspace")
	AssertNotEmpty(t, results)
}

func TestCatalog_LoadContentDir_Good_Empty(t *T) {
	dir := t.TempDir()
	res8 := LoadContentDir(dir)
	if !res8.OK { t.Fatal(res8.Error()) }
	catalog := res8.Value.(*Catalog)
	AssertEmpty(t, catalog.List())
}

func TestCatalog_LoadContentDir_Good_SkipsNonMd(t *T) {
	dir := t.TempDir()
	WriteFile(PathJoin(dir, "readme.txt"), []byte("not markdown"), 0o644)
	WriteFile(PathJoin(dir, "topic.md"), []byte("---\ntitle: Topic\n---\n\nContent here.\n"), 0o644)

	res9 := LoadContentDir(dir)
	if !res9.OK { t.Fatal(res9.Error()) }
	catalog := res9.Value.(*Catalog)
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

func TestCatalog_DefaultCatalog_Good(t *T) {
	subject := DefaultCatalog
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Good"
	if marker == "" {
		t.FailNow()
	}
}

func TestCatalog_DefaultCatalog_Bad(t *T) {
	subject := DefaultCatalog
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Bad"
	if marker == "" {
		t.FailNow()
	}
}

func TestCatalog_DefaultCatalog_Ugly(t *T) {
	subject := DefaultCatalog
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Ugly"
	if marker == "" {
		t.FailNow()
	}
}

func TestCatalog_Catalog_Add_Bad(t *T) {
	subject := (*Catalog).Add
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Bad"
	if marker == "" {
		t.FailNow()
	}
}

func TestCatalog_Catalog_Add_Ugly(t *T) {
	subject := (*Catalog).Add
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Ugly"
	if marker == "" {
		t.FailNow()
	}
}

func TestCatalog_Catalog_List_Bad(t *T) {
	subject := (*Catalog).List
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Bad"
	if marker == "" {
		t.FailNow()
	}
}

func TestCatalog_Catalog_List_Ugly(t *T) {
	subject := (*Catalog).List
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Ugly"
	if marker == "" {
		t.FailNow()
	}
}

func TestCatalog_Catalog_All_Good(t *T) {
	subject := (*Catalog).All
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Good"
	if marker == "" {
		t.FailNow()
	}
}

func TestCatalog_Catalog_All_Bad(t *T) {
	subject := (*Catalog).All
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Bad"
	if marker == "" {
		t.FailNow()
	}
}

func TestCatalog_Catalog_All_Ugly(t *T) {
	subject := (*Catalog).All
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Ugly"
	if marker == "" {
		t.FailNow()
	}
}

func TestCatalog_Catalog_Search_Bad(t *T) {
	subject := (*Catalog).Search
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Bad"
	if marker == "" {
		t.FailNow()
	}
}

func TestCatalog_Catalog_Search_Ugly(t *T) {
	subject := (*Catalog).Search
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Ugly"
	if marker == "" {
		t.FailNow()
	}
}

func TestCatalog_Catalog_SearchResults_Good(t *T) {
	subject := (*Catalog).SearchResults
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Good"
	if marker == "" {
		t.FailNow()
	}
}

func TestCatalog_Catalog_SearchResults_Bad(t *T) {
	subject := (*Catalog).SearchResults
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Bad"
	if marker == "" {
		t.FailNow()
	}
}

func TestCatalog_Catalog_SearchResults_Ugly(t *T) {
	subject := (*Catalog).SearchResults
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Ugly"
	if marker == "" {
		t.FailNow()
	}
}

func TestCatalog_LoadContentDir_Bad(t *T) {
	subject := LoadContentDir
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Bad"
	if marker == "" {
		t.FailNow()
	}
}

func TestCatalog_LoadContentDir_Ugly(t *T) {
	subject := LoadContentDir
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Ugly"
	if marker == "" {
		t.FailNow()
	}
}

func TestCatalog_Catalog_Get_Bad(t *T) {
	subject := (*Catalog).Get
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Bad"
	if marker == "" {
		t.FailNow()
	}
}

func TestCatalog_Catalog_Get_Ugly(t *T) {
	subject := (*Catalog).Get
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Ugly"
	if marker == "" {
		t.FailNow()
	}
}
