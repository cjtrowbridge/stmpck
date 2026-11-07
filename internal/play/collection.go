package play

import (
	"encoding/json"
	"errors"
	"math"
)

type Collection struct {
	Version   int
	Ready     [Count]bool
	Plays     [Count]uint64
	Merge     Merge
	Snake     Snake
	Territory Territory
	Linear    Linear
	Bounce    Bounce
}

func NewCollection() Collection { return Collection{Version: 1} }
func (c *Collection) New(id, option int, seed uint64) {
	switch id {
	case MergeID:
		c.Merge = NewMerge(option+4, seed)
	case SnakeID:
		c.Snake = NewSnake(option, seed)
	case TerritoryID:
		c.Territory = NewTerritory(seed)
	case LinearID:
		start := 1
		if c.Ready[id] {
			start = 3 - c.Linear.Starting
		}
		c.Linear = NewLinear(false, start)
	case BounceID:
		c.Bounce = NewBounce(option, seed)
	default:
		return
	}
	c.Ready[id] = true
	c.Plays[id]++
}
func (c *Collection) Pause()                  { c.Snake.Running = false; c.Bounce.Running = false }
func (c *Collection) Encode() ([]byte, error) { return json.Marshal(c) }
func Decode(data []byte) (Collection, error) {
	var c Collection
	if len(data) > 1<<20 {
		return c, errors.New("save too large")
	}
	if err := json.Unmarshal(data, &c); err != nil {
		return c, err
	}
	if err := c.Validate(); err != nil {
		return Collection{}, err
	}
	c.Linear.TwoPlayers = false
	c.Pause()
	return c, nil
}
func (c *Collection) Validate() error {
	bad := func() error { return errors.New("invalid game save") }
	if c.Version != 1 {
		return bad()
	}
	if c.Ready[0] {
		m := &c.Merge
		if m.N < 4 || m.N > 6 || m.Score < 0 {
			return bad()
		}
		for _, pos := range []MergePosition{m.MergePosition, m.UndoPosition} {
			if pos.Score < 0 {
				return bad()
			}
			for i, v := range pos.Cells {
				if v < 0 || v > 1<<30 || (v != 0 && (v < 2 || v&(v-1) != 0)) || (i >= m.N*m.N && v != 0) {
					return bad()
				}
			}
		}
		m.Over = !m.CanMove()
	}
	if c.Ready[1] {
		s := &c.Snake
		if s.Speed < 0 || s.Speed > 2 || s.Direction < 0 || s.Direction > 3 || len(s.Body) < 1 || len(s.Body) > SnakeW*SnakeH || len(s.Queue) > 2 || s.Score < 0 || s.Clock < 0 || s.Clock > 12 {
			return bad()
		}
		seen := map[int]bool{}
		for _, v := range s.Body {
			if v < 0 || v >= SnakeW*SnakeH || seen[v] {
				return bad()
			}
			seen[v] = true
		}
		if !s.Won && (s.Food < 0 || s.Food >= SnakeW*SnakeH || seen[s.Food]) {
			return bad()
		}
		d := s.Direction
		for _, v := range s.Queue {
			if v < 0 || v > 3 || v == d || (v+2)%4 == d {
				return bad()
			}
			d = v
		}
	}
	if c.Ready[2] {
		t := &c.Territory
		if t.Players == 0 {
			t.Players = 2
		}
		if t.Terrain.Version == 0 {
			t.Terrain = LegacyTerrain()
		}
		if t.Terrain.Version != 1 || t.Players < 2 || t.Players > 4 {
			return bad()
		}
		if len(t.Terrain.Polygons[0]) == 0 {
			if t.Terrain.Seed == 0 {
				t.Terrain = LegacyTerrain()
			} else {
				t.Terrain = GenerateTerrain(t.Terrain.Seed)
			}
		}
		if t.Rules == 0 {
			t.Rules = 2
			t.Source = -1
			t.Target = -1
		}
		if t.Rules != 2 || t.Source < -1 || t.Source >= 12 || t.Target < -1 || t.Target >= 12 || t.Winner < 0 || t.Winner > t.Players || t.ProductionClock < 0 || t.ProductionClock >= TerritoryProductionTicks || t.AIClock < 0 || t.AIClock >= TerritoryAITicks || len(t.Movements) > 256 {
			return bad()
		}
		for i, o := range t.Owners {
			if o < 0 || o > t.Players || t.Units[i] < 0 || t.Units[i] > 100000 || o == 0 && t.Units[i] != 0 || o > 0 && t.Units[i] < 1 {
				return bad()
			}
		}
		for _, m := range t.Movements {
			if !t.Terrain.Adjacent(m.From, m.To) || m.Owner < 1 || m.Owner > t.Players || m.Units < 1 || m.Units > 100000 || m.Duration < 1 || m.Duration > 1000 || m.Elapsed < 0 || m.Elapsed >= m.Duration {
				return bad()
			}
		}

	}
	if c.Ready[3] {
		l := &c.Linear
		if l.Turn < 1 || l.Turn > 2 || l.Human != 1 || l.Starting < 1 || l.Starting > 2 {
			return bad()
		}
		for col := 0; col < 7; col++ {
			gap := false
			for row := 5; row >= 0; row-- {
				v := l.Cells[row*7+col]
				if v < 0 || v > 2 {
					return bad()
				}
				if v == 0 {
					gap = true
				} else if gap {
					return bad()
				}
			}
		}
		l.Winner = 0
		l.Over = false
		l.Winning = [4]int{-1, -1, -1, -1}
		l.check()
	}
	if c.Ready[4] {
		b := &c.Bounce
		if b.Speed < 0 || b.Speed > 2 || b.ServeDirection != 1 && b.ServeDirection != -1 {
			return bad()
		}
		for _, v := range []float64{b.X, b.Y, b.VX, b.VY, b.Player, b.Computer, b.AITarget} {
			if math.IsNaN(v) || math.IsInf(v, 0) || math.Abs(v) > 2 {
				return bad()
			}
		}
		if b.X < -.1 || b.X > 1.1 || b.Y < 0 || b.Y > 1 || b.Player < .12 || b.Player > .88 || b.Computer < .12 || b.Computer > .88 {
			return bad()
		}
		for _, p := range b.Points {
			if p < 0 || p > 5 {
				return bad()
			}
		}
	}
	return nil
}
