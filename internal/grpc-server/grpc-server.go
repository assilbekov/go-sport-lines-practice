package grpcserver

import (
	"go-sport-lines-practice/api/proto/pkg/lines"
	"go-sport-lines-practice/internal/storage"
	"google.golang.org/grpc"
	"log/slog"
	"net"
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

func NewServer(store *storage.Storage, logger *slog.Logger) *Server {
	return &Server{
		store:  store,
		logger: logger,
	}
}

func (s *Server) MustStart(addr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		s.logger.Error("failed to listen", "error", err)
		panic(err)
	}

	s.logger.Info("grpc server started", "addr", addr)
	grpcServer := grpc.NewServer()
	lines.RegisterLinesServiceServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		s.logger.Error("failed to serve", "error", err)
		panic(err)
	}
}
