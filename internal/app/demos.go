package app

import "ebe-boilerplate/internal/play"

func (g *Game) makeDemos() {
	g.demos = play.NewDemonstrations()
	g.demoClock = 0
	g.demoMergeAnimation = 0
}
func (g *Game) updateDemos() {
	g.demoClock++
	if g.demoClock >= 720 {
		g.makeDemos()
		return
	}
	if g.demoMergeAnimation > 0 {
		g.demoMergeAnimation--
	}
	before := g.demos.Merge.Cells
	g.demos.Demonstrate(g.demoClock)
	if before != g.demos.Merge.Cells {
		g.demoMergeAnimation = 12
	}
}
