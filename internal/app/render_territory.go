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

func (g *Game) drawTerritory(dst *ebiten.Image, t *play.Territory, x, y, w, h float64, p palette, active bool) {
	// Fill all regions before drawing shared borders so adjacent fills cannot hide lines.
	for i, poly := range t.Terrain.Polygons {
		var path vector.Path
		for j, v := range poly {
			px, py := float32(x+v.X*w), float32(y+v.Y*h)
			if j == 0 {
				path.MoveTo(px, py)
			} else {
				path.LineTo(px, py)
			}
		}
		path.Close()
		if g.Theme != menu.Wireframe {
			fill := mixColor(p.panel, ownerColor(t.Owners[i], p), .17)
			vector.FillPath(dst, &path, nil, &vector.DrawPathOptions{ColorScale: colorScale(fill), AntiAlias: true})
		}
	}
	for _, border := range t.Terrain.Borders {
		ink := rgb(0xf2f5f7)
		width := 1.7
		if border.B >= 0 {
			a, b := t.Owners[border.A], t.Owners[border.B]
			switch {
			case a > 0 && b > 0 && a != b:
				ink = rgb(0xff555f)
				width = 3.2
			case a > 0 && a == b:
				ink = rgb(0x5fe1cd)
				width = 1.7
			default:
				ink = rgb(0xf4ce76)
				width = 2.2
			}
		}
		if g.Theme == menu.Wireframe {
			ink = p.ink
		}
		for i := 1; i < len(border.Points); i++ {
			a, b := border.Points[i-1], border.Points[i]
			line(dst, x+a.X*w, y+a.Y*h, x+b.X*w, y+b.Y*h, width, ink)
		}
	}
	for _, m := range t.Movements {
		a, b := t.Terrain.Centers[m.From], t.Terrain.Centers[m.To]
		u := float64(m.Elapsed) / float64(m.Duration)
		ink := ownerColor(m.Owner, p)
		if g.Theme == menu.Wireframe {
			ink = p.ink
		}
		line(dst, x+a.X*w, y+a.Y*h, x+b.X*w, y+b.Y*h, 1, ink)
		px, py := x+(a.X+(b.X-a.X)*u)*w, y+(a.Y+(b.Y-a.Y)*u)*h
		dot(dst, px, py, 7, rgb(0x080c10))
		dot(dst, px, py, 5, ink)
	}
	baseRadius := math.Max(11, math.Min(18, math.Min(w, h)*.065))
	for i, center := range t.Terrain.Centers {
		radius := baseRadius
		for j, other := range t.Terrain.Centers {
			if i != j {
				radius = math.Min(radius, math.Hypot((center.X-other.X)*w, (center.Y-other.Y)*h)*.42-2)
			}
		}
		radius = math.Max(5, radius)
		px, py := x+center.X*w, y+center.Y*h
		ink := ownerColor(t.Owners[i], p)
		if g.Theme == menu.Wireframe {
			ink = p.ink
		}
		// A dark outer halo and pale ring remain visible over every theme and border.
		dot(dst, px, py, radius+2, rgb(0x080c10))
		dot(dst, px, py, radius, ink)
		vector.StrokeCircle(dst, float32(px), float32(py), float32(radius+1), 1, rgb(0xffffff), true)
		count := fmt.Sprint(t.Units[i])
		drawUnitCount(dst, count, px, py, radius, presentation.TileInk(ink))

		sx, sy := px, py-radius-6
		switch t.Owners[i] {
		case 1:
			dot(dst, sx, sy, 2.5, ink)
		case 2:
			rect(dst, sx-2.5, sy-2.5, 5, 5, ink)
		case 3:
			triangle(dst, sx-3, sy-3, 6, ink)
		case 4:
			line(dst, sx-3, sy-3, sx+3, sy+3, 2, ink)
			line(dst, sx-3, sy+3, sx+3, sy-3, 2, ink)
		}
		if active && (i == t.Source || i == g.cursor) {
			vector.StrokeCircle(dst, float32(px), float32(py), float32(radius+4), 2, p.ink, true)
		}
	}
}

// Tiny complete bitmap digits avoid downsampling font-atlas glyphs at fractional scales.
var unitDigits = [10][5]string{
	{"111", "101", "101", "101", "111"}, {"010", "110", "010", "010", "111"},
	{"111", "001", "111", "100", "111"}, {"111", "001", "111", "001", "111"},
	{"101", "101", "111", "001", "001"}, {"111", "100", "111", "001", "111"},
	{"111", "100", "111", "101", "111"}, {"111", "001", "010", "010", "010"},
	{"111", "101", "111", "101", "111"}, {"111", "101", "111", "001", "111"},
}

func drawUnitCount(dst *ebiten.Image, count string, x, y, radius float64, ink color.Color) {
	pixel := math.Min(2, radius*1.6/float64(len(count)*4-1))
	left := x - float64(len(count)*4-1)*pixel/2
	top := y - 5*pixel/2
	for i, d := range count {
		if d < '0' || d > '9' {
			continue
		}
		for row, line := range unitDigits[d-'0'] {
			for col, v := range line {
				if v == '1' {
					rect(dst, left+float64(i*4+col)*pixel, top+float64(row)*pixel, pixel, pixel, ink)
				}
			}
		}
	}
}
