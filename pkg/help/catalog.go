package help

import (
	"fmt"
	"io/fs"
	"iter"
	"maps"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

// Catalog manages help topics.
type Catalog struct {
	topics map[string]*Topic
	index  *searchIndex
}

// DefaultCatalog returns a catalog with built-in topics.
func DefaultCatalog() *Catalog {
	c := &Catalog{
		topics: make(map[string]*Topic),
		index:  newSearchIndex(),
	}

	// Add default topics
	c.Add(&Topic{
		ID:    "getting-started",
		Title: "Getting Started",
		Content: `# Getting Started

Welcome to Core! This CLI tool helps you manage development workflows.

## Common Commands

- core dev: Development workflows
- core setup: Setup repository
- core doctor: Check environment health
- core test: Run tests

## Next Steps

Run 'core help <topic>' to learn more about a specific topic.
`,
	})
	c.Add(&Topic{
		ID:    "config",
		Title: "Configuration",
		Content: `# Configuration

Core is configured via environment variables and config files.

## Environment Variables

- CORE_DEBUG: Enable debug logging
- GITHUB_TOKEN: GitHub API token

## Config Files

Config is stored in ~/.core/config.yaml
`,
	})
	return c
}

// Add adds a topic to the catalog.
func (c *Catalog) Add(t *Topic) {
	c.topics[t.ID] = t
	c.index.Add(t)
}

// List returns all topics.
func (c *Catalog) List() []*Topic {
	var list []*Topic
	for _, t := range c.topics {
		list = append(list, t)
	}
	return list
}

// All returns an iterator for all topics.
func (c *Catalog) All() iter.Seq[*Topic] {
	return maps.Values(c.topics)
}

// Search searches for topics.
func (c *Catalog) Search(query string) []*SearchResult {
	return c.index.Search(query)
}

// SearchResults returns an iterator for search results.
func (c *Catalog) SearchResults(query string) iter.Seq[*SearchResult] {
	return slices.Values(c.Search(query))
}

// LoadContentDir recursively loads all .md files from a directory into a Catalog.
func LoadContentDir(dir string) (*Catalog, error) {
	c := &Catalog{
		topics: make(map[string]*Topic),
		index:  newSearchIndex(),
	}

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if !strings.HasSuffix(strings.ToLower(d.Name()), ".md") {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("reading %s: %w", path, err)
		}

		topic, err := ParseTopic(path, content)
		if err != nil {
			return fmt.Errorf("parsing %s: %w", path, err)
		}

		c.Add(topic)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("walking directory %s: %w", dir, err)
	}

	return c, nil
}

// Get returns a topic by ID.
func (c *Catalog) Get(id string) (*Topic, error) {
	t, ok := c.topics[id]
	if !ok {
		return nil, fmt.Errorf("topic not found: %s", id)
	}
	return t, nil
}
