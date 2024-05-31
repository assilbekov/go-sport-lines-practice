package worker

import (
	"fmt"
	"go-sport-lines-practice/internal/fetcher"
	"go-sport-lines-practice/internal/storage"
	"time"
)

func StartWorker(
	sport string,
	interval time.Duration,
	storage *storage.Storage,
	quitCh chan struct{},
) {
	f := fetcher.NewFetcher("http://localhost:8080/lines/")
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			go func() {
				line, err := f.FetchSportLines(sport)
				if err != nil {
					fmt.Printf("failed to fetch sport line: %v\n", err)
				}

				storage.UpdateLines(*line)
			}()
		case <-quitCh:
			return
		}
	}
}
