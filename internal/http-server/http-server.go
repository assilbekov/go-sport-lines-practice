package httpserver

import (
	"go-sport-lines-practice/internal/storage"
	"net/http"
)

type Server struct {
	addr  string
	store *storage.Storage
}

func NewServer(addr string, store *storage.Storage) *Server {
	return &Server{
		addr:  addr,
		store: store,
	}
}

func (s *Server) MustStart() {
	http.HandleFunc("/ready", s.readyHandler)
	if err := http.ListenAndServe(s.addr, nil); err != nil {
		panic(err)
	}
}

func (s *Server) readyHandler(w http.ResponseWriter, _ *http.Request) {
	if s.store.IsSynced() {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("ready"))

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
		_, err := w.Write([]byte("not ready"))

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
