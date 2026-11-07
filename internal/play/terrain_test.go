package play

import (
	"encoding/json"
	"math"
	"reflect"
	"testing"
)

func TestOrganicTerrainGeometry(t *testing.T) {
	for seed := uint64(1); seed <= 100; seed++ {
		terrain := GenerateTerrain(seed)
		seen := [12]bool{true}
		queue := []int{0}
		minArea, maxArea := 1.0, 0.0
		for i, poly := range terrain.Polygons {
			if len(poly) < 3 {
				t.Fatalf("empty region seed %d id %d", seed, i)
			}
			center := terrain.Centers[i]
			if terrain.At(center.X, center.Y) != i {
				t.Fatalf("center not inside seed %d region %d", seed, i)
			}
			area := 0.0
			for n, a := range poly {
				b := poly[(n+1)%len(poly)]
				area += a.X*b.Y - b.X*a.Y
				if a.X < 0 || a.Y < 0 || a.X > 1 || a.Y > 1 {
					t.Fatal("out of bounds")
				}
			}
			area = math.Abs(area) / 2
			minArea = math.Min(minArea, area)
			maxArea = math.Max(maxArea, area)
			for j := 0; j < 12; j++ {
				if terrain.Adjacent(i, j) != terrain.Adjacent(j, i) || i == j && terrain.Adjacent(i, j) {
					t.Fatal("invalid adjacency")
				}
			}
		}
		for len(queue) > 0 {
			a := queue[0]
			queue = queue[1:]
			for b := 0; b < 12; b++ {
				if terrain.Adjacent(a, b) && !seen[b] {
					seen[b] = true
					queue = append(queue, b)
				}
			}
		}
		for _, v := range seen {
			if !v {
				t.Fatalf("disconnected seed %d", seed)
			}
		}
		if maxArea/minArea < 1.5 {
			t.Fatalf("regions too uniform: %v", seed)
		}
	}
}
func TestTerritoryPlayersAndIndependentScrambling(t *testing.T) {
	for players := 2; players <= 4; players++ {
		tr := NewTerritoryMatch(42, 77, players)
		var counts, units [5]int
		for i, owner := range tr.Owners {
			counts[owner]++
			units[owner] += tr.Units[i]
		}
		for player := 1; player <= players; player++ {
			if counts[player] != 1 || units[player] != 10 {
				t.Fatal("unbalanced starts")
			}
		}
		geometry := NewTerritoryMatch(43, 77, players)
		starts := NewTerritoryMatch(42, 78, players)
		if tr.Owners != geometry.Owners || reflect.DeepEqual(tr.Terrain, geometry.Terrain) {
			t.Fatal("map scramble changed starts or not map")
		}
		if tr.Owners == starts.Owners || !reflect.DeepEqual(tr.Terrain, starts.Terrain) {
			t.Fatal("start scramble changed map or not starts")
		}

	}
}
func TestTerritorySaveCompatibility(t *testing.T) {
	c := NewCollection()
	c.New(TerritoryID, 0, 9)
	c.Territory = NewTerritoryMatch(45, 66, 4)
	data, _ := c.Encode()
	got, err := Decode(data)
	if err != nil || !reflect.DeepEqual(c.Territory, got.Territory) {
		t.Fatal("generated map round trip", err)
	}
	c.Territory.Players = 0
	c.Territory.Rules = 0
	c.Territory.Terrain = Terrain{}
	for i := range c.Territory.Owners {
		c.Territory.Owners[i] = i%2 + 1
		c.Territory.Units[i] = 3
	}
	data, _ = c.Encode()
	got, err = Decode(data)
	if err != nil || got.Territory.Players != 2 || !reflect.DeepEqual(got.Territory.Terrain, LegacyTerrain()) {
		t.Fatal("legacy map changed", err)
	}
	c.New(LinearID, 0, 1)
	c.Linear.TwoPlayers = true
	c.Linear.Drop(3)
	before := c.Linear.Cells
	data, _ = json.Marshal(c)
	got, err = Decode(data)
	if err != nil || got.Linear.TwoPlayers || got.Linear.Cells != before {
		t.Fatal("old PvP board lost", err)
	}
}
