// Geometry shared by embedded toolbar rendering, pointer activation and host focus.
package app

import (
	"ebe-boilerplate/internal/menu"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

// ToolbarRect returns the local target for intensity down, the four gallery
// icons, or intensity up. The title occupies a non-interactive slot before star.
func (g *Game) ToolbarRect(i int) image.Rectangle {
	w := g.width
	slot := min(44, max(1, (w-48)/6))
	groupLeft := (w - 48 - 4*slot) / 2
	x := groupLeft + 48 + (i-1)*slot
	if i == 0 {
		return image.Rect(0, 0, min(groupLeft, 76), 44)
	}
	if i == 5 {
		return image.Rect(max(w-76, groupLeft+48+4*slot), 0, w, 44)
	}
	return image.Rect(x, 0, x+slot, 44)
}

func (g *Game) drawGalleryToolbar(dst *ebiten.Image, p palette) {
	rect(dst, 0, 0, float64(g.width), float64(g.HeaderHeight()), p.background)
	label(dst, "STMPCK", g.ToolbarRect(1).Min.X-46, 27, p.ink)
	for i, name := range []string{"2b50", "1f465", "2728", "2699"} {
		r := g.ToolbarRect(i + 1)
		cx := float64(r.Min.X+r.Max.X) / 2
		emoji(dst, name, cx, 22, float64(min(28, r.Dx())))
		if i == g.Tab {
			line(dst, cx-12, 40, cx+12, 40, 2, p.a)
		}
	}
	for _, i := range []int{0, 5} {
		r := g.ToolbarRect(i)
		cx := float64(r.Min.X+r.Max.X) / 2
		d := 1
		name := "1f929"
		if i == 0 {
			d = -1
			name = "1f642"
		}
		ink := p.ink
		if g.Theme+d < 0 || g.Theme+d >= len(menu.IntensityLabels) {
			ink = p.muted
		}
		if r.Dx() >= 56 {
			emoji(dst, name, cx-float64(d)*13, 22, 24)
			cx += float64(d) * 15
		}
		chevron(dst, cx, 22, float64(d), ink)
	}
	line(dst, 8, 46, float64(g.width-8), 46, 1, p.edge)
	if g.Tab == menu.Worldwide && !g.HasGlobalRanking {
		rect(dst, 0, float64(g.height-menu.Footer), float64(g.width), menu.Footer, p.background)
		label(dst, "Popularity estimate", 16, g.height-7, p.muted)
	}
}
