package app

import (
	"ebe-boilerplate/internal/play"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

func (g *Game) additionAction(slot int) bool {
	if g.restartConfirm {
		if g.Selected == play.BloomID && slot == 0 {
			g.games.Bloom.Clear()
			g.paused = true
			g.restartConfirm = false
			return true
		}
		return false
	}
	switch slot {
	case 0:
		if g.Selected == play.BrickID {
			if !g.paused && g.games.Brick.Running {
				g.pauseGame()
			} else {
				g.paused = false
				g.games.Brick.Serve()
			}
			return true
		}
	case 2:
		switch g.Selected {
		case play.FlowID:
			g.games.Flow.Replay()
			g.paused = false
		case play.CalibanID:
			g.games.Caliban.Next()
		case play.BloomID:
			g.paused = true
			g.games.Bloom.Step()
		case play.EchoID:
			g.games.Echo.Retry()
			g.paused = false
		case play.BrickID:
			if g.games.Brick.Won && g.games.Brick.Level < 2 {
				g.games.Brick = play.NewBrick(g.games.Brick.Level + 1)
				g.paused = false
			}
		case play.PlanetsID:
			if g.games.Planets.Winner == 1 && g.games.Planets.System < 2 {
				g.games.Planets = play.NewPlanets(g.games.Planets.System + 1)
				g.paused = false
			}
		}
		return true
	case 3:
		if g.Selected == play.BloomID {
			g.games.Bloom.Speed = (g.games.Bloom.Speed + 1) % 3
			g.games.Bloom.Clock = 0
		}
		return true
	}
	return false
}
func (g *Game) additionDirection(d int) {
	n := 0
	switch g.Selected {
	case play.FlowID:
		n = 5
	case play.BloomID:
		n = play.BloomN
	case play.EchoID:
		n = 2
	case play.CalibanID:
		delta := 1
		if d < 2 {
			delta = -1
		}
		g.games.Caliban.Move(delta)
	case play.BrickID:
		if d == 0 {
			g.games.Brick.Move(g.games.Brick.Paddle - .055)
		}
		if d == 2 {
			g.games.Brick.Move(g.games.Brick.Paddle + .055)
		}
	case play.PlanetsID:
		delta := 1
		if d < 2 {
			delta = -1
		}
		g.cursor = max(0, min(len(g.games.Planets.Bodies)-1, g.cursor+delta))
	}
	if n > 0 {
		x, y := g.cursor%n, g.cursor/n
		switch d {
		case 0:
			x--
		case 1:
			y--
		case 2:
			x++
		case 3:
			y++
		}
		g.cursor = max(0, min(n-1, y))*n + max(0, min(n-1, x))
	}
}
func (g *Game) additionConfirm() {
	switch g.Selected {
	case play.FlowID:
		g.games.Flow.Rotate(g.cursor)
	case play.EchoID:
		g.games.Echo.Press(g.cursor)
	case play.BrickID:
		g.games.Brick.Serve()
	case play.PlanetsID:
		g.games.Planets.Choose(g.cursor)
	}
}
func (g *Game) additionTap(x, y float64) {
	if g.restartConfirm || g.paused && g.Selected != play.BloomID {
		return
	}
	bx, by, bw, bh := g.gameRect()
	u, v := (x-bx)/bw, (y-by)/bh
	if u < 0 || u >= 1 || v < 0 || v >= 1 {
		return
	}
	switch g.Selected {
	case play.FlowID:
		g.cursor = int(v*5)*5 + int(u*5)
		g.games.Flow.Rotate(g.cursor)
	case play.BloomID:
		if g.paused {
			g.cursor = int(v*play.BloomN)*play.BloomN + int(u*play.BloomN)
			g.games.Bloom.Toggle(g.cursor)
		}
	case play.EchoID:
		g.cursor = int(v*2)*2 + int(u*2)
		g.games.Echo.Press(g.cursor)
	case play.CalibanID:
		g.games.Caliban.Lane = laneAt(u, v)
	case play.PlanetsID:
		i := g.games.Planets.At(u, v)
		if i >= 0 {
			g.cursor = i
			g.games.Planets.Choose(i)
		}
	}
	g.saveGames()
}
func laneAt(x, y float64) int {
	angle := math.Atan2(y-.5, x-.5) + math.Pi/2
	return (int(math.Round(angle*play.CalibanLanes/(2*math.Pi))) + play.CalibanLanes) % play.CalibanLanes
}
func (g *Game) updateAddition() {
	switch g.Selected {
	case play.CalibanID:
		g.games.Caliban.Tick()
	case play.BloomID:
		g.games.Bloom.Tick()
	case play.EchoID:
		g.games.Echo.Tick()
	case play.BrickID:
		left, right := ebiten.IsKeyPressed(ebiten.KeyArrowLeft), ebiten.IsKeyPressed(ebiten.KeyArrowRight)
		for _, id := range ebiten.AppendGamepadIDs(nil) {
			if ebiten.IsStandardGamepadLayoutAvailable(id) {
				left = left || ebiten.IsStandardGamepadButtonPressed(id, ebiten.StandardGamepadButtonLeftLeft)
				right = right || ebiten.IsStandardGamepadButtonPressed(id, ebiten.StandardGamepadButtonLeftRight)
			}
		}
		if left {
			g.games.Brick.Move(g.games.Brick.Paddle - .012)
		}
		if right {
			g.games.Brick.Move(g.games.Brick.Paddle + .012)
		}
		g.games.Brick.Tick()
	case play.PlanetsID:
		g.games.Planets.Tick()
	}
}
func (g *Game) additionStatus() string {
	switch g.Selected {
	case play.FlowID:
		if g.games.Flow.Won {
			return "Connected!"
		}
		return fmt.Sprintf("%d", g.games.Flow.Moves)
	case play.CalibanID:
		c := g.games.Caliban
		if c.Over {
			return "Round over"
		}
		if c.Cleared {
			return "Wave clear / Next"
		}
		return fmt.Sprintf("%d / %d", c.Wave, c.Lives)
	case play.BloomID:
		return fmt.Sprintf("%d", g.games.Bloom.Generation)
	case play.EchoID:
		e := g.games.Echo
		switch e.Phase {
		case play.EchoShow:
			return "Watch"
		case play.EchoInput:
			return "Repeat"
		case play.EchoSuccess:
			return "Yes!"
		case play.EchoFailed:
			return "Retry"
		case play.EchoWon:
			return "32!"
		}
	case play.BrickID:
		b := g.games.Brick
		if b.Won {
			return "Board clear!"
		}
		if b.Over {
			return "Round over"
		}
		if !b.Running {
			return "Play to serve"
		}
		return fmt.Sprintf("%d / %d", b.Level+1, b.Lives)
	case play.PlanetsID:
		if g.games.Planets.Winner == 1 {
			return "System yours!"
		}
		if g.games.Planets.Winner > 1 {
			return "Opponent wins"
		}
		return "Source / destination"
	}
	return ""
}
func (g *Game) additionFooter(names []string) {
	switch g.Selected {
	case play.FlowID:
		names[2] = "Replay"
	case play.CalibanID:
		if g.games.Caliban.Cleared {
			names[2] = "Next"
		}
	case play.BloomID:
		names[1] = "Clear"
		names[2] = "Step"
		names[3] = fmt.Sprintf("%d / s", [3]int{2, 6, 12}[g.games.Bloom.Speed])
	case play.EchoID:
		names[2] = "Retry"
	case play.BrickID:
		if !g.games.Brick.Running {
			names[0] = "Play"
		}
		if g.games.Brick.Won && g.games.Brick.Level < 2 {
			names[2] = "Next"
		}
	case play.PlanetsID:
		if g.games.Planets.Winner == 1 && g.games.Planets.System < 2 {
			names[2] = "Next"
		}
	}
}

func (g *Game) additionDrag(x, y float64) {
	if g.Settings || g.setup || g.restartConfirm {
		return
	}
	bx, by, bw, bh := g.gameRect()
	if g.Selected == play.BrickID && !g.paused && g.gesture == "horizontal" {
		delta := x - g.lx
		if math.Hypot(g.lx-g.sx, g.ly-g.sy) < 12 {
			delta = x - g.sx
		}
		g.games.Brick.Move(g.games.Brick.Paddle + delta/bw)
	}
	if g.Selected == play.CalibanID && !g.paused && x >= bx && x <= bx+bw && y >= by && y <= by+bh {
		g.games.Caliban.Lane = laneAt((x-bx)/bw, (y-by)/bh)
	}
	if g.Selected == play.BloomID && g.paused && g.travel >= 12 {
		// Paint the stroke with one value, including cells crossed between frames.
		steps := max(1, int(math.Hypot(x-g.lx, y-g.ly)/(bw/play.BloomN/2)))
		for i := 0; i <= steps; i++ {
			u := (g.lx + (x-g.lx)*float64(i)/float64(steps) - bx) / bw
			v := (g.ly + (y-g.ly)*float64(i)/float64(steps) - by) / bh
			if u >= 0 && u < 1 && v >= 0 && v < 1 {
				g.games.Bloom.Cells[int(v*play.BloomN)*play.BloomN+int(u*play.BloomN)] = g.bloomPaint
			}
		}
	}
}
