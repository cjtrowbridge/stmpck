package app

import (
	"ebe-boilerplate/internal/menu"
	"ebe-boilerplate/internal/play"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"math"
)

func (g *Game) toggleBloomPicker() {
	if g.setup || g.restartConfirm {
		return
	}
	if g.bloomPicker {
		g.bloomPicker = false
		return
	}
	g.pauseGame()
	g.bloomPlacing = false
	g.bloomPicker = true
}
func (g *Game) moveBloomChoice(d int) {
	x, y := g.bloomChoice%2, g.bloomChoice/2
	x += [4]int{-1, 0, 1, 0}[d]
	y += [4]int{0, -1, 0, 1}[d]
	g.bloomChoice = max(0, min(2, y))*2 + max(0, min(1, x))
}
func (g *Game) confirmBloomPattern() {
	if g.bloomPicker {
		g.bloomPicker = false
		g.bloomPlacing = true
		g.cursor = play.BloomN/2*play.BloomN + play.BloomN/2
		return
	}
	if g.bloomPlacing {
		g.games.Bloom.Stamp(g.bloomChoice, g.cursor)
		g.bloomPlacing = false
		g.saveGames()
	}
}
func (g *Game) bloomPatternAction(slot int) {
	switch slot {
	case 0:
		g.confirmBloomPattern()
	case 1:
		g.bloomPicker = false
		g.bloomPlacing = false
	}
}
func (g *Game) bloomButtonRect() (x, y, w, h float64) {
	bx, by, bw, bh := g.gameRect()
	w = 112
	h = 40
	x = bx + (bw-w)/2
	y = by + bh + 42
	return
}
func (g *Game) bloomPickerRect() (x, y, w, h float64) {
	w = math.Min(440, float64(g.width)-24)
	h = math.Min(420, float64(g.height-menu.Header-86))
	h = math.Max(120, h)
	x = (float64(g.width) - w) / 2
	y = menu.Header + 8
	return
}
func (g *Game) bloomChoiceRect(i int) (x, y, w, h float64) {
	bx, by, bw, bh := g.bloomPickerRect()
	w = (bw - 24) / 2
	h = (bh - 32) / 3
	x = bx + 8 + float64(i%2)*(w+8)
	y = by + 8 + float64(i/2)*(h+8)
	return
}
func (g *Game) bloomPickerTap(x, y float64) {
	if y >= float64(g.height-62) {
		g.gameAction(int(x * 4 / float64(g.width)))
		return
	}
	for i := range play.BloomPatterns {
		rx, ry, rw, rh := g.bloomChoiceRect(i)
		if x >= rx && x < rx+rw && y >= ry && y < ry+rh {
			g.bloomChoice = i
			g.confirmBloomPattern()
			return
		}
	}
}
func (g *Game) drawBloomButton(dst *ebiten.Image, p palette) {
	if g.bloomPicker {
		return
	}
	x, y, w, h := g.bloomButtonRect()
	rect(dst, x, y, w, h, p.panel)
	centered(dst, "Patterns", x+w/2, y+h/2, p.ink)
}
func (g *Game) drawBloomPlacement(dst *ebiten.Image, p palette) {
	if !g.bloomPlacing {
		return
	}
	x, y, w, _ := g.gameRect()
	cell := w / play.BloomN
	for _, i := range play.BloomPatternCells(g.bloomChoice, g.cursor) {
		cx, cy := x+float64(i%play.BloomN)*cell, y+float64(i/play.BloomN)*cell
		vector.StrokeRect(dst, float32(cx+1), float32(cy+1), float32(cell-2), float32(cell-2), 1.5, p.b, true)
		dot(dst, cx+cell/2, cy+cell/2, math.Max(1, cell*.12), p.ink)
	}
}
func (g *Game) drawBloomPicker(dst *ebiten.Image, p palette) {
	x, y, w, h := g.bloomPickerRect()
	rect(dst, x-4, y-4, w+8, h+8, p.background)
	rect(dst, x, y, w, h, p.edge)
	for i, pattern := range play.BloomPatterns {
		cx, cy, cw, ch := g.bloomChoiceRect(i)
		rect(dst, cx, cy, cw, ch, p.panel)
		if i == g.bloomChoice {
			vector.StrokeRect(dst, float32(cx), float32(cy), float32(cw), float32(ch), 2, p.a, true)
		}
		cols, rows := len(pattern.Rows[0]), len(pattern.Rows)
		cell := math.Min(10, math.Min((cw-20)/float64(cols), (ch-34)/float64(rows)))
		left := cx + (cw-float64(cols)*cell)/2
		top := cy + (ch-24-float64(rows)*cell)/2
		for row, line := range pattern.Rows {
			for col, v := range line {
				if v == 'O' {
					g.gameBox(dst, left+float64(col)*cell, top+float64(row)*cell, math.Max(1, cell-1), math.Max(1, cell-1), p.a)
				}
			}
		}
		centered(dst, pattern.Name, cx+cw/2, cy+ch-13, p.ink)
	}
}
