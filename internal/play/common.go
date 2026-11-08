// Package play contains deterministic game rules, independent of rendering and sensory settings.
package play

import "math"

const (
	MergeID = iota
	SnakeID
	TerritoryID
	LinearID
	BounceID
	FlowID
	CalibanID
	BloomID
	EchoID
	BrickID
	PlanetsID
	Count
)

var IDs = [Count]string{"merge", "snake", "territory", "linear", "bounce", "flow", "caliban", "bloom", "echo", "brick", "planets"}

// RNG's serialized state makes undo and save/resume reproducible.
type RNG uint64

func (r *RNG) Next(n int) int {
	x := uint64(*r)
	if x == 0 {
		x = 0x9e3779b97f4a7c15
	}
	x ^= x << 13
	x ^= x >> 7
	x ^= x << 17
	*r = RNG(x)
	if n <= 0 {
		return 0
	}
	return int(x % uint64(n))
}
func clamp(v, lo, hi float64) float64 { return math.Max(lo, math.Min(hi, v)) }

type Point struct{ X, Y float64 }
