package play

import (
	"encoding/json"
	"math"
	"reflect"
	"testing"
)

func TestShiftSolvabilityMovesAndUndo(t *testing.T) {
	for seed := uint64(1); seed <= 100; seed++ {
		s := NewShift(seed)
		if s.Won || s.Solved() {
			t.Fatal("solved initial puzzle")
		}
		inversions := 0
		for i, a := range s.Cells {
			for _, b := range s.Cells[i+1:] {
				if a != 0 && b != 0 && a > b {
					inversions++
				}
			}
		}
		if (inversions+4-s.Empty()/4)%2 != 1 {
			t.Fatal("unsolvable permutation", seed)
		}
		before := s.Cells
		for d := 0; d < 4; d++ {
			if s.Direction(d) {
				break
			}
		}
		if s.Cells == before {
			t.Fatal("no valid move")
		}
		s.Undo()
		if s.Cells != before || s.Moves != 0 {
			t.Fatal("undo")
		}
		if s.Move(s.Empty()) || s.Move(-1) {
			t.Fatal("illegal move")
		}
	}
	s := NewShift(1)
	for i := range s.Cells {
		s.Cells[i] = (i + 1) % 16
	}
	s.Cells[14], s.Cells[15] = 0, 15
	s.Move(15)
	if !s.Won {
		t.Fatal("missing completion")
	}
	s.Replay()
	if s.Won || s.Cells != s.Initial {
		t.Fatal("replay")
	}
}
func TestDeductionUniquePuzzlesAndImmutableClues(t *testing.T) {
	for seed := uint64(1); seed <= 20; seed++ {
		d := NewDeduction(seed)
		if CountSudokuSolutions(d.Clues, 2) != 1 {
			t.Fatal("nonunique", seed)
		}
		clues := 0
		for i, v := range d.Clues {
			if v != 0 {
				clues++
				if d.Set(i, 0) {
					t.Fatal("clue erased")
				}
			} else {
				before := d.Cells
				if !d.Set(i, d.Solution[i]) {
					t.Fatal("entry refused")
				}
				d.Undo()
				if d.Cells != before {
					t.Fatal("undo")
				}
			}
		}
		if clues < 36 || clues > 45 {
			t.Fatal("unexpected clue count", clues)
		}
		for i, v := range d.Solution {
			d.Set(i, v)
		}
		if !d.Won {
			t.Fatal("solution not recognized")
		}
	}
	board := [81]int{}
	board[0] = 1
	board[1] = 1
	if CountSudokuSolutions(board, 2) != 0 {
		t.Fatal("invalid clues accepted")
	}
}
func TestStackBagRowsRotationAndTopOut(t *testing.T) {
	s := NewStack(42)
	seen := [7]bool{}
	seen[s.Piece] = true
	seen[s.Next] = true
	for i := 0; i < 5; i++ {
		v := s.take()
		if seen[v] {
			t.Fatal("bag repeated")
		}
		seen[v] = true
	}
	s = NewStack(4)
	for i := 180; i < 200; i++ {
		s.Cells[i] = 1
	}
	s.Cells[179] = 2
	s.clearRows()
	if s.Lines != 2 || s.Score != 400 || s.Cells[199] != 2 {
		t.Fatal("rows did not compact")
	}
	s = NewStack(5)
	s.Piece = 0
	s.Rotation = 1
	s.X = -2
	s.Y = 2
	if !s.Fits(s.X, s.Y, s.Rotation) || !s.Rotate() || !s.Fits(s.X, s.Y, s.Rotation) {
		t.Fatal("wall kick")
	}
	s = NewStack(6)
	for i := 0; i < 40; i++ {
		s.Cells[i] = 1
	}
	s.spawn()
	if !s.Over {
		t.Fatal("blocked spawn")
	}
	for pace, ticks := range []int{90, 60, 30} {
		s = NewStack(2)
		s.Pace = pace
		for i := 0; i < ticks-1; i++ {
			s.Tick()
		}
		if s.Y != 0 {
			t.Fatal("gravity early")
		}
		s.Tick()
		if s.Y != 1 {
			t.Fatal("gravity late")
		}
	}
}
func TestDriftWrapSplitCollisionAndWaves(t *testing.T) {
	d := NewDrift(2)
	d.X = .999
	d.VX = .3
	d.Tick()
	if d.X > .01 {
		t.Fatal("ship failed to wrap")
	}
	d.Rocks = []SpaceRock{{X: .2, Y: .2, Size: 2}}
	d.Shots = []SpaceShot{{X: .2, Y: .2, Life: 10}}
	d.Tick()
	if len(d.Rocks) != 2 || d.Rocks[0].Size != 1 || d.Hits != 1 {
		t.Fatal("splitting")
	}
	d.Shield = 0
	d.Rocks = []SpaceRock{{X: d.X, Y: d.Y, Size: 0}}
	d.Tick()
	if d.Lives != 2 || d.Shield == 0 {
		t.Fatal("collision")
	}
	d.Tick()
	if d.Lives != 2 {
		t.Fatal("respawn protection")
	}
	d.Rocks = nil
	d.Tick()
	if !d.Cleared {
		t.Fatal("wave completion")
	}
	d.NextWave()
	if d.Level != 2 || d.Cleared || len(d.Rocks) != 4 {
		t.Fatal("wave advance")
	}
}
func TestClusterMatchDetachBounceAndLoss(t *testing.T) {
	c := NewCluster(1)
	c.Cells = [80]int{}
	c.Cells[0] = 1
	c.Cells[1] = 1
	c.Cells[8] = 1
	c.Cells[16] = 2
	c.Cells[24] = 3
	c.resolve(8)
	if !c.Won || c.Score != 5 {
		t.Fatal("match did not detach hanging cluster", c.Cells, c.Score)
	}
	c = NewCluster(2)
	c.Fire()
	if c.Fire() {
		t.Fatal("second shot accepted")
	}
	c.X = .056
	c.Y = .7
	c.VX = -.5
	c.VY = -.1
	c.Tick()
	if c.VX <= 0 {
		t.Fatal("wall reflection")
	}
	c = NewCluster(3)
	c.Cells = [80]int{}
	c.Cells[72] = 1
	c.resolve(72)
	if !c.Over {
		t.Fatal("bottom row must lose")
	}
	for i := 0; i < 80; i++ {
		for _, j := range ClusterNeighbors(i) {
			found := false
			for _, k := range ClusterNeighbors(j) {
				found = found || k == i
			}
			if !found {
				t.Fatal("asymmetric hex adjacency", i, j)
			}
		}
	}
}
func TestWaveHitsLivesAndDescent(t *testing.T) {
	w := NewWave(1)
	w.Enemies = [32]bool{}
	w.Enemies[0] = true
	p := w.Position(0)
	w.Shots = []WaveShot{{X: p.X, Y: p.Y + .01}}
	w.Tick()
	if !w.Cleared || w.Hits != 1 {
		t.Fatal("last hit")
	}
	w.NextWave()
	if w.Level != 2 || w.Cleared {
		t.Fatal("next wave")
	}
	w.Shield = 0
	w.Shots = []WaveShot{{X: w.Player, Y: .89, Enemy: true}}
	w.Tick()
	if w.Lives != 2 {
		t.Fatal("player hit")
	}
	w.X = .30
	y := w.Y
	w.Tick()
	if w.Direction != -1 || w.Y <= y {
		t.Fatal("boundary descent")
	}
	w.Y = .82
	w.Tick()
	if !w.Over {
		t.Fatal("invasion should end round")
	}
}
func TestCrossingCarsWaterLogsHomesAndPauseState(t *testing.T) {
	c := NewCrossing(1)
	c.Shield = 0
	c.Y = 9
	c.X = .5
	c.Phases[9] = 0
	c.check()
	if c.Lives != 2 || c.Y != 10 {
		t.Fatal("car collision")
	}
	c.Shield = 0
	c.Y = 1
	c.X = 2.6
	c.Phases[1] = 0
	c.check()
	if c.Lives != 1 {
		t.Fatal("water collision")
	}
	c = NewCrossing(1)
	c.Shield = 0
	c.Y = 1
	c.X = 1
	c.Phases[1] = 0
	x := c.X
	c.Tick()
	if c.X == x || c.Lives != 3 {
		t.Fatal("log carry")
	}
	for i := 0; i < 3; i++ {
		c.X = 1.5 + float64(i)*3
		c.Y = 1
		c.Move(1)
	}
	if !c.Won {
		t.Fatal("home completion")
	}
	c.NextLevel()
	if c.Won || c.Level != 2 || c.Homes != [3]bool{} {
		t.Fatal("next level")
	}
}
func TestExpansionSavesLongRunsAndOldCatalog(t *testing.T) {
	c := NewDemonstrations()
	for tick := 1; tick <= 3000; tick++ {
		c.demonstrateExpansion(tick)
		if tick%60 == 0 {
			if err := c.Validate(); err != nil {
				t.Fatalf("tick %d: %v", tick, err)
			}
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
	if !reflect.DeepEqual(c, loaded) {
		t.Fatal("save changed state")
	}
	for tick := 3001; tick <= 3100; tick++ {
		c.demonstrateExpansion(tick)
		loaded.demonstrateExpansion(tick)
	}
	if !reflect.DeepEqual(c, loaded) {
		t.Fatal("continuation changed")
	}
	legacy := struct {
		Version int
		Ready   [11]bool
		Plays   [11]uint64
		Flow    Flow
	}{Version: 1, Flow: NewFlow(123)}
	legacy.Ready[FlowID] = true
	legacy.Plays[FlowID] = 7
	data, _ = json.Marshal(legacy)
	old, err := Decode(data)
	if err != nil || old.Flow.Cells != legacy.Flow.Cells || old.Plays[FlowID] != 7 || old.Ready[DriftID] {
		t.Fatal("eleven-game compatibility", err)
	}
}
func TestExpansionRejectsCorruptSaves(t *testing.T) {
	for _, corrupt := range []func(*Collection){func(c *Collection) { c.Drift.Rocks[0].Size = 9 }, func(c *Collection) { c.Stack.Piece = 99 }, func(c *Collection) { c.Deduction.Solution[0] = 0 }, func(c *Collection) { c.Cluster.Color = 7 }, func(c *Collection) { c.Wave.Shots = []WaveShot{{X: math.NaN()}} }, func(c *Collection) { c.Crossing.Y = 12 }, func(c *Collection) { c.Shift.Cells[0] = c.Shift.Cells[1] }} {
		c := NewDemonstrations()
		corrupt(&c)
		if c.Validate() == nil {
			t.Fatal("corrupt save accepted")
		}
	}
}
