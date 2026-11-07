package play

type Linear struct {
	Cells        [42]int
	Turn, Winner int
	TwoPlayers   bool
	Human        int
	Over         bool
	Winning      [4]int
	Last         int
	Starting     int
}

func NewLinear(two bool, starting int) Linear {
	if starting != 2 {
		starting = 1
	}
	return Linear{Turn: starting, Starting: starting, Human: 1, TwoPlayers: two, Last: -1, Winning: [4]int{-1, -1, -1, -1}}
}
func (l *Linear) Landing(c int) int {
	if c < 0 || c >= 7 {
		return -1
	}
	for r := 5; r >= 0; r-- {
		if l.Cells[r*7+c] == 0 {
			return r*7 + c
		}
	}
	return -1
}
func (l *Linear) Drop(c int) bool {
	if l.Over {
		return false
	}
	i := l.Landing(c)
	if i < 0 {
		return false
	}
	l.Cells[i] = l.Turn
	l.Last = i
	l.check()
	if !l.Over {
		l.Turn = 3 - l.Turn
	}
	return true
}
func (l *Linear) check() {
	for r := 0; r < 6; r++ {
		for c := 0; c < 7; c++ {
			v := l.Cells[r*7+c]
			if v == 0 {
				continue
			}
			for _, d := range [][2]int{{1, 0}, {0, 1}, {1, 1}, {1, -1}} {
				var win [4]int
				ok := true
				for k := 0; k < 4; k++ {
					x, y := c+k*d[0], r+k*d[1]
					if x < 0 || x >= 7 || y < 0 || y >= 6 || l.Cells[y*7+x] != v {
						ok = false
						break
					}
					win[k] = y*7 + x
				}
				if ok {
					l.Winner = v
					l.Over = true
					l.Winning = win
					return
				}
			}
		}
	}
	for _, v := range l.Cells {
		if v == 0 {
			return
		}
	}
	l.Over = true
}
func (l *Linear) BestColumn() int {
	order := []int{3, 2, 4, 1, 5, 0, 6}
	for _, player := range []int{l.Turn, 3 - l.Turn} {
		for _, c := range order {
			if l.Landing(c) < 0 {
				continue
			}
			copy := *l
			copy.Turn = player
			copy.Drop(c)
			if copy.Winner == player {
				return c
			}
		}
	}
	for _, c := range order {
		if l.Landing(c) >= 0 {
			return c
		}
	}
	return -1
}
