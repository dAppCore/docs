package help

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefaultCatalog_Good(t *testing.T) {
	c := DefaultCatalog()

	require.NotNil(t, c)
	require.NotNil(t, c.topics)
	require.NotNil(t, c.index)

	t.Run("contains built-in topics", func(t *testing.T) {
		topics := c.List()
		assert.GreaterOrEqual(t, len(topics), 2, "should have at least 2 default topics")
	})

	t.Run("getting-started topic exists", func(t *testing.T) {
		topic, err := c.Get("getting-started")
		require.NoError(t, err)
		assert.Equal(t, "Getting Started", topic.Title)
		assert.Contains(t, topic.Content, "Common Commands")
	})

	t.Run("config topic exists", func(t *testing.T) {
		topic, err := c.Get("config")
		require.NoError(t, err)
		assert.Equal(t, "Configuration", topic.Title)
		assert.Contains(t, topic.Content, "Environment Variables")
	})
}

func TestCatalog_Add_Good(t *testing.T) {
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

	t.Run("topic is retrievable after add", func(t *testing.T) {
		got, err := c.Get("test-topic")
		require.NoError(t, err)
		assert.Equal(t, topic, got)
	})

	t.Run("topic is searchable after add", func(t *testing.T) {
		results := c.Search("test")
		assert.NotEmpty(t, results)
	})

	t.Run("overwrite existing topic", func(t *testing.T) {
		replacement := &Topic{
			ID:      "test-topic",
			Title:   "Replaced Topic",
			Content: "Replacement content.",
		}
		c.Add(replacement)

		got, err := c.Get("test-topic")
		require.NoError(t, err)
		assert.Equal(t, "Replaced Topic", got.Title)
	})
}

func TestCatalog_List_Good(t *testing.T) {
	c := &Catalog{
		topics: make(map[string]*Topic),
		index:  newSearchIndex(),
	}

	t.Run("empty catalog returns empty list", func(t *testing.T) {
		list := c.List()
		assert.Empty(t, list)
	})

	t.Run("returns all added topics", func(t *testing.T) {
		c.Add(&Topic{ID: "alpha", Title: "Alpha"})
		c.Add(&Topic{ID: "beta", Title: "Beta"})
		c.Add(&Topic{ID: "gamma", Title: "Gamma"})

		list := c.List()
		assert.Len(t, list, 3)

		// Collect IDs (order is not guaranteed from map)
		ids := make(map[string]bool)
		for _, t := range list {
			ids[t.ID] = true
		}
		assert.True(t, ids["alpha"])
		assert.True(t, ids["beta"])
		assert.True(t, ids["gamma"])
	})
}

func TestCatalog_Search_Good(t *testing.T) {
	c := DefaultCatalog()

	t.Run("finds default topics", func(t *testing.T) {
		results := c.Search("configuration")
		assert.NotEmpty(t, results)
	})

	t.Run("empty query returns nil", func(t *testing.T) {
		results := c.Search("")
		assert.Nil(t, results)
	})

	t.Run("no match returns empty", func(t *testing.T) {
		results := c.Search("zzzyyyxxx")
		assert.Empty(t, results)
	})
}

func TestCatalog_Get_Good(t *testing.T) {
	c := &Catalog{
		topics: make(map[string]*Topic),
		index:  newSearchIndex(),
	}

	c.Add(&Topic{ID: "exists", Title: "Existing Topic"})

	t.Run("existing topic", func(t *testing.T) {
		topic, err := c.Get("exists")
		require.NoError(t, err)
		assert.Equal(t, "Existing Topic", topic.Title)
	})

	t.Run("missing topic returns error", func(t *testing.T) {
		topic, err := c.Get("does-not-exist")
		assert.Nil(t, topic)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "topic not found")
		assert.Contains(t, err.Error(), "does-not-exist")
	})
}

func TestCatalog_Search_Good_ScoreTiebreaking(t *testing.T) {
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
	require.Len(t, results, 2)

	// With equal scores, results should be sorted alphabetically by title.
	assert.Equal(t, "Alpha", results[0].Topic.Title)
	assert.Equal(t, "Zebra", results[1].Topic.Title)
	assert.Equal(t, results[0].Score, results[1].Score,
		"scores should be equal for tie-breaking to apply")
}

func TestCatalog_LoadContentDir_Good(t *testing.T) {
	dir := t.TempDir()
	os.MkdirAll(filepath.Join(dir, "cli"), 0o755)
	os.WriteFile(filepath.Join(dir, "cli", "dev-work.md"), []byte("---\ntitle: Dev Work\ntags: [cli, dev]\n---\n\n## Usage\n\ncore dev work syncs your workspace.\n"), 0o644)
	os.WriteFile(filepath.Join(dir, "cli", "setup.md"), []byte("---\ntitle: Setup\ntags: [cli]\n---\n\n## Installation\n\nRun core setup to get started.\n"), 0o644)

	catalog, err := LoadContentDir(dir)
	require.NoError(t, err)
	assert.Len(t, catalog.List(), 2)

	devWork, err := catalog.Get("dev-work")
	require.NoError(t, err)
	assert.Equal(t, "Dev Work", devWork.Title)
	assert.Contains(t, devWork.Tags, "cli")

	results := catalog.Search("workspace")
	assert.NotEmpty(t, results)
}

func TestCatalog_LoadContentDir_Good_Empty(t *testing.T) {
	dir := t.TempDir()
	catalog, err := LoadContentDir(dir)
	require.NoError(t, err)
	assert.Empty(t, catalog.List())
}

func TestCatalog_LoadContentDir_Good_SkipsNonMd(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "readme.txt"), []byte("not markdown"), 0o644)
	os.WriteFile(filepath.Join(dir, "topic.md"), []byte("---\ntitle: Topic\n---\n\nContent here.\n"), 0o644)

	catalog, err := LoadContentDir(dir)
	require.NoError(t, err)
	assert.Len(t, catalog.List(), 1)
}

func BenchmarkSearch(b *testing.B) {
	// Build a catalog with 100+ topics for benchmarking.
	c := &Catalog{
		topics: make(map[string]*Topic),
		index:  newSearchIndex(),
	}

	for i := range 150 {
		c.Add(&Topic{
			ID:      fmt.Sprintf("topic-%d", i),
			Title:   fmt.Sprintf("Topic Number %d About Various Subjects", i),
			Content: fmt.Sprintf("This is the content of topic %d. It covers installation, configuration, deployment, and testing of the system.", i),
			Tags:    []string{"generated", fmt.Sprintf("tag%d", i%10)},
			Sections: []Section{
				{
					ID:      fmt.Sprintf("section-%d-a", i),
					Title:   "Overview",
					Content: "An overview of the topic and its purpose.",
				},
				{
					ID:      fmt.Sprintf("section-%d-b", i),
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
