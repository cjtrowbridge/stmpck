package play

const SnakeW = 16
const SnakeH = 12

type Snake struct {
	Body               []int
	Food               int
	Direction          int
	Queue              []int
	Speed              int
	Random             RNG
	Score              int
	Over, Won, Running bool
	Clock              int
}

func NewSnake(speed int, seed uint64) Snake {
	s := Snake{Body: []int{6*SnakeW + 7, 6*SnakeW + 6, 6*SnakeW + 5}, Direction: 2, Speed: min(2, max(0, speed)), Random: RNG(seed)}
	s.food()
	return s
}
func (s *Snake) food() {
	empty := []int{}
	for i := 0; i < SnakeW*SnakeH; i++ {
		found := false
		for _, b := range s.Body {
			if b == i {
				found = true
				break
			}
		}
		if !found {
			empty = append(empty, i)
		}
	}
	if len(empty) == 0 {
		s.Won = true
		s.Over = true
		s.Running = false
		s.Food = -1
		return
	}
	s.Food = empty[s.Random.Next(len(empty))]
}
func (s *Snake) Turn(d int) {
	if s.Over || d < 0 || d > 3 {
		return
	}
	last := s.Direction
	if len(s.Queue) > 0 {
		last = s.Queue[len(s.Queue)-1]
	}
	if d == last || (d+2)%4 == last {
		return
	}
	if len(s.Queue) < 2 {
		s.Queue = append(s.Queue, d)
	}
}
func (s *Snake) Tick() bool {
	if !s.Running || s.Over {
		return false
	}
	s.Clock++
	if s.Clock < [3]int{12, 8, 5}[s.Speed] {
		return false
	}
	s.Clock = 0
	if len(s.Queue) > 0 {
		s.Direction = s.Queue[0]
		s.Queue = s.Queue[1:]
	}
	h := s.Body[0]
	x, y := h%SnakeW, h/SnakeW
	switch s.Direction {
	case 0:
		x--
	case 1:
		y--
	case 2:
		x++
	case 3:
		y++
	}
	next := y*SnakeW + x
	grow := next == s.Food
	collision := x < 0 || x >= SnakeW || y < 0 || y >= SnakeH
	end := len(s.Body) - 1
	if grow {
		end++
	}
	for i := 0; i < end; i++ {
		if s.Body[i] == next {
			collision = true
		}
	}
	if collision {
		s.Over = true
		s.Running = false
		return true
	}
	s.Body = append([]int{next}, s.Body...)
	if grow {
		s.Score++
		s.food()
	} else {
		s.Body = s.Body[:len(s.Body)-1]
	}
	return true
}
