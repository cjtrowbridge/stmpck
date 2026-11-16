package app

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestUpdateBridgeLifecycle(t *testing.T) {
	g := &Game{}
	if g.updateMessage() != "Updates are available in the Android app." {
		t.Fatal("desktop state must be honest")
	}
	g.checkForUpdate() // Unsupported platforms have no native callback.
	var called atomic.Int32
	g.SetUpdateCheck(func() { called.Add(1); g.SetUpdateStatus("Checking for an update...") })
	g.checkForUpdate()
	if called.Load() != 1 || g.updateMessage() != "Checking for an update..." {
		t.Fatal("manual update did not reach native bridge")
	}
	g.SetUpdateCheck(nil)
	g.checkForUpdate()
	if called.Load() != 1 {
		t.Fatal("detached Activity received request")
	}
}

func TestUpdateBridgeConcurrentDetachAndStatus(t *testing.T) {
	g := &Game{}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			g.SetUpdateCheck(func() {})
			g.SetUpdateStatus("Up to date.")
			g.SetUpdateCheck(nil)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			g.checkForUpdate()
			_ = g.updateMessage()
		}
	}()
	wg.Wait()
}
