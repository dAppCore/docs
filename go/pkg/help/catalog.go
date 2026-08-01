package help

import (
	"iter"
	"maps"
	"slices"

	core "dappco.re/go"
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
func LoadContentDir(dir string) core.Result {
	c := &Catalog{
		topics: make(map[string]*Topic),
		index:  newSearchIndex(),
	}

	err := core.PathWalkDir(dir, func(path string, d core.FsDirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if !core.HasSuffix(core.Lower(d.Name()), ".md") {
			return nil
		}

		r := core.ReadFile(path)
		if !r.OK {
			return core.E("catalog.LoadContentDir", core.Sprintf("reading %s", path), r.Value.(error))
		}
		content := r.Value.([]byte)

		topicResult := ParseTopic(path, content)
		if !topicResult.OK {
			return core.E("catalog.LoadContentDir", core.Sprintf("parsing %s", path), topicResult.Value.(error))
		}

		c.Add(topicResult.Value.(*Topic))
		return nil
	})
	if err != nil {
		return core.Fail(core.E("catalog.LoadContentDir", core.Sprintf("walking directory %s", dir), err))
	}
	return core.Ok(c)
}

// Get returns a topic by ID.
func (c *Catalog) Get(id string) core.Result {
	t, ok := c.topics[id]
	if !ok {
		return core.Fail(core.E("catalog.Get", core.Sprintf("topic not found: %s", id), nil))
	}
	return core.Ok(t)
}
