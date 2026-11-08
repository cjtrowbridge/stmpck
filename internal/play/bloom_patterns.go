package play

// Compact, shared definitions for the picker, placement preview and gallery.
type LifePattern struct {
	Name string
	Rows []string
}

var BloomPatterns = []LifePattern{
	{"Glider", []string{".O.", "..O", "OOO"}},
	{"Blinker", []string{"OOO"}},
	{"Toad", []string{".OOO", "OOO."}},
	{"Beacon", []string{"OO..", "OO..", "..OO", "..OO"}},
	{"Block", []string{"OO", "OO"}},
	{"Pulsar", []string{"..OOO...OOO..", ".............", "O....O.O....O", "O....O.O....O", "O....O.O....O", "..OOO...OOO..", ".............", "..OOO...OOO..", "O....O.O....O", "O....O.O....O", "O....O.O....O", ".............", "..OOO...OOO.."}},
}

// Anchor is the pattern's center cell. The simulation and stamps both wrap.
func BloomPatternCells(pattern, anchor int) []int {
	if pattern < 0 || pattern >= len(BloomPatterns) || anchor < 0 || anchor >= BloomN*BloomN {
		return nil
	}
	p := BloomPatterns[pattern]
	x, y := anchor%BloomN-len(p.Rows[0])/2, anchor/BloomN-len(p.Rows)/2
	cells := []int{}
	for dy, row := range p.Rows {
		for dx, v := range row {
			if v == 'O' {
				cells = append(cells, ((y+dy+BloomN)%BloomN)*BloomN+(x+dx+BloomN)%BloomN)
			}
		}
	}
	return cells
}
func (b *Bloom) Stamp(pattern, anchor int) bool {
	cells := BloomPatternCells(pattern, anchor)
	if len(cells) == 0 {
		return false
	}
	for _, i := range cells {
		b.Cells[i] = true
	}
	b.Clock = 0
	return true
}
func BloomDemonstration(pattern int) Bloom {
	b := Bloom{Speed: 1}
	b.Stamp(pattern, BloomN/2*BloomN+BloomN/2)
	return b
}
