package play

// Demonstrations use real rules and never touch the player's saved sessions.
func NewDemonstrations() Collection {
	c := NewCollection()
	for id := 0; id < Count; id++ {
		c.New(id, 0, uint64(id+11))
	}
	c.Snake.Running = true
	c.Bounce.Serve()
	c.Bloom = BloomDemonstration(0)
	return c
}
func (c *Collection) Demonstrate(n int) {
	if n%40 == 0 {
		c.Merge.Move([4]int{0, 3, 2, 1}[(n/40)%4])
	}
	s := &c.Snake
	if s.Over {
		*s = NewSnake(0, 12)
		s.Running = true
	}
	if n%12 == 0 {
		h := s.Body[0]
		food := s.Food
		if h%SnakeW < food%SnakeW {
			s.Turn(2)
		} else if h%SnakeW > food%SnakeW {
			s.Turn(0)
		} else if h/SnakeW < food/SnakeW {
			s.Turn(3)
		} else {
			s.Turn(1)
		}
	}
	s.Tick()
	t := &c.Territory
	if t.Winner != 0 {
		*t = NewTerritory(31)
	}
	if n%TerritoryAITicks == 0 {
		t.ComputerOrder(1)
	}
	t.Tick()

	if n%45 == 0 {
		l := &c.Linear
		step := (n/45 - 1) % 9
		if step == 0 {
			*l = NewLinear(true, 1)
		}
		if step < 7 {
			l.Drop([]int{0, 0, 1, 1, 2, 2, 3}[step])
		}
	}
	b := &c.Bounce
	if n < 240 {
		b.Move(b.Y + .02)
	} else {
		b.Move(.15)
	}
	if b.Over {
		*b = NewBounce(0, 15)
	}
	if !b.Running && n%60 == 0 {
		b.Serve()
	}
	b.Tick()
	c.demonstrateAdditions(n)
	c.demonstrateExpansion(n)
}
