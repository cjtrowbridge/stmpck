package play

import (
	"encoding/json"
	"math"
	"reflect"
	"testing"
)

func TestMergeRulesAndUndo(t *testing.T) {
	m := NewMerge(4, 17)
	m.Cells = [36]int{2, 2, 2, 2}
	m.Score = 0
	before := m.MergePosition
	if !m.Move(0) || m.Cells[0] != 4 || m.Cells[1] != 4 || m.Score != 8 {
		t.Fatal(m)
	}
	after := m.MergePosition
	if !m.Undo() || m.MergePosition != before {
		t.Fatal("undo lost state")
	}
	m.Move(0)
	if m.MergePosition != after {
		t.Fatal("undo changed random spawn")
	}
	m.Cells = [36]int{2}
	m.Score = 0
	rng := m.Random
	if m.Move(0) || m.Random != rng {
		t.Fatal("invalid move spawned tile")
	}
	m.Cells = [36]int{1024, 1024}
	m.Move(0)
	if !m.Won || !m.CanMove() {
		t.Fatal("goal must permit continued play")
	}
	for _, n := range []int{4, 5, 6} {
		m = NewMerge(n, 1)
		nonzero := 0
		for _, v := range m.Cells {
			if v != 0 {
				nonzero++
			}
		}
		if nonzero != 2 {
			t.Fatal("starting tiles")
		}
		for i := 0; i < 1000; i++ {
			m.Move(i % 4)
			for _, v := range m.Cells {
				if v != 0 && v&(v-1) != 0 {
					t.Fatal("non-power tile")
				}
			}
		}
	}
}
func TestMergeEndAndSingleMerge(t *testing.T) {
	m := NewMerge(4, 9)
	m.Cells = [36]int{2, 2, 4}
	m.Move(0)
	if m.Cells[0] != 4 || m.Cells[1] != 4 {
		t.Fatal("merged a tile twice")
	}
	for i := 0; i < 16; i++ {
		m.Cells[i] = 2 << ((i + i/4) % 2)
	}
	if m.CanMove() {
		t.Fatal("checkerboard should end")
	}
	m.Cells[1] = m.Cells[0]
	if !m.CanMove() {
		t.Fatal("full board with merge is playable")
	}
}
func TestSnakeQueueAndCollision(t *testing.T) {
	s := NewSnake(0, 1)
	s.Turn(0)
	if len(s.Queue) != 0 {
		t.Fatal("reverse accepted")
	}
	s.Turn(1)
	s.Turn(0)
	s.Turn(3)
	if !reflect.DeepEqual(s.Queue, []int{1, 0}) {
		t.Fatal(s.Queue)
	}
	s.Running = true
	for i := 0; i < 12; i++ {
		s.Tick()
	}
	if s.Direction != 1 {
		t.Fatal("first turn skipped")
	}
	for i := 0; i < 12; i++ {
		s.Tick()
	}
	if s.Direction != 0 {
		t.Fatal("second turn skipped")
	}
	s = NewSnake(2, 1)
	s.Body = []int{0, 1, 2}
	s.Direction = 0
	s.Running = true
	for i := 0; i < 5; i++ {
		s.Tick()
	}
	if !s.Over || s.Running {
		t.Fatal("wall collision")
	}
	s = NewSnake(0, 1)
	for _, b := range s.Body {
		if b == s.Food {
			t.Fatal("food inside snake")
		}
	}
	// Moving into the departing tail is legal, and filling the board wins.
	s.Body = []int{17, 18, 2, 1}
	s.Direction = 1
	s.Food = 40
	s.Running = true
	s.Clock = 11
	s.Tick()
	if s.Over {
		t.Fatal("departing tail collision")
	}
	s.Body = nil
	for i := 0; i < SnakeW*SnakeH; i++ {
		s.Body = append(s.Body, i)
	}
	s.food()
	if !s.Won || s.Food != -1 {
		t.Fatal("full board")
	}
}
func TestLinearWinsAndComputer(t *testing.T) {
	l := NewLinear(true, 1)
	for _, c := range []int{0, 0, 1, 1, 2, 2, 3} {
		if !l.Drop(c) {
			t.Fatal("drop failed")
		}
	}
	if l.Winner != 1 || !l.Over {
		t.Fatal("horizontal win")
	}
	l = NewLinear(true, 1)
	for i := 0; i < 6; i++ {
		l.Drop(0)
	}
	turn := l.Turn
	if l.Drop(0) || l.Turn != turn {
		t.Fatal("full column consumes turn")
	}
	l = NewLinear(false, 1)
	for c := 0; c < 3; c++ {
		l.Cells[35+c] = 1
	}
	l.Turn = 2
	if l.BestColumn() != 3 {
		t.Fatal("must block")
	}
	for r := 3; r < 6; r++ {
		l.Cells[r*7+6] = 2
	}
	if l.BestColumn() != 6 {
		t.Fatal("must take immediate win before block")
	}
	for _, cells := range [][4]int{{0, 8, 16, 24}, {3, 9, 15, 21}, {0, 7, 14, 21}} {
		l = NewLinear(true, 1)
		for _, i := range cells {
			l.Cells[i] = 2
		}
		l.check()
		if l.Winner != 2 {
			t.Fatal("diagonal/vertical win", cells)
		}
	}
}
func TestBouncePhysicsAndServe(t *testing.T) {
	b := NewBounce(2, 12)
	x := b.X
	b.Tick()
	if b.X != x {
		t.Fatal("moved before serve")
	}
	b.Serve()
	b.X = .066
	b.Y = b.Player
	b.VX = -.9
	b.VY = 0
	b.Tick()
	if b.VX <= 0 || b.Contacts != 1 {
		t.Fatal("paddle tunneling", b)
	}
	if math.Abs(b.VY) < .07 {
		t.Fatal("flat trap")
	}
	b.Y = .019
	b.VY = -.8
	b.X = .5
	b.Tick()
	if b.VY < 0 {
		t.Fatal("wall bounce")
	}
	b.X = 1.019
	b.VX = .9
	b.Points[0] = 4
	b.Tick()
	if !b.Over || b.Points[0] != 5 || b.Running {
		t.Fatal("first to five", b)
	}
	b = NewBounce(0, 3)
	b.Serve()
	for i := 0; i < 2000; i++ {
		before := b.Computer
		b.Tick()
		if math.Abs(b.Computer-before) > .42/60+.000001 {
			t.Fatal("AI teleport")
		}
		if !b.Running {
			b.Serve()
		}
		if b.Over {
			break
		}
	}
}
func TestSaveRoundTripAndValidation(t *testing.T) {
	c := NewCollection()
	for i := 0; i < Count; i++ {
		c.New(i, 0, uint64(1+i))
	}
	c.Merge.Move(0)
	c.Snake.Running = true
	c.Snake.Turn(1)
	c.Territory.Tick()
	c.Linear.Drop(3)
	c.Bounce.Serve()
	c.Bounce.Tick()
	data, err := c.Encode()
	if err != nil {
		t.Fatal(err)
	}
	got, err := Decode(data)
	if err != nil {
		t.Fatal(err)
	}
	c.Pause()
	c.Merge.Travel = [36]int{}
	if !reflect.DeepEqual(c, got) {
		t.Fatalf("roundtrip differs\n%+v\n%+v", c, got)
	}
	if !got.Merge.Undo() {
		t.Fatal("lost undo")
	}
	c.Merge.N = 999
	data, _ = json.Marshal(c)
	if _, err = Decode(data); err == nil {
		t.Fatal("accepted invalid dimensions")
	}
	if _, err = Decode([]byte("garbage")); err == nil {
		t.Fatal("accepted corrupt save")
	}
}
func TestEveryGameRemainsValidDuringPlay(t *testing.T) {
	c := NewCollection()
	for i := 0; i < Count; i++ {
		c.New(i, 2, uint64(i+100))
	}
	c.Snake.Running = true
	c.Bounce.Serve()
	for tick := 0; tick < 5000; tick++ {
		if tick%10 == 0 {
			c.Merge.Move(tick / 10 % 4)
		}
		c.Snake.Tick()
		c.Bounce.Tick()
		if tick%30 == 0 && !c.Linear.Over {
			c.Linear.Drop(c.Linear.BestColumn())
		}
		c.Territory.Tick()
		if tick%120 == 0 {
			c.Territory.ComputerOrder(1)
		}

		if err := c.Validate(); err != nil {
			t.Fatalf("invalid at tick %d: %v", tick, err)
		}
	}
}

func TestDemonstrationsShowGameEvents(t *testing.T) {
	c := NewDemonstrations()
	initialOwners := c.Territory.Owners
	var merged, ate, captured, connected, scored bool
	for tick := 1; tick < 720; tick++ {
		c.Demonstrate(tick)
		merged = merged || c.Merge.Score > 0
		ate = ate || c.Snake.Score > 0
		captured = captured || c.Territory.Owners != initialOwners
		connected = connected || c.Linear.Winner != 0
		scored = scored || c.Bounce.Points[0]+c.Bounce.Points[1] > 0
	}
	if !merged || !ate || !captured || !connected || !scored {
		t.Fatalf("demo events merge=%v snake=%v territory=%v linear=%v bounce=%v", merged, ate, captured, connected, scored)
	}
}
