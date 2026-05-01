// SPDX-Licence-Identifier: EUPL-1.2
package help

import (
	. "dappco.re/go"
	"io"
	"net/http"
	"net/http/httptest"
)

// testServer creates a test catalog with topics and returns an httptest.Server.
func testServer(t *T) *httptest.Server {
	t.Helper()
	c := &Catalog{
		topics: make(map[string]*Topic),
		index:  newSearchIndex(),
	}
	c.Add(&Topic{
		ID:      "getting-started",
		Title:   "Getting Started",
		Content: "# Getting Started\n\nWelcome to the **guide**.\n\n## Installation\n\nInstall the tool.\n",
		Tags:    []string{"intro", "setup"},
		Sections: []Section{
			{ID: "getting-started", Title: "Getting Started", Level: 1},
			{ID: "installation", Title: "Installation", Level: 2, Content: "Install the tool."},
		},
		Related: []string{"config"},
	})
	c.Add(&Topic{
		ID:      "config",
		Title:   "Configuration",
		Content: "# Configuration\n\nConfigure your environment.\n",
		Tags:    []string{"setup"},
	})

	srv := NewServer(c, ":0")
	return httptest.NewServer(srv)
}

func TestServer_HandleIndex_Good(t *T) {
	ts := testServer(t)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/")
	RequireNoError(t, err)
	defer resp.Body.Close()

	AssertEqual(t, http.StatusOK, resp.StatusCode)
	AssertContains(t, resp.Header.Get("Content-Type"), "text/html")
	AssertEqual(t, "nosniff", resp.Header.Get("X-Content-Type-Options"))

	var buf [64 * 1024]byte
	n, _ := resp.Body.Read(buf[:])
	body := string(buf[:n])
	AssertContains(t, body, "Getting Started")
	AssertContains(t, body, "Configuration")
}

func TestServer_HandleTopic_Good(t *T) {
	ts := testServer(t)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/topics/getting-started")
	RequireNoError(t, err)
	defer resp.Body.Close()

	AssertEqual(t, http.StatusOK, resp.StatusCode)
	AssertContains(t, resp.Header.Get("Content-Type"), "text/html")

	var buf [64 * 1024]byte
	n, _ := resp.Body.Read(buf[:])
	body := string(buf[:n])
	AssertContains(t, body, "Getting Started")
	AssertContains(t, body, "<strong>guide</strong>")
}

func TestServer_HandleTopic_Bad_NotFound(t *T) {
	ts := testServer(t)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/topics/nonexistent")
	RequireNoError(t, err)
	defer resp.Body.Close()

	AssertEqual(t, http.StatusNotFound, resp.StatusCode)
	AssertContains(t, resp.Header.Get("Content-Type"), "text/html")
}

func TestServer_HandleSearch_Good(t *T) {
	ts := testServer(t)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/search?q=install")
	RequireNoError(t, err)
	defer resp.Body.Close()

	AssertEqual(t, http.StatusOK, resp.StatusCode)
	AssertContains(t, resp.Header.Get("Content-Type"), "text/html")

	var buf [64 * 1024]byte
	n, _ := resp.Body.Read(buf[:])
	body := string(buf[:n])
	AssertContains(t, body, "install")
}

func TestServer_HandleSearch_Bad_NoQuery(t *T) {
	ts := testServer(t)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/search")
	RequireNoError(t, err)
	defer resp.Body.Close()

	AssertEqual(t, http.StatusBadRequest, resp.StatusCode)
}

func TestServer_HandleAPITopics_Good(t *T) {
	ts := testServer(t)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/api/topics")
	RequireNoError(t, err)
	defer resp.Body.Close()

	AssertEqual(t, http.StatusOK, resp.StatusCode)
	AssertContains(t, resp.Header.Get("Content-Type"), "application/json")
	AssertEqual(t, "nosniff", resp.Header.Get("X-Content-Type-Options"))

	var topics []Topic
	body, err := io.ReadAll(resp.Body)
	RequireNoError(t, err)
	if r := JSONUnmarshal(body, &topics); !r.OK {
		t.Fatal(r.Error())
	}
	AssertLen(t, topics, 2)
}

func TestServer_HandleAPITopic_Good(t *T) {
	ts := testServer(t)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/api/topics/getting-started")
	RequireNoError(t, err)
	defer resp.Body.Close()

	AssertEqual(t, http.StatusOK, resp.StatusCode)
	AssertContains(t, resp.Header.Get("Content-Type"), "application/json")

	var topic Topic
	body, err := io.ReadAll(resp.Body)
	RequireNoError(t, err)
	if r := JSONUnmarshal(body, &topic); !r.OK {
		t.Fatal(r.Error())
	}
	AssertEqual(t, "Getting Started", topic.Title)
	AssertEqual(t, "getting-started", topic.ID)
}

func TestServer_HandleAPITopic_Bad_NotFound(t *T) {
	ts := testServer(t)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/api/topics/nonexistent")
	RequireNoError(t, err)
	defer resp.Body.Close()

	AssertEqual(t, http.StatusNotFound, resp.StatusCode)
	AssertContains(t, resp.Header.Get("Content-Type"), "application/json")
}

func TestServer_HandleAPISearch_Good(t *T) {
	ts := testServer(t)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/api/search?q=config")
	RequireNoError(t, err)
	defer resp.Body.Close()

	AssertEqual(t, http.StatusOK, resp.StatusCode)
	AssertContains(t, resp.Header.Get("Content-Type"), "application/json")

	var results []SearchResult
	body, err := io.ReadAll(resp.Body)
	RequireNoError(t, err)
	if r := JSONUnmarshal(body, &results); !r.OK {
		t.Fatal(r.Error())
	}
	AssertNotEmpty(t, results)
}

func TestServer_HandleAPISearch_Bad_NoQuery(t *T) {
	ts := testServer(t)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/api/search")
	RequireNoError(t, err)
	defer resp.Body.Close()

	AssertEqual(t, http.StatusBadRequest, resp.StatusCode)
	AssertContains(t, resp.Header.Get("Content-Type"), "application/json")
}

func TestServer_ContentTypeHeaders_Good(t *T) {
	ts := testServer(t)
	defer ts.Close()

	tests := []struct {
		name        string
		path        string
		contentType string
	}{
		{"index HTML", "/", "text/html"},
		{"topic HTML", "/topics/getting-started", "text/html"},
		{"search HTML", "/search?q=test", "text/html"},
		{"API topics JSON", "/api/topics", "application/json"},
		{"API topic JSON", "/api/topics/getting-started", "application/json"},
		{"API search JSON", "/api/search?q=test", "application/json"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *T) {
			resp, err := http.Get(ts.URL + tt.path)
			RequireNoError(t, err)
			defer resp.Body.Close()

			AssertContains(t, resp.Header.Get("Content-Type"), tt.contentType,
				"Content-Type for %s should contain %s", tt.path, tt.contentType)
		})
	}
}

func TestNewServer_Good(t *T) {
	c := DefaultCatalog()
	srv := NewServer(c, ":8080")

	AssertNotNil(t, srv)
	AssertEqual(t, ":8080", srv.addr)
	AssertNotNil(t, srv.mux)
	AssertEqual(t, c, srv.catalog)
}
