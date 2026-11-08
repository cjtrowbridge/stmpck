package play

import "math/bits"

type Deduction struct {
	Clues, Cells, Solution, Previous [81]int
	CanUndo, Won                     bool
}

func sudokuAllowed(board *[81]int, i, v int) bool {
	for k := 0; k < 9; k++ {
		if k != i%9 && board[i/9*9+k] == v || k != i/9 && board[k*9+i%9] == v {
			return false
		}
	}
	x, y := i%9/3*3, i/9/3*3
	for dy := 0; dy < 3; dy++ {
		for dx := 0; dx < 3; dx++ {
			j := (y+dy)*9 + x + dx
			if j != i && board[j] == v {
				return false
			}
		}
	}
	return true
}

// CountSudokuSolutions stops at limit and uses the smallest candidate set first.
func CountSudokuSolutions(board [81]int, limit int) int {
	if limit < 1 {
		return 0
	}
	for i, v := range board {
		if v < 0 || v > 9 || v != 0 && !sudokuAllowed(&board, i, v) {
			return 0
		}
	}
	var search func() int
	search = func() int {
		best, mask, fewest := -1, 0, 10
		for i, v := range board {
			if v != 0 {
				continue
			}
			m := 0
			for n := 1; n <= 9; n++ {
				if sudokuAllowed(&board, i, n) {
					m |= 1 << n
				}
			}
			count := bits.OnesCount(uint(m))
			if count == 0 {
				return 0
			}
			if count < fewest {
				best, mask, fewest = i, m, count
			}
		}
		if best < 0 {
			return 1
		}
		total := 0
		for n := 1; n <= 9; n++ {
			if mask&(1<<n) == 0 {
				continue
			}
			board[best] = n
			total += search()
			board[best] = 0
			if total >= limit {
				return limit
			}
		}
		return total
	}
	return search()
}
func NewDeduction(seed uint64) Deduction {
	r := RNG(seed)
	shuffle := func(values []int) {
		for i := len(values) - 1; i > 0; i-- {
			j := r.Next(i + 1)
			values[i], values[j] = values[j], values[i]
		}
	}
	axis := func() []int {
		bands := []int{0, 1, 2}
		shuffle(bands)
		out := []int{}
		for _, b := range bands {
			rows := []int{0, 1, 2}
			shuffle(rows)
			for _, v := range rows {
				out = append(out, 3*b+v)
			}
		}
		return out
	}
	rows, cols := axis(), axis()
	digits := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	shuffle(digits)
	d := Deduction{}
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			d.Solution[y*9+x] = digits[(rows[y]*3+rows[y]/3+cols[x])%9]
		}
	}
	d.Clues = d.Solution
	order := make([]int, 81)
	for i := range order {
		order[i] = i
	}
	shuffle(order)
	remaining := 81
	for _, i := range order {
		if remaining <= 36 {
			break
		}
		v := d.Clues[i]
		d.Clues[i] = 0
		if CountSudokuSolutions(d.Clues, 2) != 1 {
			d.Clues[i] = v
		} else {
			remaining--
		}
	}
	d.Cells = d.Clues
	return d
}
func (d *Deduction) Set(i, v int) bool {
	if i < 0 || i >= 81 || v < 0 || v > 9 || d.Clues[i] != 0 || d.Cells[i] == v || d.Won {
		return false
	}
	d.Previous = d.Cells
	d.CanUndo = true
	d.Cells[i] = v
	d.Won = d.Cells == d.Solution
	return true
}
func (d *Deduction) Conflict(i int) bool {
	return i >= 0 && i < 81 && d.Cells[i] != 0 && !sudokuAllowed(&d.Cells, i, d.Cells[i])
}
func (d *Deduction) Undo() {
	if d.CanUndo {
		d.Cells = d.Previous
		d.CanUndo = false
		d.Won = d.Cells == d.Solution
	}
}
