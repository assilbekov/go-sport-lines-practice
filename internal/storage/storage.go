package storage

import "sync"

type Line struct {
	Soccer   float64
	Football float64
	Baseball float64
}

type Storage struct {
	mu    sync.RWMutex
	lines Line

	// Is there a better way to do this?
	synced bool
}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) UpdateLines(newLines Line) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if newLines.Soccer != 0 {
		s.lines.Soccer = newLines.Soccer
	}
	if newLines.Football != 0 {
		s.lines.Football = newLines.Football
	}
	if newLines.Baseball != 0 {
		s.lines.Baseball = newLines.Baseball
	}

	s.synced = true
}

func (s *Storage) GetLines() (Line, error) {
	// Is this going to block the execution thread?
	// Will I be able to read the lines while they are being updated?
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.lines, nil
}

func (s *Storage) IsSynced() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.synced
}
