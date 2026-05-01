// SPDX-Licence-Identifier: EUPL-1.2
package help

import (
	"encoding/json"
	"net/http"
)

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
func (s *Server) ListenAndServe() error {
	srv := &http.Server{
		Addr:    s.addr,
		Handler: s.mux,
	}
	return srv.ListenAndServe()
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

	topic, err := s.catalog.Get(id)
	if err != nil {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(Render404Page()))
		return
	}

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

	if err := json.NewEncoder(w).Encode(topics); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func (s *Server) handleAPITopic(w http.ResponseWriter, r *http.Request) {
	setSecurityHeaders(w)
	id := r.PathValue("id")

	topic, err := s.catalog.Get(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(map[string]string{"error": "topic not found"}); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(topic); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func (s *Server) handleAPISearch(w http.ResponseWriter, r *http.Request) {
	setSecurityHeaders(w)
	query := r.URL.Query().Get("q")

	if query == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(map[string]string{"error": "missing query parameter 'q'"}); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	results := s.catalog.Search(query)

	if err := json.NewEncoder(w).Encode(results); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
