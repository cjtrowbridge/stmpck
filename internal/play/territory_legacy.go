package play

import "math"

// Legacy geometry is retained only to preserve existing saved matches.
// Fixed geographic mesh: shared, bent boundaries prevent gaps and ambiguous adjacency.
var MapNodes = [4][5]Point{
	{{.04, .08}, {.24, .025}, {.49, .08}, {.72, .035}, {.96, .10}},
	{{.015, .35}, {.31, .30}, {.45, .42}, {.81, .29}, {.985, .38}},
	{{.05, .76}, {.20, .65}, {.56, .71}, {.69, .68}, {.95, .74}},
	{{.08, .94}, {.34, .985}, {.48, .92}, {.80, .975}, {.92, .92}},
}

func boundary(a, b Point) []Point {
	reverse := a.X > b.X || (a.X == b.X && a.Y > b.Y)
	if reverse {
		a, b = b, a
	}
	dx, dy := b.X-a.X, b.Y-a.Y
	length := math.Hypot(dx, dy)
	p := []Point{a}
	for i := 1; i < 4; i++ {
		t := float64(i) / 4
		bend := math.Sin(t*math.Pi*2+.7) * math.Min(.014, length*.08)
		p = append(p, Point{a.X + t*dx - dy/length*bend, a.Y + t*dy + dx/length*bend})
	}
	p = append(p, b)
	if reverse {
		for i, j := 0, len(p)-1; i < j; i, j = i+1, j-1 {
			p[i], p[j] = p[j], p[i]
		}
	}
	return p
}
func Polygon(id int) []Point {
	if id < 0 || id >= 12 {
		return nil
	}
	r, c := id/4, id%4
	corners := []Point{MapNodes[r][c], MapNodes[r][c+1], MapNodes[r+1][c+1], MapNodes[r+1][c]}
	p := []Point{}
	for i := 0; i < 4; i++ {
		edge := boundary(corners[i], corners[(i+1)%4])
		p = append(p, edge[:len(edge)-1]...)
	}
	return p
}
func Center(id int) Point {
	p := Polygon(id)
	var c Point
	for _, v := range p {
		c.X += v.X
		c.Y += v.Y
	}
	c.X /= float64(len(p))
	c.Y /= float64(len(p))
	return c
}
func TerritoryAt(x, y float64) int {
	for id := 0; id < 12; id++ {
		p := Polygon(id)
		inside := false
		for i, j := 0, len(p)-1; i < len(p); j, i = i, i+1 {
			a, b := p[i], p[j]
			if (a.Y > y) != (b.Y > y) && x < (b.X-a.X)*(y-a.Y)/(b.Y-a.Y)+a.X {
				inside = !inside
			}
		}
		if inside {
			return id
		}
	}
	return -1
}
func Adjacent(a, b int) bool {
	return a >= 0 && a < 12 && b >= 0 && b < 12 && ((a/4 == b/4 && abs(a%4-b%4) == 1) || (a%4 == b%4 && abs(a/4-b/4) == 1))
}
func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
