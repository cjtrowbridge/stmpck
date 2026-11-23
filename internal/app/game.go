// STMPCK shares deterministic game models between its gallery and live games.
package app

import (
	"image"
	"image/color"
	"math"
	"math/rand/v2"
	"sync"
	"sync/atomic"

	"ebe-boilerplate/internal/menu"
	"ebe-boilerplate/internal/play"
	"ebe-boilerplate/internal/storage"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/basicfont"
)

type Game struct {
	updateMu                                                        sync.Mutex
	updateCheck                                                     func()
	updateStatus                                                    atomic.Pointer[string]
	embedded                                                        bool
	OffsetX, OffsetY                                                int
	InputScale                                                      float64
	territoryDrafts                                                 [3]play.Territory
	territoryMaps                                                   [3]uint64
	territoryStarts                                                 uint64
	opponents                                                       int
	galleryOrder                                                    []int
	mergeBefore                                                     [36]int
	mergeTargets                                                    [36]int
	mergeAnimation                                                  int
	mergeQueue                                                      []int
	demoMergeAnimation                                              int
	pendingBack                                                     atomic.Uint32
	pendingSuspend                                                  atomic.Bool
	saveSnapshot                                                    atomic.Pointer[[]byte]
	flushMu                                                         sync.Mutex
	saveFailed                                                      atomic.Bool
	store                                                           *storage.Store
	games                                                           play.Collection
	previews                                                        galleryImages
	resumeMenu, setup, paused, restartConfirm, symbols, wasSettings bool
	option, cursor, aiClock, flash, saveClock                       int
	menu.State
	mockOptions               [4]bool
	width, height             int
	tick                      float64
	sparkleSeed               uint64
	pressed                   bool
	touch                     bool
	touchID                   ebiten.TouchID
	sx, sy, lx, ly            float64
	travel                    float64
	wheelCooldown             int
	gesture                   string
	planetDragFrom            int
	bloomPaint                bool
	calibanDragStart          int
	bloomPicker, bloomPlacing bool
	bloomChoice               int
}

func NewGame() *Game { return newGame("") }
func NewEmbedded(root string) *Game {
	g := newGame(root)
	g.embedded = true
	g.CompactGallery = true
	return g
}
func (g *Game) local(x, y int) (int, int) {
	if g.InputScale == 0 {
		return x, y
	}
	return int(float64(x-g.OffsetX) / g.InputScale), int(float64(y-g.OffsetY) / g.InputScale)
}
func newGame(root string) *Game {
	g := &Game{State: menu.New(), mockOptions: [4]bool{false, true, false, false}, sparkleSeed: rand.Uint64(), width: 480, height: 800}
	g.initGames(root)
	return g
}

// RequestBack may be called from the Android UI thread. Navigation stays on the game thread.
func (g *Game) RequestBack() { g.pendingBack.Add(1) }

func (g *Game) Update() error {
	for n := g.pendingBack.Swap(0); n > 0; n-- {
		g.Back()
		g.pressed = false
	}
	if g.pendingSuspend.Swap(false) {
		g.pauseGame()
		g.pressed = false
	}
	if !ebiten.IsFocused() {
		if !g.paused && g.Selected >= 0 {
			g.pauseGame()
		}
		g.pressed = false
		return nil
	}
	if g.Settings && !g.wasSettings && g.Selected >= 0 {
		g.pauseGame()
	}
	g.wasSettings = g.Settings
	if g.wheelCooldown > 0 {
		g.wheelCooldown--
	}
	if g.Autoplay() {
		g.tick++
		if g.Selected < 0 && !g.Settings {
			g.updatePreviews()
		}
	}
	up := inpututil.IsKeyJustPressed(ebiten.KeyArrowUp)
	down := inpututil.IsKeyJustPressed(ebiten.KeyArrowDown)
	left := inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft)
	right := inpututil.IsKeyJustPressed(ebiten.KeyArrowRight)
	enter := inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsKeyJustPressed(ebiten.KeyNumpadEnter)
	back := inpututil.IsKeyJustPressed(ebiten.KeyEscape)
	for _, id := range ebiten.AppendGamepadIDs(nil) {
		if !ebiten.IsStandardGamepadLayoutAvailable(id) {
			continue
		}
		up = up || inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonLeftTop)
		down = down || inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonLeftBottom)
		left = left || inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonLeftLeft)
		right = right || inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonLeftRight)
		enter = enter || inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonRightBottom)
		back = back || inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonRightRight)
		if g.Selected >= 0 && !g.Settings && inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonRightTop) {
			g.gameAction(0)
		}
		if g.Selected >= play.FlowID && !g.Settings {
			if g.Selected == play.BloomID && inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonFrontTopRight) {
				g.toggleBloomPicker()
			}
			if inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonRightLeft) {
				g.gameAction(2)
			}
			if inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonLeftStick) {
				g.gameAction(3)
			}
			if inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonRightStick) {
				g.gameAction(1)
			}
		}
		if inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonCenterRight) {
			g.Settings = !g.Settings
		}
	}
	if g.embedded && (g.Selected < 0 || g.Settings) {
		up, down, left, right, enter, back = false, false, false, false, false, false
	}
	if back {
		g.Back()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyH) || inpututil.IsKeyJustPressed(ebiten.KeyF1) {
		g.Settings = !g.Settings
	}
	if !g.embedded && g.Settings && (enter || inpututil.IsKeyJustPressed(ebiten.KeyU)) {
		g.checkForUpdate()
	}
	if !g.Settings && g.Selected < 0 && (!g.embedded || !ebiten.IsKeyPressed(ebiten.KeyAlt)) {
		if inpututil.IsKeyJustPressed(ebiten.Key1) {
			g.SetTab(menu.Personal)
		}
		if inpututil.IsKeyJustPressed(ebiten.Key2) {
			g.SetTab(menu.Worldwide)
		}
		if inpututil.IsKeyJustPressed(ebiten.Key3) {
			g.SetTab(menu.Discover)
		}
		if !g.embedded && inpututil.IsKeyJustPressed(ebiten.KeyTab) {
			g.SetTab((g.Tab + 1) % 3)
		}
		for _, id := range ebiten.AppendGamepadIDs(nil) {
			if !g.embedded && ebiten.IsStandardGamepadLayoutAvailable(id) {
				if inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonFrontTopRight) {
					g.SetTab((g.Tab + 1) % 3)
				}
				if inpututil.IsStandardGamepadButtonJustPressed(id, ebiten.StandardGamepadButtonFrontTopLeft) {
					g.SetTab((g.Tab + 2) % 3)
				}
			}
		}
	}
	if !g.Settings && g.Selected >= 0 {
		g.updatePlay(up, down, left, right, enter)
	}
	if !g.Settings && g.Selected < 0 {
		if left {
			g.Cycle(-1)
		}
		if right {
			g.Cycle(1)
		}
		if g.Selected < 0 {
			if up {
				g.Move(-1, g.height)
			}
			if down {
				g.Move(1, g.height)
			}
			if enter {
				g.Open(g.Focus)
			}
			horizontal, wheel := ebiten.Wheel()
			if math.Abs(horizontal) > math.Abs(wheel) {
				if g.wheelCooldown == 0 && horizontal != 0 {
					if horizontal < 0 {
						g.Cycle(1)
					} else {
						g.Cycle(-1)
					}
					g.wheelCooldown = 15
				}
			} else {
				g.Scroll -= wheel * 42
			}
			g.Clamp(g.height)
		}
	}
	// One pointer at a time. Ignore extra fingers and synthetic mouse events.
	if !g.pressed {
		ids := inpututil.AppendJustPressedTouchIDs(nil)
		if len(ids) > 0 {
			g.touch = true
			g.touchID = ids[0]
			x, y := g.local(ebiten.TouchPosition(g.touchID))
			g.begin(float64(x), float64(y))
		} else if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			g.touch = false
			x, y := g.local(ebiten.CursorPosition())
			g.begin(float64(x), float64(y))
		}
	}
	if g.pressed {
		x, y := g.local(ebiten.CursorPosition())
		held := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
		if g.touch {
			held = false
			for _, id := range ebiten.AppendTouchIDs(nil) {
				if id == g.touchID {
					held = true
				}
			}
			if held {
				x, y = g.local(ebiten.TouchPosition(g.touchID))
			} else {
				x, y = int(g.lx), int(g.ly)
			}
		}
		fx, fy := float64(x), float64(y)
		g.travel = math.Max(g.travel, math.Hypot(fx-g.sx, fy-g.sy))
		if g.gesture == "" && g.travel >= 12 {
			g.gesture = menu.Gesture(fx-g.sx, fy-g.sy)
		}
		if held && g.gesture == "vertical" && g.Selected < 0 && !g.Settings && g.sy >= float64(g.HeaderHeight()) {
			g.Scroll -= fy - g.ly
			g.Clamp(g.height)
		}
		if held && g.Selected == play.BounceID && !g.Settings && !g.setup && !g.paused && g.gesture == "vertical" {
			_, _, _, bh := g.gameRect()
			delta := fy - g.ly
			if math.Hypot(g.lx-g.sx, g.ly-g.sy) < 12 {
				delta = fy - g.sy
			}
			g.games.Bounce.Move(g.games.Bounce.Player + delta/bh)
		}

		if held && g.Selected >= play.FlowID {
			g.additionDrag(fx, fy)
		}
		g.lx, g.ly = fx, fy
		if !held {
			g.pressed = false
			if g.travel < 12 {
				g.activate(fx, fy)
			} else if g.Selected >= 0 && !g.Settings {
				if g.Selected >= play.DriftID {
					g.expansionSwipe(fx, fy)
				}
				if !g.setup && !g.paused && !g.restartConfirm && g.Selected == play.PlanetsID {
					bx, by, bw, bh := g.gameRect()
					if g.games.Planets.Order(1, g.planetDragFrom, g.games.Planets.At((fx-bx)/bw, (fy-by)/bh)) {
						g.games.Planets.Source = -1
						g.saveGames()
					}
				}
				if g.Selected == play.BloomID {
					g.saveGames()
				}
				dx, dy := fx-g.sx, fy-g.sy
				if g.Selected == play.TerritoryID {
					g.territorySwipe(g.sx, g.sy, fx, fy)
				}
				if g.Selected == play.MergeID || g.Selected == play.SnakeID {
					if math.Abs(dx) > math.Abs(dy) {
						if dx < 0 {
							g.direction(0)
						} else {
							g.direction(2)
						}
					} else {
						if dy < 0 {
							g.direction(1)
						} else {
							g.direction(3)
						}
					}
				}
			} else if g.gesture == "horizontal" && math.Abs(fx-g.sx) > 48 && !g.Settings && g.sy >= float64(g.HeaderHeight()) {
				if fx < g.sx {
					g.Cycle(1)
				} else {
					g.Cycle(-1)
				}
			}
		}
	}
	return nil
}

func (g *Game) begin(x, y float64) {
	if y < 0 || (g.embedded && (x < 0 || x >= float64(g.width) || y >= float64(g.height))) {
		return
	}
	g.pressed = true
	g.sx = x
	g.sy = y
	g.lx = x
	g.ly = y
	g.travel = 0
	g.gesture = ""
	g.planetDragFrom = -1
	g.calibanDragStart = g.games.Caliban.Lane
	if g.Selected >= play.FlowID && !g.setup && !g.Settings {
		bx, by, bw, bh := g.gameRect()
		u, v := (x-bx)/bw, (y-by)/bh
		if u >= 0 && u < 1 && v >= 0 && v < 1 {
			if g.Selected == play.PlanetsID {
				g.planetDragFrom = g.games.Planets.At(u, v)
			}
			if g.Selected == play.BloomID {
				g.bloomPaint = !g.games.Bloom.Cells[int(v*play.BloomN)*play.BloomN+int(u*play.BloomN)]
			}
		}
	}
}
func (g *Game) activate(x, y float64) {
	if g.Settings {
		if y < 52 && (x < 64 || x >= float64(g.width-56)) {
			g.Back()
			return
		}
		if !g.embedded && x >= 16 && x <= float64(g.width-16) && y >= 430 && y < 482 {
			g.checkForUpdate()
			return
		}
		row := int((y - 100) / 58)
		if y >= 100 && row >= 0 && row < len(g.mockOptions) && x >= 16 && x <= float64(g.width-16) {
			g.mockOptions[row] = !g.mockOptions[row]
		}
		return
	}
	if g.embedded && g.Selected < 0 && y < float64(g.HeaderHeight()) {
		for i := 0; i < 6; i++ {
			if image.Pt(int(x), int(y)).In(g.ToolbarRect(i)) {
				switch i {
				case 0:
					g.Cycle(-1)
				case 5:
					g.Cycle(1)
				case 4:
					g.Settings = true
				default:
					g.SetTab(i - 1)
				}
				break
			}
		}
		return
	}
	if y < float64(g.HeaderHeight()) {
		if y < 44 {
			switch {
			case x >= float64(g.width-56):
				g.Settings = true
			case x < 64 && g.Selected >= 0:
				g.Back()
			}
		} else if y < menu.Header {
			if x >= 16 && x <= 84 {
				g.Cycle(-1)
			} else if x >= float64(g.width-84) && x <= float64(g.width-16) {
				g.Cycle(1)
			} else if g.Selected < 0 {
				for i := 0; i < 3; i++ {
					if math.Abs(x-(float64(g.width)/2+float64(i-1)*52)) <= 22 {
						g.SetTab(i)
					}
				}
			}
		}

		return
	}
	if g.Selected >= 0 {
		g.playTap(x, y)
		return
	}
	if g.Selected < 0 {
		if index := g.TileAt(x, y, g.width); index >= 0 {
			g.Open(index)
		}
	}

}
func (g *Game) cardBounds() (float64, float64) {
	w := math.Max(1, float64(g.width)-32)
	return (float64(g.width) - w) / 2, w
}
func (g *Game) Layout(w, h int) (int, int) {
	if w < 1 || h < 1 {
		return g.width, g.height
	}
	scale := math.Max(1, float64(w)/800)
	g.width = int(float64(w) / scale)
	g.height = int(float64(h) / scale)
	g.RowHeight = int(math.Ceil(menu.TileWidth(g.width))) + 14
	g.Clamp(g.height)
	return g.width, g.height
}

func rect(dst *ebiten.Image, x, y, w, h float64, c color.Color) {
	vector.DrawFilledRect(dst, float32(x), float32(y), float32(w), float32(h), c, false)
}
func line(dst *ebiten.Image, x, y, x2, y2, width float64, c color.Color) {
	vector.StrokeLine(dst, float32(x), float32(y), float32(x2), float32(y2), float32(width), c, true)
}
func dot(dst *ebiten.Image, x, y, r float64, c color.Color) {
	vector.DrawFilledCircle(dst, float32(x), float32(y), float32(r), c, true)
}
func label(dst *ebiten.Image, s string, x, y int, c color.Color) {
	text.Draw(dst, s, basicfont.Face7x13, x, y, c)
}

// ActivateMenu applies the same hit testing as a pointer release in the embedded menu.
func (g *Game) ActivateMenu(x, y float64) { g.activate(x, y) }
