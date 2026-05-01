// SPDX-Licence-Identifier: EUPL-1.2
package help

import (
	"net/http"

	core "dappco.re/go"
)

// writeJSON marshals v as JSON and writes it to w. If marshalling fails, an
// HTTP 500 is written. Replaces json.NewEncoder(w).Encode(v) without the
// streaming-encoder banned import.
func writeJSON(w http.ResponseWriter, v any) {
	r := core.JSONMarshal(v)
	if !r.OK {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(r.Value.([]byte)); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

// Server serves the help catalog over HTTP.
type Server struct {
	catalog *Catalog
	addr    string
	mux     *http.ServeMux
}

// NewServer creates an HTTP server for the given catalog.
// Routes are registered on creation; the caller can use ServeHTTP as
// an http.Handler or call ListenAndServe to start listening.
func NewServer(catalog *Catalog, addr string) *Server {
	s := &Server{
		catalog: catalog,
		addr:    addr,
		mux:     http.NewServeMux(),
	}

	// HTML routes
	s.mux.HandleFunc("GET /", s.handleIndex)
	s.mux.HandleFunc("GET /topics/{id}", s.handleTopic)
	s.mux.HandleFunc("GET /search", s.handleSearch)

	// JSON API routes
	s.mux.HandleFunc("GET /api/topics", s.handleAPITopics)
	s.mux.HandleFunc("GET /api/topics/{id}", s.handleAPITopic)
	s.mux.HandleFunc("GET /api/search", s.handleAPISearch)

	return s
}

// ServeHTTP implements http.Handler, delegating to the internal mux.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

// ListenAndServe starts the HTTP server.
func (s *Server) ListenAndServe() core.Result {
	srv := &http.Server{
		Addr:    s.addr,
		Handler: s.mux,
	}
	if err := srv.ListenAndServe(); err != nil {
		return core.Fail(core.E("help.Server.ListenAndServe", "listen and serve", err))
	}
	return core.Ok(nil)
}

// setSecurityHeaders sets common security headers.
func setSecurityHeaders(w http.ResponseWriter) {
	w.Header().Set("X-Content-Type-Options", "nosniff")
}

// --- HTML handlers ---

func (s *Server) handleIndex(w http.ResponseWriter, _ *http.Request) {
	setSecurityHeaders(w)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	topics := s.catalog.List()
	_, _ = w.Write([]byte(RenderIndexPage(topics)))
}

func (s *Server) handleTopic(w http.ResponseWriter, r *http.Request) {
	setSecurityHeaders(w)
	id := r.PathValue("id")

	res := s.catalog.Get(id)
	if !res.OK {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(Render404Page()))
		return
	}
	topic := res.Value.(*Topic)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write([]byte(RenderTopicPage(topic, s.catalog.List())))
}

func (s *Server) handleSearch(w http.ResponseWriter, r *http.Request) {
	setSecurityHeaders(w)
	query := r.URL.Query().Get("q")

	if query == "" {
		http.Error(w, "Missing search query parameter 'q'", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	results := s.catalog.Search(query)
	_, _ = w.Write([]byte(RenderSearchPage(query, results)))
}

// --- JSON API handlers ---

func (s *Server) handleAPITopics(w http.ResponseWriter, _ *http.Request) {
	setSecurityHeaders(w)
	w.Header().Set("Content-Type", "application/json")
	topics := s.catalog.List()

	writeJSON(w, topics)
}

func (s *Server) handleAPITopic(w http.ResponseWriter, r *http.Request) {
	setSecurityHeaders(w)
	id := r.PathValue("id")

	res := s.catalog.Get(id)
	if !res.OK {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		writeJSON(w, map[string]string{"error": "topic not found"})
		return
	}
	topic := res.Value.(*Topic)

	w.Header().Set("Content-Type", "application/json")
	writeJSON(w, topic)
}

func (s *Server) handleAPISearch(w http.ResponseWriter, r *http.Request) {
	setSecurityHeaders(w)
	query := r.URL.Query().Get("q")

	if query == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		writeJSON(w, map[string]string{"error": "missing query parameter 'q'"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	results := s.catalog.Search(query)

	writeJSON(w, results)
}
