package play

const (
	StackW = 10
	StackH = 20
)

type GridCell struct{ X, Y int }

var stackShapes = [7][4]GridCell{
	{{0, 1}, {1, 1}, {2, 1}, {3, 1}}, {{1, 0}, {2, 0}, {1, 1}, {2, 1}},
	{{1, 0}, {0, 1}, {1, 1}, {2, 1}}, {{1, 0}, {2, 0}, {0, 1}, {1, 1}},
	{{0, 0}, {1, 0}, {1, 1}, {2, 1}}, {{0, 0}, {0, 1}, {1, 1}, {2, 1}}, {{2, 0}, {0, 1}, {1, 1}, {2, 1}},
}

type Stack struct {
	Cells                                                            [StackW * StackH]int
	Piece, Next, Rotation, X, Y, Clock, Pace, Lines, Score, BagIndex int
	Bag                                                              [7]int
	RNG                                                              RNG
	Won, Over                                                        bool
}

func StackShape(piece, rotation int) [4]GridCell {
	cells := stackShapes[piece]
	if piece == 1 {
		return cells
	}
	size := 3
	if piece == 0 {
		size = 4
	}
	for k := 0; k < rotation; k++ {
		for i, p := range cells {
			cells[i] = GridCell{size - 1 - p.Y, p.X}
		}
	}
	return cells
}
func NewStack(seed uint64) Stack {
	s := Stack{RNG: RNG(seed), BagIndex: 7, Pace: 1}
	s.Next = s.take()
	s.spawn()
	return s
}
func (s *Stack) take() int {
	if s.BagIndex >= 7 {
		for i := range s.Bag {
			s.Bag[i] = i
		}
		for i := 6; i > 0; i-- {
			j := s.RNG.Next(i + 1)
			s.Bag[i], s.Bag[j] = s.Bag[j], s.Bag[i]
		}
		s.BagIndex = 0
	}
	v := s.Bag[s.BagIndex]
	s.BagIndex++
	return v
}
func (s *Stack) Fits(x, y, rot int) bool {
	for _, p := range StackShape(s.Piece, rot) {
		xx, yy := x+p.X, y+p.Y
		if xx < 0 || xx >= StackW || yy < 0 || yy >= StackH || s.Cells[yy*StackW+xx] != 0 {
			return false
		}
	}
	return true
}
func (s *Stack) spawn() {
	s.Piece = s.Next
	s.Next = s.take()
	s.X = 3
	s.Y = 0
	s.Rotation = 0
	s.Clock = 0
	if !s.Fits(s.X, s.Y, s.Rotation) {
		s.Over = true
	}
}
func (s *Stack) Move(dx int) bool {
	if s.Over || s.Won || !s.Fits(s.X+dx, s.Y, s.Rotation) {
		return false
	}
	s.X += dx
	return true
}
func (s *Stack) Rotate() bool {
	if s.Over || s.Won {
		return false
	}
	rot := (s.Rotation + 1) % 4
	for _, dx := range []int{0, -1, 1, -2, 2} {
		if s.Fits(s.X+dx, s.Y, rot) {
			s.X += dx
			s.Rotation = rot
			return true
		}
	}
	return false
}
func (s *Stack) lock() {
	for _, p := range StackShape(s.Piece, s.Rotation) {
		s.Cells[(s.Y+p.Y)*StackW+s.X+p.X] = s.Piece + 1
	}
	s.clearRows()
	if s.Lines >= 40 {
		s.Won = true
		return
	}
	s.spawn()
}
func (s *Stack) clearRows() {
	dest := StackH - 1
	cleared := 0
	for row := StackH - 1; row >= 0; row-- {
		full := true
		for x := 0; x < StackW; x++ {
			full = full && s.Cells[row*StackW+x] != 0
		}
		if full {
			cleared++
			continue
		}
		for x := 0; x < StackW; x++ {
			s.Cells[dest*StackW+x] = s.Cells[row*StackW+x]
		}
		dest--
	}
	for i := 0; i < (dest+1)*StackW; i++ {
		s.Cells[i] = 0
	}
	s.Lines += cleared
	s.Score += cleared * cleared * 100
}
func (s *Stack) Down() {
	if s.Over || s.Won {
		return
	}
	if s.Fits(s.X, s.Y+1, s.Rotation) {
		s.Y++
	} else {
		s.lock()
	}
	s.Clock = 0
}
func (s *Stack) Drop() {
	if s.Over || s.Won {
		return
	}
	for s.Fits(s.X, s.Y+1, s.Rotation) {
		s.Y++
	}
	s.lock()
}
func (s *Stack) Tick() {
	if s.Over || s.Won {
		return
	}
	s.Clock++
	if s.Clock >= [3]int{90, 60, 30}[s.Pace] {
		s.Down()
	}
}
