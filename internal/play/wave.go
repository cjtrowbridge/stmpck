package play

type WaveShot struct {
	X, Y  float64
	Enemy bool
}
type Wave struct {
	Enemies                           [32]bool
	X, Y, Direction, Player           float64
	Shots                             []WaveShot
	Level, Lives, Clock, Shield, Hits int
	RNG                               RNG
	Over, Cleared                     bool
}

func NewWave(seed uint64) Wave {
	w := Wave{Level: 1, Lives: 3, RNG: RNG(seed)}
	w.startWave()
	return w
}
func (w *Wave) startWave() {
	for i := range w.Enemies {
		w.Enemies[i] = true
	}
	w.X = .08
	w.Y = .12
	w.Direction = 1
	w.Player = .5
	w.Clock = 0
	w.Shield = 120
	w.Shots = nil
	w.Cleared = false
}
func (w *Wave) Position(i int) Point { return Point{w.X + float64(i%8)*.095, w.Y + float64(i/8)*.075} }
func (w *Wave) Move(x float64)       { w.Player = clamp(x, .04, .96) }
func (w *Wave) NextWave() {
	if w.Cleared && w.Level < 3 {
		w.Level++
		w.startWave()
	}
}
func (w *Wave) Tick() {
	if w.Over || w.Cleared {
		return
	}
	w.Clock = (w.Clock + 1) % 180
	if w.Shield > 0 {
		w.Shield--
	}
	w.X += w.Direction * (.045 + float64(w.Level)*.015) / 60
	if w.X < .035 {
		w.X = .035
		w.Direction = 1
		w.Y += .035
	}
	if w.X > .30 {
		w.X = .30
		w.Direction = -1
		w.Y += .035
	}
	if w.Clock%18 == 0 {
		w.Shots = append(w.Shots, WaveShot{w.Player, .87, false})
	}
	if w.Clock%max(40, 105-w.Level*15) == 0 {
		bottom := []int{}
		for col := 0; col < 8; col++ {
			for row := 3; row >= 0; row-- {
				i := row*8 + col
				if w.Enemies[i] {
					bottom = append(bottom, i)
					break
				}
			}
		}
		if len(bottom) > 0 {
			p := w.Position(bottom[w.RNG.Next(len(bottom))])
			w.Shots = append(w.Shots, WaveShot{p.X, p.Y + .02, true})
		}
	}
	shots := w.Shots[:0]
	for _, s := range w.Shots {
		hit := false
		if s.Enemy {
			s.Y += (.2 + float64(w.Level)*.03) / 60
			if s.Y >= .875 && s.Y <= .925 && absFloat(s.X-w.Player) < .04 {
				hit = true
				if w.Shield == 0 {
					w.Lives--
					w.Shield = 90
					w.Over = w.Lives == 0
				}
			}
		} else {
			s.Y -= .65 / 60
			for i, alive := range w.Enemies {
				p := w.Position(i)
				if alive && absFloat(s.X-p.X) < .031 && absFloat(s.Y-p.Y) < .026 {
					w.Enemies[i] = false
					w.Hits++
					hit = true
					break
				}
			}
		}
		if !hit && s.Y >= 0 && s.Y <= 1 {
			shots = append(shots, s)
		}
	}
	w.Shots = shots
	remaining := false
	for i, alive := range w.Enemies {
		if alive {
			remaining = true
			if w.Position(i).Y >= .82 {
				w.Over = true
			}
		}
	}
	w.Cleared = !w.Over && !remaining
}
func absFloat(v float64) float64 {
	if v < 0 {
		return -v
	}
	return v
}
