package app

import (
	"ebe-boilerplate/internal/menu"
	"ebe-boilerplate/internal/play"
	"ebe-boilerplate/internal/presentation"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"math"
)

func (g *Game) renderAddition(dst *ebiten.Image, c *play.Collection, id int, x, y, w, h float64, p palette, active bool) {
	if id >= play.DriftID {
		g.renderExpansion(dst, c, id, x, y, w, h, p, active)
		return
	}
	switch id {
	case play.FlowID:
		f := &c.Flow
		connected := f.Connected()
		cell := w / 5
		for i, ports := range f.Cells {
			cx, cy := x+(float64(i%5)+.5)*cell, y+(float64(i/5)+.5)*cell
			ink := p.muted
			if connected[i] {
				ink = p.a
			}
			for d := 0; d < 4; d++ {
				if ports&(1<<d) != 0 {
					line(dst, cx, cy, cx+float64([4]int{0, 1, 0, -1}[d])*cell*.5, cy+float64([4]int{-1, 0, 1, 0}[d])*cell*.5, math.Max(2, cell*.07), ink)
				}
			}
			if connected[i] {
				dot(dst, cx, cy, cell*.12, ink)
			} else {
				vector.StrokeCircle(dst, float32(cx), float32(cy), float32(cell*.12), 1.5, ink, true)
			}
			if i == 12 {
				vector.StrokeCircle(dst, float32(cx), float32(cy), float32(cell*.2), 2, p.ink, true)
			}
			if active && g.cursor == i {
				vector.StrokeRect(dst, float32(cx-cell*.44), float32(cy-cell*.44), float32(cell*.88), float32(cell*.88), 1, p.ink, true)
			}
		}
	case play.CalibanID:
		a := &c.Caliban
		pos := func(lane int, depth float64) (float64, float64) {
			angle := float64(lane)*2*math.Pi/play.CalibanLanes - math.Pi/2
			r := w * (.08 + depth*.36)
			return x + w/2 + math.Cos(angle)*r, y + h/2 + math.Sin(angle)*r
		}
		for lane := 0; lane < play.CalibanLanes; lane++ {
			ax, ay := pos(lane, 0)
			bx, by := pos(lane, 1)
			cx, cy := pos((lane+1)%play.CalibanLanes, 1)
			line(dst, ax, ay, bx, by, 1, p.edge)
			line(dst, bx, by, cx, cy, 2, p.edge)
		}
		px, py := pos(a.Lane, 1)
		triangle(dst, px-w*.025, py-w*.025, w*.05, p.a)
		for _, e := range a.Enemies {
			px, py := pos(e.Lane, e.Depth)
			g.gameBox(dst, px-w*.016, py-w*.016, w*.032, w*.032, p.b)
		}
		for _, s := range a.Shots {
			px, py := pos(s.Lane, s.Depth)
			dot(dst, px, py, math.Max(1.5, w*.006), p.ink)
		}
		for i := 0; i < a.Lives; i++ {
			dot(dst, x+w*.46+float64(i)*w*.04, y+h*.5, w*.012, p.ink)
		}
	case play.BloomID:
		cell := w / play.BloomN
		for i, live := range c.Bloom.Cells {
			cx, cy := x+float64(i%play.BloomN)*cell, y+float64(i/play.BloomN)*cell
			if live {
				g.gameBox(dst, cx+.5, cy+.5, cell-1, cell-1, p.a)
			} else if active {
				dot(dst, cx+cell/2, cy+cell/2, .6, p.edge)
			}
		}
		if active && g.paused {
			vector.StrokeRect(dst, float32(x+float64(g.cursor%play.BloomN)*cell), float32(y+float64(g.cursor/play.BloomN)*cell), float32(cell), float32(cell), 1, p.ink, true)
		}
	case play.EchoID:
		e := &c.Echo
		size := w * .44
		for i := 0; i < 4; i++ {
			px, py := x+w*.03+float64(i%2)*w*.5, y+h*.03+float64(i/2)*h*.5
			ink := p.edge
			if e.Cue() == i {
				ink = p.a
			}
			g.gameBox(dst, px, py, size, size, ink)
			cx, cy := px+size/2, py+size*.43
			r := size * .16
			switch i {
			case 0:
				vector.StrokeCircle(dst, float32(cx), float32(cy), float32(r), 2, p.ink, true)
			case 1:
				vector.StrokeRect(dst, float32(cx-r), float32(cy-r), float32(r*2), float32(r*2), 2, p.ink, true)
			case 2:
				triangle(dst, cx-r, cy-r, r*2, p.ink)
			case 3:
				line(dst, cx-r, cy-r, cx+r, cy+r, 2, p.ink)
				line(dst, cx-r, cy+r, cx+r, cy-r, 2, p.ink)
			}
			drawUnitCount(dst, fmt.Sprint(i+1), cx, py+size*.8, size*.08, p.ink)
			if active && g.cursor == i {
				vector.StrokeRect(dst, float32(px+3), float32(py+3), float32(size-6), float32(size-6), 1, p.ink, true)
			}
		}
		for i := range e.Sequence {
			ink := p.edge
			if i < e.Input {
				ink = p.a
			}
			dot(dst, x+w*(.04+.029*float64(i)), y+h*.49, math.Max(1, w*.006), ink)
		}
	case play.BrickID:
		b := &c.Brick
		for i, live := range b.Bricks {
			if !live {
				continue
			}
			bx, by, bw, bh := play.BrickRect(i)
			ink := p.a
			if i/8%2 == 1 {
				ink = p.b
			}
			g.gameBox(dst, x+bx*w, y+by*h, bw*w, bh*h, ink)
		}
		g.gameBox(dst, x+(b.Paddle-.1)*w, y+.89*h, .2*w, .025*h, p.a)
		dot(dst, x+b.X*w, y+b.Y*h, math.Max(2, w*.012), p.ink)
		for i := 0; i < b.Lives; i++ {
			dot(dst, x+w*.45+float64(i)*w*.05, y+h*.05, w*.012, p.ink)
		}
	case play.PlanetsID:
		g.drawPlanets(dst, &c.Planets, x, y, w, h, p, active)
	}
}
func (g *Game) drawPlanets(dst *ebiten.Image, s *play.Planets, x, y, w, h float64, p palette, active bool) {
	dot(dst, x+w/2, y+h/2, w*.025, p.b)
	for _, b := range s.Bodies {
		center := play.Point{X: .5, Y: .5}
		if b.Parent >= 0 {
			center = s.Position(b.Parent)
		}
		vector.StrokeCircle(dst, float32(x+center.X*w), float32(y+center.Y*h), float32(b.Radius*w), 1, p.edge, true)
	}
	for _, f := range s.Fleets {
		pos := s.Position(f.Target)
		ink := ownerColor(f.Owner, p)
		if g.Theme == menu.Wireframe {
			ink = p.ink
		}
		line(dst, x+f.X*w, y+f.Y*h, x+pos.X*w, y+pos.Y*h, 1, p.edge)
		dot(dst, x+f.X*w, y+f.Y*h, math.Max(2, w*.009), ink)
	}
	for i, b := range s.Bodies {
		pos := s.Position(i)
		px, py := x+pos.X*w, y+pos.Y*h
		r := w * .027
		if b.Parent >= 0 {
			r = w * .019
		}
		r = math.Max(3, r)
		ink := ownerColor(b.Owner, p)
		if g.Theme == menu.Wireframe {
			ink = p.ink
		}
		dot(dst, px, py, r+2, p.panel)
		dot(dst, px, py, r, ink)
		vector.StrokeCircle(dst, float32(px), float32(py), float32(r+1), 1, p.ink, true)
		drawUnitCount(dst, fmt.Sprint(b.Units), px, py, r*.82, presentation.TileInk(ink))
		if b.Owner > 0 {
			drawUnitCount(dst, fmt.Sprint(b.Owner), px, py-r-5, math.Max(2, r*.35), p.ink)
		}
		if active && (i == s.Source || i == g.cursor) {
			vector.StrokeCircle(dst, float32(px), float32(py), float32(r+5), 2, p.ink, true)
		}
	}
}
func (g *Game) planetChoiceRect(i int) (x, y, w, h float64) {
	w = math.Min(float64(g.width)-32, (float64(g.height)-menu.Header-36)/3)
	h = w
	x = (float64(g.width) - w) / 2
	y = menu.Header + 8 + float64(i)*(h+8)
	return
}
func (g *Game) drawPlanetChoices(dst *ebiten.Image, p palette) {
	for i := 0; i < 3; i++ {
		x, y, w, h := g.planetChoiceRect(i)
		s := play.NewPlanets(i)
		rect(dst, x, y, w, h, p.panel)
		g.drawPlanets(dst, &s, x, y, w, h, p, false)
		drawUnitCount(dst, fmt.Sprint(i+1), x+12, y+12, 8, p.ink)
		if i == g.option {
			vector.StrokeRect(dst, float32(x), float32(y), float32(w), float32(h), 2, p.a, true)
		}
	}
}
