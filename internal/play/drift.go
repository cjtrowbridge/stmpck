package play

import "math"

type SpaceRock struct {
	X, Y, VX, VY float64
	Size         int
}
type SpaceShot struct {
	X, Y, VX, VY float64
	Life         int
}
type Drift struct {
	X, Y, VX, VY, Angle               float64
	Rocks                             []SpaceRock
	Shots                             []SpaceShot
	Level, Lives, Clock, Shield, Hits int
	Thrust, Over, Cleared             bool
	RNG                               RNG
}

func wrapUnit(v float64) float64 { return v - math.Floor(v) }
func spaceDistance(x, y, xx, yy float64) float64 {
	dx, dy := math.Abs(x-xx), math.Abs(y-yy)
	return math.Hypot(math.Min(dx, 1-dx), math.Min(dy, 1-dy))
}
func RockRadius(size int) float64 { return .018 + float64(size)*.018 }
func NewDrift(seed uint64) Drift {
	d := Drift{Level: 1, Lives: 3, RNG: RNG(seed)}
	d.startWave()
	return d
}
func (d *Drift) startWave() {
	d.X = .5
	d.Y = .5
	d.VX = 0
	d.VY = 0
	d.Angle = -math.Pi / 2
	d.Shield = 120
	d.Clock = 0
	d.Thrust = false
	d.Cleared = false
	d.Shots = nil
	d.Rocks = nil
	for i := 0; i < 2+d.Level; i++ {
		a := float64(i) * 2 * math.Pi / float64(2+d.Level)
		heading := float64(d.RNG.Next(628)) / 100
		speed := .04 + float64(d.Level)*.012
		d.Rocks = append(d.Rocks, SpaceRock{.5 + .37*math.Cos(a), .5 + .37*math.Sin(a), speed * math.Cos(heading), speed * math.Sin(heading), 2})
	}
}
func (d *Drift) NextWave() {
	if d.Cleared && d.Level < 3 {
		d.Level++
		d.startWave()
	}
}
func (d *Drift) Turn(delta float64) { d.Angle = math.Remainder(d.Angle+delta, 2*math.Pi) }
func (d *Drift) Tick() {
	if d.Over || d.Cleared {
		return
	}
	d.Clock = (d.Clock + 1) % 18
	if d.Shield > 0 {
		d.Shield--
	}
	if d.Thrust {
		d.VX += math.Cos(d.Angle) * .003
		d.VY += math.Sin(d.Angle) * .003
	}
	speed := math.Hypot(d.VX, d.VY)
	if speed > .35 {
		d.VX *= .35 / speed
		d.VY *= .35 / speed
	}
	d.X = wrapUnit(d.X + d.VX/60)
	d.Y = wrapUnit(d.Y + d.VY/60)
	if d.Clock == 0 {
		d.Shots = append(d.Shots, SpaceShot{d.X, d.Y, .7 * math.Cos(d.Angle), .7 * math.Sin(d.Angle), 75})
	}
	for i := range d.Rocks {
		r := &d.Rocks[i]
		r.X = wrapUnit(r.X + r.VX/60)
		r.Y = wrapUnit(r.Y + r.VY/60)
	}
	shots := d.Shots[:0]
	for _, s := range d.Shots {
		s.X = wrapUnit(s.X + s.VX/60)
		s.Y = wrapUnit(s.Y + s.VY/60)
		s.Life--
		hit := -1
		for i, r := range d.Rocks {
			if spaceDistance(s.X, s.Y, r.X, r.Y) < RockRadius(r.Size)+.006 {
				hit = i
				break
			}
		}
		if hit >= 0 {
			r := d.Rocks[hit]
			d.Rocks = append(d.Rocks[:hit], d.Rocks[hit+1:]...)
			d.Hits++
			if r.Size > 0 {
				for _, sign := range []float64{-1, 1} {
					angle := math.Atan2(r.VY, r.VX) + sign*.7
					speed := math.Hypot(r.VX, r.VY) * 1.3
					d.Rocks = append(d.Rocks, SpaceRock{r.X, r.Y, speed * math.Cos(angle), speed * math.Sin(angle), r.Size - 1})
				}
			}
		} else if s.Life > 0 {
			shots = append(shots, s)
		}
	}
	d.Shots = shots
	if d.Shield == 0 {
		for _, r := range d.Rocks {
			if spaceDistance(d.X, d.Y, r.X, r.Y) < RockRadius(r.Size)+.014 {
				d.Lives--
				d.X = .5
				d.Y = .5
				d.VX = 0
				d.VY = 0
				d.Shield = 120
				d.Over = d.Lives == 0
				break
			}
		}
	}
	d.Cleared = !d.Over && len(d.Rocks) == 0
}
