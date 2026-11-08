package app

import (
	"ebe-boilerplate/internal/play"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"math"
)

func (g *Game) expansionAction(slot int) bool {
	if g.restartConfirm {
		return false
	}
	if slot == 2 {
		switch g.Selected {
		case play.DriftID:
			if g.games.Drift.Cleared {
				g.games.Drift.NextWave()
			} else {
				g.games.Drift.Thrust = !g.games.Drift.Thrust
			}
		case play.StackID:
			if !g.paused {
				g.games.Stack.Drop()
			}
		case play.DeductionID:
			g.games.Deduction.Undo()
		case play.ClusterID:
			if !g.paused {
				g.games.Cluster.Fire()
			}
		case play.WaveID:
			g.games.Wave.NextWave()
		case play.CrossingID:
			g.games.Crossing.NextLevel()
		case play.ShiftID:
			g.games.Shift.Undo()
		}
		return true
	}
	if slot == 3 {
		switch g.Selected {
		case play.StackID:
			g.games.Stack.Pace = (g.games.Stack.Pace + 1) % 3
			g.games.Stack.Clock = 0
		case play.DeductionID:
			g.games.Deduction.Set(g.cursor, 0)
		case play.ShiftID:
			g.games.Shift.Replay()
		}
		return true
	}
	return false
}
func (g *Game) expansionDirection(d int) {
	switch g.Selected {
	case play.DriftID:
		if d == 0 {
			g.games.Drift.Turn(-.12)
		}
		if d == 2 {
			g.games.Drift.Turn(.12)
		}
		if d == 3 {
			g.games.Drift.VX *= .8
			g.games.Drift.VY *= .8
		}
	case play.StackID:
		s := &g.games.Stack
		switch d {
		case 0:
			s.Move(-1)
		case 1:
			s.Rotate()
		case 2:
			s.Move(1)
		case 3:
			s.Down()
		}
	case play.DeductionID:
		x, y := g.cursor%9, g.cursor/9
		x += [4]int{-1, 0, 1, 0}[d]
		y += [4]int{0, -1, 0, 1}[d]
		g.cursor = max(0, min(8, y))*9 + max(0, min(8, x))
	case play.ClusterID:
		delta := .08
		if d < 2 {
			delta = -delta
		}
		g.games.Cluster.SetAim(g.games.Cluster.Aim + delta)
	case play.WaveID:
		if d == 0 {
			g.games.Wave.Move(g.games.Wave.Player - .05)
		}
		if d == 2 {
			g.games.Wave.Move(g.games.Wave.Player + .05)
		}
	case play.CrossingID:
		g.games.Crossing.Move(d)
	case play.ShiftID:
		g.games.Shift.Direction(d)
	}
}
func (g *Game) expansionConfirm() {
	switch g.Selected {
	case play.DriftID:
		g.games.Drift.Thrust = !g.games.Drift.Thrust
	case play.StackID:
		g.games.Stack.Rotate()
	case play.DeductionID:
		g.games.Deduction.Set(g.cursor, (g.games.Deduction.Cells[g.cursor]+1)%10)
	case play.ClusterID:
		g.games.Cluster.Fire()
	case play.WaveID:
		g.games.Wave.NextWave()
	case play.CrossingID:
		g.games.Crossing.NextLevel()
	}
}
func (g *Game) expansionTap(x, y float64) {
	bx, by, bw, bh := g.gameRect()
	u, v := (x-bx)/bw, (y-by)/bh
	if g.Selected == play.DeductionID {
		for digit := 1; digit <= 9; digit++ {
			rx, ry, rw, rh := g.deductionDigitRect(digit)
			if x >= rx && x < rx+rw && y >= ry && y < ry+rh {
				g.games.Deduction.Set(g.cursor, digit)
				return
			}
		}
	}
	if u < 0 || u >= 1 || v < 0 || v >= 1 {
		return
	}
	switch g.Selected {
	case play.DriftID:
		g.games.Drift.Angle = math.Atan2(v-g.games.Drift.Y, u-g.games.Drift.X)
	case play.StackID:
		g.games.Stack.Rotate()
	case play.DeductionID:
		g.cursor = int(v*9)*9 + int(u*9)
	case play.ClusterID:
		g.games.Cluster.SetAim(math.Atan2(u-.5, .93-v))
		g.games.Cluster.Fire()
	case play.WaveID:
		g.games.Wave.Move(u)
	case play.CrossingID:
		dx, dy := u*9-g.games.Crossing.X, v*11-float64(g.games.Crossing.Y)-.5
		d := 3
		if math.Abs(dx) > math.Abs(dy) {
			d = 2
			if dx < 0 {
				d = 0
			}
		} else if dy < 0 {
			d = 1
		}
		g.games.Crossing.Move(d)
	case play.ShiftID:
		g.games.Shift.Move(int(v*4)*4 + int(u*4))
	}
}
func (g *Game) expansionDrag(x, y float64) {
	if g.paused || g.travel < 12 {
		return
	}
	_, _, bw, _ := g.gameRect()
	switch g.Selected {
	case play.DriftID:
		dx, dy := x-g.lx, y-g.ly
		if math.Hypot(dx, dy) > 1 {
			g.games.Drift.Angle = math.Atan2(dy, dx)
			g.games.Drift.Thrust = true
		}
	case play.WaveID:
		delta := x - g.lx
		if math.Hypot(g.lx-g.sx, g.ly-g.sy) < 12 {
			delta = x - g.sx
		}
		g.games.Wave.Move(g.games.Wave.Player + delta/bw)
	case play.ClusterID:
		g.games.Cluster.SetAim(g.games.Cluster.Aim + (x-g.lx)*2/bw)
	}
}
func (g *Game) expansionSwipe(x, y float64) {
	if g.setup || g.paused || g.restartConfirm {
		return
	}
	switch g.Selected {
	case play.DriftID:
		g.games.Drift.Thrust = false
	case play.ClusterID:
		g.games.Cluster.Fire()
	case play.StackID, play.CrossingID, play.ShiftID:
		dx, dy := x-g.sx, y-g.sy
		d := 3
		if math.Abs(dx) > math.Abs(dy) {
			d = 2
			if dx < 0 {
				d = 0
			}
		} else if dy < 0 {
			d = 1
		}
		g.expansionDirection(d)
	}
	g.saveGames()
}
func heldDirection() (left, right, up, down bool) {
	left = ebiten.IsKeyPressed(ebiten.KeyArrowLeft)
	right = ebiten.IsKeyPressed(ebiten.KeyArrowRight)
	up = ebiten.IsKeyPressed(ebiten.KeyArrowUp)
	down = ebiten.IsKeyPressed(ebiten.KeyArrowDown)
	for _, id := range ebiten.AppendGamepadIDs(nil) {
		if !ebiten.IsStandardGamepadLayoutAvailable(id) {
			continue
		}
		left = left || ebiten.IsStandardGamepadButtonPressed(id, ebiten.StandardGamepadButtonLeftLeft)
		right = right || ebiten.IsStandardGamepadButtonPressed(id, ebiten.StandardGamepadButtonLeftRight)
		up = up || ebiten.IsStandardGamepadButtonPressed(id, ebiten.StandardGamepadButtonLeftTop)
		down = down || ebiten.IsStandardGamepadButtonPressed(id, ebiten.StandardGamepadButtonLeftBottom)
	}
	return
}
func (g *Game) updateExpansion() {
	switch g.Selected {
	case play.DriftID:
		d := &g.games.Drift
		left, right, up, down := heldDirection()
		if left {
			d.Turn(-.045)
		}
		if right {
			d.Turn(.045)
		}
		if down {
			d.VX *= .98
			d.VY *= .98
		}
		thrust := d.Thrust
		d.Thrust = d.Thrust || up
		d.Tick()
		d.Thrust = thrust
	case play.StackID:
		g.games.Stack.Tick()
	case play.ClusterID:
		g.games.Cluster.Tick()
	case play.WaveID:
		left, right, _, _ := heldDirection()
		if left {
			g.games.Wave.Move(g.games.Wave.Player - .012)
		}
		if right {
			g.games.Wave.Move(g.games.Wave.Player + .012)
		}
		g.games.Wave.Tick()
	case play.CrossingID:
		g.games.Crossing.Tick()
	}
}
func (g *Game) deductionKeys() {
	for i, key := range []ebiten.Key{ebiten.Key1, ebiten.Key2, ebiten.Key3, ebiten.Key4, ebiten.Key5, ebiten.Key6, ebiten.Key7, ebiten.Key8, ebiten.Key9} {
		if inpututil.IsKeyJustPressed(key) {
			g.games.Deduction.Set(g.cursor, i+1)
			g.saveGames()
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) || inpututil.IsKeyJustPressed(ebiten.KeyDelete) || inpututil.IsKeyJustPressed(ebiten.Key0) {
		g.games.Deduction.Set(g.cursor, 0)
		g.saveGames()
	}
}
func (g *Game) deductionDigitRect(digit int) (x, y, w, h float64) {
	bx, by, bw, bh := g.gameRect()
	w = bw / 9
	h = 40
	x = bx + float64(digit-1)*w
	y = by + bh + 36
	return
}
func (g *Game) expansionFooter(names []string) {
	switch g.Selected {
	case play.DriftID:
		if g.games.Drift.Cleared {
			if g.games.Drift.Level < 3 {
				names[2] = "Next"
			}
		} else {
			names[2] = "Thrust"
			if g.games.Drift.Thrust {
				names[2] = "Coast"
			}
		}
	case play.StackID:
		names[2] = "Drop"
		names[3] = []string{"Slow", "Medium", "Fast"}[g.games.Stack.Pace]
	case play.DeductionID:
		names[2] = "Undo"
		names[3] = "Erase"
	case play.ClusterID:
		names[2] = "Fire"
	case play.WaveID:
		if g.games.Wave.Cleared && g.games.Wave.Level < 3 {
			names[2] = "Next"
		}
	case play.CrossingID:
		if g.games.Crossing.Won && g.games.Crossing.Level < 3 {
			names[2] = "Next"
		}
	case play.ShiftID:
		names[2] = "Undo"
		names[3] = "Replay"
	}
}
func (g *Game) expansionStatus() string {
	switch g.Selected {
	case play.DriftID:
		d := g.games.Drift
		if d.Over {
			return "Round over"
		}
		if d.Cleared {
			return "Wave clear!"
		}
		return fmt.Sprintf("%d / 3     %d", d.Level, d.Lives)
	case play.StackID:
		s := g.games.Stack
		if s.Won {
			return "40!"
		}
		if s.Over {
			return "No room"
		}
		return fmt.Sprintf("%d / 40", s.Lines)
	case play.DeductionID:
		if g.games.Deduction.Won {
			return "Complete!"
		}
		return ""
	case play.ClusterID:
		c := g.games.Cluster
		if c.Won {
			return "Clear!"
		}
		if c.Over {
			return "No room"
		}
		return fmt.Sprint(c.Score)
	case play.WaveID:
		w := g.games.Wave
		if w.Over {
			return "Round over"
		}
		if w.Cleared {
			return "Wave clear!"
		}
		return fmt.Sprintf("%d / 3     %d", w.Level, w.Lives)
	case play.CrossingID:
		c := g.games.Crossing
		if c.Over {
			return "Round over"
		}
		if c.Won {
			return "Across!"
		}
		return fmt.Sprintf("%d / 3     %d", c.Level, c.Lives)
	case play.ShiftID:
		if g.games.Shift.Won {
			return "Complete!"
		}
		return fmt.Sprint(g.games.Shift.Moves)
	}
	return ""
}
