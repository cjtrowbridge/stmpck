package storage

import (
	"ebe-boilerplate/internal/play"
	"errors"
	"sync"
)

// Store serializes writes from the game loop and native lifecycle callbacks.
type Store struct {
	Root string
	mu   sync.Mutex
}

func (s *Store) Load() (play.Collection, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	var last error
	for _, key := range []string{"games.json", "games.backup.json"} {
		data, err := s.read(key)
		if err != nil {
			last = err
			continue
		}
		c, err := play.Decode(data)
		if err == nil {
			return c, nil
		}
		last = err
	}
	return play.NewCollection(), last
}
func (s *Store) Save(data []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.Root == "" {
		return errors.New("save directory unavailable")
	}
	if _, err := play.Decode(data); err != nil {
		return err
	}
	if old, err := s.read("games.json"); err == nil {
		if _, err = play.Decode(old); err == nil {
			if err = s.write("games.backup.json", old); err != nil {
				return err
			}
		}
	}
	return s.write("games.json", data)
}
