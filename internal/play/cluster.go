package play

import "math"

const (
	ClusterW = 8
	ClusterH = 10
)

type Cluster struct {
	Cells                     [ClusterW * ClusterH]int
	Color, Next, Score, Shots int
	Aim, X, Y, VX, VY         float64
	Flying, Won, Over         bool
	RNG                       RNG
}

func ClusterPosition(i int) Point {
	return Point{.065 + .115*(float64(i%ClusterW)+.5*float64(i/ClusterW%2)), .06 + .10*float64(i/ClusterW)}
}
func ClusterNeighbors(i int) []int {
	row, col := i/ClusterW, i%ClusterW
	out := []int{}
	add := func(x, y int) {
		if x >= 0 && x < ClusterW && y >= 0 && y < ClusterH {
			out = append(out, y*ClusterW+x)
		}
	}
	add(col-1, row)
	add(col+1, row)
	shift := -1
	if row%2 == 1 {
		shift = 0
	}
	for _, dy := range []int{-1, 1} {
		add(col+shift, row+dy)
		add(col+shift+1, row+dy)
	}
	return out
}
func NewCluster(seed uint64) Cluster {
	c := Cluster{RNG: RNG(seed), X: .5, Y: .93}
	for i := 0; i < 5*ClusterW; i++ {
		c.Cells[i] = 1 + c.RNG.Next(4)
	}
	c.Color = c.pick()
	c.Next = c.pick()
	return c
}
func (c *Cluster) pick() int {
	available := [5]bool{}
	for _, v := range c.Cells {
		available[v] = true
	}
	colors := []int{}
	for v := 1; v <= 4; v++ {
		if available[v] {
			colors = append(colors, v)
		}
	}
	if len(colors) == 0 {
		return 1
	}
	return colors[c.RNG.Next(len(colors))]
}
func (c *Cluster) SetAim(a float64) {
	if !c.Flying {
		c.Aim = clamp(a, -1.1, 1.1)
	}
}
func (c *Cluster) Fire() bool {
	if c.Flying || c.Over || c.Won {
		return false
	}
	c.X = .5
	c.Y = .93
	c.VX = .85 * math.Sin(c.Aim)
	c.VY = -.85 * math.Cos(c.Aim)
	c.Flying = true
	c.Shots++
	return true
}
func (c *Cluster) component(start int, same bool) []int {
	seen := [80]bool{}
	seen[start] = true
	queue := []int{start}
	for k := 0; k < len(queue); k++ {
		for _, i := range ClusterNeighbors(queue[k]) {
			if !seen[i] && c.Cells[i] != 0 && (!same || c.Cells[i] == c.Cells[start]) {
				seen[i] = true
				queue = append(queue, i)
			}
		}
	}
	return queue
}
func (c *Cluster) resolve(at int) {
	group := c.component(at, true)
	if len(group) >= 3 {
		for _, i := range group {
			c.Cells[i] = 0
			c.Score++
		}
		connected := [80]bool{}
		for i := 0; i < ClusterW; i++ {
			if c.Cells[i] != 0 && !connected[i] {
				for _, j := range c.component(i, false) {
					connected[j] = true
				}
			}
		}
		for i, v := range c.Cells {
			if v != 0 && !connected[i] {
				c.Cells[i] = 0
				c.Score++
			}
		}
	}
	remaining := 0
	for i, v := range c.Cells {
		if v != 0 {
			remaining++
			if i/ClusterW == ClusterH-1 {
				c.Over = true
			}
		}
	}
	c.Won = remaining == 0
	c.Flying = false
	c.X = .5
	c.Y = .93
	c.Color = c.Next
	present := false
	for _, v := range c.Cells {
		present = present || v == c.Color
	}
	if !present {
		c.Color = c.pick()
	}
	c.Next = c.pick()
}
func (c *Cluster) attach(hit int) {
	candidates := []int{}
	if hit < 0 {
		for i := 0; i < ClusterW; i++ {
			candidates = append(candidates, i)
		}
	} else {
		candidates = ClusterNeighbors(hit)
	}
	best := -1
	distance := math.Inf(1)
	for _, i := range candidates {
		if c.Cells[i] != 0 {
			continue
		}
		p := ClusterPosition(i)
		d := math.Hypot(c.X-p.X, c.Y-p.Y)
		if d < distance {
			best = i
			distance = d
		}
	}
	if best < 0 {
		c.Over = true
		c.Flying = false
		return
	}
	c.Cells[best] = c.Color
	c.resolve(best)
}
func (c *Cluster) Tick() {
	if !c.Flying || c.Over || c.Won {
		return
	}
	for step := 0; step < 4; step++ {
		c.X += c.VX / 240
		c.Y += c.VY / 240
		if c.X < .055 {
			c.X = .11 - c.X
			c.VX = -c.VX
		}
		if c.X > .945 {
			c.X = 1.89 - c.X
			c.VX = -c.VX
		}
		hit := -1
		best := .109
		for i, v := range c.Cells {
			if v == 0 {
				continue
			}
			p := ClusterPosition(i)
			dist := math.Hypot(c.X-p.X, c.Y-p.Y)
			if dist < best {
				hit = i
				best = dist
			}
		}
		if hit >= 0 || c.Y <= .06 {
			c.attach(hit)
			return
		}
	}
}
