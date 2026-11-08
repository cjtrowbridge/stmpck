package play

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestPlanetsWaitsForFirstValidHumanOrder(t *testing.T) {
	p := NewPlanets(1)
	owners := append([]OrbitalBody(nil), p.Bodies...)
	phase := p.Position(0)
	p.Choose(0)
	if p.Order(1, 0, 0) || p.Order(1, -1, 2) || p.Order(2, 2, 1) {
		t.Fatal("invalid or premature order accepted")
	}
	for tick := 0; tick < 1000; tick++ {
		p.ComputerOrder(2)
		p.Tick()
	}
	if p.Started || p.ActiveTicks != 0 || p.AIClock != 0 || p.ProductionClock != 0 || len(p.Fleets) != 0 {
		t.Fatal("waiting match progressed", p)
	}
	for i, b := range p.Bodies {
		if b.Owner != owners[i].Owner || b.Units != owners[i].Units {
			t.Fatal("units changed before first order")
		}
	}
	if p.Position(0) == phase {
		t.Fatal("waiting orbits should remain alive")
	}
	if !p.Order(1, 0, 1) || !p.Started || p.Bodies[0].Units != 1 {
		t.Fatal("first dispatch did not start match")
	}
	for tick := 0; tick < 359; tick++ {
		p.Tick()
		for _, f := range p.Fleets {
			if f.Owner != 1 {
				t.Fatal("opponent moved before initial six-second delay")
			}
		}
	}
	p.Tick()
	enemy := false
	for _, f := range p.Fleets {
		enemy = enemy || f.Owner > 1
	}
	if !enemy {
		t.Fatal("opponent did not act when initial delay elapsed")
	}
}

func TestPlanetsReactionRampAndSaveContinuity(t *testing.T) {
	p := NewPlanets(2)
	for _, tt := range []struct{ age, want int }{{0, 360}, {5400, 240}, {10800, 120}, {20000, 120}} {
		p.ActiveTicks = tt.age
		if p.ReactionTicks() != tt.want {
			t.Fatal("reaction curve", tt, p.ReactionTicks())
		}
	}
	c := NewCollection()
	c.New(PlanetsID, 1, 1)
	c.Planets.Order(1, 0, 1)
	for tick := 0; tick < 240; tick++ {
		c.Planets.Tick()
	}
	data, _ := c.Encode()
	loaded, err := Decode(data)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(c.Planets, loaded.Planets) {
		t.Fatal("in-progress timing not preserved")
	}
	for tick := 0; tick < 900; tick++ {
		c.Planets.Tick()
		loaded.Planets.Tick()
	}
	if !reflect.DeepEqual(c.Planets, loaded.Planets) {
		t.Fatal("resumed simulation diverged")
	}
	c.New(PlanetsID, 0, 1)
	data, _ = c.Encode()
	waiting, err := Decode(data)
	if err != nil || waiting.Planets.Started || waiting.Planets.AIInterval != 360 {
		t.Fatal("waiting state did not survive reload", err)
	}
}

func TestPlanetsLegacyTimingMigration(t *testing.T) {
	c := NewCollection()
	c.New(PlanetsID, 0, 1)
	c.Planets.Order(1, 0, 4)
	c.Planets.AIClock = 59
	data, _ := c.Encode()
	var old map[string]json.RawMessage
	json.Unmarshal(data, &old)
	var planets map[string]json.RawMessage
	json.Unmarshal(old["Planets"], &planets)
	for _, field := range []string{"TimingVersion", "Started", "ActiveTicks", "AIInterval"} {
		delete(planets, field)
	}
	old["Planets"], _ = json.Marshal(planets)
	data, _ = json.Marshal(old)
	loaded, err := Decode(data)
	if err != nil {
		t.Fatal(err)
	}
	p := loaded.Planets
	if !p.Started || p.ActiveTicks != PlanetsRampTicks || p.AIInterval != 120 || p.AIClock != 59 || !reflect.DeepEqual(p.Fleets, c.Planets.Fleets) {
		t.Fatal("legacy match altered", p)
	}
	for _, corrupt := range []func(*Planets){func(p *Planets) { p.AIInterval = 0 }, func(p *Planets) { p.ActiveTicks = -1 }, func(p *Planets) { p.Started = false }, func(p *Planets) { p.AIClock = p.AIInterval }} {
		bad := loaded
		corrupt(&bad.Planets)
		if bad.Validate() == nil {
			t.Fatal("invalid timing accepted")
		}
	}
}

func TestLifeCatalogPeriodsAndGliderTravel(t *testing.T) {
	for i, period := range []int{0, 2, 2, 2, 1, 3} {
		b := BloomDemonstration(i)
		before := b.Cells
		if period == 0 {
			for step := 0; step < 4; step++ {
				b.Step()
			}
			translated := [BloomN * BloomN]bool{}
			for j, live := range before {
				if live {
					translated[((j/BloomN+1)%BloomN)*BloomN+(j%BloomN+1)%BloomN] = true
				}
			}
			if b.Cells != translated {
				t.Fatal("glider did not translate after four generations")
			}
			continue
		}
		for step := 0; step < period; step++ {
			b.Step()
			if step+1 < period && b.Cells == before {
				t.Fatal("unexpected shorter period", BloomPatterns[i].Name)
			}
		}
		if b.Cells != before {
			t.Fatal("wrong oscillator period", BloomPatterns[i].Name)
		}
	}
}

func TestBloomStampWrapOverlayAndPersistence(t *testing.T) {
	for pattern := range BloomPatterns {
		b := Bloom{Speed: 1}
		b.Cells[100] = true
		b.Generation = 20
		b.Clock = 7
		expected := b.Cells
		for _, i := range BloomPatternCells(pattern, 0) {
			expected[i] = true
		}
		if !b.Stamp(pattern, 0) || b.Cells != expected || !b.Cells[100] || b.Generation != 20 || b.Clock != 0 {
			t.Fatal("stamp changed unrelated cells or history", pattern)
		}
		before := b.Cells
		b.Stamp(pattern, 0)
		if b.Cells != before {
			t.Fatal("overlay toggled cells")
		}
		if b.Stamp(-1, 0) || b.Stamp(pattern, -1) || b.Stamp(pattern, 576) {
			t.Fatal("invalid placement accepted")
		}
		c := NewCollection()
		c.Ready[BloomID] = true
		c.Bloom = b
		data, _ := c.Encode()
		saved, err := Decode(data)
		if err != nil || saved.Bloom != b {
			t.Fatal("stamp not saved", err)
		}
	}
}

func TestBloomPreviewUsesCatalog(t *testing.T) {
	c := NewDemonstrations()
	if c.Bloom.Cells != BloomDemonstration(0).Cells {
		t.Fatal("first preview not glider")
	}
	for i := 1; i < len(BloomPatterns); i++ {
		c.demonstrateAdditions(i * 120)
		if c.Bloom.Cells != BloomDemonstration(i).Cells {
			t.Fatal("missing preview pattern", i)
		}
	}
}
