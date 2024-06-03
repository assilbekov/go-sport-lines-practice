package main

import (
	"fmt"
	"go-sport-lines-practice/internal/configs"
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

	log := setupLogger(cfg.LogLevel)
	log.Info("starting sports line processor")

	quitCh := make(chan struct{})
	go worker.StartWorker("SOCCER", cfg.SportsSyncIntervals.Soccer, store, quitCh)
	go worker.StartWorker("FOOTBALL", cfg.SportsSyncIntervals.Football, store, quitCh)
	go worker.StartWorker("BASEBALL", cfg.SportsSyncIntervals.Baseball, store, quitCh)

	select {
	// wait forever
	}
}

func setupLogger(logLevel string) *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{Level: slog.LevelDebug},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
