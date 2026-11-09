//go:build previewcapture && js && wasm

package main

import (
	"ebe-boilerplate/internal/app"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	if err := ebiten.RunGame(app.NewPreviewCapture()); err != nil {
		panic(err)
	}
}
