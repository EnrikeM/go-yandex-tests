package inmem

import (
	"fmt"
	"sync"
)

type Storage struct {
	urls map[string]string
	mu   sync.RWMutex
}

func (s *Storage) Set(k, v string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.urls[k] = v
}

func (s *Storage) Get(k string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	v, ok := s.urls[k]
	if !ok {
		return v, fmt.Errorf("short URL not found")
	}
	return v, nil
}

func New() *Storage {
	return &Storage{
		urls: make(map[string]string, 10),
		mu:   sync.RWMutex{},
	}
}
