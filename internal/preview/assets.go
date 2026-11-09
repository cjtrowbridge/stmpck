// Package preview defines the bundled, offline gallery image contract.
package preview

import (
	"embed"
	"fmt"
	"image"
	"image/png"

	"ebe-boilerplate/internal/play"
)

const (
	Size       = 96
	FPS        = 12
	Columns    = 8
	CacheBytes = 64 << 20
)

//go:embed assets/*
var Files embed.FS

func Frames(id int) int {
	if id == play.BloomID {
		return 144
	}
	return 96
}

func Name(id, style int, sheet bool) string {
	kind := "still"
	if sheet {
		kind = "sheet"
	}
	return fmt.Sprintf("assets/%s-%d-%s.png", play.IDs[id], style, kind)
}

func Decode(id, style int, sheet bool) (image.Image, error) {
	f, err := Files.Open(Name(id, style, sheet))
	if err != nil {
		return nil, err
	}
	defer f.Close()
	im, err := png.Decode(f)
	if err != nil {
		return nil, err
	}
	w, h := Size, Size
	if sheet {
		w, h = Size*Columns, Size*(Frames(id)/Columns)
	}
	if im.Bounds().Dx() != w || im.Bounds().Dy() != h {
		return nil, fmt.Errorf("invalid preview dimensions: %s", Name(id, style, sheet))
	}
	return im, nil
}

func FrameRect(id, frame int) image.Rectangle {
	frame %= Frames(id)
	x, y := frame%Columns*Size, frame/Columns*Size
	return image.Rect(x, y, x+Size, y+Size)
}
