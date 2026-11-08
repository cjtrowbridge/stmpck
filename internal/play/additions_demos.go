package play

func (c *Collection) demonstrateAdditions(n int) {
	if n%12 == 0 {
		for i := range c.Flow.Cells {
			if c.Flow.Cells[i] != c.Flow.Solution[i] {
				c.Flow.Rotate(i)
				break
			}
		}
	}
	if c.Flow.Won && n%180 == 0 {
		c.Flow = NewFlow(uint64(n))
	}
	a := &c.Caliban
	if a.Over {
		*a = NewCaliban(17)
	}
	if a.Cleared {
		a.Next()
	}
	if len(a.Enemies) > 0 {
		a.Lane = a.Enemies[0].Lane
	}
	a.Tick()
	if n%120 == 0 {
		c.Bloom = BloomDemonstration((n / 120) % len(BloomPatterns))
	}
	c.Bloom.Tick()
	e := &c.Echo
	e.Tick()
	if e.Phase == EchoInput && n%30 == 0 {
		e.Press(e.Sequence[e.Input])
	}
	if e.Phase == EchoWon {
		*e = NewEcho(19)
	}
	b := &c.Brick
	if b.Over {
		*b = NewBrick(0)
	}
	if b.Won {
		*b = NewBrick((b.Level + 1) % 3)
	}
	b.Move(b.X)
	if !b.Running && n%60 == 0 {
		b.Serve()
	}
	b.Tick()
	p := &c.Planets
	if p.Winner != 0 {
		*p = NewPlanets(1)
	}
	if n%120 == 0 {
		p.ComputerOrder(1)
	}
	p.Tick()
}
