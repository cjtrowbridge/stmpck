package main

import (
	"ebe-boilerplate/internal/app"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	ebiten.SetWindowSize(480, 800)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("STMPCK")
	game := app.NewGame()
	defer game.Suspend()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
