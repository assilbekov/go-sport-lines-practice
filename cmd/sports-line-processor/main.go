package main

import (
	"fmt"
	"go-sport-lines-practice/configs"
	"go-sport-lines-practice/internal/storage"
	"go-sport-lines-practice/internal/worker"
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
