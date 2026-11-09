package app

import (
	"fmt"
	"image"
	"image/color"
	"math"

	"ebe-boilerplate/internal/menu"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type palette struct{ background, panel, edge, ink, muted, a, b color.RGBA }

var palettes = [3]palette{
	{rgb(0x111314), rgb(0x1b1e20), rgb(0x393e41), rgb(0xf3f1e9), rgb(0x909b9b), rgb(0xe0e5df), rgb(0x778486)},
	{rgb(0x101a1a), rgb(0x192827), rgb(0x324b45), rgb(0xf2f0db), rgb(0x9daf9d), rgb(0xc3e88a), rgb(0xf6b69c)},
	{rgb(0x180e29), rgb(0x281b3d), rgb(0x59416f), rgb(0xfff0fa), rgb(0xc1a9cd), rgb(0x67efd8), rgb(0xff83c4)},
}

func rgb(v uint32) color.RGBA { return color.RGBA{uint8(v >> 16), uint8(v >> 8), uint8(v), 255} }

func (g *Game) Draw(screen *ebiten.Image) {
	p := palettes[min(2, max(0, g.Theme-1))]
	if g.Theme == menu.Still {
		p = palettes[1]
	}
	p = g.effectPalette(p)
	if g.Settings {
		g.drawSettings(screen, palettes[1])
		return
	}
	screen.Fill(p.background)
	if g.Selected < 0 {
		body := screen.SubImage(image.Rect(0, min(menu.Header, g.height), g.width, g.height)).(*ebiten.Image)
		order := g.galleryOrder
		if len(order) == 0 {
			centered(body, "All games explored", float64(g.width)/2, menu.Header+60, p.ink)
		}
		for i, id := range order {
			x, y, w, _ := g.TileRect(i, g.width)
			if y+float64(g.RowHeight) < menu.Header || y > float64(g.height) {
				continue
			}
			rect(body, x, y, w, float64(g.RowHeight-14), p.panel)
			if i == g.Focus {
				vector.StrokeRect(body, float32(x+1), float32(y+1), float32(w-2), float32(g.RowHeight-16), 2, p.a, false)
			}
			g.breathingBorder(body, x, y, w, float64(g.RowHeight-14), p)
			g.preview(body, id, x+4, y+4, w-8, float64(g.RowHeight-22), p)
			// Four corner marks distinguish preview cards from play buttons.
			line(body, x+w-22, y+18, x+w-14, y+18, 1, p.muted)
			line(body, x+w-14, y+18, x+w-14, y+26, 1, p.muted)
		}
		if float64(g.Rows()*g.RowHeight+menu.Header+menu.Footer) > float64(g.height) {
			track := float64(g.height - menu.Header - menu.Footer)
			thumb := math.Max(20, track*track/float64(g.Rows()*g.RowHeight))
			y := menu.Header + g.Scroll/(float64(g.Rows()*g.RowHeight)-track)*(track-thumb)
			rect(screen, float64(g.width)-6, y, 2, thumb, p.muted)
		}
	} else {
		g.drawPlay(screen, p)
	}

	// Header is fixed while the gallery scrolls.
	rect(screen, 0, 0, float64(g.width), menu.Header, p.background)
	if g.Selected >= 0 {
		chevron(screen, 28, 26, -1, p.ink)
	} else {
		if g.Tab == menu.Worldwide && !g.HasGlobalRanking {
			label(screen, "Popularity estimate", 16, 31, p.muted)
		} else {
			label(screen, "STMPCK", 20, 31, p.ink)
		}
	}
	emoji(screen, "2699", float64(g.width-28), 26, 26)
	cx := float64(g.width) / 2
	if g.Selected < 0 {
		for i, name := range []string{"2b50", "1f465", "2728"} {
			tx := cx + float64(i-1)*52
			emoji(screen, name, tx, 67, 28)
			if i == g.Tab {
				line(screen, tx-16, 87, tx+16, 87, 2, p.a)
			}
		}
	}
	emoji(screen, "1f642", 60, 67, 28)
	emoji(screen, "1f929", float64(g.width-60), 67, 28)
	for _, d := range []int{-1, 1} {
		x := 28.0
		if d > 0 {
			x = float64(g.width) - 28
		}
		ink := p.ink
		if g.Theme+d < 0 || g.Theme+d >= len(menu.IntensityLabels) {
			ink = p.muted
		}
		chevron(screen, x, 67, float64(d), ink)
		line(screen, x-float64(d)*9, 67, x+float64(d)*3, 67, 2, ink)
	}
	line(screen, 16, menu.Header-9, float64(g.width-16), menu.Header-9, 1, p.edge)
	if !g.Settings {
		g.drawSparkles(screen, p)
	}
	if g.Settings {
		g.drawSettings(screen, p)
	}
}

func chevron(dst *ebiten.Image, x, y, d float64, c color.Color) {
	line(dst, x-d*4, y-7, x+d*3, y, 2, c)
	line(dst, x+d*3, y, x-d*4, y+7, 2, c)
}
func triangle(dst *ebiten.Image, x, y, size float64, c color.Color) {
	var path vector.Path
	path.MoveTo(float32(x), float32(y))
	path.LineTo(float32(x+size), float32(y+size*.6))
	path.LineTo(float32(x), float32(y+size*1.2))
	path.Close()
	vector.FillPath(dst, &path, nil, &vector.DrawPathOptions{ColorScale: colorScale(c), AntiAlias: true})
}
func colorScale(c color.Color) ebiten.ColorScale {
	var s ebiten.ColorScale
	s.ScaleWithColor(c)
	return s
}

func (g *Game) drawSettings(dst *ebiten.Image, p palette) {
	dst.Fill(p.background)
	chevron(dst, 28, 26, -1, p.ink)
	emoji(dst, "2699", float64(g.width-28), 26, 26)
	label(dst, "Settings", 60, 31, p.ink)
	label(dst, "Preview only - switches do not apply", 20, 66, p.muted)
	label(dst, "or save settings yet.", 20, 84, p.muted)
	names := [...]string{"Disable menu animations", "Disable all audio", "Disable haptics", "Reduce decorative effects"}
	for i, name := range names {
		y := float64(100 + i*58)
		rect(dst, 16, y, float64(g.width-32), 50, p.panel)
		label(dst, name, 26, int(y+30), p.ink)
		x := float64(g.width - 76)
		ink := p.muted
		knob := x + 10
		if g.mockOptions[i] {
			ink = p.a
			knob = x + 34
		}
		rect(dst, x, y+15, 44, 20, p.edge)
		dot(dst, knob, y+25, 8, ink)
	}
	rect(dst, 16, 338, float64(g.width-32), 78, p.panel)
	label(dst, "Games wide (preview only)", 26, 360, p.ink)
	choiceWidth := float64(g.width-52) / 8
	for i := 1; i <= 8; i++ {
		x := 26 + float64(i-1)*choiceWidth
		if i == menu.Columns {
			vector.StrokeRect(dst, float32(x), 370, float32(choiceWidth-4), 32, 2, p.a, true)
		}
		centered(dst, fmt.Sprint(i), x+(choiceWidth-4)/2, 386, p.ink)
	}
	lines := []string{
		"Menu arrows: intensity / browse.",
		"Game arrows: move. Enter: confirm.",
		"Space / pad Y: pause. R: new game.",
		"Merge U: undo. S: symbols.",
		"Enter: open. Back / Esc: up one level.",
		"1/2/3: sorting. H/F1 or Start: settings.",
		"Emoji: Twemoji, Twitter + contributors.",
		"creativecommons.org/licenses/by/4.0/",
	}
	for i, text := range lines {
		label(dst, text, 20, 440+i*18, p.muted)
	}
}
