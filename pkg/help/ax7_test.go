package help

import (
	. "dappco.re/go"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"slices"
)

func ax7EmptyCatalog() *Catalog {
	return &Catalog{
		topics: make(map[string]*Topic),
		index:  newSearchIndex(),
	}
}

func ax7Topic() *Topic {
	return &Topic{
		ID:      "agent-guide",
		Title:   "Agent Guide",
		Content: "# Agent Guide\n\nRun the **agent** from the terminal.\n",
		Tags:    []string{"cli", "agent"},
		Sections: []Section{
			{ID: "agent-guide", Title: "Agent Guide", Level: 1, Content: "Run the agent."},
		},
		Related: []string{"config"},
	}
}

func ax7Catalog() *Catalog {
	c := ax7EmptyCatalog()
	c.Add(ax7Topic())
	c.Add(&Topic{ID: "config", Title: "Configuration", Content: "Configure the agent.", Tags: []string{"setup"}})
	return c
}

func TestAX7_RenderMarkdown_Good(t *T) {
	html, err := RenderMarkdown("# Agent\n\nRun the **agent**.")
	RequireNoError(t, err)
	AssertContains(t, html, "<h1>Agent</h1>")
	AssertContains(t, html, "<strong>agent</strong>")
}

func TestAX7_RenderMarkdown_Bad(t *T) {
	html, err := RenderMarkdown("")
	RequireNoError(t, err)
	AssertEqual(t, "", html)
	AssertNotContains(t, html, "<h1>")
}

func TestAX7_RenderMarkdown_Ugly(t *T) {
	html, err := RenderMarkdown(`<div class="raw">agent</div>`)
	RequireNoError(t, err)
	AssertContains(t, html, `<div class="raw">agent</div>`)
	AssertNotContains(t, html, "&lt;div")
}

func TestAX7_Generate_Good(t *T) {
	dir := t.TempDir()
	err := Generate(ax7Catalog(), dir)
	RequireNoError(t, err)
	_, err = os.Stat(filepath.Join(dir, "index.html"))
	AssertNoError(t, err)
}

func TestAX7_Generate_Bad(t *T) {
	dir := t.TempDir()
	outputFile := filepath.Join(dir, "site")
	RequireNoError(t, os.WriteFile(outputFile, []byte("not a directory"), 0o644))
	err := Generate(ax7Catalog(), outputFile)
	AssertError(t, err)
}

func TestAX7_Generate_Ugly(t *T) {
	dir := t.TempDir()
	err := Generate(ax7EmptyCatalog(), dir)
	RequireNoError(t, err)
	data, err := os.ReadFile(filepath.Join(dir, "search-index.json"))
	RequireNoError(t, err)
	AssertContains(t, string(data), "[]")
}

func TestAX7_NewServer_Good(t *T) {
	c := ax7Catalog()
	srv := NewServer(c, ":8080")
	AssertNotNil(t, srv)
	AssertEqual(t, c, srv.catalog)
	AssertNotNil(t, srv.mux)
}

func TestAX7_NewServer_Bad(t *T) {
	srv := NewServer(nil, "")
	AssertNotNil(t, srv)
	AssertNil(t, srv.catalog)
	AssertEqual(t, "", srv.addr)
}

func TestAX7_NewServer_Ugly(t *T) {
	srv := NewServer(ax7Catalog(), "127.0.0.1:-1")
	AssertNotNil(t, srv)
	AssertEqual(t, "127.0.0.1:-1", srv.addr)
	AssertNotNil(t, srv.mux)
}

func TestAX7_Server_ServeHTTP_Good(t *T) {
	srv := NewServer(ax7Catalog(), ":0")
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	srv.ServeHTTP(rec, req)
	AssertEqual(t, http.StatusOK, rec.Code)
}

func TestAX7_Server_ServeHTTP_Bad(t *T) {
	srv := NewServer(ax7Catalog(), ":0")
	req := httptest.NewRequest(http.MethodGet, "/topics/missing", nil)
	rec := httptest.NewRecorder()
	srv.ServeHTTP(rec, req)
	AssertEqual(t, http.StatusNotFound, rec.Code)
}

func TestAX7_Server_ServeHTTP_Ugly(t *T) {
	srv := NewServer(ax7Catalog(), ":0")
	req := httptest.NewRequest(http.MethodGet, "/api/search", nil)
	rec := httptest.NewRecorder()
	srv.ServeHTTP(rec, req)
	AssertEqual(t, http.StatusBadRequest, rec.Code)
}

func TestAX7_Server_ListenAndServe_Good(t *T) {
	srv := NewServer(ax7Catalog(), "127.0.0.1:-1")
	err := srv.ListenAndServe()
	AssertError(t, err)
	AssertNotNil(t, srv.mux)
}

func TestAX7_Server_ListenAndServe_Bad(t *T) {
	srv := NewServer(ax7Catalog(), "bad addr")
	err := srv.ListenAndServe()
	AssertError(t, err)
	AssertEqual(t, "bad addr", srv.addr)
}

func TestAX7_Server_ListenAndServe_Ugly(t *T) {
	var srv *Server
	AssertPanics(t, func() {
		_ = srv.ListenAndServe()
	})
	AssertNil(t, srv)
}

func TestAX7_ParseHelpText_Good(t *T) {
	topic := ParseHelpText("agent run", "Usage:\n  agent run [flags]\n\nFlags:\n  --fast")
	AssertEqual(t, "agent-run", topic.ID)
	AssertContains(t, topic.Content, "## Usage")
	AssertContains(t, topic.Content, "--fast")
}

func TestAX7_ParseHelpText_Bad(t *T) {
	topic := ParseHelpText("empty", "")
	AssertEqual(t, "empty", topic.ID)
	AssertEqual(t, "", topic.Content)
	AssertEqual(t, []string{"cli", "empty"}, topic.Tags)
}

func TestAX7_ParseHelpText_Ugly(t *T) {
	topic := ParseHelpText("cli", "Root command.\nSee also: agent run, config")
	AssertEqual(t, []string{"cli"}, topic.Tags)
	AssertEqual(t, []string{"agent-run", "config"}, topic.Related)
	AssertNotContains(t, topic.Content, "See also")
}

func TestAX7_IngestCLIHelp_Good(t *T) {
	c := IngestCLIHelp(map[string]string{"agent run": "Run agent.", "agent stop": "Stop agent."})
	topics := c.List()
	AssertLen(t, topics, 2)
	AssertNotEmpty(t, c.Search("agent"))
}

func TestAX7_IngestCLIHelp_Bad(t *T) {
	c := IngestCLIHelp(map[string]string{})
	topics := c.List()
	AssertEmpty(t, topics)
	AssertNil(t, c.Search(""))
}

func TestAX7_IngestCLIHelp_Ugly(t *T) {
	c := IngestCLIHelp(map[string]string{"cli": "CLI root command."})
	topic, err := c.Get("cli")
	RequireNoError(t, err)
	AssertEqual(t, []string{"cli"}, topic.Tags)
	AssertEqual(t, "Cli", topic.Title)
}

func TestAX7_DefaultCatalog_Good(t *T) {
	c := DefaultCatalog()
	topics := c.List()
	AssertGreaterOrEqual(t, len(topics), 2)
	AssertNotEmpty(t, c.Search("configuration"))
}

func TestAX7_DefaultCatalog_Bad(t *T) {
	c := DefaultCatalog()
	topic, err := c.Get("missing")
	AssertNil(t, topic)
	AssertError(t, err, "topic not found")
}

func TestAX7_DefaultCatalog_Ugly(t *T) {
	c := DefaultCatalog()
	list := c.List()
	list[0] = nil
	topic, err := c.Get("getting-started")
	RequireNoError(t, err)
	AssertEqual(t, "Getting Started", topic.Title)
}

func TestAX7_Catalog_Add_Good(t *T) {
	c := ax7EmptyCatalog()
	topic := ax7Topic()
	c.Add(topic)
	got, err := c.Get("agent-guide")
	RequireNoError(t, err)
	AssertEqual(t, topic, got)
}

func TestAX7_Catalog_Add_Bad(t *T) {
	c := ax7EmptyCatalog()
	c.Add(&Topic{ID: "agent", Title: "Old"})
	c.Add(&Topic{ID: "agent", Title: "New"})
	got, err := c.Get("agent")
	RequireNoError(t, err)
	AssertEqual(t, "New", got.Title)
}

func TestAX7_Catalog_Add_Ugly(t *T) {
	c := ax7EmptyCatalog()
	c.Add(&Topic{ID: "", Title: "Empty ID", Content: "edge"})
	got, err := c.Get("")
	RequireNoError(t, err)
	AssertEqual(t, "Empty ID", got.Title)
}

func TestAX7_Catalog_List_Good(t *T) {
	c := ax7Catalog()
	topics := c.List()
	AssertLen(t, topics, 2)
	AssertNotNil(t, topics[0])
}

func TestAX7_Catalog_List_Bad(t *T) {
	c := ax7EmptyCatalog()
	topics := c.List()
	AssertEmpty(t, topics)
	AssertNotNil(t, c.index)
}

func TestAX7_Catalog_List_Ugly(t *T) {
	c := ax7EmptyCatalog()
	c.Add(&Topic{ID: "dup", Title: "One"})
	c.Add(&Topic{ID: "dup", Title: "Two"})
	topics := c.List()
	AssertLen(t, topics, 1)
	AssertEqual(t, "Two", topics[0].Title)
}

func TestAX7_Catalog_All_Good(t *T) {
	c := ax7Catalog()
	topics := slices.Collect(c.All())
	AssertLen(t, topics, 2)
	AssertNotNil(t, topics[0])
}

func TestAX7_Catalog_All_Bad(t *T) {
	c := ax7EmptyCatalog()
	topics := slices.Collect(c.All())
	AssertEmpty(t, topics)
	AssertNotNil(t, c.topics)
}

func TestAX7_Catalog_All_Ugly(t *T) {
	c := ax7EmptyCatalog()
	c.Add(&Topic{ID: "", Title: "Empty"})
	topics := slices.Collect(c.All())
	AssertLen(t, topics, 1)
	AssertEqual(t, "", topics[0].ID)
}

func TestAX7_Catalog_Search_Good(t *T) {
	c := ax7Catalog()
	results := c.Search("agent")
	AssertNotEmpty(t, results)
	AssertEqual(t, "agent-guide", results[0].Topic.ID)
}

func TestAX7_Catalog_Search_Bad(t *T) {
	c := ax7Catalog()
	results := c.Search("")
	AssertNil(t, results)
	AssertNotNil(t, c.index)
}

func TestAX7_Catalog_Search_Ugly(t *T) {
	c := ax7Catalog()
	results := c.Search("!@#$")
	AssertEmpty(t, results)
	AssertNotNil(t, c.topics)
}

func TestAX7_Catalog_SearchResults_Good(t *T) {
	c := ax7Catalog()
	results := slices.Collect(c.SearchResults("agent"))
	AssertNotEmpty(t, results)
	AssertEqual(t, "agent-guide", results[0].Topic.ID)
}

func TestAX7_Catalog_SearchResults_Bad(t *T) {
	c := ax7Catalog()
	results := slices.Collect(c.SearchResults(""))
	AssertEmpty(t, results)
	AssertNotNil(t, c.index)
}

func TestAX7_Catalog_SearchResults_Ugly(t *T) {
	c := ax7Catalog()
	results := slices.Collect(c.SearchResults("zzzz"))
	AssertEmpty(t, results)
	AssertNotNil(t, c.topics)
}

func TestAX7_LoadContentDir_Good(t *T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "agent.md")
	RequireNoError(t, os.WriteFile(path, []byte("---\ntitle: Agent\n---\n\n# Agent\n"), 0o644))
	c, err := LoadContentDir(dir)
	RequireNoError(t, err)
	AssertLen(t, c.List(), 1)
}

func TestAX7_LoadContentDir_Bad(t *T) {
	c, err := LoadContentDir(filepath.Join(t.TempDir(), "missing"))
	AssertNil(t, c)
	AssertError(t, err)
	AssertContains(t, err.Error(), "walking directory")
}

func TestAX7_LoadContentDir_Ugly(t *T) {
	dir := t.TempDir()
	RequireNoError(t, os.WriteFile(filepath.Join(dir, "skip.txt"), []byte("# Skip\n"), 0o644))
	RequireNoError(t, os.WriteFile(filepath.Join(dir, "KEEP.MD"), []byte("# Keep\n"), 0o644))
	c, err := LoadContentDir(dir)
	RequireNoError(t, err)
	AssertLen(t, c.List(), 1)
}

func TestAX7_Catalog_Get_Good(t *T) {
	c := ax7Catalog()
	topic, err := c.Get("agent-guide")
	RequireNoError(t, err)
	AssertEqual(t, "Agent Guide", topic.Title)
}

func TestAX7_Catalog_Get_Bad(t *T) {
	c := ax7Catalog()
	topic, err := c.Get("missing")
	AssertNil(t, topic)
	AssertError(t, err)
	AssertContains(t, err.Error(), "missing")
}

func TestAX7_Catalog_Get_Ugly(t *T) {
	c := ax7EmptyCatalog()
	c.Add(&Topic{ID: "", Title: "Empty ID"})
	topic, err := c.Get("")
	RequireNoError(t, err)
	AssertEqual(t, "Empty ID", topic.Title)
}

func TestAX7_ParseTopic_Good(t *T) {
	topic, err := ParseTopic("agent.md", []byte("---\ntitle: Agent\n---\n\n# Agent\n\nBody"))
	RequireNoError(t, err)
	AssertEqual(t, "Agent", topic.Title)
	AssertEqual(t, "agent", topic.ID)
}

func TestAX7_ParseTopic_Bad(t *T) {
	topic, err := ParseTopic("bad.md", []byte("---\n: bad\n---\n# Bad\n"))
	RequireNoError(t, err)
	AssertEqual(t, "Bad", topic.Title)
	AssertContains(t, topic.Content, "---")
}

func TestAX7_ParseTopic_Ugly(t *T) {
	topic, err := ParseTopic("empty.md", nil)
	RequireNoError(t, err)
	AssertEqual(t, "empty", topic.ID)
	AssertEmpty(t, topic.Sections)
}

func TestAX7_ExtractFrontmatter_Good(t *T) {
	fm, body := ExtractFrontmatter("---\ntitle: Agent\norder: 2\n---\n# Body")
	RequireNotEmpty(t, body)
	AssertNotNil(t, fm)
	AssertEqual(t, "Agent", fm.Title)
}

func TestAX7_ExtractFrontmatter_Bad(t *T) {
	content := "---\n: bad\n---\n# Body"
	fm, body := ExtractFrontmatter(content)
	AssertNil(t, fm)
	AssertEqual(t, content, body)
}

func TestAX7_ExtractFrontmatter_Ugly(t *T) {
	fm, body := ExtractFrontmatter("---\r\n\r\n---\r\n# Body")
	AssertNotNil(t, fm)
	AssertEqual(t, "", fm.Title)
	AssertContains(t, body, "# Body")
}

func TestAX7_ExtractSections_Good(t *T) {
	sections := ExtractSections("# Agent\n\nIntro\n## Run\n\nSteps")
	AssertLen(t, sections, 2)
	AssertEqual(t, "agent", sections[0].ID)
	AssertContains(t, sections[1].Content, "Steps")
}

func TestAX7_ExtractSections_Bad(t *T) {
	sections := ExtractSections("No markdown heading here.")
	AssertEmpty(t, sections)
	AssertLen(t, sections, 0)
	AssertFalse(t, len(sections) > 0)
}

func TestAX7_ExtractSections_Ugly(t *T) {
	sections := ExtractSections("# One\n## Two\n### Three")
	AssertLen(t, sections, 3)
	AssertEqual(t, "", sections[0].Content)
	AssertEqual(t, "", sections[1].Content)
}

func TestAX7_AllSections_Good(t *T) {
	sections := slices.Collect(AllSections("# Agent\n\nIntro\n## Run\n\nSteps"))
	AssertLen(t, sections, 2)
	AssertEqual(t, "Agent", sections[0].Title)
	AssertEqual(t, "Run", sections[1].Title)
}

func TestAX7_AllSections_Bad(t *T) {
	sections := slices.Collect(AllSections("plain text"))
	AssertEmpty(t, sections)
	AssertLen(t, sections, 0)
	AssertFalse(t, len(sections) > 0)
}

func TestAX7_AllSections_Ugly(t *T) {
	count := 0
	for section := range AllSections("# One\n## Two") {
		AssertEqual(t, "One", section.Title)
		count++
		break
	}
	AssertEqual(t, 1, count)
}

func TestAX7_GenerateID_Good(t *T) {
	id := GenerateID("Agent Guide")
	AssertEqual(t, "agent-guide", id)
	AssertNotContains(t, id, " ")
	AssertContains(t, id, "-")
}

func TestAX7_GenerateID_Bad(t *T) {
	id := GenerateID("!@#$%^&*()")
	AssertEqual(t, "", id)
	AssertEmpty(t, id)
	AssertNotContains(t, id, "-")
}

func TestAX7_GenerateID_Ugly(t *T) {
	id := GenerateID("日本語 Agent 🚀")
	AssertContains(t, id, "日本語")
	AssertContains(t, id, "agent")
	AssertNotContains(t, id, "🚀")
}

func TestAX7_RenderIndexPage_Good(t *T) {
	html := RenderIndexPage(ax7Catalog().List())
	AssertContains(t, html, `role="banner"`)
	AssertContains(t, html, "Agent Guide")
	AssertContains(t, html, `role="main"`)
}

func TestAX7_RenderIndexPage_Bad(t *T) {
	html := RenderIndexPage(nil)
	AssertContains(t, html, "No topics available")
	AssertContains(t, html, `0 topics`)
	AssertContains(t, html, `role="contentinfo"`)
}

func TestAX7_RenderIndexPage_Ugly(t *T) {
	html := RenderIndexPage([]*Topic{{ID: "x", Title: `<script>alert("x")</script>`, Content: "safe"}})
	AssertNotContains(t, html, `<script>alert`)
	AssertContains(t, html, "&lt;script&gt;")
	AssertContains(t, html, "safe")
}

func TestAX7_RenderTopicPage_Good(t *T) {
	html := RenderTopicPage(ax7Topic(), ax7Catalog().List())
	AssertContains(t, html, `role="complementary"`)
	AssertContains(t, html, "<strong>")
	AssertContains(t, html, "Agent Guide")
}

func TestAX7_RenderTopicPage_Bad(t *T) {
	html := RenderTopicPage(&Topic{ID: "solo", Title: "Solo", Content: "Plain content."}, nil)
	AssertContains(t, html, `role="main"`)
	AssertNotContains(t, html, `role="complementary"`)
	AssertContains(t, html, "Plain content")
}

func TestAX7_RenderTopicPage_Ugly(t *T) {
	html := RenderTopicPage(&Topic{ID: "edge", Title: `<script>`, Content: "Edge content."}, nil)
	AssertNotContains(t, html, `<script>`)
	AssertContains(t, html, "Edge content")
	AssertContains(t, html, `role="banner"`)
}

func TestAX7_RenderSearchPage_Good(t *T) {
	result := &SearchResult{Topic: ax7Topic(), Score: 12.5, Snippet: "Run the **agent**."}
	html := RenderSearchPage("agent", []*SearchResult{result})
	AssertContains(t, html, "Agent Guide")
	AssertContains(t, html, "12.5")
}

func TestAX7_RenderSearchPage_Bad(t *T) {
	html := RenderSearchPage("missing", nil)
	AssertContains(t, html, "No results")
	AssertContains(t, html, "missing")
	AssertContains(t, html, `role="main"`)
}

func TestAX7_RenderSearchPage_Ugly(t *T) {
	html := RenderSearchPage(`<agent>`, nil)
	AssertNotContains(t, html, `<agent>`)
	AssertContains(t, html, "&lt;agent&gt;")
	AssertContains(t, html, `role="banner"`)
}

func TestAX7_Render404Page_Good(t *T) {
	html := Render404Page()
	AssertContains(t, html, "404")
	AssertContains(t, html, "Not Found")
	AssertContains(t, html, `role="main"`)
}

func TestAX7_Render404Page_Bad(t *T) {
	html := Render404Page()
	AssertContains(t, html, "Browse all topics")
	AssertContains(t, html, `role="contentinfo"`)
	AssertNotContains(t, html, "No results")
}

func TestAX7_Render404Page_Ugly(t *T) {
	html := Render404Page()
	AssertContains(t, html, "<!DOCTYPE html>")
	AssertContains(t, html, `role="banner"`)
	AssertContains(t, html, `href="/"`)
}

func TestAX7_Index_Add_Good(t *T) {
	idx := newSearchIndex()
	idx.Add(ax7Topic())
	AssertNotNil(t, idx.topics["agent-guide"])
	AssertContains(t, idx.index["agent"], "agent-guide")
}

func TestAX7_Index_Add_Bad(t *T) {
	idx := newSearchIndex()
	idx.Add(&Topic{ID: "repeat", Title: "Repeat Repeat"})
	AssertLen(t, idx.index["repeat"], 1)
	AssertEqual(t, "Repeat Repeat", idx.topics["repeat"].Title)
}

func TestAX7_Index_Add_Ugly(t *T) {
	idx := newSearchIndex()
	idx.Add(&Topic{ID: "", Title: "Empty", Content: "edge"})
	AssertNotNil(t, idx.topics[""])
	AssertContains(t, idx.index["empty"], "")
}

func TestAX7_Index_Search_Good(t *T) {
	idx := newSearchIndex()
	idx.Add(ax7Topic())
	results := idx.Search("agent")
	AssertNotEmpty(t, results)
	AssertEqual(t, "agent-guide", results[0].Topic.ID)
}

func TestAX7_Index_Search_Bad(t *T) {
	idx := newSearchIndex()
	results := idx.Search("")
	AssertNil(t, results)
	AssertNotNil(t, idx.index)
}

func TestAX7_Index_Search_Ugly(t *T) {
	idx := newSearchIndex()
	idx.topics["ghost"] = nil
	idx.index["ghost"] = []string{"ghost"}
	results := idx.Search("ghost")
	AssertEmpty(t, results)
}

func TestAX7_Tokens_Good(t *T) {
	tokens := slices.Collect(Tokens("Hello, agent world!"))
	AssertEqual(t, []string{"hello", "agent", "world"}, tokens)
	AssertContains(t, tokens, "agent")
	AssertNotContains(t, tokens, "!")
}

func TestAX7_Tokens_Bad(t *T) {
	tokens := slices.Collect(Tokens("a b c"))
	AssertEmpty(t, tokens)
	AssertLen(t, tokens, 0)
	AssertFalse(t, len(tokens) > 0)
}

func TestAX7_Tokens_Ugly(t *T) {
	tokens := slices.Collect(Tokens("running"))
	AssertContains(t, tokens, "running")
	AssertContains(t, tokens, "runn")
	AssertLen(t, tokens, 2)
}
