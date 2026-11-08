package app

import (
	"ebe-boilerplate/internal/menu"
	"ebe-boilerplate/internal/play"
	"ebe-boilerplate/internal/presentation"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"math"
)

func (g *Game) renderExpansion(dst *ebiten.Image, c *play.Collection, id int, x, y, w, h float64, p palette, active bool) {
	switch id {
	case play.DriftID:
		d := &c.Drift
		for _, rock := range d.Rocks {
			r := play.RockRadius(rock.Size) * w
			cx, cy := x+rock.X*w, y+rock.Y*h
			for j := 0; j < 8; j++ {
				a, b := float64(j)*math.Pi/4, float64(j+1)*math.Pi/4
				ra, rb := r, r
				if j%2 == 0 {
					ra *= .8
				} else {
					rb *= .8
				}
				line(dst, cx+math.Cos(a)*ra, cy+math.Sin(a)*ra, cx+math.Cos(b)*rb, cy+math.Sin(b)*rb, 1.5, p.b)
			}
		}
		cx, cy := x+d.X*w, y+d.Y*h
		r := w * .022
		angles := [3]float64{d.Angle, d.Angle + 2.5, d.Angle - 2.5}
		for i, a := range angles {
			b := angles[(i+1)%3]
			line(dst, cx+math.Cos(a)*r, cy+math.Sin(a)*r, cx+math.Cos(b)*r, cy+math.Sin(b)*r, 2, p.a)
		}
		if d.Thrust {
			line(dst, cx-math.Cos(d.Angle)*r, cy-math.Sin(d.Angle)*r, cx-math.Cos(d.Angle)*r*1.8, cy-math.Sin(d.Angle)*r*1.8, 2, p.b)
		}
		if d.Shield > 0 {
			vector.StrokeCircle(dst, float32(cx), float32(cy), float32(r*1.6), 1, p.edge, true)
		}
		for _, s := range d.Shots {
			dot(dst, x+s.X*w, y+s.Y*h, math.Max(1.5, w*.004), p.ink)
		}
	case play.StackID:
		s := &c.Stack
		cw, ch := w/play.StackW, h/play.StackH
		for i, v := range s.Cells {
			cx, cy := x+float64(i%play.StackW)*cw, y+float64(i/play.StackW)*ch
			if v == 0 {
				vector.StrokeRect(dst, float32(cx), float32(cy), float32(cw), float32(ch), 1, p.edge, true)
				continue
			}
			ink := p.a
			if v%2 == 0 {
				ink = p.b
			}
			g.gameBox(dst, cx+1, cy+1, cw-2, ch-2, ink)
		}
		if !s.Over && !s.Won {
			ghost := s.Y
			for s.Fits(s.X, ghost+1, s.Rotation) {
				ghost++
			}
			for _, cell := range play.StackShape(s.Piece, s.Rotation) {
				cx := x + float64(s.X+cell.X)*cw
				gy := y + float64(ghost+cell.Y)*ch
				vector.StrokeRect(dst, float32(cx+2), float32(gy+2), float32(cw-4), float32(ch-4), 1, p.muted, true)
				g.gameBox(dst, cx+1, y+float64(s.Y+cell.Y)*ch+1, cw-2, ch-2, p.ink)
			}
		}
	case play.DeductionID:
		d := &c.Deduction
		cell := w / 9
		for i, v := range d.Cells {
			cx, cy := x+float64(i%9)*cell, y+float64(i/9)*cell
			if d.Clues[i] != 0 {
				rect(dst, cx+1, cy+1, cell-2, cell-2, p.edge)
			}
			ink := p.a
			if d.Clues[i] != 0 {
				ink = p.ink
			}
			if v != 0 {
				boardNumber(dst, v, cx+cell/2, cy+cell/2, math.Min(24, cell*.44), ink)
			}
			if d.Conflict(i) {
				line(dst, cx+cell-9, cy+3, cx+cell-3, cy+9, 1, p.b)
				line(dst, cx+cell-3, cy+3, cx+cell-9, cy+9, 1, p.b)
			}
			if active && g.cursor == i {
				vector.StrokeRect(dst, float32(cx+2), float32(cy+2), float32(cell-4), float32(cell-4), 2, p.a, true)
			}
		}
		for i := 0; i <= 9; i++ {
			thick := .7
			if i%3 == 0 {
				thick = 2
			}
			line(dst, x+float64(i)*cell, y, x+float64(i)*cell, y+h, thick, p.muted)
			line(dst, x, y+float64(i)*cell, x+w, y+float64(i)*cell, thick, p.muted)
		}
	case play.ClusterID:
		v := &c.Cluster
		for i, color := range v.Cells {
			if color != 0 {
				pos := play.ClusterPosition(i)
				g.clusterBall(dst, color, x+pos.X*w, y+pos.Y*h, w*.05, p)
			}
		}
		g.clusterBall(dst, v.Next, x+w*.85, y+h*.94, w*.035, p)
		if !v.Flying {
			line(dst, x+w*.5, y+h*.93, x+w*(.5+.18*math.Sin(v.Aim)), y+h*(.93-.18*math.Cos(v.Aim)), 2, p.ink)
		}
		g.clusterBall(dst, v.Color, x+v.X*w, y+v.Y*h, w*.05, p)
	case play.WaveID:
		v := &c.Wave
		for i, alive := range v.Enemies {
			if !alive {
				continue
			}
			pos := v.Position(i)
			cx, cy := x+pos.X*w, y+pos.Y*h
			ink := p.a
			if i/8%2 == 0 {
				ink = p.b
			}
			g.gameBox(dst, cx-w*.025, cy-h*.018, w*.05, h*.032, ink)
			line(dst, cx-w*.03, cy+h*.026, cx-w*.015, cy+h*.005, 1.5, ink)
			line(dst, cx+w*.03, cy+h*.026, cx+w*.015, cy+h*.005, 1.5, ink)
			dot(dst, cx-w*.01, cy, w*.004, p.panel)
			dot(dst, cx+w*.01, cy, w*.004, p.panel)
		}
		uprightTriangle(dst, x+v.Player*w, y+h*.9, w*.06, p.ink)
		if v.Shield > 0 {
			vector.StrokeCircle(dst, float32(x+v.Player*w), float32(y+h*.9), float32(w*.045), 1, p.edge, true)
		}
		for _, s := range v.Shots {
			ink := p.a
			if s.Enemy {
				ink = p.b
			}
			line(dst, x+s.X*w, y+s.Y*h, x+s.X*w, y+(s.Y+.015)*h, 2, ink)
		}
	case play.CrossingID:
		v := &c.Crossing
		cw, ch := w/9, h/11
		for row := 0; row < 11; row++ {
			cy := y + float64(row)*ch
			if row == 0 || row == 6 || row == 10 {
				rect(dst, x, cy, w, ch, p.edge)
				continue
			}
			line(dst, x, cy, x+w, cy, 1, p.edge)
			size := 1.05
			ink := p.b
			if row <= 5 {
				size = 2.1
				ink = p.a
			}
			for k := -1; k <= 3; k++ {
				start := v.Phases[row] + float64(k)*3
				left, right := math.Max(0, start), math.Min(9, start+size)
				if left >= right {
					continue
				}
				g.gameBox(dst, x+left*cw, cy+ch*.2, (right-left)*cw, ch*.6, ink)
				if row >= 7 {
					for _, wx := range []float64{start + .2, start + .8} {
						if wx > 0 && wx < 9 {
							dot(dst, x+wx*cw, cy+ch*.8, ch*.08, p.ink)
						}
					}
				}
			}
		}
		for i, home := range v.Homes {
			cx := x + (1.5+float64(i)*3)*cw
			vector.StrokeRect(dst, float32(cx-cw*.4), float32(y+ch*.1), float32(cw*.8), float32(ch*.8), 2, p.ink, true)
			if home {
				dot(dst, cx, y+ch/2, ch*.25, p.a)
			}
		}
		cx, cy := x+v.X*cw, y+(float64(v.Y)+.5)*ch
		dot(dst, cx, cy, ch*.33, p.panel)
		uprightTriangle(dst, cx, cy, ch*.5, p.ink)
	case play.ShiftID:
		s := &c.Shift
		cell := w / 4
		for i, v := range s.Cells {
			if v == 0 {
				continue
			}
			cx, cy := x+float64(i%4)*cell, y+float64(i/4)*cell
			ink := p.edge
			if v == i+1 {
				ink = p.a
			}
			g.gameBox(dst, cx+3, cy+3, cell-6, cell-6, ink)
			textInk := presentation.TileInk(ink)
			if g.Theme == menu.Wireframe {
				textInk = p.ink
			}
			boardNumber(dst, v, cx+cell/2, cy+cell/2, math.Min(30, cell*.35), textInk)
		}
	}
}
func (g *Game) clusterBall(dst *ebiten.Image, value int, x, y, r float64, p palette) {
	ink := ownerColor(value, p)
	if g.Theme == menu.Wireframe {
		ink = p.ink
		vector.StrokeCircle(dst, float32(x), float32(y), float32(r), 1.5, ink, true)
	} else {
		dot(dst, x, y, r, ink)
	}
	mark := presentation.TileInk(ink)
	if g.Theme == menu.Wireframe {
		mark = p.ink
	}
	switch value {
	case 1:
		vector.StrokeCircle(dst, float32(x), float32(y), float32(r*.35), 1.5, mark, true)
	case 2:
		vector.StrokeRect(dst, float32(x-r*.32), float32(y-r*.32), float32(r*.64), float32(r*.64), 1.5, mark, true)
	case 3:
		triangle(dst, x-r*.4, y-r*.4, r*.8, mark)
	case 4:
		line(dst, x-r*.3, y-r*.3, x+r*.3, y+r*.3, 1.5, mark)
		line(dst, x-r*.3, y+r*.3, x+r*.3, y-r*.3, 1.5, mark)
	}
}
func (g *Game) drawDeductionDigits(dst *ebiten.Image, p palette) {
	for digit := 1; digit <= 9; digit++ {
		x, y, w, h := g.deductionDigitRect(digit)
		rect(dst, x+1, y, w-2, h, p.panel)
		boardNumber(dst, digit, x+w/2, y+h/2, 20, p.ink)
	}
}
func (g *Game) drawStackNext(dst *ebiten.Image, p palette) {
	x, y, w, h := g.gameRect()
	cell := math.Min(15, w/6)
	for _, pos := range play.StackShape(g.games.Stack.Next, 0) {
		g.gameBox(dst, x+w/2+float64(pos.X-2)*cell, y+h+36+float64(pos.Y)*cell, cell-2, cell-2, p.a)
	}
}

// Reuse the tiny procedural digit patterns, scaled to the available cell size.
func boardNumber(dst *ebiten.Image, value int, x, y, height float64, ink color.Color) {
	text := fmt.Sprint(value)
	pixel := math.Max(1, math.Floor(height/5))
	left := x - float64(len(text)*4-1)*pixel/2
	top := y - 5*pixel/2
	for i, digit := range text {
		if digit < '0' || digit > '9' {
			continue
		}
		for row, pattern := range unitDigits[digit-'0'] {
			for col, v := range pattern {
				if v == '1' {
					rect(dst, left+float64(i*4+col)*pixel, top+float64(row)*pixel, pixel, pixel, ink)
				}
			}
		}
	}
}
func uprightTriangle(dst *ebiten.Image, x, y, size float64, ink color.Color) {
	var path vector.Path
	path.MoveTo(float32(x), float32(y-size/2))
	path.LineTo(float32(x+size/2), float32(y+size/2))
	path.LineTo(float32(x-size/2), float32(y+size/2))
	path.Close()
	vector.FillPath(dst, &path, nil, &vector.DrawPathOptions{ColorScale: colorScale(ink), AntiAlias: true})
}
