package app

import (
	"fmt"
	"math"
	"math/rand/v2"

	"ebe-boilerplate/internal/menu"
	"ebe-boilerplate/internal/play"
	"ebe-boilerplate/internal/storage"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) initGames() {
	g.games = play.NewCollection()
	g.store = &storage.Store{Root: storage.DefaultRoot()}
	g.loadGames()
	g.makeDemos()
	installLifecycle(g)
}

// SetStorage is called before Android starts its update loop.
func (g *Game) SetStorage(root string) { g.store = &storage.Store{Root: root}; g.loadGames() }
func (g *Game) loadGames() {
	if c, err := g.store.Load(); err == nil {
		g.games = c
	}
	g.Plays = g.games.Plays
	g.galleryOrder = g.State.Order()
	g.snapshot()
}
func (g *Game) snapshot() {
	data, err := g.games.Encode()
	if err == nil {
		g.saveSnapshot.Store(&data)
	}
}
func (g *Game) saveGames() { g.snapshot(); g.Flush() }

// Flush writes an immutable snapshot, so Android can save while its view is suspending.
func (g *Game) Flush() {
	if data := g.saveSnapshot.Load(); data != nil {
		if err := g.store.Save(*data); err != nil {
			g.saveFailed.Store(true)
		} else {
			g.saveFailed.Store(false)
		}
	}
}
func (g *Game) Suspend()   { g.pendingSuspend.Store(true); g.Flush() }
func (g *Game) pauseGame() { g.games.Pause(); g.paused = true; g.aiClock = 0; g.saveGames() }
func (g *Game) Back() {
	if g.Settings {
		g.Settings = false
		return
	}
	if g.Selected >= 0 {
		g.pauseGame()
		g.mergeQueue = nil
		g.mergeAnimation = 0
		g.setup = false
		g.restartConfirm = false
	}
	g.State.Back()
	if g.Tab == menu.Discover {
		g.galleryOrder = g.State.Order()
		g.Focus = min(g.Focus, max(0, len(g.galleryOrder)-1))
		g.Clamp(g.height)
	}
}
func (g *Game) Open(row int) {
	if row < 0 || row >= len(g.galleryOrder) {
		return
	}
	g.Focus = row
	g.Selected = g.galleryOrder[row]
	if g.Selected < 0 {
		return
	}
	id := g.Selected
	g.setup = !g.games.Ready[id]
	created := false
	if id == play.LinearID && g.setup {
		g.games.New(id, 0, rand.Uint64())
		created = true
		g.setup = false
	}
	if id == play.TerritoryID && g.setup {
		g.prepareTerritorySetup()
	}
	g.option = 0
	if g.games.Ready[id] {
		switch id {
		case play.MergeID:
			g.option = g.games.Merge.N - 4
		case play.SnakeID:
			g.option = g.games.Snake.Speed

		case play.BounceID:
			g.option = g.games.Bounce.Speed
		}
	}
	g.cursor = 0
	g.paused = id == play.SnakeID || id == play.BounceID || id == play.TerritoryID
	g.restartConfirm = false
	g.pressed = false
	if !g.setup {
		if !created {
			g.games.Plays[id]++
		}
		g.Plays = g.games.Plays
		g.saveGames()
	}
}
func (g *Game) SetTab(tab int) { g.State.SetTab(tab); g.galleryOrder = g.State.Order() }

func (g *Game) startGame() {
	if g.Selected < 0 {
		return
	}
	if g.Selected == play.TerritoryID {
		g.games.Territory = g.territoryDrafts[g.option]
		g.games.Ready[play.TerritoryID] = true
		g.games.Plays[play.TerritoryID]++
	} else {
		g.games.New(g.Selected, g.option, rand.Uint64())
	}
	g.mergeAnimation = 0
	g.mergeQueue = nil
	g.Plays = g.games.Plays
	g.setup = false
	g.paused = false
	g.restartConfirm = false
	g.cursor = 0
	g.aiClock = 0
	g.saveGames()
}
func (g *Game) gameRect() (x, y, w, h float64) {
	x, w = g.cardBounds()
	y = menu.Header + 8
	h = math.Max(80, float64(g.height)-y-155)
	switch g.Selected {
	case play.MergeID:
		w = math.Min(w, h)
		h = w
	case play.SnakeID:
		h = math.Min(h, w*float64(play.SnakeH)/play.SnakeW)
		w = h * play.SnakeW / play.SnakeH
	case play.LinearID:
		h = math.Min(h, w*6/7)
		w = h * 7 / 6
	default:
		h = math.Min(h, w/1.2)
		w = h * 1.2
	}
	x = (float64(g.width) - w) / 2
	return
}
func (g *Game) gameAction(slot int) {
	if g.setup {
		switch slot {
		case 0:
			g.startGame()
		case 1:
			g.setup = false
			g.restartConfirm = false
			if !g.games.Ready[g.Selected] {
				g.Back()
			}
		}
		return
	}
	switch slot {
	case 0:
		if g.restartConfirm {
			g.setup = true
			g.restartConfirm = false
			g.option = 0
			if g.Selected == play.TerritoryID {
				g.prepareTerritorySetup()
			}
			if g.Selected == play.LinearID {
				g.startGame()
			}
			return
		}
		switch g.Selected {
		case play.SnakeID:
			g.paused = false
			if g.games.Snake.Running {
				g.pauseGame()
			} else if !g.games.Snake.Over {
				g.games.Snake.Running = true
			}
		case play.BounceID:
			g.paused = false
			if g.games.Bounce.Running {
				g.pauseGame()
			} else {
				g.games.Bounce.Serve()
			}
		default:
			g.paused = !g.paused
		}
	case 1:
		if g.restartConfirm {
			g.restartConfirm = false
		} else {
			g.restartConfirm = true
			g.pauseGame()
		}
	case 2:
		if g.restartConfirm {
			g.restartConfirm = false
			return
		}
		switch g.Selected {
		case play.MergeID:
			g.games.Merge.Undo()
			g.mergeAnimation = 0
			g.mergeQueue = nil
			g.flash = 12
		case play.LinearID:
			if g.games.Linear.Over {
				g.startGame()
			}
		}
	case 3:
		if g.Selected == play.MergeID {
			g.symbols = !g.symbols
		}
	}
	g.saveGames()
}
func (g *Game) direction(d int) {
	if g.setup {
		if d == 0 || d == 1 {
			g.option = max(0, g.option-1)
		} else {
			g.option = min(g.optionCount()-1, g.option+1)
		}
		return
	}
	if g.paused || g.restartConfirm {
		return
	}
	switch g.Selected {
	case play.MergeID:
		if g.mergeAnimation > 0 {
			if len(g.mergeQueue) < 4 {
				g.mergeQueue = append(g.mergeQueue, d)
			}
			return
		}
		before := g.games.Merge.Cells
		if g.games.Merge.Move(d) {
			g.mergeBefore = before
			g.mergeTargets = g.games.Merge.Travel
			g.mergeAnimation = 12
			g.flash = 14
			g.saveGames()
		}
	case play.SnakeID:
		g.games.Snake.Turn(d)
		if !g.games.Snake.Over {
			g.games.Snake.Running = true
		}
		g.saveGames()
	case play.LinearID:
		if d == 0 {
			g.cursor = max(0, g.cursor-1)
		}
		if d == 2 {
			g.cursor = min(6, g.cursor+1)
		}
	case play.TerritoryID:
		g.cursor = g.games.Territory.Terrain.Neighbor(g.cursor, d)
	case play.BounceID:
		if d == 1 {
			g.games.Bounce.Move(g.games.Bounce.Player - .055)
		}
		if d == 3 {
			g.games.Bounce.Move(g.games.Bounce.Player + .055)
		}
	}
}
func (g *Game) confirmGame() {
	if g.setup {
		g.startGame()
		return
	}
	if g.restartConfirm {
		g.gameAction(0)
		return
	}
	if g.paused {
		g.paused = false
		g.saveGames()
		return
	}
	switch g.Selected {
	case play.SnakeID, play.BounceID:
		g.gameAction(0)
	case play.LinearID:
		g.dropLinear(g.cursor)
	case play.TerritoryID:
		if !g.paused {
			g.games.Territory.Choose(g.cursor)
			g.saveGames()
		}
	}
}
func (g *Game) dropLinear(col int) {
	l := &g.games.Linear
	if !g.paused && !g.restartConfirm && (l.TwoPlayers || l.Turn == l.Human) {
		if l.Drop(col) {
			g.flash = 12
			g.aiClock = 0
			g.saveGames()
		}
	}
}
func (g *Game) optionCount() int {
	switch g.Selected {
	case play.MergeID, play.SnakeID, play.BounceID:
		return 3
	case play.TerritoryID:
		return 3
	}
	return 1
}
func (g *Game) playTap(x, y float64) {
	if g.setup && g.Selected == play.MergeID {
		for i := 0; i < 3; i++ {
			bx, by, bw, bh := g.mergeChoiceRect(i)
			if x >= bx && x <= bx+bw && y >= by && y <= by+bh {
				g.option = i
				g.startGame()
				return
			}
		}
		return
	}
	if g.setup && g.Selected == play.TerritoryID {
		g.territorySetupTap(x, y)
		return
	}
	if y >= float64(g.height-62) {
		slot := int(x * 4 / float64(g.width))
		if slot >= 0 && slot < 4 {
			g.gameAction(slot)
		}
		return
	}
	if g.setup {
		if y >= menu.Header+45 && y < menu.Header+145 {
			g.option = min(g.optionCount()-1, max(0, int(x*float64(g.optionCount())/float64(g.width))))
		}
		return
	}
	if g.paused || g.restartConfirm {
		return
	}
	bx, by, bw, bh := g.gameRect()

	if x < bx || x > bx+bw || y < by || y > by+bh {
		return
	}
	switch g.Selected {
	case play.LinearID:
		g.cursor = min(6, int((x-bx)*7/bw))
		g.dropLinear(g.cursor)
	case play.TerritoryID:
		g.cursor = g.games.Territory.Terrain.At((x-bx)/bw, (y-by)/bh)
		if !g.paused {
			g.games.Territory.Choose(g.cursor)
			g.saveGames()
		}

	}
}
func (g *Game) updatePlay(up, down, left, right, enter bool) {
	if g.setup && g.Selected == play.TerritoryID {
		g.territorySetupKeys()
	}
	if g.mergeAnimation > 0 && !g.paused {
		g.mergeAnimation--
	}
	if g.mergeAnimation == 0 && len(g.mergeQueue) > 0 && !g.paused && !g.setup && !g.restartConfirm {
		d := g.mergeQueue[0]
		g.mergeQueue = g.mergeQueue[1:]
		g.direction(d)
	}

	if g.flash > 0 {
		g.flash--
	}
	if left {
		g.direction(0)
	}
	if up {
		g.direction(1)
	}
	if right {
		g.direction(2)
	}
	if down {
		g.direction(3)
	}
	if enter {
		g.confirmGame()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.gameAction(0)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.gameAction(1)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyU) && g.Selected == play.MergeID {
		g.gameAction(2)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) && g.Selected == play.MergeID {
		g.symbols = !g.symbols
	}

	if g.setup || g.paused || g.restartConfirm {
		return
	}
	changed := false
	switch g.Selected {
	case play.SnakeID:
		changed = g.games.Snake.Tick()
	case play.TerritoryID:
		changed = g.games.Territory.Tick()

	case play.LinearID:
		l := &g.games.Linear
		if !l.TwoPlayers && !l.Over && l.Turn != l.Human {
			g.aiClock++
			if g.aiClock >= 30 {
				l.Drop(l.BestColumn())
				g.aiClock = 0
				g.flash = 12
				changed = true
			}
		}
	case play.BounceID:
		b := &g.games.Bounce
		if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
			b.Move(b.Player - .012)
		}
		if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
			b.Move(b.Player + .012)
		}
		for _, id := range ebiten.AppendGamepadIDs(nil) {
			if ebiten.IsStandardGamepadLayoutAvailable(id) {
				if ebiten.IsStandardGamepadButtonPressed(id, ebiten.StandardGamepadButtonLeftTop) {
					b.Move(b.Player - .012)
				}
				if ebiten.IsStandardGamepadButtonPressed(id, ebiten.StandardGamepadButtonLeftBottom) {
					b.Move(b.Player + .012)
				}
			}
		}
		changed = b.Tick()
	}
	if changed {
		g.snapshot()
		g.saveClock++
		if g.Selected == play.LinearID || g.saveClock >= 30 || g.games.Snake.Over || !g.games.Bounce.Running && g.Selected == play.BounceID {
			g.saveClock = 0
			g.Flush()
		}
	}
}
func (g *Game) statusText() string {
	if g.restartConfirm {
		return "New game? Confirm or cancel."
	}
	if g.setup {
		return "Choose, then start"
	}
	if g.paused {
		return "Paused - press play to resume"
	}
	switch g.Selected {
	case play.MergeID:
		m := g.games.Merge
		if m.Over {
			return fmt.Sprintf("%d   No moves - undo or new game", m.Score)
		}
		if m.Won {
			return fmt.Sprintf("%d   2048! Keep merging", m.Score)
		}
		return fmt.Sprintf("%d   /   2048", m.Score)
	case play.SnakeID:
		s := g.games.Snake
		if s.Won {
			return "Board complete!"
		}
		if s.Over {
			return fmt.Sprintf("%d   Round over", s.Score)
		}
		if !s.Running {
			return "Direction or play to start"
		}
		return fmt.Sprint(s.Score)
	case play.TerritoryID:
		t := g.games.Territory
		if t.Winner == 1 {
			return "All territories yours"
		}
		if t.Winner > 1 {
			return fmt.Sprintf("Opponent %d controls the map", t.Winner-1)
		}
		if !t.Alive(1) {
			return "Your forces are defeated"
		}
		return "Choose source, then adjacent destination"

	case play.LinearID:
		l := g.games.Linear
		if l.Winner != 0 {
			return fmt.Sprintf("Player %d wins", l.Winner)
		}
		if l.Over {
			return "Draw"
		}
		return fmt.Sprintf("Player %d", l.Turn)
	case play.BounceID:
		b := g.games.Bounce
		if b.Over {
			if b.Points[0] >= 5 {
				return "You win"
			}
			return "Opponent wins"
		}
		if !b.Running {
			return "Play to serve"
		}
	}
	return ""
}

func (g *Game) territorySwipe(x, y, endX, endY float64) {
	if g.setup || g.paused || g.restartConfirm || g.Settings {
		return
	}
	bx, by, bw, bh := g.gameRect()
	from := play.Point{X: (x - bx) / bw, Y: (y - by) / bh}
	to := play.Point{X: (endX - bx) / bw, Y: (endY - by) / bh}
	if g.games.Territory.OrderBetween(1, from, to) {
		g.saveGames()
	}
}
