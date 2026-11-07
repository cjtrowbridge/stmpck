package play

import "math"

const TerritoryProductionTicks = 180
const TerritoryAITicks = 120
const TerritorySpeed = .12 // normalized map distance per second, at 60 simulation ticks/sec

type Movement struct {
	From, To, Owner, Units int
	Elapsed, Duration      int
}
type Territory struct {
	Terrain                  Terrain
	Players                  int
	Owners                   [12]int
	Units                    [12]int
	Rules                    int
	Source, Target, Winner   int
	Movements                []Movement
	ProductionClock, AIClock int
	Random                   RNG
	Dice                     [2]int
	LastAttack               [2]int
	Battles                  int
}

func NewTerritory(seed uint64) Territory { return NewTerritoryMatch(seed, seed^0x9e3779b97f4a7c15, 2) }
func NewTerritoryMatch(mapSeed, startSeed uint64, players int) Territory {
	t := Territory{Terrain: GenerateTerrain(mapSeed), Players: min(4, max(2, players)), Rules: 2, Source: -1, Target: -1, Random: RNG(startSeed), LastAttack: [2]int{-1, -1}}
	order := [12]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	for i := 11; i > 0; i-- {
		j := t.Random.Next(i + 1)
		order[i], order[j] = order[j], order[i]
	}
	for player := 1; player <= t.Players; player++ {
		id := order[player-1]
		t.Owners[id] = player
		t.Units[id] = 10
	}
	return t
}
func (t *Territory) Alive(player int) bool {
	for _, o := range t.Owners {
		if o == player {
			return true
		}
	}
	for _, m := range t.Movements {
		if m.Owner == player {
			return true
		}
	}
	return false
}
func (t *Territory) CanOrder(player, a, b int) bool {
	return player >= 1 && player <= t.Players && t.Winner == 0 && t.Terrain.Adjacent(a, b) && t.Owners[a] == player && t.Units[a] > 1 && len(t.Movements) < 256
}
func (t *Territory) Order(player, a, b int) bool {
	if !t.CanOrder(player, a, b) {
		return false
	}
	from, to := t.Terrain.Centers[a], t.Terrain.Centers[b]
	duration := max(1, int(math.Ceil(math.Hypot(to.X-from.X, to.Y-from.Y)/TerritorySpeed*60)))
	t.Movements = append(t.Movements, Movement{From: a, To: b, Owner: player, Units: t.Units[a] - 1, Duration: duration})
	t.Units[a] = 1
	return true
}
func (t *Territory) Choose(id int) bool {
	if id < 0 || id >= 12 || t.Winner != 0 {
		return false
	}
	if t.Order(1, t.Source, id) {
		t.Source = -1
		t.Target = -1
		return true
	}
	if t.Owners[id] == 1 && t.Units[id] > 1 {
		t.Source = id
		t.Target = -1
		return true
	}
	return false
}
func (t *Territory) arrive(m Movement) {
	if t.Owners[m.To] == m.Owner {
		t.Units[m.To] = min(100000, t.Units[m.To]+m.Units)
		return
	}
	force := m.Units
	for force > 0 && t.Units[m.To] > 0 {
		t.Dice = [2]int{t.Random.Next(6) + 1, t.Random.Next(6) + 1}
		if t.Dice[0] > t.Dice[1] {
			t.Units[m.To]--
		} else {
			force--
		}
	}
	if force > 0 {
		t.Owners[m.To] = m.Owner
		t.Units[m.To] = force
	}
	t.LastAttack = [2]int{m.From, m.To}
	t.Battles++
}

// ComputerOrder is also used for the human side in isolated preview simulations.
func (t *Territory) ComputerOrder(player int) {
	a, b, best := -1, -1, -1000000
	for i, owner := range t.Owners {
		if owner != player {
			continue
		}
		for j, o := range t.Owners {
			if !t.CanOrder(player, i, j) {
				continue
			}
			score := t.Units[i] - 1 - t.Units[j]
			if o == player {
				// Feed weaker friendly frontier regions instead of bouncing forces endlessly.
				frontier := false
				for k, enemy := range t.Owners {
					if enemy != player && t.Terrain.Adjacent(j, k) {
						frontier = true
					}
				}
				if !frontier || t.Units[i] <= t.Units[j]+2 {
					continue
				}
				score -= 10
			} else if score <= 1 {
				continue
			} else {
				score += 20
			}
			if score > best {
				a, b, best = i, j, score
			}
		}
	}
	if a >= 0 {
		t.Order(player, a, b)
	}
}
func (t *Territory) Tick() bool {
	if t.Winner != 0 {
		return false
	}
	t.ProductionClock++
	if t.ProductionClock >= TerritoryProductionTicks {
		t.ProductionClock = 0
		for i, o := range t.Owners {
			if o > 0 && t.Units[i] < 100000 {
				t.Units[i]++
			}
		}
	}
	pending := t.Movements[:0]
	for _, m := range t.Movements {
		m.Elapsed++
		if m.Elapsed >= m.Duration {
			t.arrive(m)
		} else {
			pending = append(pending, m)
		}
	}
	t.Movements = pending
	t.AIClock++
	if t.AIClock >= TerritoryAITicks {
		t.AIClock = 0
		for player := 2; player <= t.Players; player++ {
			t.ComputerOrder(player)
		}
	}
	winner := t.Owners[0]
	if winner > 0 {
		for _, o := range t.Owners {
			if o != winner {
				winner = 0
				break
			}
		}
		if winner > 0 {
			for _, m := range t.Movements {
				if m.Owner != winner {
					winner = 0
					break
				}
			}
		}
	}
	t.Winner = winner
	return true
}
