//go:build previewcapture && js && wasm

package app

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
	"syscall/js"

	"ebe-boilerplate/internal/menu"
	"ebe-boilerplate/internal/play"
	"ebe-boilerplate/internal/preview"
	"github.com/hajimehoshi/ebiten/v2"
)

// Capture uses real rules/rendering and never initializes player storage.
type PreviewCapture struct {
	request   js.Func
	id, style int
	pending   bool
}

func NewPreviewCapture() *PreviewCapture {
	ebiten.SetRunnableOnUnfocused(true)
	c := &PreviewCapture{}
	c.request = js.FuncOf(func(_ js.Value, args []js.Value) any {
		if c.pending {
			return false
		}
		c.id, c.style, c.pending = args[0].Int(), args[1].Int(), true
		js.Global().Set("previewResult", js.Null())
		return true
	})
	js.Global().Set("capturePreview", c.request)
	ids := make([]any, play.Count)
	for i, id := range play.IDs {
		ids[i] = id
	}
	js.Global().Set("previewIDs", js.ValueOf(ids))
	return c
}
func (c *PreviewCapture) Update() error              { return nil }
func (c *PreviewCapture) Layout(_, _ int) (int, int) { return 96, 96 }
func (c *PreviewCapture) Draw(screen *ebiten.Image) {
	if !c.pending {
		return
	}
	c.pending = false
	if c.id < 0 || c.id >= play.Count || c.style < 0 || c.style > 2 {
		panic("invalid capture request")
	}
	g := &Game{}
	g.Theme = []int{menu.Wireframe, menu.DefaultIntensity, menu.Vivid}[c.style]
	p := palettes[c.style]
	demos := play.NewDemonstrations()
	frames := preview.Frames(c.id)
	if c.style == 0 {
		frames = 1
	}
	w, h := preview.Size*preview.Columns, preview.Size*((frames+preview.Columns-1)/preview.Columns)
	sheet := ebiten.NewImage(w, h)
	framePixels := make([]byte, preview.Size*preview.Size*4)
	tick := 0
	// Start with established play so static cards are informative too.
	for tick < 120 {
		tick++
		demos.Demonstrate(tick)
	}
	for frame := 0; frame < frames; frame++ {
		if frame > 0 {
			for j := 0; j < 60/preview.FPS; j++ {
				tick++
				if g.demoMergeAnimation > 0 {
					g.demoMergeAnimation--
				}
				before := demos.Merge.Cells
				demos.Demonstrate(tick)
				if demos.Merge.Cells != before {
					g.demoMergeAnimation = 12
				}
			}
		}
		r := preview.FrameRect(c.id, frame)
		tile := sheet.SubImage(r).(*ebiten.Image)
		g.renderGame(tile, &demos, c.id, float64(r.Min.X), float64(r.Min.Y), preview.Size, preview.Size, p, false)
		// Flush each frame rather than retaining an entire animation's draw
		// commands, which can exhaust WebAssembly memory on software renderers.
		tile.ReadPixels(framePixels)
	}
	pixels := image.NewRGBA(image.Rect(0, 0, w, h))
	sheet.ReadPixels(pixels.Pix)
	encode := func(im image.Image) string {
		var b bytes.Buffer
		if err := png.Encode(&b, im); err != nil {
			panic(err)
		}
		return base64.StdEncoding.EncodeToString(b.Bytes())
	}
	result := map[string]any{"still": encode(pixels.SubImage(image.Rect(0, 0, preview.Size, preview.Size)))}
	if c.style > 0 {
		result["sheet"] = encode(pixels)
	}
	js.Global().Set("previewResult", js.ValueOf(result))
	sheet.Deallocate()
}
