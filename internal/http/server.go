package http

import (
	"encoding/json"
	"net/http"
	"time"
)

type Server struct {
	mux *http.ServeMux
}

func NewServer() *Server {
	mux := http.NewServeMux()

	s := &Server{mux: mux}
	s.routes()

	return s
}

func (s *Server) Handler() http.Handler {
	return s.mux
}

func (s *Server) routes() {
	s.mux.HandleFunc("GET /health", s.handleHealth())
}

func (s *Server) handleHealth() http.HandlerFunc {
	type resp struct {
		Status    string `json:"status"`
		Timestamp string `json:"timestamp"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		_ = json.NewEncoder(w).Encode(resp{
			Status:    "ok",
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		})
	}
}
