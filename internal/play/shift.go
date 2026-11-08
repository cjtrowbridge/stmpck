package play

type Shift struct {
	Cells, Initial, Previous [16]int
	Moves, PreviousMoves     int
	CanUndo, Won             bool
}

func NewShift(seed uint64) Shift {
	s := Shift{}
	for i := 0; i < 15; i++ {
		s.Cells[i] = i + 1
	}
	r := RNG(seed)
	last := -1
	for step := 0; step < 240; step++ {
		empty := s.Empty()
		choices := []int{}
		for _, d := range []int{-4, -1, 1, 4} {
			i := empty + d
			if i >= 0 && i < 16 && i != last && absInt(i%4-empty%4)+absInt(i/4-empty/4) == 1 {
				choices = append(choices, i)
			}
		}
		i := choices[r.Next(len(choices))]
		s.Cells[empty], s.Cells[i] = s.Cells[i], 0
		last = empty
	}
	if s.Solved() {
		s.Cells[14], s.Cells[15] = s.Cells[15], s.Cells[14]
	}
	s.Initial = s.Cells
	return s
}
func absInt(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
func (s *Shift) Empty() int {
	for i, v := range s.Cells {
		if v == 0 {
			return i
		}
	}
	return -1
}
func (s *Shift) Solved() bool {
	for i, v := range s.Cells {
		if v != (i+1)%16 {
			return false
		}
	}
	return true
}
func (s *Shift) Move(i int) bool {
	if s.Won || i < 0 || i >= 16 {
		return false
	}
	empty := s.Empty()
	if absInt(i%4-empty%4)+absInt(i/4-empty/4) != 1 {
		return false
	}
	s.Previous = s.Cells
	s.PreviousMoves = s.Moves
	s.CanUndo = true
	s.Cells[empty], s.Cells[i] = s.Cells[i], 0
	s.Moves++
	s.Won = s.Solved()
	return true
}
func (s *Shift) Direction(d int) bool {
	if d < 0 || d > 3 {
		return false
	}
	return s.Move(s.Empty() + [4]int{-1, -4, 1, 4}[d])
}
func (s *Shift) Undo() {
	if s.CanUndo {
		s.Cells = s.Previous
		s.Moves = s.PreviousMoves
		s.CanUndo = false
		s.Won = s.Solved()
	}
}
func (s *Shift) Replay() { s.Cells = s.Initial; s.Moves = 0; s.CanUndo = false; s.Won = s.Solved() }
