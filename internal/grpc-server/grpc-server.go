package grpcserver

import (
	linespb "go-sport-lines-practice/api/proto/pkg/lines"
	"go-sport-lines-practice/internal/storage"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"sync"
	"time"
)

type Server struct {
	linespb.UnimplementedLinesServiceServer

	store  *storage.Storage
	logger *slog.Logger

	// How do I know that the storage is synced?
	synced bool
	mu     sync.Mutex

	subs []*Subscription
}

type Subscription struct {
	stream linespb.LinesService_SubscribeOnSportLinesServer

	lines    []string
	interval time.Duration
}

func NewServer(store *storage.Storage, logger *slog.Logger) *Server {
	return &Server{
		store:  store,
		logger: logger,
	}
}

func (s *Server) SubscribeOnSportsLines(stream linespb.LinesService_SubscribeOnSportLinesServer) error {
	for {
		in, err := stream.Recv()
		if err != nil {
			s.logger.Error("failed to receive", "error", err)
			return err
		}

		go func() {
			ticker := time.NewTicker(time.Duration(in.Interval) * time.Second)
			defer ticker.Stop()

			for {
				select {
				case <-ticker.C:
					_, err := s.store.GetLines()
					if err != nil {
						s.logger.Error("failed to get lines", "error", err)
						return
					}
				}
			}
		}()
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
	linespb.RegisterLinesServiceServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		s.logger.Error("failed to serve", "error", err)
		panic(err)
	}
}
