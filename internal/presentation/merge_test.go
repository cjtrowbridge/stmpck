package presentation

import (
	"image/color"
	"testing"
)

func TestUniqueMergeColors(t *testing.T) {
	for _, gray := range []bool{false, true} {
		seen := map[color.RGBA]bool{}
		for e := 1; e <= 30; e++ {
			c := MergeColor(1<<e, gray)
			if seen[c] {
				t.Fatalf("repeated color at exponent %d", e)
			}
			seen[c] = true
			if gray && (c.R != c.G || c.G != c.B) {
				t.Fatal("wireframe lost monochrome")
			}
		}
	}
}
