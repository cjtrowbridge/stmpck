// Package surface exposes the shared STMPCK game to the private host.
// It retains the standalone game rules and save format.
package surface

import (
	"ebe-boilerplate/internal/app"
	"ebe-boilerplate/internal/menu"
)

type Game = app.Game

func New(root string) *Game { return app.NewEmbedded(root) }

func GameNames() [menu.Count]string {
	return menu.Names
}
