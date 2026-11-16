package storage

import (
	"ebe-boilerplate/internal/play"
	"testing"
)

func populatedCollection() play.Collection {
	c := play.NewCollection()
	for id := 0; id < play.Count; id++ {
		c.New(id, 0, uint64(id+1))
	}
	return c
}

func BenchmarkCollectionEncode(b *testing.B) {
	c := populatedCollection()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := c.Encode(); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkStoreSave(b *testing.B) {
	c := populatedCollection()
	data, err := c.Encode()
	if err != nil {
		b.Fatal(err)
	}
	s := Store{Root: b.TempDir()}
	if err := s.Save(data); err != nil {
		b.Fatal(err)
	}
	b.ReportAllocs()
	b.SetBytes(int64(len(data)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := s.Save(data); err != nil {
			b.Fatal(err)
		}
	}
}
