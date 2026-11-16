// Package stmpckmobile exposes the same menu to the native Android view.
package stmpckmobile

import (
	"ebe-boilerplate/internal/app"
	"github.com/hajimehoshi/ebiten/v2/mobile"
	"sync"
)

var start sync.Once
var game = app.NewGame()

// Start initializes the menu with fresh session defaults once per process.
func Start(path string) {
	start.Do(func() {
		game.SetStorage(path)
		mobile.SetGame(game)
	})
}

// Back requests one upward navigation step from the Android UI thread.
func Back() { game.RequestBack() }

// Suspend saves the last immutable game snapshot before the native view stops.
func Suspend() { game.Suspend() }

// Updater is implemented by the Android Activity; it returns without blocking.
type Updater interface{ CheckForUpdate() }

func RegisterUpdater(updater Updater) {
	if updater == nil {
		game.SetUpdateCheck(nil)
		return
	}
	game.SetUpdateCheck(func() { updater.CheckForUpdate() })
}

func ReportUpdateStatus(status string) { game.SetUpdateStatus(status) }
