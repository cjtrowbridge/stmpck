package play

const CalibanLanes = 12

type TubeObject struct {
	Lane  int
	Depth float64
}
type Caliban struct {
	Lane, Wave, Lives, Clock, Spawned, Hits int
	Enemies, Shots                          []TubeObject
	RNG                                     RNG
	Cleared, Over                           bool
}

func NewCaliban(seed uint64) Caliban { return Caliban{Wave: 1, Lives: 3, RNG: RNG(seed)} }
func (c *Caliban) Move(delta int) {
	if !c.Over {
		c.Lane = (c.Lane + delta + CalibanLanes) % CalibanLanes
	}
}
func (c *Caliban) Next() {
	if c.Cleared {
		c.Wave++
		c.Clock = 0
		c.Spawned = 0
		c.Shots = nil
		c.Cleared = false
	}
}
func (c *Caliban) Tick() {
	if c.Over || c.Cleared {
		return
	}
	c.Clock++
	if c.Clock%12 == 0 {
		c.Shots = append(c.Shots, TubeObject{c.Lane, 1})
	}
	quota := 5 + min(c.Wave, 20)*2
	if c.Clock%max(24, 90-c.Wave*3) == 0 && c.Spawned < quota {
		c.Enemies = append(c.Enemies, TubeObject{c.RNG.Next(CalibanLanes), 0})
		c.Spawned++
	}
	for i := range c.Enemies {
		c.Enemies[i].Depth += (.1 + float64(min(c.Wave, 20))*.008) / 60
	}
	shots := c.Shots[:0]
	for _, s := range c.Shots {
		old := s.Depth
		s.Depth -= 1.5 / 60
		hit := -1
		for i, e := range c.Enemies {
			if e.Lane == s.Lane && e.Depth >= s.Depth-.015 && e.Depth <= old+.02 {
				hit = i
				break
			}
		}
		if hit >= 0 {
			c.Enemies = append(c.Enemies[:hit], c.Enemies[hit+1:]...)
			c.Hits++
		} else if s.Depth >= 0 {
			shots = append(shots, s)
		}
	}
	c.Shots = shots
	enemies := c.Enemies[:0]
	for _, e := range c.Enemies {
		if e.Depth >= 1 {
			c.Lives--
		} else {
			enemies = append(enemies, e)
		}
	}
	c.Enemies = enemies
	if c.Lives <= 0 {
		c.Lives = 0
		c.Over = true
	}
	c.Cleared = !c.Over && c.Spawned == quota && len(c.Enemies) == 0
}
