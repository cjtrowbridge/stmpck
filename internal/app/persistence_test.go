package app

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"ebe-boilerplate/internal/menu"
	"ebe-boilerplate/internal/play"
	"ebe-boilerplate/internal/storage"
)

func TestFlushSelectsLatestSnapshotAfterSerialization(t *testing.T) {
	g := &Game{store: &storage.Store{Root: t.TempDir()}, games: play.NewCollection()}
	g.games.Plays[0] = 1
	g.snapshot()
	// Hold the serialization boundary while a lifecycle flush starts. It must
	// not select its bytes until the boundary is released.
	g.flushMu.Lock()
	started, done := make(chan struct{}), make(chan struct{})
	go func() { close(started); g.Flush(); close(done) }()
	<-started
	early := false
	select {
	case <-done:
		early = true
	case <-time.After(50 * time.Millisecond):
	}
	g.games.Plays[0] = 2
	g.snapshot()
	g.flushMu.Unlock()
	<-done
	if early {
		t.Fatal("flush bypassed serialization boundary")
	}
	got, err := g.store.Load()
	if err != nil || got.Plays[0] != 2 {
		t.Fatalf("stale flush: plays=%d err=%v", got.Plays[0], err)
	}
	if g.saveFailed.Load() {
		t.Fatal("successful flush reported failure")
	}
}

func TestFlushFailureAndRecovery(t *testing.T) {
	g := &Game{store: &storage.Store{}, games: play.NewCollection()}
	g.snapshot()
	g.Flush()
	if !g.saveFailed.Load() {
		t.Fatal("failed write not reported")
	}
	g.store = &storage.Store{Root: t.TempDir()}
	g.Flush()
	if g.saveFailed.Load() {
		t.Fatal("successful recovery not reported")
	}
}

func idleGame(t testing.TB, id int) *Game {
	t.Helper()
	g := &Game{store: &storage.Store{Root: t.TempDir()}, games: play.NewCollection(), State: menu.New()}
	for game := 0; game < play.Count; game++ {
		g.games.New(game, 0, uint64(game+1))
	}
	g.Selected = id
	if id == play.EchoID {
		g.games.Echo.Phase = play.EchoInput
		g.games.Echo.Clock = 0
	}
	g.snapshot()
	return g
}

func TestStationaryGamesDoNotSnapshotOrWrite(t *testing.T) {
	for _, id := range []int{play.EchoID, play.ClusterID} {
		t.Run(play.IDs[id], func(t *testing.T) {
			g := idleGame(t, id)
			before := g.saveSnapshot.Load()
			for tick := 0; tick < 600; tick++ {
				g.updatePlay(false, false, false, false, false)
			}
			if g.saveSnapshot.Load() != before || g.saveClock != 0 {
				t.Fatal("stationary game advanced persistence")
			}
			if _, err := g.store.Load(); err == nil {
				t.Fatal("stationary game wrote a save")
			}
		})
	}
}

func TestChangedTickStillSnapshotsAndFlushes(t *testing.T) {
	for _, id := range []int{play.EchoID, play.ClusterID} {
		t.Run(play.IDs[id], func(t *testing.T) {
			g := idleGame(t, id)
			if id == play.EchoID {
				g.games.Echo.Retry()
			} else {
				g.games.Cluster.Fire()
			}
			before := g.saveSnapshot.Load()
			g.saveClock = 29
			g.updatePlay(false, false, false, false, false)
			if g.saveSnapshot.Load() == before || g.saveClock != 0 {
				t.Fatal("changed tick skipped persistence")
			}
			if _, err := g.store.Load(); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func BenchmarkIdleAdditions(b *testing.B) {
	for _, id := range []int{play.EchoID, play.ClusterID} {
		b.Run(play.IDs[id], func(b *testing.B) {
			g := idleGame(b, id)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				g.updatePlay(false, false, false, false, false)
			}
		})
	}
}

func TestIdleTransitionFlushesLastChangedSnapshotOnce(t *testing.T) {
	for _, id := range []int{play.EchoID, play.ClusterID} {
		t.Run(play.IDs[id], func(t *testing.T) {
			g := idleGame(t, id)
			g.games.Plays[id] = 99
			g.snapshot()
			latest := g.saveSnapshot.Load()
			g.saveClock = 1
			g.updatePlay(false, false, false, false, false)
			got, err := g.store.Load()
			if err != nil || got.Plays[id] != 99 || g.saveClock != 0 {
				t.Fatalf("final snapshot not flushed: %v", err)
			}
			for tick := 0; tick < 60; tick++ {
				g.updatePlay(false, false, false, false, false)
			}
			if g.saveSnapshot.Load() != latest || g.saveClock != 0 {
				t.Fatal("idle state did not stay clean")
			}
			if _, err := os.Stat(filepath.Join(g.store.Root, "games.backup.json")); !os.IsNotExist(err) {
				t.Fatalf("idle state wrote a second save: %v", err)
			}
		})
	}
}
