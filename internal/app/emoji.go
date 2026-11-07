package app

import (
	"bytes"
	"embed"
	"github.com/hajimehoshi/ebiten/v2"
	"image/png"
)

// Tiny locally bundled graphics avoid platform font differences and network access.
// Attribution and CC BY 4.0 terms are in emoji/ATTRIBUTION.md and LICENSE-GRAPHICS.
//
//go:embed emoji/*.png emoji/LICENSE-GRAPHICS emoji/ATTRIBUTION.md
var emojiFiles embed.FS
var emojiImages = map[string]*ebiten.Image{}

func emoji(dst *ebiten.Image, name string, cx, cy, size float64) {
	img := emojiImages[name]
	if img == nil {
		data, err := emojiFiles.ReadFile("emoji/" + name + ".png")
		if err != nil {
			panic(err)
		}
		decoded, err := png.Decode(bytes.NewReader(data))
		if err != nil {
			panic(err)
		}
		img = ebiten.NewImageFromImage(decoded)
		emojiImages[name] = img
	}
	var op ebiten.DrawImageOptions
	op.GeoM.Scale(size/float64(img.Bounds().Dx()), size/float64(img.Bounds().Dy()))
	op.GeoM.Translate(cx-size/2, cy-size/2)
	op.Filter = ebiten.FilterLinear
	dst.DrawImage(img, &op)
}
