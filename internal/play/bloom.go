package play

const BloomN = 24

type Bloom struct {
	Cells        [BloomN * BloomN]bool
	Generation   uint64
	Speed, Clock int
}

func NewBloom() Bloom {
	b := Bloom{Speed: 1}
	for _, i := range []int{8*24 + 9, 9*24 + 10, 10*24 + 8, 10*24 + 9, 10*24 + 10, 15*24 + 15, 15*24 + 16, 15*24 + 17} {
		b.Cells[i] = true
	}
	return b
}
func (b *Bloom) Step() {
	next := [BloomN * BloomN]bool{}
	for i, v := range b.Cells {
		n := 0
		x, y := i%BloomN, i/BloomN
		for dy := -1; dy <= 1; dy++ {
			for dx := -1; dx <= 1; dx++ {
				if dx == 0 && dy == 0 {
					continue
				}
				if b.Cells[((y+dy+BloomN)%BloomN)*BloomN+(x+dx+BloomN)%BloomN] {
					n++
				}
			}
		}
		next[i] = n == 3 || v && n == 2
	}
	b.Cells = next
	b.Generation++
	b.Clock = 0
}
func (b *Bloom) Tick() {
	b.Clock++
	if b.Clock >= [3]int{30, 10, 5}[b.Speed] {
		b.Step()
	}
}
func (b *Bloom) Toggle(i int) {
	if i >= 0 && i < len(b.Cells) {
		b.Cells[i] = !b.Cells[i]
	}
}
func (b *Bloom) Clear() { b.Cells = [BloomN * BloomN]bool{}; b.Generation = 0; b.Clock = 0 }
