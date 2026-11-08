package app

import (
	"fmt"
	"image/color"
	"math"
	"math/bits"

	"ebe-boilerplate/internal/menu"
	"ebe-boilerplate/internal/play"
	"ebe-boilerplate/internal/presentation"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (g *Game) gameBox(dst *ebiten.Image, x, y, w, h float64, c color.Color) {
	if g.Theme == menu.Wireframe {
		vector.StrokeRect(dst, float32(x), float32(y), float32(w), float32(h), 1.4, c, true)
	} else {
		rect(dst, x, y, w, h, c)
	}
}
func centered(dst *ebiten.Image, s string, x, y float64, c color.Color) {
	label(dst, s, int(x)-len(s)*7/2, int(y)+5, c)
}
func (g *Game) renderGame(dst *ebiten.Image, c *play.Collection, id int, x, y, w, h float64, p palette, active bool) {
	if w <= 0 || h <= 0 {
		return
	}
	ratio := 1.0
	switch id {
	case play.SnakeID:
		ratio = float64(play.SnakeW) / play.SnakeH
	case play.LinearID:
		ratio = 7.0 / 6
	case play.TerritoryID, play.BounceID, play.BrickID:
		ratio = 1.2
	case play.StackID:
		ratio = .5
	case play.CrossingID:
		ratio = 9.0 / 11
	}
	rw := math.Min(w, h*ratio)
	rh := rw / ratio
	x += (w - rw) / 2
	y += (h - rh) / 2
	w, h = rw, rh
	rect(dst, x, y, w, h, p.panel)
	if id >= play.FlowID {
		g.renderAddition(dst, c, id, x, y, w, h, p, active)
		return
	}
	switch id {
	case play.MergeID:
		m := &c.Merge
		animation := g.mergeAnimation
		before, targets := g.mergeBefore, g.mergeTargets
		if !active {
			animation = g.demoMergeAnimation
			before = m.UndoPosition.Cells
			targets = m.Travel
			if !g.Autoplay() {
				animation = 0
			}
		}

		cell := w / float64(m.N)
		for i := 0; i < m.N*m.N; i++ {
			cx, cy := x+float64(i%m.N)*cell, y+float64(i/m.N)*cell
			gap := math.Max(2, cell*.04)
			g.gameBox(dst, cx+gap, cy+gap, cell-2*gap, cell-2*gap, p.edge)
			v := m.Cells[i]
			if animation > 0 {
				continue
			}
			if v == 0 {
				continue
			}
			exponent := bits.Len(uint(v)) - 1
			ink := presentation.MergeColor(v, g.Theme == menu.Wireframe)
			g.gameBox(dst, cx+gap+1, cy+gap+1, cell-2*gap-2, cell-2*gap-2, ink)
			textInk := presentation.TileInk(ink)
			if g.Theme == menu.Wireframe {
				textInk = p.ink
			}
			if active && g.symbols {
				for k := 0; k < exponent; k++ {
					px := cx + cell*.2 + float64(k%5)*cell*.14
					py := cy + cell*.25 + float64(k/5)*cell*.15
					dot(dst, px, py, math.Max(1, cell*.035), textInk)
				}
			} else {
				centered(dst, fmt.Sprint(v), cx+cell/2, cy+cell/2, textInk)
			}
			if active && g.flash > 0 && m.Changed[i] {
				vector.StrokeRect(dst, float32(cx+gap), float32(cy+gap), float32(cell-2*gap), float32(cell-2*gap), 2, p.ink, true)
			}
			if active && i == m.Spawn {
				dot(dst, cx+cell-gap*2, cy+gap*2, 2, p.ink)
			}
		}
		if animation > 0 {
			progress := 1 - float64(animation)/12
			progress = 1 - (1-progress)*(1-progress)
			for i, v := range before {
				if v == 0 || i >= m.N*m.N {
					continue
				}
				dest := targets[i]
				cx := x + (float64(i%m.N)*(1-progress)+float64(dest%m.N)*progress)*cell
				cy := y + (float64(i/m.N)*(1-progress)+float64(dest/m.N)*progress)*cell
				gap := math.Max(2, cell*.04)
				ink := presentation.MergeColor(v, g.Theme == menu.Wireframe)
				tileInk := presentation.TileInk(ink)
				if g.Theme == menu.Wireframe {
					tileInk = p.ink
				}
				g.gameBox(dst, cx+gap, cy+gap, cell-2*gap, cell-2*gap, ink)
				if g.symbols {
					for k := 0; k < bits.Len(uint(v))-1; k++ {
						dot(dst, cx+cell*.2+float64(k%5)*cell*.14, cy+cell*.25+float64(k/5)*cell*.15, math.Max(1, cell*.035), tileInk)
					}
				} else {
					centered(dst, fmt.Sprint(v), cx+cell/2, cy+cell/2, tileInk)
				}
			}
		}

	case play.SnakeID:
		s := &c.Snake
		cw, ch := w/play.SnakeW, h/play.SnakeH
		vector.StrokeRect(dst, float32(x), float32(y), float32(w), float32(h), 2, p.edge, true)
		for i, b := range s.Body {
			cx, cy := x+float64(b%play.SnakeW)*cw, y+float64(b/play.SnakeW)*ch
			ink := p.a
			if i == 0 {
				ink = p.b
			}
			g.gameBox(dst, cx+1, cy+1, cw-2, ch-2, ink)
			if i == 0 {
				dot(dst, cx+cw*.35, cy+ch*.35, math.Max(1, cw*.08), p.ink)
				dot(dst, cx+cw*.65, cy+ch*.35, math.Max(1, cw*.08), p.ink)
			}
		}
		if s.Food >= 0 {
			dot(dst, x+(float64(s.Food%play.SnakeW)+.5)*cw, y+(float64(s.Food/play.SnakeW)+.5)*ch, math.Min(cw, ch)*.32, p.b)
		}
	case play.TerritoryID:
		g.drawTerritory(dst, &c.Territory, x, y, w, h, p, active)

	case play.LinearID:
		l := &c.Linear
		cw, ch := w/7, h/6
		for i, v := range l.Cells {
			cx, cy := x+(float64(i%7)+.5)*cw, y+(float64(i/7)+.5)*ch
			r := math.Min(cw, ch) * .38
			ink := p.edge
			if v == 1 {
				ink = p.a
			}
			if v == 2 {
				ink = p.b
			}
			if g.Theme == menu.Wireframe {
				vector.StrokeCircle(dst, float32(cx), float32(cy), float32(r), 1.5, ink, true)
			} else {
				dot(dst, cx, cy, r, ink)
			}
			if v == 2 {
				vector.StrokeRect(dst, float32(cx-r*.45), float32(cy-r*.45), float32(r*.9), float32(r*.9), 2, p.ink, true)
			}
			if v == 1 {
				dot(dst, cx, cy, math.Max(1, r*.15), p.ink)
			}
		}
		if active && !l.Over {
			landing := l.Landing(g.cursor)
			if landing >= 0 {
				cx, cy := x+(float64(g.cursor)+.5)*cw, y+(float64(landing/7)+.5)*ch
				vector.StrokeCircle(dst, float32(cx), float32(cy), float32(cw*.44), 2, p.ink, true)
			}
		}
		if l.Winner != 0 {
			a, b := l.Winning[0], l.Winning[3]
			line(dst, x+(float64(a%7)+.5)*cw, y+(float64(a/7)+.5)*ch, x+(float64(b%7)+.5)*cw, y+(float64(b/7)+.5)*ch, 3, p.ink)
		}
	case play.BounceID:
		b := &c.Bounce
		vector.StrokeRect(dst, float32(x), float32(y), float32(w), float32(h), 2, p.edge, true)
		for i := 0; i < 10; i++ {
			line(dst, x+w/2, y+float64(i)*h/10, x+w/2, y+(float64(i)+.4)*h/10, 1, p.edge)
		}
		g.gameBox(dst, x+w*.04, y+(b.Player-.12)*h, w*.012, h*.24, p.a)
		g.gameBox(dst, x+w*.948, y+(b.Computer-.12)*h, w*.012, h*.24, p.b)
		if g.Theme >= menu.Vivid {
			for i := 1; i < 5; i++ {
				dot(dst, x+(b.X-b.VX*float64(i)/120)*w, y+(b.Y-b.VY*float64(i)/120)*h, math.Max(1, w*.01-float64(i)*.5), p.edge)
			}
		}
		dot(dst, x+b.X*w, y+b.Y*h, w*.012, p.ink)
		for side, score := range b.Points {
			for i := 0; i < 5; i++ {
				ink := p.edge
				if i < score {
					ink = p.ink
				}
				dot(dst, x+w*(.25+float64(side)*.5)+float64(i-2)*10, y+12, 3, ink)
			}
		}
	}
}

func (g *Game) drawPlay(dst *ebiten.Image, p palette) {
	if g.setup && g.Selected == play.PlanetsID {
		g.drawPlanetChoices(dst, p)
		return
	}
	if g.setup && g.Selected == play.MergeID {
		g.drawMergeChoices(dst, p)
		return
	}
	if g.setup && g.Selected == play.TerritoryID {
		g.drawTerritorySetup(dst, p)
		return
	}
	if g.setup {
		centered(dst, menu.Names[g.Selected], float64(g.width)/2, menu.Header+20, p.ink)
		names := []string{"Start"}
		switch g.Selected {
		case play.MergeID:
			names = []string{"4 x 4", "5 x 5", "6 x 6"}
		case play.SnakeID, play.BounceID:
			names = []string{"Slow", "Medium", "Fast"}
		case play.LinearID:
			names = []string{"Computer"}
		}
		width := float64(g.width) / float64(len(names))
		for i, name := range names {
			x := float64(i)*width + 8
			y := float64(menu.Header + 50)
			rect(dst, x, y, width-16, 82, p.panel)
			if i == g.option {
				vector.StrokeRect(dst, float32(x), float32(y), float32(width-16), 82, 2, p.a, true)
			}
			centered(dst, name, x+(width-16)/2, y+41, p.ink)
		}
	} else {
		x, y, w, h := g.gameRect()
		g.breathingBorder(dst, x, y, w, h, p)
		g.renderGame(dst, &g.games, g.Selected, x, y, w, h, p, true)
		if g.Selected == play.BloomID {
			g.drawBloomPlacement(dst, p)
		}
		centered(dst, g.statusText(), float64(g.width)/2, y+h+18, p.ink)
		if g.Selected == play.DeductionID {
			g.drawDeductionDigits(dst, p)
		}
		if g.Selected == play.StackID {
			g.drawStackNext(dst, p)
		}
		if g.Selected == play.BloomID {
			g.drawBloomButton(dst, p)
		}

		if g.Selected == play.TerritoryID && g.games.Territory.Dice[0] > 0 {
			t := g.games.Territory
			centered(dst, fmt.Sprintf("Dice: %d / %d", t.Dice[0], t.Dice[1]), float64(g.width)/2, y+h+45, p.muted)
		}
		if g.Selected == play.MergeID && g.symbols {
			centered(dst, "Dots count powers of 2", float64(g.width)/2, y+h+91, p.muted)
		}
	}
	y := float64(g.height - 58)
	names := []string{"Play", "New", "", ""}
	if !g.setup && !g.paused {
		names[0] = "Pause"
		if g.Selected == play.SnakeID && !g.games.Snake.Running || g.Selected == play.BounceID && !g.games.Bounce.Running {
			names[0] = "Play"
		}
	}
	if g.setup {
		names[0] = "Start"
		names[1] = "Back"
	}
	if !g.setup {
		if g.Selected >= play.FlowID {
			g.additionFooter(names)
		}
		switch g.Selected {
		case play.MergeID:
			names[2] = "Undo"
			names[3] = "123 / dots"
		case play.LinearID:
			if g.games.Linear.Over {
				names[2] = "Replay"
			}
		}
	}
	if g.restartConfirm {
		names = []string{"Confirm", "Cancel", "", ""}
	}
	for i, name := range names {
		if name == "" {
			continue
		}
		x := float64(i*g.width) / 4
		w := float64(g.width) / 4
		rect(dst, x+4, y, w-8, 48, p.panel)
		centered(dst, name, x+w/2, y+24, p.ink)
	}
	if g.saveFailed.Load() {
		label(dst, "Save unavailable - keep app open", 18, menu.Header+12, p.ink)
	}
	if g.Selected == play.BloomID && g.bloomPicker {
		g.drawBloomPicker(dst, p)
	}
}

func ownerColor(owner int, p palette) color.RGBA {
	switch owner {
	case 0:
		return rgb(0xc5ccd3)
	case 1:
		return p.a
	case 2:
		return p.b
	case 3:
		return rgb(0x8bbcff)
	default:
		return rgb(0xe7bbff)
	}
}
