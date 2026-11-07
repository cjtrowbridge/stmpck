//go:build js && wasm

package app

import "syscall/js"

func installLifecycle(g *Game) {
	// These callbacks have the lifetime of the running game.
	hidden := js.FuncOf(func(this js.Value, args []js.Value) any {
		if js.Global().Get("document").Get("hidden").Bool() {
			g.Suspend()
		}
		return nil
	})
	closing := js.FuncOf(func(this js.Value, args []js.Value) any { g.Suspend(); return nil })
	js.Global().Get("document").Call("addEventListener", "visibilitychange", hidden)
	js.Global().Call("addEventListener", "pagehide", closing)
}
