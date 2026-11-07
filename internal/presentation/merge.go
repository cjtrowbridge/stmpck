package presentation

import (
	"image/color"
	"math"
	"math/bits"
)

// MergeColor samples a continuous hue/saturation ramp at each exponent, rather
// than alternating a small palette. The nonintegral hue step avoids repeated
// colors within the game's supported exponents 1..30. Gray mode remains gray.
func MergeColor(value int, gray bool) color.RGBA {
	exponent := max(1, bits.Len(uint(max(2, value)))-1)
	if gray {
		v := uint8(235 - 5*exponent)
		return color.RGBA{v, v, v, 255}
	}
	h := math.Mod(float64(exponent-1)*.087, 1) * 6
	saturation := .5 + .01*float64(exponent)
	v := .94
	c := v * saturation
	x := c * (1 - math.Abs(math.Mod(h, 2)-1))
	m := v - c
	var r, g, b float64
	switch int(h) {
	case 0:
		r, g, b = c, x, 0
	case 1:
		r, g, b = x, c, 0
	case 2:
		r, g, b = 0, c, x
	case 3:
		r, g, b = 0, x, c
	case 4:
		r, g, b = x, 0, c
	default:
		r, g, b = c, 0, x
	}
	return color.RGBA{uint8((r + m) * 255), uint8((g + m) * 255), uint8((b + m) * 255), 255}
}
func TileInk(c color.RGBA) color.RGBA {
	channel := func(v uint8) float64 {
		x := float64(v) / 255
		if x <= .04045 {
			return x / 12.92
		}
		return math.Pow((x+.055)/1.055, 2.4)
	}
	luminance := .2126*channel(c.R) + .7152*channel(c.G) + .0722*channel(c.B)
	if luminance > .179 {
		return color.RGBA{0, 0, 0, 255}
	}
	return color.RGBA{255, 255, 255, 255}
}
