package play

import "math"

// A versioned seed reproduces immutable geometry without bloating every game save.
// Version 1 must remain stable for saved matches. Seed zero represents legacy maps.
type Border struct {
	A, B   int
	Points []Point
}

type Terrain struct {
	Borders  []Border `json:"-"`
	Version  int
	Seed     uint64
	Polygons [12][]Point  `json:"-"`
	Centers  [12]Point    `json:"-"`
	Links    [12][12]bool `json:"-"`
	Rivers   [][]Point    `json:"-"`
}

func GenerateTerrain(seed uint64) Terrain {
	if seed == 0 {
		seed = 1
	}
	t := Terrain{Version: 1, Seed: seed}
	r := RNG(seed)
	var sites [12]Point
	for i := range sites {
		for attempt := 0; ; attempt++ {
			a := float64(r.Next(1000000)) / 1000000 * 2 * math.Pi
			radius := math.Sqrt(float64(r.Next(1000000))/1000000) * .36
			p := Point{.5 + math.Cos(a)*radius, .5 + math.Sin(a)*radius}
			ok := true
			for j := 0; j < i; j++ {
				if math.Hypot(p.X-sites[j].X, p.Y-sites[j].Y) < .115 {
					ok = false
				}
			}
			if ok || attempt > 5000 {
				sites[i] = p
				break
			}
		}
	}
	phase := float64(r.Next(10000)) / 10000 * 2 * math.Pi
	warp := func(p Point) Point {
		return Point{p.X + .025*math.Sin(8*p.Y+phase) + .012*math.Sin(17*p.Y+4*p.X+phase), p.Y + .025*math.Sin(7*p.X+phase) + .012*math.Sin(14*p.X+3*p.Y+phase)}
	}
	coast := []Point{}
	for i := 0; i < 36; i++ {
		angle := float64(i) * 2 * math.Pi / 36
		coast = append(coast, Point{.5 + .43*math.Cos(angle), .5 + .43*math.Sin(angle)})
	}
	for i, site := range sites {
		poly := append([]Point(nil), coast...)
		for j, other := range sites {
			if i == j {
				continue
			}
			nx, ny := other.X-site.X, other.Y-site.Y
			limit := (other.X*other.X + other.Y*other.Y - site.X*site.X - site.Y*site.Y) / 2
			poly = clipHalfPlane(poly, nx, ny, limit)
		}
		for k, a := range poly {
			b := poly[(k+1)%len(poly)]
			length := math.Hypot(a.X-b.X, a.Y-b.Y)
			parts := max(1, int(math.Ceil(length/.018)))
			edge := make([]Point, 0, parts+1)
			for n := 0; n <= parts; n++ {
				u := float64(n) / float64(parts)
				edge = append(edge, warp(Point{a.X + (b.X-a.X)*u, a.Y + (b.Y-a.Y)*u}))
			}
			t.Polygons[i] = append(t.Polygons[i], edge[:len(edge)-1]...)
			neighbor := -1
			if length > 1e-6 {
				for j, other := range sites {
					if i == j {
						continue
					}
					nx, ny := other.X-site.X, other.Y-site.Y
					limit := (other.X*other.X + other.Y*other.Y - site.X*site.X - site.Y*site.Y) / 2
					if math.Abs(a.X*nx+a.Y*ny-limit) < 1e-8 && math.Abs(b.X*nx+b.Y*ny-limit) < 1e-8 {
						t.Links[i][j] = true
						neighbor = j
						if i < j && (i+j)%3 == 0 {
							t.Rivers = append(t.Rivers, edge)
						}
					}
				}
			}
			if neighbor < 0 || i < neighbor {
				t.Borders = append(t.Borders, Border{i, neighbor, edge})
			}
		}
		t.Centers[i] = warp(site)
	}
	return t
}
func clipHalfPlane(poly []Point, nx, ny, limit float64) []Point {
	result := []Point{}
	if len(poly) == 0 {
		return result
	}
	for i, a := range poly {
		b := poly[(i+1)%len(poly)]
		da := a.X*nx + a.Y*ny - limit
		db := b.X*nx + b.Y*ny - limit
		if da <= 1e-12 {
			result = append(result, a)
		}
		if (da < 0) != (db < 0) {
			u := da / (da - db)
			result = append(result, Point{a.X + (b.X-a.X)*u, a.Y + (b.Y-a.Y)*u})
		}
	}
	return result
}
func LegacyTerrain() Terrain {
	t := Terrain{Version: 1}
	for i := 0; i < 12; i++ {
		t.Polygons[i] = Polygon(i)
		r, c := i/4, i%4
		corners := []Point{MapNodes[r][c], MapNodes[r][c+1], MapNodes[r+1][c+1], MapNodes[r+1][c]}
		neighbors := []int{-1, -1, -1, -1}
		if r > 0 {
			neighbors[0] = i - 4
		}
		if c < 3 {
			neighbors[1] = i + 1
		}
		if r < 2 {
			neighbors[2] = i + 4
		}
		if c > 0 {
			neighbors[3] = i - 1
		}
		for side, b := range neighbors {
			if b < 0 || i < b {
				t.Borders = append(t.Borders, Border{i, b, boundary(corners[side], corners[(side+1)%4])})
			}
		}
		t.Centers[i] = Center(i)
		for j := 0; j < 12; j++ {
			t.Links[i][j] = Adjacent(i, j)
		}
	}
	return t
}
func (t Terrain) Adjacent(a, b int) bool {
	return a >= 0 && a < 12 && b >= 0 && b < 12 && t.Links[a][b]
}
func (t Terrain) At(x, y float64) int {
	for id, p := range t.Polygons {
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
func (t Terrain) Neighbor(id, dir int) int {
	if id < 0 || id >= 12 {
		return 0
	}
	dx, dy := [4]float64{-1, 0, 1, 0}[dir], [4]float64{0, -1, 0, 1}[dir]
	best := id
	cost := math.Inf(1)
	a := t.Centers[id]
	for i, b := range t.Centers {
		vx, vy := b.X-a.X, b.Y-a.Y
		forward := vx*dx + vy*dy
		if forward <= .01 {
			continue
		}
		score := (vx*vx + vy*vy) / forward
		if score < cost {
			best = i
			cost = score
		}
	}
	return best
}
