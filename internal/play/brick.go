package play

type Brick struct {
	Bricks               [40]bool
	X, Y, VX, VY, Paddle float64
	Level, Lives, Hits   int
	Running, Won, Over   bool
}

func NewBrick(level int) Brick {
	b := Brick{Level: max(0, min(2, level)), Lives: 3, Paddle: .5, X: .5, Y: .85}
	for i := range b.Bricks {
		b.Bricks[i] = b.Level == 0 || b.Level == 1 && (i/8+i%8)%2 == 0 || b.Level == 2 && (i%8 == 0 || i%8 == 7 || i/8%2 == 0)
	}
	return b
}
func (b *Brick) Move(x float64) {
	b.Paddle = clamp(x, .1, .9)
	if !b.Running && b.VX == 0 {
		b.X = b.Paddle
	}
}
func (b *Brick) Serve() {
	if b.Over || b.Won {
		return
	}
	if b.VX == 0 {
		b.X = b.Paddle
		b.Y = .85
		b.VX = .25
		b.VY = -.55
	}
	b.Running = true
}
func BrickRect(i int) (x, y, w, h float64) {
	return .04 + float64(i%8)*.115, .12 + float64(i/8)*.065, .105, .05
}
func (b *Brick) Tick() {
	if !b.Running || b.Over || b.Won {
		return
	}
	for step := 0; step < 4; step++ {
		ox, oy := b.X, b.Y
		b.X += b.VX / 240
		b.Y += b.VY / 240
		if b.X < .012 {
			b.X = .012
			b.VX = -b.VX
		}
		if b.X > .988 {
			b.X = .988
			b.VX = -b.VX
		}
		if b.Y < .012 {
			b.Y = .012
			b.VY = -b.VY
		}
		if b.VY > 0 && oy < .878 && b.Y >= .878 && b.X >= b.Paddle-.112 && b.X <= b.Paddle+.112 {
			b.Y = .878
			b.VX = (b.X - b.Paddle) * 4
			b.VY = -.6
		}
		for i, live := range b.Bricks {
			if !live {
				continue
			}
			x, y, w, h := BrickRect(i)
			if b.X+.012 < x || b.X-.012 > x+w || b.Y+.012 < y || b.Y-.012 > y+h {
				continue
			}
			b.Bricks[i] = false
			b.Hits++
			if ox+.012 <= x || ox-.012 >= x+w {
				b.VX = -b.VX
				b.X = ox
			} else {
				b.VY = -b.VY
				b.Y = oy
			}
			break
		}
		if b.Y > 1.02 {
			b.Lives--
			b.Running = false
			b.VX = 0
			b.VY = 0
			b.X = b.Paddle
			b.Y = .85
			b.Over = b.Lives == 0
			return
		}
		remaining := false
		for _, v := range b.Bricks {
			remaining = remaining || v
		}
		if !remaining {
			b.Won = true
			b.Running = false
			return
		}
	}
}
