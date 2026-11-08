package play

import "math"

type Crossing struct {
	X                       float64
	Y, Level, Lives, Shield int
	Phases                  [11]float64
	Homes                   [3]bool
	Over, Won               bool
}

func NewCrossing(seed uint64) Crossing {
	c := Crossing{X: 4.5, Y: 10, Level: 1, Lives: 3, Shield: 45}
	r := RNG(seed)
	for i := range c.Phases {
		c.Phases[i] = float64(r.Next(300)) / 100
	}
	return c
}
func (c *Crossing) LaneSpeed(row int) float64 {
	speed := (.32 + float64(row%3)*.10) * (1 + float64(c.Level-1)*.25)
	if row%2 == 0 {
		speed = -speed
	}
	return speed
}

// Repeated objects use a three-cell period, avoiding hidden discontinuities at wrap.
func (c *Crossing) OnObject(x float64, row int) bool {
	u := math.Mod(x-c.Phases[row]+30, 3)
	width := 1.05
	if row <= 5 {
		width = 2.1
	}
	return u < width
}
func (c *Crossing) die() {
	if c.Shield > 0 {
		return
	}
	c.Lives--
	c.Over = c.Lives == 0
	c.X = 4.5
	c.Y = 10
	c.Shield = 45
}
func (c *Crossing) check() {
	if c.Over || c.Won {
		return
	}
	if c.X < 0 || c.X >= 9 {
		c.die()
		return
	}
	if c.Y == 0 {
		home := int(c.X) / 3
		if absFloat(c.X-(1.5+float64(home)*3)) < .75 && !c.Homes[home] {
			c.Homes[home] = true
			c.X = 4.5
			c.Y = 10
			c.Shield = 45
			c.Won = c.Homes[0] && c.Homes[1] && c.Homes[2]
		} else {
			c.die()
		}
		return
	}
	if c.Y >= 7 && c.Y <= 9 && c.OnObject(c.X, c.Y) {
		c.die()
	}
	if c.Y >= 1 && c.Y <= 5 && !c.OnObject(c.X, c.Y) {
		c.die()
	}
}
func (c *Crossing) Move(d int) {
	if c.Over || c.Won || d < 0 || d > 3 {
		return
	}
	c.X = clamp(c.X+float64([4]int{-1, 0, 1, 0}[d]), .5, 8.5)
	c.Y = max(0, min(10, c.Y+[4]int{0, -1, 0, 1}[d]))
	c.Shield = 0
	c.check()
}
func (c *Crossing) Tick() {
	if c.Over || c.Won {
		return
	}
	if c.Shield > 0 {
		c.Shield--
	}
	for row := 1; row < 10; row++ {
		if row == 6 {
			continue
		}
		delta := c.LaneSpeed(row) / 60
		if c.Y == row && row <= 5 && c.OnObject(c.X, row) {
			c.X += delta
		}
		c.Phases[row] = math.Mod(c.Phases[row]+delta+3, 3)
	}
	c.check()
}
func (c *Crossing) NextLevel() {
	if c.Won && c.Level < 3 {
		c.Level++
		c.Homes = [3]bool{}
		c.Won = false
		c.X = 4.5
		c.Y = 10
		c.Shield = 45
	}
}
