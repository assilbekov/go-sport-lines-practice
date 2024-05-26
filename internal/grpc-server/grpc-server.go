package grpcserver

import (
	"go-sport-lines-practice/api/proto/pkg/lines"
	"go-sport-lines-practice/internal/storage"
	"log/slog"
	"net"
	"sync"
)

type Server struct {
	lines.UnimplementedLinesServiceServer
	store  *storage.Storage
	synced bool
	mu     sync.Mutex
	subs   map[lines.LinesService_SubscribeOnSportLinesClient][]string
}

func NewServer(store *storage.Storage) *Server {
	return &Server{
		store: store,
		subs:  make(map[lines.LinesService_SubscribeOnSportLinesClient][]string),
	}
}

func (s *Server) Start(addr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		slog.Error("failed to listen", "err", err)
		return
	}

	grpcServer := lines.LinesServiceServer(s)
	lines.RegisterLinesServiceServer(grpcServer, s)
	slog.Info("grpc server started", "addr", addr)
	if err = grpcServer.Serve(lis); err != nil {
		slog.Error("failed to serve", "err", err)
	}
}
