package main

import (
	"context"
	"fmt"
	"go-sport-lines-practice/internal/configs"
	"go-sport-lines-practice/internal/fetcher"
	grpcserver "go-sport-lines-practice/internal/grpc-server"
	httpserver "go-sport-lines-practice/internal/http-server"
	"go-sport-lines-practice/internal/lib/slogpretty"
	"go-sport-lines-practice/internal/storage"
	"go-sport-lines-practice/internal/worker"
	"log/slog"
	"os"
)

func main() {
	cfg := configs.LoadConfig()
	fmt.Printf("config: %+v\n", cfg)

	store := storage.NewStorage()

	logger := setupLogger(cfg.LogLevel)
	logger.Info("starting sports line processor")

	ctx := context.Background()
	f := fetcher.NewFetcher(cfg.BaseURL, logger)
	w := worker.NewWorker(f, store, logger)

	go w.Start(ctx, "SOCCER", cfg.SportsSyncIntervals.Soccer)
	go w.Start(ctx, "FOOTBALL", cfg.SportsSyncIntervals.Football)
	go w.Start(ctx, "BASEBALL", cfg.SportsSyncIntervals.Baseball)

	httpServer := httpserver.NewServer(cfg.HTTPServerAddr, store)
	grpcServer := grpcserver.NewServer(store, logger)

	go httpServer.MustStart()
	go grpcServer.MustStart(cfg.GRPCServerAddr)

	// graceful shutdown
	stop := make(chan os.Signal, 1)
	sign := <-stop

	logger.Info("stopping sports line processor", "signal", sign)
}

func setupLogger(logLevel string) *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{Level: slog.LevelDebug},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
