package app

import (
	"image/color"
	"math"

	"ebe-boilerplate/internal/menu"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// Four opt-in effect levels; at most 29 sparkles are alive at once.
var sparkleRates = [...]float64{2, 5, 10, 18}

// One full expansion/contraction per beat, from a resting-style breath to dance tempo.
var breathBPM = [...]float64{15, 45, 120, 180}
var rainbowPeriods = [...]float64{0, 24, 16, 10}
var rainbowAmounts = [...]float64{0, .25, .55, .85}

func mixColor(a, b color.RGBA, amount float64) color.RGBA {
	mix := func(x, y uint8) uint8 { return uint8(float64(x)*(1-amount) + float64(y)*amount) }
	return color.RGBA{mix(a.R, b.R), mix(a.G, b.G), mix(a.B, b.B), 255}
}

func rainbow(phase float64) color.RGBA {
	channel := func(offset float64) uint8 { return uint8(155 + 100*math.Sin(2*math.Pi*(phase+offset))) }
	return color.RGBA{channel(0), channel(1.0 / 3), channel(2.0 / 3), 255}
}

func (g *Game) effectPalette(p palette) palette {
	if g.Theme <= menu.Breathe {
		return p
	}
	level := g.Theme - menu.Breathe
	phase := g.tick / 60 / rainbowPeriods[level]
	p.a = mixColor(p.a, rainbow(phase), rainbowAmounts[level])
	p.b = mixColor(p.b, rainbow(phase+.35), rainbowAmounts[level])
	return p
}

func (g *Game) breathingBorder(dst *ebiten.Image, x, y, w, h float64, p palette) {
	if g.Theme < menu.Breathe {
		return
	}
	level := g.Theme - menu.Breathe
	breath := .5 - .5*math.Cos(2*math.Pi*g.tick/3600*breathBPM[level])
	ink := mixColor(p.edge, p.a, .2+.65*breath)
	width := 1 + breath*(1+float64(level)*.5)
	vector.StrokeRect(dst, float32(x+3), float32(y+3), float32(w-6), float32(h-6), float32(width), ink, true)
}

// A deterministic hash gives each birth its own position without per-frame randomness.
func sparkleRandom(seed uint64) float64 {
	seed ^= seed >> 30
	seed *= 0xbf58476d1ce4e5b9
	seed ^= seed >> 27
	seed *= 0x94d049bb133111eb
	seed ^= seed >> 31
	return float64(seed>>11) / float64(uint64(1)<<53)
}

func (g *Game) drawSparkles(dst *ebiten.Image, p palette) {
	if g.Theme < menu.Breathe {
		return
	}
	rate := sparkleRates[g.Theme-menu.Breathe]
	now := g.tick / 60
	const life = 1.6
	for birth := max(0, int(math.Ceil((now-life)*rate))); birth <= int(now*rate); birth++ {
		age := now - float64(birth)/rate
		if age < 0 || age >= life {
			continue
		}
		seed := g.sparkleSeed + uint64(birth)*3
		x := 8 + sparkleRandom(seed)*math.Max(0, float64(g.width)-16)
		y := 8 + sparkleRandom(seed+1)*math.Max(0, float64(g.height)-16)
		// Preserve the header controls' hit targets and legibility.
		if y < float64(g.HeaderHeight()) && (y < 92 || x < 90 || x > float64(g.width-90)) {
			continue
		}
		envelope := math.Pow(math.Sin(math.Pi*age/life), 2)
		size := (2 + 3*sparkleRandom(seed+2)) * envelope
		ink := mixColor(p.a, p.b, sparkleRandom(seed+2))
		alpha := .7 * envelope
		ink = color.RGBA{uint8(float64(ink.R) * alpha), uint8(float64(ink.G) * alpha), uint8(float64(ink.B) * alpha), uint8(255 * alpha)}
		line(dst, x-size, y, x+size, y, 1, ink)
		line(dst, x, y-size, x, y+size, 1, ink)
	}
}
