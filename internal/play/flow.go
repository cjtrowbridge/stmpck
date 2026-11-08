package play

// Ports run clockwise from north. A generated spanning tree guarantees a solution.
type Flow struct {
	Cells, Initial, Solution [25]uint8
	Moves                    int
	Won                      bool
}

func rotatePorts(v uint8) uint8 { return (v<<1 | v>>3) & 15 }
func flowNeighbor(i, d int) int {
	x, y := i%5, i/5
	x += [4]int{0, 1, 0, -1}[d]
	y += [4]int{-1, 0, 1, 0}[d]
	if x < 0 || x >= 5 || y < 0 || y >= 5 {
		return -1
	}
	return y*5 + x
}
func NewFlow(seed uint64) Flow {
	f := Flow{}
	r := RNG(seed)
	seen := [25]bool{}
	stack := []int{12}
	seen[12] = true
	for len(stack) > 0 {
		i := stack[len(stack)-1]
		choices := []int{}
		for d := 0; d < 4; d++ {
			n := flowNeighbor(i, d)
			if n >= 0 && !seen[n] {
				choices = append(choices, d)
			}
		}
		if len(choices) == 0 {
			stack = stack[:len(stack)-1]
			continue
		}
		d := choices[r.Next(len(choices))]
		n := flowNeighbor(i, d)
		f.Solution[i] |= 1 << d
		f.Solution[n] |= 1 << ((d + 2) % 4)
		seen[n] = true
		stack = append(stack, n)
	}
	f.Cells = f.Solution
	for i := range f.Cells {
		for k := r.Next(4); k > 0; k-- {
			f.Cells[i] = rotatePorts(f.Cells[i])
		}
	}
	if f.Complete() {
		for i, v := range f.Cells {
			if rotatePorts(v) != v {
				f.Cells[i] = rotatePorts(v)
				break
			}
		}
	}
	f.Initial = f.Cells
	return f
}
func (f *Flow) Connected() [25]bool {
	seen := [25]bool{}
	seen[12] = true
	queue := []int{12}
	for len(queue) > 0 {
		i := queue[0]
		queue = queue[1:]
		for d := 0; d < 4; d++ {
			n := flowNeighbor(i, d)
			if n >= 0 && !seen[n] && f.Cells[i]&(1<<d) != 0 && f.Cells[n]&(1<<((d+2)%4)) != 0 {
				seen[n] = true
				queue = append(queue, n)
			}
		}
	}
	return seen
}
func (f *Flow) Complete() bool {
	seen := f.Connected()
	for i, v := range f.Cells {
		if !seen[i] {
			return false
		}
		for d := 0; d < 4; d++ {
			if v&(1<<d) == 0 {
				continue
			}
			n := flowNeighbor(i, d)
			if n < 0 || f.Cells[n]&(1<<((d+2)%4)) == 0 {
				return false
			}
		}
	}
	return true
}
func (f *Flow) Rotate(i int) {
	if i < 0 || i >= 25 || f.Won {
		return
	}
	f.Cells[i] = rotatePorts(f.Cells[i])
	f.Moves++
	f.Won = f.Complete()
}
func (f *Flow) Replay() { f.Cells = f.Initial; f.Moves = 0; f.Won = false }
