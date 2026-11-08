package play

import (
	"errors"
	"math"
)

func (c *Collection) validateAdditions() error {
	bad := func() error { return errors.New("invalid additional game save") }
	finite := func(v float64) bool { return !math.IsNaN(v) && !math.IsInf(v, 0) }
	if c.Ready[FlowID] {
		f := &c.Flow
		if f.Moves < 0 {
			return bad()
		}
		for _, board := range [][25]uint8{f.Cells, f.Initial, f.Solution} {
			for _, v := range board {
				if v == 0 || v > 15 {
					return bad()
				}
			}
		}
		f.Won = f.Complete()
	}
	if c.Ready[BloomID] {
		b := &c.Bloom
		if b.Speed < 0 || b.Speed > 2 || b.Clock < 0 || b.Clock >= 30 {
			return bad()
		}
	}
	if c.Ready[EchoID] {
		e := &c.Echo
		if len(e.Sequence) < 1 || len(e.Sequence) > 32 || e.Phase < 0 || e.Phase > EchoWon || e.Clock < 0 || e.Clock > 32*45 || e.Input < 0 || e.Input > len(e.Sequence) || e.Flash < 0 || e.Flash > 12 || e.Pad < -1 || e.Pad > 3 {
			return bad()
		}
		if e.Phase == EchoInput && e.Input >= len(e.Sequence) {
			return bad()
		}
		for _, v := range e.Sequence {
			if v < 0 || v > 3 {
				return bad()
			}
		}
	}
	if c.Ready[CalibanID] {
		v := &c.Caliban
		if v.Lane < 0 || v.Lane >= CalibanLanes || v.Wave < 1 || v.Wave > 100000 || v.Lives < 0 || v.Lives > 3 || v.Clock < 0 || v.Spawned < 0 || v.Spawned > 45 || v.Hits < 0 || len(v.Enemies) > 45 || len(v.Shots) > 50 {
			return bad()
		}
		for _, objects := range [][]TubeObject{v.Enemies, v.Shots} {
			for _, o := range objects {
				if o.Lane < 0 || o.Lane >= CalibanLanes || !finite(o.Depth) || o.Depth < 0 || o.Depth > 1 {
					return bad()
				}
			}
		}
	}
	if c.Ready[BrickID] {
		b := &c.Brick
		if b.Level < 0 || b.Level > 2 || b.Lives < 0 || b.Lives > 3 || b.Hits < 0 || b.Hits > 40 {
			return bad()
		}
		for _, v := range []float64{b.X, b.Y, b.VX, b.VY, b.Paddle} {
			if !finite(v) || math.Abs(v) > 2 {
				return bad()
			}
		}
		if b.X < 0 || b.X > 1 || b.Y < 0 || b.Y > 1.03 || b.Paddle < .1 || b.Paddle > .9 {
			return bad()
		}
	}
	if c.Ready[PlanetsID] {
		p := &c.Planets
		if p.System < 0 || p.System > 2 || p.Players != 2+p.System || len(p.Bodies) != 5+2*p.System || p.Source < -1 || p.Source >= len(p.Bodies) || p.Winner < 0 || p.Winner > p.Players || p.ProductionClock < 0 || p.ProductionClock >= 180 || p.AIClock < 0 || p.AIClock >= 120 || len(p.Fleets) > 256 {
			return bad()
		}
		for i, b := range p.Bodies {
			if b.Parent < -1 || b.Parent >= i || b.Parent >= 0 && p.Bodies[b.Parent].Parent != -1 || b.Owner < 0 || b.Owner > p.Players || b.Units < 0 || b.Units > 100000 || b.Owner == 0 && b.Units != 0 || b.Owner > 0 && b.Units < 1 || !finite(b.Radius) || b.Radius <= 0 || b.Radius > .45 || !finite(b.Phase) || b.Phase < 0 || b.Phase >= 2*math.Pi || !finite(b.Rate) || b.Rate < 0 || b.Rate > .2 {
				return bad()
			}
		}
		for _, f := range p.Fleets {
			if f.Target < 0 || f.Target >= len(p.Bodies) || f.Owner < 1 || f.Owner > p.Players || f.Units < 1 || f.Units > 100000 || !finite(f.X) || !finite(f.Y) || f.X < 0 || f.X > 1 || f.Y < 0 || f.Y > 1 {
				return bad()
			}
		}
	}
	return nil
}
