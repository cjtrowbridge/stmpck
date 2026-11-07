package app

import (
	"fmt"
	"math"
	"math/rand/v2"

	"ebe-boilerplate/internal/menu"
	"ebe-boilerplate/internal/play"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (g *Game) mergeChoiceRect(i int) (x, y, w, h float64) {
	x, w = g.cardBounds()
	h = math.Min(140, math.Max(60, (float64(g.height-menu.Header)-64)/3))
	y = float64(menu.Header+28) + float64(i)*(h+10)
	return
}
func (g *Game) drawMergeChoices(dst *ebiten.Image, p palette) {
	for i := 0; i < 3; i++ {
		x, y, w, h := g.mergeChoiceRect(i)
		rect(dst, x, y, w, h, p.panel)
		if i == g.option {
			vector.StrokeRect(dst, float32(x), float32(y), float32(w), float32(h), 2, p.a, true)
		}
		n := i + 4
		size := math.Min(h-24, 72)
		cell := size / float64(n)
		gx := x + 18
		gy := y + (h-size)/2
		for row := 0; row < n; row++ {
			for col := 0; col < n; col++ {
				rect(dst, gx+float64(col)*cell, gy+float64(row)*cell, cell-2, cell-2, p.edge)
			}
		}
		centered(dst, fmt.Sprintf("%d x %d", n, n), x+w*.65, y+h/2, p.ink)
	}
}
func (g *Game) prepareTerritorySetup() {
	g.opponents = 1
	g.territoryStarts = rand.Uint64()
	for i := range g.territoryMaps {
		g.territoryMaps[i] = rand.Uint64()
	}
	g.rebuildTerritoryDrafts()
}
func (g *Game) rebuildTerritoryDrafts() {
	for i, seed := range g.territoryMaps {
		g.territoryDrafts[i] = play.NewTerritoryMatch(seed, g.territoryStarts, g.opponents+1)
	}
}
func (g *Game) scrambleMaps() {
	for i := range g.territoryMaps {
		g.territoryMaps[i] = rand.Uint64()
	}
	g.rebuildTerritoryDrafts()
}
func (g *Game) scrambleStarts() { g.territoryStarts = rand.Uint64(); g.rebuildTerritoryDrafts() }
func (g *Game) territoryChoiceRect(i int) (x, y, w, h float64) {
	x, w = g.cardBounds()
	top := float64(menu.Header + 112)
	h = math.Min(240, math.Max(60, (float64(g.height)-top-28)/3))
	y = top + float64(i)*(h+6)
	return
}
func (g *Game) territorySetupTap(x, y float64) {
	if y >= menu.Header+8 && y < menu.Header+48 {
		if x >= 16 && x <= 80 {
			g.opponents = max(1, g.opponents-1)
			g.rebuildTerritoryDrafts()
		}
		if x >= float64(g.width-80) && x <= float64(g.width-16) {
			g.opponents = min(3, g.opponents+1)
			g.rebuildTerritoryDrafts()
		}
		return
	}
	if y >= menu.Header+54 && y < menu.Header+98 {
		if x < float64(g.width)/2 {
			g.scrambleMaps()
		} else {
			g.scrambleStarts()
		}
		return
	}
	for i := 0; i < 3; i++ {
		bx, by, bw, bh := g.territoryChoiceRect(i)
		if x >= bx && x <= bx+bw && y >= by && y <= by+bh {
			g.option = i
			g.startGame()
			return
		}
	}
}
func (g *Game) territorySetupKeys() {
	if inpututil.IsKeyJustPressed(ebiten.KeyMinus) {
		g.opponents = max(1, g.opponents-1)
		g.rebuildTerritoryDrafts()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEqual) {
		g.opponents = min(3, g.opponents+1)
		g.rebuildTerritoryDrafts()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyM) {
		g.scrambleMaps()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		g.scrambleStarts()
	}
}
func (g *Game) drawTerritorySetup(dst *ebiten.Image, p palette) {
	y := float64(menu.Header + 8)
	rect(dst, 16, y, 64, 40, p.panel)
	centered(dst, "-", 48, y+20, p.ink)
	rect(dst, float64(g.width-80), y, 64, 40, p.panel)
	centered(dst, "+", float64(g.width-48), y+20, p.ink)
	centered(dst, fmt.Sprintf("%d opponents", g.opponents), float64(g.width)/2, y+20, p.ink)
	half := float64(g.width) / 2
	for i, name := range []string{"Shuffle maps", "Shuffle starts"} {
		x := float64(i)*half + 16
		rect(dst, x, y+46, half-32, 42, p.panel)
		centered(dst, name, x+(half-32)/2, y+67, p.ink)
	}
	for i := 0; i < 3; i++ {
		x, y, w, h := g.territoryChoiceRect(i)
		rect(dst, x, y, w, h, p.panel)
		if i == g.option {
			vector.StrokeRect(dst, float32(x), float32(y), float32(w), float32(h), 2, p.a, true)
		}
		collection := play.Collection{Territory: g.territoryDrafts[i]}
		g.renderGame(dst, &collection, play.TerritoryID, x+8, y+4, w-16, h-8, p, false)
	}
}
