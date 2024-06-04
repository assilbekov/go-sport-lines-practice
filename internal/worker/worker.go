package worker

import (
	"context"
	"fmt"
	"go-sport-lines-practice/internal/fetcher"
	"go-sport-lines-practice/internal/storage"
	"log/slog"
	"time"
)

type Worker struct {
	fetcher *fetcher.Fetcher
	storage *storage.Storage
	logger  *slog.Logger
}

func NewWorker(
	fetcher *fetcher.Fetcher,
	storage *storage.Storage,
	logger *slog.Logger,
) *Worker {
	return &Worker{
		fetcher: fetcher,
		storage: storage,
		logger:  logger,
	}
}

func (w *Worker) Start(ctx context.Context, sport string, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			go func() {
				line, err := w.fetcher.FetchSportLines(sport)
				if err != nil {
					w.logger.Error("failed to fetch sport line", "err", err)
					fmt.Printf("failed to fetch sport line: %v\n", err)
				}

				w.storage.UpdateLines(*line)
			}()
		case <-ctx.Done():
			w.logger.Info("stopping worker", "sport", sport)
		}
	}
}
