package play

import "math"

func (c *Collection) demonstrateExpansion(n int) {
	d := &c.Drift
	if d.Over {
		*d = NewDrift(31)
	}
	if d.Cleared {
		if d.Level < 3 {
			d.NextWave()
		} else {
			*d = NewDrift(31)
		}
	}
	if len(d.Rocks) > 0 {
		d.Angle = math.Atan2(d.Rocks[0].Y-d.Y, d.Rocks[0].X-d.X)
	}
	d.Tick()
	s := &c.Stack
	if s.Over || s.Won {
		*s = NewStack(32)
	}
	if n%30 == 0 {
		bestX, bestRot, bestScore := s.X, s.Rotation, math.Inf(1)
		for rot := 0; rot < 4; rot++ {
			for x := -2; x < 10; x++ {
				if !s.Fits(x, 0, rot) {
					continue
				}
				y := 0
				for s.Fits(x, y+1, rot) {
					y++
				}
				trial := *s
				trial.X = x
				trial.Y = y
				trial.Rotation = rot
				trial.lock()
				height, holes := 0, 0
				for col := 0; col < 10; col++ {
					seen := false
					for row := 0; row < 20; row++ {
						if trial.Cells[row*10+col] != 0 {
							if !seen {
								height += 20 - row
							}
							seen = true
						} else if seen {
							holes++
						}
					}
				}
				score := float64(height + holes*12 - (trial.Lines-s.Lines)*20)
				if score < bestScore {
					bestX, bestRot, bestScore = x, rot, score
				}
			}
		}
		for s.Rotation != bestRot {
			if !s.Rotate() {
				break
			}
		}
		for s.X != bestX {
			dx := 1
			if s.X > bestX {
				dx = -1
			}
			if !s.Move(dx) {
				break
			}
		}
		s.Drop()
	}
	if n%18 == 0 {
		for i, v := range c.Deduction.Cells {
			if v == 0 {
				c.Deduction.Set(i, c.Deduction.Solution[i])
				break
			}
		}
	}
	cl := &c.Cluster
	if cl.Over || cl.Won {
		*cl = NewCluster(uint64(n + 34))
	}
	if !cl.Flying && n%24 == 0 {
		best, score := 0.0, -1
		for step := -8; step <= 8; step += 4 {
			trial := *cl
			trial.SetAim(float64(step) / 10)
			trial.Fire()
			for ticks := 0; ticks < 300 && trial.Flying; ticks++ {
				trial.Tick()
			}
			value := trial.Score - cl.Score
			if trial.Over {
				value -= 100
			}
			if value > score {
				score = value
				best = float64(step) / 10
			}
		}
		cl.SetAim(best)
		cl.Fire()
	}
	cl.Tick()
	w := &c.Wave
	if w.Over {
		*w = NewWave(35)
	}
	if w.Cleared {
		if w.Level < 3 {
			w.NextWave()
		} else {
			*w = NewWave(35)
		}
	}
	for i, alive := range w.Enemies {
		if alive {
			w.Move(w.Position(i).X)
			break
		}
	}
	w.Tick()
	cr := &c.Crossing
	if cr.Over {
		*cr = NewCrossing(36)
	}
	if cr.Won {
		if cr.Level < 3 {
			cr.NextLevel()
		} else {
			*cr = NewCrossing(36)
		}
	}
	if n%24 == 0 {
		for _, dir := range []int{1, 0, 2, 3} {
			trial := *cr
			trial.Shield = 0
			trial.Move(dir)
			for k := 0; k < 8; k++ {
				trial.Tick()
			}
			if trial.Lives == cr.Lives {
				cr.Move(dir)
				break
			}
		}
	}
	cr.Tick()
	if n%20 == 0 {
		sh := &c.Shift
		if sh.Won {
			*sh = NewShift(37)
		}
		for k := 0; k < 4; k++ {
			if sh.Direction((n/20 + k) % 4) {
				break
			}
		}
	}
}
