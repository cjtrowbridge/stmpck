//go:build !js

package storage

import (
	"os"
	"path/filepath"
)

func DefaultRoot() string {
	root, err := os.UserConfigDir()
	if err != nil {
		return ""
	}
	return filepath.Join(root, "stmpck")
}
func (s *Store) read(name string) ([]byte, error) { return os.ReadFile(filepath.Join(s.Root, name)) }
func (s *Store) write(name string, data []byte) error {
	if err := os.MkdirAll(s.Root, 0700); err != nil {
		return err
	}
	f, err := os.CreateTemp(s.Root, ".games-*")
	if err != nil {
		return err
	}
	temp := f.Name()
	defer os.Remove(temp)
	if _, err = f.Write(data); err == nil {
		err = f.Sync()
	}
	closeErr := f.Close()
	if err != nil {
		return err
	}
	if closeErr != nil {
		return closeErr
	}
	if err = os.Rename(temp, filepath.Join(s.Root, name)); err != nil {
		return err
	}
	if dir, err := os.Open(s.Root); err == nil {
		defer dir.Close()
		_ = dir.Sync()
	}
	return nil
}
