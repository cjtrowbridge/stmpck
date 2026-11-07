package play

// Merge retains the exact pre-move random state in its single undo slot.
type MergePosition struct {
	Cells  [36]int
	Score  int
	Random RNG
	Won    bool
}
type Merge struct {
	MergePosition
	N            int
	UndoPosition MergePosition
	HasUndo      bool
	Over         bool
	Spawn        int
	Changed      [36]bool
	Travel       [36]int `json:"-"`
}

func NewMerge(n int, seed uint64) Merge {
	if n < 4 || n > 6 {
		n = 4
	}
	m := Merge{N: n, MergePosition: MergePosition{Random: RNG(seed)}}
	m.spawn()
	m.spawn()
	return m
}
func (m *Merge) spawn() {
	empty := []int{}
	for i := 0; i < m.N*m.N; i++ {
		if m.Cells[i] == 0 {
			empty = append(empty, i)
		}
	}
	if len(empty) == 0 {
		return
	}
	m.Spawn = empty[m.Random.Next(len(empty))]
	v := 2
	if m.Random.Next(10) == 0 {
		v = 4
	}
	m.Cells[m.Spawn] = v
}

// Directions: left, up, right, down.
func (m *Merge) Move(dir int) bool {
	if dir < 0 || dir > 3 || m.Over {
		return false
	}
	before := m.MergePosition
	at := func(line, pos int) int {
		switch dir {
		case 0:
			return line*m.N + pos
		case 1:
			return pos*m.N + line
		case 2:
			return line*m.N + m.N - 1 - pos
		default:
			return (m.N-1-pos)*m.N + line
		}
	}
	for row := 0; row < m.N; row++ {
		values := []int{}
		origins := []int{}
		for c := 0; c < m.N; c++ {
			if v := m.Cells[at(row, c)]; v != 0 {
				values = append(values, v)
				origins = append(origins, at(row, c))
			}
		}
		merged := []int{}
		for i := 0; i < len(values); i++ {
			v := values[i]
			m.Travel[origins[i]] = at(row, len(merged))
			if i+1 < len(values) && v == values[i+1] && v < 1<<30 {
				m.Travel[origins[i+1]] = at(row, len(merged))
				v *= 2
				m.Score += v
				i++
			}
			merged = append(merged, v)
		}
		for c := 0; c < m.N; c++ {
			v := 0
			if c < len(merged) {
				v = merged[c]
			}
			m.Cells[at(row, c)] = v
		}
	}
	if before.Cells == m.Cells {
		m.MergePosition = before
		return false
	}
	m.UndoPosition = before
	m.HasUndo = true
	m.spawn()
	for i := 0; i < m.N*m.N; i++ {
		m.Changed[i] = m.Cells[i] != before.Cells[i]
		if m.Cells[i] >= 2048 {
			m.Won = true
		}
	}
	m.Over = !m.CanMove()
	return true
}
func (m *Merge) CanMove() bool {
	for i := 0; i < m.N*m.N; i++ {
		v := m.Cells[i]
		if v == 0 {
			return true
		}
		if v < 1<<30 && ((i%m.N+1 < m.N && v == m.Cells[i+1]) || (i/m.N+1 < m.N && v == m.Cells[i+m.N])) {
			return true
		}
	}
	return false
}
func (m *Merge) Undo() bool {
	if !m.HasUndo {
		return false
	}
	m.MergePosition = m.UndoPosition
	m.HasUndo = false
	m.Over = !m.CanMove()
	m.Changed = [36]bool{}
	m.Spawn = -1
	return true
}
