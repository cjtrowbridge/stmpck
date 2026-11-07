package storage

import (
	"os"
	"path/filepath"
	"testing"

	"ebe-boilerplate/internal/play"
)

func TestAtomicSaveAndCorruptRecovery(t *testing.T) {
	s := Store{Root: t.TempDir()}
	c := play.NewCollection()
	c.New(0, 0, 1)
	first, _ := c.Encode()
	if err := s.Save(first); err != nil {
		t.Fatal(err)
	}
	c.Merge.Move(0)
	second, _ := c.Encode()
	if err := s.Save(second); err != nil {
		t.Fatal(err)
	}
	got, err := s.Load()
	if err != nil || got.Merge.Cells != c.Merge.Cells {
		t.Fatal("roundtrip", err)
	}
	if err = os.WriteFile(filepath.Join(s.Root, "games.json"), []byte("truncated"), 0600); err != nil {
		t.Fatal(err)
	}
	got, err = s.Load()
	want, _ := play.Decode(first)
	if err != nil || got.Merge.Cells != want.Merge.Cells {
		t.Fatal("backup recovery", err)
	}
	if err = s.Save([]byte("invalid")); err == nil {
		t.Fatal("invalid data overwrote saves")
	}
	if err = s.Save(second); err != nil {
		t.Fatal(err)
	}
	if _, err = s.Load(); err != nil {
		t.Fatal(err)
	}
}
