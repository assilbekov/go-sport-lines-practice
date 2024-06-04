package grpcserver

import (
	"go-sport-lines-practice/api/proto/pkg/lines"
	"go-sport-lines-practice/internal/storage"
	"log/slog"
	"sync"
)

type Server struct {
	lines.UnimplementedLinesServiceServer

	store  *storage.Storage
	logger *slog.Logger

	// How do I know that the storage is synced?
	synced bool
	mu     sync.Mutex
}

func NewServer(store *storage.Storage) *Server {
	return &Server{
		store: store,
	}
}
