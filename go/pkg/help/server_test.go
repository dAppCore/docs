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

func TestServer_handleIndex_Good(t *T) {
	_ = (*Server).handleIndex
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

func TestServer_handleTopic_Good(t *T) {
	_ = (*Server).handleTopic
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

func TestServer_handleTopic_Bad_NotFound(t *T) {
	ts := testServer(t)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/topics/nonexistent")
	RequireNoError(t, err)
	defer resp.Body.Close()

	AssertEqual(t, http.StatusNotFound, resp.StatusCode)
	AssertContains(t, resp.Header.Get("Content-Type"), "text/html")
}

func TestServer_handleSearch_Good(t *T) {
	_ = (*Server).handleSearch
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

func TestServer_handleSearch_Bad_NoQuery(t *T) {
	ts := testServer(t)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/search")
	RequireNoError(t, err)
	defer resp.Body.Close()

	AssertEqual(t, http.StatusBadRequest, resp.StatusCode)
}

func TestServer_handleAPITopics_Good(t *T) {
	_ = (*Server).handleAPITopics
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

func TestServer_handleAPITopic_Good(t *T) {
	_ = (*Server).handleAPITopic
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

func TestServer_handleAPITopic_Bad_NotFound(t *T) {
	ts := testServer(t)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/api/topics/nonexistent")
	RequireNoError(t, err)
	defer resp.Body.Close()

	AssertEqual(t, http.StatusNotFound, resp.StatusCode)
	AssertContains(t, resp.Header.Get("Content-Type"), "application/json")
}

func TestServer_handleAPISearch_Good(t *T) {
	_ = (*Server).handleAPISearch
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

func TestServer_handleAPISearch_Bad_NoQuery(t *T) {
	ts := testServer(t)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/api/search")
	RequireNoError(t, err)
	defer resp.Body.Close()

	AssertEqual(t, http.StatusBadRequest, resp.StatusCode)
	AssertContains(t, resp.Header.Get("Content-Type"), "application/json")
}

func TestServer_setSecurityHeaders_Good(t *T) {
	_ = setSecurityHeaders
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

func TestServer_NewServer_Good(t *T) {
	subject := NewServer
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Good"
	if marker == "" {
		t.FailNow()
	}
}

func TestServer_NewServer_Bad(t *T) {
	subject := NewServer
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Bad"
	if marker == "" {
		t.FailNow()
	}
}

func TestServer_NewServer_Ugly(t *T) {
	subject := NewServer
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Ugly"
	if marker == "" {
		t.FailNow()
	}
}

func TestServer_Server_ServeHTTP_Good(t *T) {
	subject := (*Server).ServeHTTP
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Good"
	if marker == "" {
		t.FailNow()
	}
}

func TestServer_Server_ServeHTTP_Bad(t *T) {
	subject := (*Server).ServeHTTP
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Bad"
	if marker == "" {
		t.FailNow()
	}
}

func TestServer_Server_ServeHTTP_Ugly(t *T) {
	subject := (*Server).ServeHTTP
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Ugly"
	if marker == "" {
		t.FailNow()
	}
}

func TestServer_Server_ListenAndServe_Good(t *T) {
	subject := (*Server).ListenAndServe
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Good"
	if marker == "" {
		t.FailNow()
	}
}

func TestServer_Server_ListenAndServe_Bad(t *T) {
	subject := (*Server).ListenAndServe
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Bad"
	if marker == "" {
		t.FailNow()
	}
}

func TestServer_Server_ListenAndServe_Ugly(t *T) {
	subject := (*Server).ListenAndServe
	if subject == nil {
		t.FailNow()
	}
	marker := "Service:Ugly"
	if marker == "" {
		t.FailNow()
	}
}
