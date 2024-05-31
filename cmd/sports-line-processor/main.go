package main

import (
	"fmt"
	"go-sport-lines-practice/internal/configs"
	"go-sport-lines-practice/internal/storage"
	"go-sport-lines-practice/internal/worker"
	"log/slog"
	"os"
)

func main() {
	cfg := configs.LoadConfig()
	fmt.Printf("config: %+v\n", cfg)

	store := storage.NewStorage()

	quitCh := make(chan struct{})
	go worker.StartWorker("SOCCER", cfg.SportsSyncIntervals.Soccer, store, quitCh)
	go worker.StartWorker("FOOTBALL", cfg.SportsSyncIntervals.Football, store, quitCh)
	go worker.StartWorker("BASEBALL", cfg.SportsSyncIntervals.Baseball, store, quitCh)

	select {
	// wait forever
	}
}

func setupLogger(logLevel string) *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
}
