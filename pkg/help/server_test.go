// SPDX-Licence-Identifier: EUPL-1.2
package help

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// testServer creates a test catalog with topics and returns an httptest.Server.
func testServer(t *testing.T) *httptest.Server {
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

func TestServer_HandleIndex_Good(t *testing.T) {
	ts := testServer(t)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/")
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Contains(t, resp.Header.Get("Content-Type"), "text/html")
	assert.Equal(t, "nosniff", resp.Header.Get("X-Content-Type-Options"))

	var buf [64 * 1024]byte
	n, _ := resp.Body.Read(buf[:])
	body := string(buf[:n])
	assert.Contains(t, body, "Getting Started")
	assert.Contains(t, body, "Configuration")
}

func TestServer_HandleTopic_Good(t *testing.T) {
	ts := testServer(t)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/topics/getting-started")
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Contains(t, resp.Header.Get("Content-Type"), "text/html")

	var buf [64 * 1024]byte
	n, _ := resp.Body.Read(buf[:])
	body := string(buf[:n])
	assert.Contains(t, body, "Getting Started")
	assert.Contains(t, body, "<strong>guide</strong>")
}

func TestServer_HandleTopic_Bad_NotFound(t *testing.T) {
	ts := testServer(t)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/topics/nonexistent")
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	assert.Contains(t, resp.Header.Get("Content-Type"), "text/html")
}

func TestServer_HandleSearch_Good(t *testing.T) {
	ts := testServer(t)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/search?q=install")
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Contains(t, resp.Header.Get("Content-Type"), "text/html")

	var buf [64 * 1024]byte
	n, _ := resp.Body.Read(buf[:])
	body := string(buf[:n])
	assert.Contains(t, body, "install")
}

func TestServer_HandleSearch_Bad_NoQuery(t *testing.T) {
	ts := testServer(t)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/search")
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestServer_HandleAPITopics_Good(t *testing.T) {
	ts := testServer(t)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/api/topics")
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Contains(t, resp.Header.Get("Content-Type"), "application/json")
	assert.Equal(t, "nosniff", resp.Header.Get("X-Content-Type-Options"))

	var topics []Topic
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&topics))
	assert.Len(t, topics, 2)
}

func TestServer_HandleAPITopic_Good(t *testing.T) {
	ts := testServer(t)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/api/topics/getting-started")
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Contains(t, resp.Header.Get("Content-Type"), "application/json")

	var topic Topic
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&topic))
	assert.Equal(t, "Getting Started", topic.Title)
	assert.Equal(t, "getting-started", topic.ID)
}

func TestServer_HandleAPITopic_Bad_NotFound(t *testing.T) {
	ts := testServer(t)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/api/topics/nonexistent")
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	assert.Contains(t, resp.Header.Get("Content-Type"), "application/json")
}

func TestServer_HandleAPISearch_Good(t *testing.T) {
	ts := testServer(t)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/api/search?q=config")
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Contains(t, resp.Header.Get("Content-Type"), "application/json")

	var results []SearchResult
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&results))
	assert.NotEmpty(t, results)
}

func TestServer_HandleAPISearch_Bad_NoQuery(t *testing.T) {
	ts := testServer(t)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/api/search")
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	assert.Contains(t, resp.Header.Get("Content-Type"), "application/json")
}

func TestServer_ContentTypeHeaders_Good(t *testing.T) {
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
		t.Run(tt.name, func(t *testing.T) {
			resp, err := http.Get(ts.URL + tt.path)
			require.NoError(t, err)
			defer resp.Body.Close()

			assert.Contains(t, resp.Header.Get("Content-Type"), tt.contentType,
				"Content-Type for %s should contain %s", tt.path, tt.contentType)
		})
	}
}

func TestNewServer_Good(t *testing.T) {
	c := DefaultCatalog()
	srv := NewServer(c, ":8080")

	assert.NotNil(t, srv)
	assert.Equal(t, ":8080", srv.addr)
	assert.NotNil(t, srv.mux)
	assert.Equal(t, c, srv.catalog)
}
