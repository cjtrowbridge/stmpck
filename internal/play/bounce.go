package play

import "math"

// Normalized arena coordinates keep physics identical at every screen size.
type Bounce struct {
	X, Y, VX, VY, Player, Computer float64
	Points                         [2]int
	Speed                          int
	Running, Over                  bool
	ServeDirection                 int
	Random                         RNG
	Clock                          int
	AITarget                       float64
	Contacts                       int
}

func NewBounce(speed int, seed uint64) Bounce {
	b := Bounce{X: .5, Y: .5, Player: .5, Computer: .5, Speed: min(2, max(0, speed)), ServeDirection: 1, Random: RNG(seed), AITarget: .5}
	return b
}
func (b *Bounce) Serve() {
	if b.Over {
		return
	}
	if b.VX == 0 {
		speed := [3]float64{.55, .72, .9}[b.Speed]
		b.VX = speed * float64(b.ServeDirection)
		b.VY = (float64(b.Random.Next(2))*2 - 1) * speed * .35
	}
	b.Running = true
}
func (b *Bounce) Move(y float64) { b.Player = clamp(y, .12, .88) }
func (b *Bounce) Tick() bool {
	if !b.Running || b.Over {
		return false
	}
	b.Clock++
	if b.Clock%12 == 0 {
		b.AITarget = clamp(b.Y+(float64(b.Random.Next(21))-10)/180, .12, .88)
	}
	b.Computer += clamp(b.AITarget-b.Computer, -.42/60, .42/60)
	// Substeps and crossing tests prevent tunneling at the highest speed.
	for step := 0; step < 4; step++ {
		oldX := b.X
		b.X += b.VX / 240
		b.Y += b.VY / 240
		if b.Y < .018 {
			b.Y = .036 - b.Y
			b.VY = math.Abs(b.VY)
		}
		if b.Y > .982 {
			b.Y = 1.964 - b.Y
			b.VY = -math.Abs(b.VY)
		}
		hit := false
		paddle := b.Player
		direction := 1.0
		if b.VX < 0 && oldX >= .065 && b.X <= .065 && math.Abs(b.Y-b.Player) <= .13 {
			b.X = .065
			hit = true
		}
		if b.VX > 0 && oldX <= .935 && b.X >= .935 && math.Abs(b.Y-b.Computer) <= .13 {
			b.X = .935
			hit = true
			paddle = b.Computer
			direction = -1
		}
		if hit {
			speed := [3]float64{.55, .72, .9}[b.Speed]
			angle := clamp((b.Y-paddle)/.12, -1, 1) * .85
			b.VX = direction * speed * math.Cos(angle)
			b.VY = speed * math.Sin(angle)
			if math.Abs(b.VY) < .08 {
				b.VY = .08
			}
			b.Contacts++
		}
		if b.X < -.02 || b.X > 1.02 {
			scorer := 0
			if b.X < 0 {
				scorer = 1
			}
			b.Points[scorer]++
			b.ServeDirection = 1
			if scorer == 1 {
				b.ServeDirection = -1
			}
			b.X = .5
			b.Y = .5
			b.VX = 0
			b.VY = 0
			b.Running = false
			b.Over = b.Points[scorer] >= 5
			return true
		}
	}
	return true
}
