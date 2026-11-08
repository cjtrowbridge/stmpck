package play

import "math"

// Parent -1 denotes a planet orbiting the star; otherwise it is a moon.
type OrbitalBody struct {
	Parent, Owner, Units int
	Radius, Phase, Rate  float64
}
type Fleet struct {
	Target, Owner, Units int
	X, Y                 float64
}
type Planets struct {
	System, Players, Source, Winner, ProductionClock, AIClock int
	TimingVersion, ActiveTicks, AIInterval                    int
	Started                                                   bool
	Bodies                                                    []OrbitalBody
	Fleets                                                    []Fleet
}

const FleetSpeed = .28
const PlanetsRampTicks = 180 * 60

func NewPlanets(system int) Planets {
	system = max(0, min(2, system))
	p := Planets{System: system, Players: 2 + system, Source: -1, TimingVersion: 1, AIInterval: 360}
	for i := 0; i < 5+system; i++ {
		p.Bodies = append(p.Bodies, OrbitalBody{Parent: -1, Radius: .13 + float64(i)*.045, Phase: math.Mod(float64(i)*2.4, 2*math.Pi), Rate: .025 + float64(i)*.004})
	}
	for owner := 1; owner <= p.Players; owner++ {
		i := (owner - 1) * len(p.Bodies) / p.Players
		p.Bodies[i].Owner = owner
		p.Bodies[i].Units = 10
	}
	for i := 0; i < system; i++ {
		p.Bodies = append(p.Bodies, OrbitalBody{Parent: 1 + i*2, Radius: .055, Phase: float64(i) * 2, Rate: .17})
	}
	return p
}
func (p *Planets) Position(i int) Point {
	b := p.Bodies[i]
	center := Point{.5, .5}
	if b.Parent >= 0 {
		center = p.Position(b.Parent)
	}
	return Point{center.X + b.Radius*math.Cos(b.Phase), center.Y + b.Radius*math.Sin(b.Phase)}
}
func (p *Planets) At(x, y float64) int {
	best := -1
	distance := .045
	for i := range p.Bodies {
		v := p.Position(i)
		d := math.Hypot(v.X-x, v.Y-y)
		if d < distance {
			best = i
			distance = d
		}
	}
	return best
}
func (p *Planets) Order(owner, from, to int) bool {
	if owner != 1 && !p.Started {
		return false
	}
	if p.Winner != 0 || from < 0 || from >= len(p.Bodies) || to < 0 || to >= len(p.Bodies) || from == to || len(p.Fleets) >= 256 {
		return false
	}
	b := &p.Bodies[from]
	if b.Owner != owner || b.Units < 2 {
		return false
	}
	if !p.Started {
		p.Started = true
		p.ActiveTicks = 0
		p.AIClock = 0
		p.AIInterval = 360
	}
	pos := p.Position(from)
	units := b.Units - 1
	b.Units = 1
	p.Fleets = append(p.Fleets, Fleet{to, owner, units, pos.X, pos.Y})
	return true
}
func (p *Planets) Choose(i int) {
	if i < 0 || i >= len(p.Bodies) || p.Winner != 0 {
		return
	}
	if p.Source >= 0 && i != p.Source && p.Order(1, p.Source, i) {
		p.Source = -1
		return
	}
	p.Source = -1
	if p.Bodies[i].Owner == 1 {
		p.Source = i
	}
}
func (p *Planets) ComputerOrder(owner int) {
	if owner != 1 && !p.Started {
		return
	}
	from := -1
	for i, b := range p.Bodies {
		if b.Owner == owner && (from < 0 || b.Units > p.Bodies[from].Units) {
			from = i
		}
	}
	if from < 0 || p.Bodies[from].Units < 4 {
		return
	}
	to := -1
	best := math.Inf(1)
	a := p.Position(from)
	for i, b := range p.Bodies {
		if b.Owner == owner {
			continue
		}
		pos := p.Position(i)
		score := float64(b.Units) + math.Hypot(pos.X-a.X, pos.Y-a.Y)*5
		if score < best {
			best = score
			to = i
		}
	}
	if to >= 0 && p.Bodies[from].Units > p.Bodies[to].Units+2 {
		p.Order(owner, from, to)
	}
}
func (p *Planets) Tick() {
	if p.Winner != 0 {
		return
	}
	for i := range p.Bodies {
		b := &p.Bodies[i]
		b.Phase = math.Mod(b.Phase+b.Rate/60, 2*math.Pi)
	}
	if !p.Started {
		return
	}
	p.ActiveTicks = min(PlanetsRampTicks, p.ActiveTicks+1)
	p.ProductionClock++
	if p.ProductionClock >= 180 {
		p.ProductionClock = 0
		for i := range p.Bodies {
			b := &p.Bodies[i]
			if b.Owner > 0 {
				b.Units = min(100000, b.Units+1)
			}
		}
	}
	fleets := p.Fleets[:0]
	for _, f := range p.Fleets {
		pos := p.Position(f.Target)
		dx, dy := pos.X-f.X, pos.Y-f.Y
		dist := math.Hypot(dx, dy)
		if dist <= FleetSpeed/60 {
			b := &p.Bodies[f.Target]
			if b.Owner == f.Owner {
				b.Units = min(100000, b.Units+f.Units)
			} else if f.Units > b.Units {
				b.Owner = f.Owner
				b.Units = f.Units - b.Units
			} else {
				b.Units -= f.Units
				if b.Units == 0 {
					b.Owner = 0
				}
			}
		} else {
			f.X += dx / dist * FleetSpeed / 60
			f.Y += dy / dist * FleetSpeed / 60
			fleets = append(fleets, f)
		}
	}
	p.Fleets = fleets
	p.AIClock++
	if p.AIClock >= p.AIInterval {
		p.AIClock = 0
		p.AIInterval = p.ReactionTicks()
		for owner := 2; owner <= p.Players; owner++ {
			p.ComputerOrder(owner)
		}
	}
	alive := [5]bool{}
	for _, b := range p.Bodies {
		if b.Owner > 0 {
			alive[b.Owner] = true
		}
	}
	for _, f := range p.Fleets {
		alive[f.Owner] = true
	}
	winner, count := 0, 0
	for i := 1; i <= p.Players; i++ {
		if alive[i] {
			winner = i
			count++
		}
	}
	if count == 1 {
		p.Winner = winner
	}
	if p.Source >= 0 && p.Bodies[p.Source].Owner != 1 {
		p.Source = -1
	}
}

func (p *Planets) ReactionTicks() int {
	return 360 - 240*min(PlanetsRampTicks, max(0, p.ActiveTicks))/PlanetsRampTicks
}
