package play

import (
	"encoding/json"
	"math"
	"reflect"
	"testing"
)

func TestFlowGeneratedPuzzlesAreSolvable(t *testing.T) {
	for seed := uint64(1); seed <= 200; seed++ {
		f := NewFlow(seed)
		if f.Complete() {
			t.Fatal("puzzle starts solved", seed)
		}
		initial := f.Cells
		for i := range f.Cells {
			for rotations := 0; f.Cells[i] != f.Solution[i] && rotations < 4; rotations++ {
				f.Rotate(i)
			}
		}
		if !f.Complete() || !f.Won {
			t.Fatal("not solvable", seed)
		}
		f.Replay()
		if f.Cells != initial || f.Won || f.Moves != 0 {
			t.Fatal("replay changed puzzle")
		}
	}
}
func TestBloomBlinkerAndWrappedBirth(t *testing.T) {
	b := Bloom{}
	for _, i := range []int{12*24 + 11, 12*24 + 12, 12*24 + 13} {
		b.Cells[i] = true
	}
	before := b.Cells
	b.Step()
	if !b.Cells[11*24+12] || !b.Cells[12*24+12] || !b.Cells[13*24+12] {
		t.Fatal("blinker did not rotate")
	}
	b.Step()
	if b.Cells != before {
		t.Fatal("blinker period")
	}
	b.Clear()
	b.Cells[23*24+23] = true
	b.Cells[23*24] = true
	b.Cells[23] = true
	b.Step()
	if !b.Cells[0] {
		t.Fatal("edges must wrap")
	}
	for speed, ticks := range []int{30, 10, 5} {
		b = NewBloom()
		b.Speed = speed
		for i := 0; i < ticks-1; i++ {
			b.Tick()
		}
		if b.Generation != 0 {
			t.Fatal("too fast")
		}
		b.Tick()
		if b.Generation != 1 {
			t.Fatal("wrong cadence")
		}
	}
}
func TestEchoSilentSequenceAndRetry(t *testing.T) {
	e := NewEcho(42)
	first := e.Sequence[0]
	if e.Press(first) {
		t.Fatal("accepted during demonstration")
	}
	for i := 0; i < 45; i++ {
		e.Tick()
	}
	if e.Phase != EchoInput {
		t.Fatal("never invited input")
	}
	e.Press((first + 1) % 4)
	if e.Phase != EchoFailed {
		t.Fatal("wrong cue accepted")
	}
	e.Retry()
	if e.Sequence[0] != first || e.Input != 0 {
		t.Fatal("retry changed sequence")
	}
	for round := 1; round <= 32; round++ {
		for e.Phase == EchoShow {
			e.Tick()
		}
		for _, pad := range append([]int(nil), e.Sequence...) {
			if !e.Press(pad) {
				t.Fatal("correct input rejected")
			}
		}
		for i := 0; i < 60; i++ {
			e.Tick()
		}
		if round < 32 && len(e.Sequence) != round+1 {
			t.Fatal("sequence did not grow")
		}
	}
	if e.Phase != EchoWon {
		t.Fatal("missing completion")
	}
}
func TestCalibanShotsHitsAndLoss(t *testing.T) {
	c := NewCaliban(1)
	c.Enemies = []TubeObject{{Lane: 0, Depth: .5}}
	c.Shots = []TubeObject{{Lane: 0, Depth: .51}}
	c.Tick()
	if len(c.Enemies) != 0 || c.Hits != 1 {
		t.Fatal("shot missed overlapping enemy")
	}
	c.Enemies = []TubeObject{{Lane: 1, Depth: .999}, {Lane: 2, Depth: .999}, {Lane: 3, Depth: .999}}
	c.Tick()
	if !c.Over || c.Lives != 0 {
		t.Fatal("rim crossings must cost lives")
	}
	c = NewCaliban(1)
	c.Spawned = 7
	c.Tick()
	if !c.Cleared {
		t.Fatal("wave not complete")
	}
	c.Next()
	if c.Wave != 2 || c.Cleared || c.Spawned != 0 {
		t.Fatal("next wave")
	}
}
func TestBrickCollisionsServeAndCompletion(t *testing.T) {
	b := NewBrick(0)
	b.Serve()
	b.X = .5
	b.Y = .875
	b.VX = 0
	b.VY = .6
	b.Tick()
	if b.VY >= 0 {
		t.Fatal("paddle missed")
	}
	b.X = .01
	b.Y = .6
	b.VX = -.6
	b.VY = 0
	b.Tick()
	if b.VX <= 0 {
		t.Fatal("wall missed")
	}
	b.Bricks = [40]bool{}
	b.Bricks[0] = true
	x, y, w, h := BrickRect(0)
	b.X = x + w/2
	b.Y = y + h + .013
	b.VY = -.6
	b.VX = 0
	b.Tick()
	if !b.Won || b.Bricks[0] || b.Running {
		t.Fatal("last brick did not end board")
	}
	b = NewBrick(0)
	for loss := 0; loss < 3; loss++ {
		b.Serve()
		b.Y = 1.02
		b.VY = .6
		b.Tick()
		if b.Running {
			t.Fatal("automatic serve after loss")
		}
	}
	if !b.Over || b.Lives != 0 {
		t.Fatal("three misses must end game")
	}
	b.Serve()
	if b.Running {
		t.Fatal("served after game over")
	}
}
func TestPlanetsOrdersOrbitAndArrival(t *testing.T) {
	for system := 0; system < 3; system++ {
		p := NewPlanets(system)
		counts := [5]int{}
		for _, b := range p.Bodies {
			counts[b.Owner]++
		}
		for owner := 1; owner <= p.Players; owner++ {
			if counts[owner] != 1 {
				t.Fatal("multiple starting territories")
			}
		}
		from := 0
		to := len(p.Bodies) - 1
		start := p.Position(to)
		if !p.Order(1, from, to) || p.Bodies[from].Units != 1 || p.Fleets[0].Units != 9 {
			t.Fatal("order not deducted immediately")
		}
		original := p.Fleets[0]
		p.Tick()
		if p.Fleets[0].Target != to {
			t.Fatal("destination changed")
		}
		distance := math.Hypot(p.Fleets[0].X-original.X, p.Fleets[0].Y-original.Y)
		if math.Abs(distance-FleetSpeed/60) > 1e-10 {
			t.Fatal("wrong travel speed")
		}
		for i := 0; i < 600 && len(p.Fleets) > 0; i++ {
			p.AIClock = 0
			p.Tick()
		}
		if len(p.Fleets) != 0 || p.Bodies[to].Owner != 1 {
			t.Fatal("moving destination never reached", system, p.Fleets)
		}
		if p.Position(to) == start {
			t.Fatal("orbit did not move")
		}
	}
}
func TestPlanetsFriendlyEnemyAndInFlightVictory(t *testing.T) {
	p := NewPlanets(0)
	for i := range p.Bodies {
		p.Bodies[i].Owner = 1
		p.Bodies[i].Units = 5
	}
	target := p.Position(1)
	p.Fleets = []Fleet{{Target: 1, Owner: 2, Units: 8, X: target.X + .1, Y: target.Y}}
	p.Tick()
	if p.Winner != 0 {
		t.Fatal("ignored hostile fleet")
	}
	for i := 0; i < 60; i++ {
		p.Tick()
	}
	if p.Bodies[1].Owner != 2 || p.Bodies[1].Units != 3 {
		t.Fatal("arrival combat", p.Bodies[1])
	}
	p.Fleets = []Fleet{{Target: 1, Owner: 2, Units: 4, X: p.Position(1).X, Y: p.Position(1).Y}}
	p.Tick()
	if p.Bodies[1].Units != 7 {
		t.Fatal("friendly transfer")
	}
}
func TestAddedGamesSaveResumeAndLegacyCompatibility(t *testing.T) {
	c := NewCollection()
	for id := 0; id < Count; id++ {
		c.New(id, 0, uint64(id+1))
	}
	c.Brick.Serve()
	c.Planets.Order(1, 0, 4)
	for i := 1; i <= 400; i++ {
		c.demonstrateAdditions(i)
		if err := c.Validate(); err != nil {
			t.Fatal("invalid live game", i, err)
		}
	}
	c.Pause()
	data, err := c.Encode()
	if err != nil {
		t.Fatal(err)
	}
	loaded, err := Decode(data)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(c.Flow, loaded.Flow) || !reflect.DeepEqual(c.Caliban, loaded.Caliban) || c.Bloom != loaded.Bloom || !reflect.DeepEqual(c.Echo, loaded.Echo) || c.Brick != loaded.Brick || !reflect.DeepEqual(c.Planets, loaded.Planets) {
		t.Fatal("new game state changed on load")
	}
	for i := 0; i < 120; i++ {
		c.Planets.Tick()
		loaded.Planets.Tick()
	}
	if !reflect.DeepEqual(c.Planets, loaded.Planets) {
		t.Fatal("save changed future simulation")
	}
	legacy := struct {
		Version int
		Ready   [5]bool
		Plays   [5]uint64
		Merge   Merge
	}{Version: 1, Ready: [5]bool{true}, Plays: [5]uint64{7}, Merge: NewMerge(4, 23)}
	data, _ = json.Marshal(legacy)
	loaded, err = Decode(data)
	if err != nil || loaded.Merge.Cells != legacy.Merge.Cells || loaded.Plays[0] != 7 || loaded.Ready[FlowID] {
		t.Fatal("old five-game save incompatible", err)
	}
}
func TestNewGamesLongRunAndInvalidSaves(t *testing.T) {
	c := NewDemonstrations()
	for i := 1; i <= 10000; i++ {
		c.demonstrateAdditions(i)
		if err := c.Validate(); err != nil {
			t.Fatalf("tick %d: %v", i, err)
		}
	}
	for _, corrupt := range []func(*Collection){func(c *Collection) { c.Flow.Cells[0] = 16 }, func(c *Collection) { c.Bloom.Speed = 9 }, func(c *Collection) { c.Echo.Phase = EchoInput; c.Echo.Input = len(c.Echo.Sequence) }, func(c *Collection) { c.Caliban.Shots = []TubeObject{{Lane: 99, Depth: .5}} }, func(c *Collection) { c.Brick.Paddle = math.NaN() }, func(c *Collection) { c.Planets.Bodies[0].Parent = 0 }, func(c *Collection) { c.Planets.Fleets = []Fleet{{Target: 99, Owner: 1, Units: 2}} }} {
		fresh := NewDemonstrations()
		corrupt(&fresh)
		if fresh.Validate() == nil {
			t.Fatal("corrupt save accepted")
		}
	}
}
