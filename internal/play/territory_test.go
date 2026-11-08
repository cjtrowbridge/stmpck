package play

import (
	"math"
	"reflect"
	"testing"
)

func territoryFixture() Territory {
	t := NewTerritoryMatch(42, 77, 2)
	t.Terrain = LegacyTerrain()
	t.Owners = [12]int{1, 0, 2}
	t.Units = [12]int{10, 0, 10}
	return t
}
func TestTerritoryImmediateLockedMovement(t *testing.T) {
	tr := territoryFixture()
	if tr.Order(2, 0, 1) || tr.Order(1, 0, 11) || tr.Order(1, -1, 0) {
		t.Fatal("illegal order")
	}
	if !tr.Choose(0) || !tr.Choose(1) || tr.Units[0] != 1 || tr.Owners[1] != 0 || len(tr.Movements) != 1 {
		t.Fatal("dispatch must debit immediately without instant capture")
	}
	m := tr.Movements[0]
	if m.Units != 9 || m.Owner != 1 {
		t.Fatal("wrong commitment")
	}
	if tr.Order(1, 0, 1) {
		t.Fatal("spent units twice")
	}
	tr.Choose(2)
	if tr.Movements[0] != m {
		t.Fatal("selection changed locked order")
	}
	a, b := tr.Terrain.Centers[0], tr.Terrain.Centers[1]
	seconds := math.Hypot(a.X-b.X, a.Y-b.Y) / TerritorySpeed
	if math.Abs(float64(m.Duration)/60-seconds) > 1.0/60 {
		t.Fatal("wrong speed")
	}
	// The source may fall during travel; the dispatched army keeps its original owner.
	tr.Owners[0] = 2
	for n := 0; n < m.Duration-1; n++ {
		tr.Tick()
	}
	if tr.Owners[1] != 0 {
		t.Fatal("early arrival")
	}
	tr.Tick()
	if tr.Owners[1] != 1 || tr.Units[1] != 9 {
		t.Fatal("arrival did not occupy neutral destination")
	}
}
func TestTerritoryFriendlyArrivalProductionAndCombat(t *testing.T) {
	tr := territoryFixture()
	tr.Owners[1] = 1
	tr.Units[1] = 3
	tr.Order(1, 0, 1)
	duration := tr.Movements[0].Duration
	for n := 0; n < duration; n++ {
		tr.Tick()
	}
	if tr.Units[1] != 12 {
		t.Fatal("friendly transfer")
	}
	tr = territoryFixture()
	for n := 0; n < TerritoryProductionTicks-1; n++ {
		tr.Tick()
	}
	if tr.Units[0] != 10 || tr.Units[11] != 0 {
		t.Fatal("early or neutral production")
	}
	tr.Tick()
	if tr.Units[0] != 11 || tr.Units[11] != 0 {
		t.Fatal("production")
	}
	tr = territoryFixture()
	tr.Owners[1] = 2
	tr.Units[1] = 3
	tr.Order(1, 0, 1)
	if tr.Battles != 0 || tr.Units[1] != 3 {
		t.Fatal("combat before arrival")
	}
	for len(tr.Movements) > 0 {
		tr.Tick()
	}
	if tr.Battles < 1 || tr.Units[1] < 1 {
		t.Fatal("arrival combat")
	}
}
func TestTerritorySaveMidFlight(t *testing.T) {
	c := NewCollection()
	c.Ready[TerritoryID] = true
	c.Territory = territoryFixture()
	c.Territory.Order(1, 0, 1)
	for n := 0; n < 25; n++ {
		c.Territory.Tick()
	}
	data, err := c.Encode()
	if err != nil {
		t.Fatal(err)
	}
	loaded, err := Decode(data)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(c.Territory, loaded.Territory) {
		t.Fatal("midflight state changed on load")
	}
	for n := 0; n < 600; n++ {
		c.Territory.Tick()
		loaded.Territory.Tick()
	}
	if !reflect.DeepEqual(c.Territory, loaded.Territory) {
		t.Fatal("resumed simulation diverged")
	}
	loaded.Territory.Movements = []Movement{{From: 0, To: 11, Owner: 1, Units: 9, Duration: 10}}
	if loaded.Validate() == nil {
		t.Fatal("accepted nonadjacent order")
	}
}
func TestTerritoryInFlightArmyPreventsVictoryAndElimination(t *testing.T) {
	tr := territoryFixture()
	for i := range tr.Owners {
		tr.Owners[i] = 1
		tr.Units[i] = 10
	}
	tr.Movements = []Movement{{From: 0, To: 1, Owner: 2, Units: 2, Duration: 100}}
	if !tr.Alive(2) {
		t.Fatal("army in flight counted as eliminated")
	}
	tr.Tick()
	if tr.Winner != 0 {
		t.Fatal("premature victory")
	}
	tr.Movements = nil
	tr.Tick()
	if tr.Winner != 1 {
		t.Fatal("missing victory")
	}
}
func TestTerritoryBordersCoverEveryAdjacency(t *testing.T) {
	for seed := uint64(0); seed < 100; seed++ {
		tr := GenerateTerrain(seed)
		if seed == 0 {
			tr = LegacyTerrain()
		}
		var seen [12][12]int
		for _, b := range tr.Borders {
			if len(b.Points) < 2 {
				t.Fatal("empty border")
			}
			if b.B >= 0 {
				seen[b.A][b.B]++
				seen[b.B][b.A]++
			}
		}
		for a := 0; a < 12; a++ {
			for b := 0; b < 12; b++ {
				if tr.Adjacent(a, b) != (seen[a][b] == 1) {
					t.Fatalf("missing/duplicate border seed %d %d %d", seed, a, b)
				}
			}
		}
	}
}

func TestTerritorySwipeOrder(t *testing.T) {
	for _, owner := range []int{0, 1, 2} {
		tr := territoryFixture()
		tr.Owners[1] = owner
		if owner > 0 {
			tr.Units[1] = 3
		}
		a, b := tr.Terrain.Centers[0], tr.Terrain.Centers[1]
		tr.Source = 2
		if !tr.OrderBetween(1, a, b) || tr.Units[0] != 1 || len(tr.Movements) != 1 || tr.Source != -1 {
			t.Fatal("swipe must issue one order and immediately debit source")
		}
		movement := tr.Movements[0]
		if movement.From != 0 || movement.To != 1 || movement.Units != 9 {
			t.Fatal("wrong gesture endpoints")
		}
		if tr.OrderBetween(1, a, b) || tr.Movements[0] != movement {
			t.Fatal("repeat gesture spent unavailable units or changed order")
		}
	}
	for _, pair := range [][2]int{{0, 0}, {0, 11}, {2, 1}} {
		tr := territoryFixture()
		before := tr
		if tr.OrderBetween(1, tr.Terrain.Centers[pair[0]], tr.Terrain.Centers[pair[1]]) || !reflect.DeepEqual(before, tr) {
			t.Fatal("invalid gesture changed state", pair)
		}
	}
	tr := territoryFixture()
	a := tr.Terrain.Centers[0]
	outside := Point{-1, -1}
	if tr.OrderBetween(1, a, outside) || tr.OrderBetween(1, outside, a) {
		t.Fatal("off-map endpoint accepted")
	}
}
