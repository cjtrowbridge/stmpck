package app

// SetUpdateCheck connects the Android wrapper without networking on the game loop.
// A mutex protects the callback; immutable status strings cross threads atomically.
func (g *Game) SetUpdateCheck(check func()) {
	g.updateMu.Lock()
	g.updateCheck = check
	g.updateMu.Unlock()
}

func (g *Game) SetUpdateStatus(status string) { g.updateStatus.Store(&status) }

func (g *Game) checkForUpdate() {
	g.updateMu.Lock()
	check := g.updateCheck
	g.updateMu.Unlock()
	if check != nil {
		check()
	}
}

func (g *Game) updateMessage() string {
	if status := g.updateStatus.Load(); status != nil {
		return *status
	}
	return "Updates are available in the Android app."
}
