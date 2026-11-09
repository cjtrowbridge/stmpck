package app

import (
	"image"

	"ebe-boilerplate/internal/menu"
	"ebe-boilerplate/internal/play"
	"ebe-boilerplate/internal/preview"
	"github.com/hajimehoshi/ebiten/v2"
)

type previewKey struct{ id, style int }
type previewSheet struct {
	image *ebiten.Image
	used  uint64
}
type previewResult struct {
	key   previewKey
	image image.Image
	err   error
}
type galleryImages struct {
	posters [play.Count][3]*ebiten.Image
	sheets  map[previewKey]previewSheet
	failed  map[previewKey]bool
	result  chan previewResult
	loading bool
	bytes   int
	clock   uint64
	ticks   [play.Count]int
}

func (g *Game) initPreviews() {
	g.previews.sheets = make(map[previewKey]previewSheet)
	g.previews.failed = make(map[previewKey]bool)
	g.previews.result = make(chan previewResult, 1)
	for id := 0; id < play.Count; id++ {
		for style := 0; style < 3; style++ {
			im, err := preview.Decode(id, style, false)
			if err != nil {
				panic(err)
			} // Bundled assets are validated before builds.
			g.previews.posters[id][style] = ebiten.NewImageFromImage(im)
		}
	}
}

func (g *Game) previewStyle() int {
	if g.Theme == menu.Wireframe {
		return 0
	}
	if g.Theme >= menu.Vivid {
		return 2
	}
	return 1
}

func (g *Game) updatePreviews() {
	p := &g.previews
	p.clock++
	style := g.previewStyle()
	var visible [play.Count]bool
	for i, id := range g.galleryOrder {
		_, y, _, h := g.TileRect(i, g.width)
		visible[id] = y+h > menu.Header && y < float64(g.height)
	}
	protected := func(key previewKey) bool { return key.style == style && visible[key.id] }
	select {
	case r := <-p.result:
		p.loading = false
		if r.err != nil {
			p.failed[r.key] = true
			break
		}
		cost := r.image.Bounds().Dx() * r.image.Bounds().Dy() * 4
		for p.bytes+cost > preview.CacheBytes {
			var oldest previewKey
			age := ^uint64(0)
			for key, sheet := range p.sheets {
				if !protected(key) && sheet.used < age {
					oldest, age = key, sheet.used
				}
			}
			if age == ^uint64(0) {
				break
			}
			sheet := p.sheets[oldest]
			p.bytes -= sheet.image.Bounds().Dx() * sheet.image.Bounds().Dy() * 4
			sheet.image.Deallocate()
			delete(p.sheets, oldest)
		}
		if p.bytes+cost <= preview.CacheBytes {
			p.sheets[r.key] = previewSheet{ebiten.NewImageFromImage(r.image), p.clock}
			p.bytes += cost
		}
	default:
	}
	for _, id := range g.galleryOrder {
		if !visible[id] {
			continue
		}
		key := previewKey{id, style}
		if sheet, ok := p.sheets[key]; ok {
			sheet.used = p.clock
			p.sheets[key] = sheet
			p.ticks[id] = (p.ticks[id] + 1) % (preview.Frames(id) * 60 / preview.FPS)
		} else if !p.loading && !p.failed[key] {
			available := preview.CacheBytes - p.bytes
			for old, sheet := range p.sheets {
				if !protected(old) {
					available += sheet.image.Bounds().Dx() * sheet.image.Bounds().Dy() * 4
				}
			}
			// Keep visible sheets resident. Very large viewports use posters
			// for excess cards instead of endlessly evicting/reloading images.
			if available < preview.Size*preview.Size*preview.Frames(id)*4 {
				continue
			}
			p.loading = true
			go func() {
				im, err := preview.Decode(key.id, key.style, true)
				p.result <- previewResult{key, im, err}
			}()
		}
	}
}

func (g *Game) preview(dst *ebiten.Image, id int, x, y, w, h float64, _ palette) {
	style := g.previewStyle()
	im := g.previews.posters[id][style]
	if style > 0 {
		if sheet, ok := g.previews.sheets[previewKey{id, style}]; ok {
			im = sheet.image.SubImage(preview.FrameRect(id, g.previews.ticks[id]/(60/preview.FPS))).(*ebiten.Image)
		}
	}
	// Captures have a square canvas; fit it without distorting any game's board.
	size := min(w, h)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(size/preview.Size, size/preview.Size)
	op.GeoM.Translate(x+(w-size)/2, y+(h-size)/2)
	op.Filter = ebiten.FilterLinear
	dst.DrawImage(im, op)
}
