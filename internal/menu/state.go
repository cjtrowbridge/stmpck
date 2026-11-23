// Package menu holds the launcher's input-independent navigation state.
package menu

import (
	"ebe-boilerplate/internal/play"
	"math"
	"sort"
)

const (
	Columns = 3
	Gap     = 8
	Count   = play.Count
	Header  = 100
	Row     = 214
	Footer  = 24
)

// Synthetic familiarity weights, not measured market share. See menu-tabs plan.
var PlaceholderPopularity = [Count]int{90, 100, 60, 70, 80, 55, 50, 65, 75, 85, 58, 78, 98, 92, 72, 82, 76, 68}

var Names = [Count]string{"Merge", "Snake", "Territory", "Linear", "Bounce", "Flow", "Caliban", "Bloom", "Echo", "Brick", "Planets", "Drift", "Stack", "Deduction", "Bubble Volley", "Wave", "Crossing", "Shift"}

const (
	Personal = iota
	Worldwide
	Discover
)

var TabLabels = [3]string{"⭐", "👥", "✨"}

const (
	Wireframe = iota
	Still
	DefaultIntensity
	Vivid
	Breathe
	RainbowSoft
	RainbowRich
	RainbowFull
)

var IntensityLabels = [...]string{"🙂", "🙂", "🙂", "🤩", "🤩", "🤩", "🤩", "🤩"}

type State struct {
	CompactGallery   bool
	RowHeight        int
	Tab              int
	Sound            bool
	Plays            [Count]uint64
	GlobalOrder      [Count]int
	HasGlobalRanking bool
	Scroll           float64
	Focus            int
	Selected         int
	Theme            int
	Settings         bool
}

func New() State {
	s := State{Selected: -1, Theme: DefaultIntensity, RowHeight: Row}
	return s
}

// Order never treats browsing previews as playing games.
func (s *State) Order() []int {
	order := make([]int, Count)
	for i := range order {
		order[i] = i
	}
	switch s.Tab {
	case Personal:
		sort.SliceStable(order[:], func(i, j int) bool { return s.Plays[order[i]] > s.Plays[order[j]] })
	case Worldwide, Discover:
		if s.HasGlobalRanking {
			copy(order, s.GlobalOrder[:])
		} else {
			sort.SliceStable(order, func(i, j int) bool { return PlaceholderPopularity[order[i]] > PlaceholderPopularity[order[j]] })
		}
		if s.Tab == Discover {
			unplayed := order[:0]
			for _, id := range order {
				if s.Plays[id] == 0 {
					unplayed = append(unplayed, id)
				}
			}
			order = unplayed
		}

	}
	return order
}

// RecordPlay is reserved for real gameplay integration, not preview selection.
func (s *State) RecordPlay(id int) {
	if id >= 0 && id < Count {
		s.Plays[id]++
	}
}

// SetGlobalRanking accepts a complete externally sourced ranking, when one exists.
func (s *State) SetGlobalRanking(order []int) bool {
	if len(order) != Count {
		return false
	}
	var seen [Count]bool
	for _, id := range order {
		if id < 0 || id >= Count || seen[id] {
			return false
		}
		seen[id] = true
	}
	copy(s.GlobalOrder[:], order)
	s.HasGlobalRanking = true
	return true
}

func (s *State) SetTab(tab int) {
	if tab < 0 || tab > 2 || tab == s.Tab {
		return
	}
	s.Tab = tab
	s.Focus = 0
	s.Scroll = 0
}

func (s *State) HeaderHeight() int {
	if s.CompactGallery && s.Selected < 0 {
		return 48
	}
	return Header
}

func (s *State) Clamp(height int) {
	limit := math.Max(0, float64(s.Rows()*s.RowHeight+s.HeaderHeight()+Footer-height))
	s.Scroll = math.Max(0, math.Min(limit, s.Scroll))
}

func (s *State) Move(delta, height int) {
	s.Focus = max(0, min(len(s.Order())-1, s.Focus+delta))
	top := float64((s.Focus/Columns)*s.RowHeight) - s.Scroll
	bottom := top + float64(s.RowHeight)
	if top < 0 {
		s.Scroll += top
	}
	if bottom > float64(height-s.HeaderHeight()-Footer) {
		s.Scroll += bottom - float64(height-s.HeaderHeight()-Footer)
	}
	s.Clamp(height)
}

func (s *State) Autoplay() bool { return s.Theme >= DefaultIntensity }

func (s *State) Cycle(delta int) { s.Theme = max(0, min(len(IntensityLabels)-1, s.Theme+delta)) }
func (s *State) Open(index int) {
	if index >= 0 && index < len(s.Order()) {
		s.Focus = index
		s.Selected = s.Order()[index]
	}
}
func (s *State) Back() {
	if s.Settings {
		s.Settings = false
		return
	}
	s.Selected = -1
}

// Gesture distinguishes navigation from taps; vertical movement wins ties.
func Gesture(dx, dy float64) string {
	if math.Abs(dx) < 12 && math.Abs(dy) < 12 {
		return "tap"
	}
	if math.Abs(dx) > math.Abs(dy) {
		return "horizontal"
	}
	return "vertical"
}

// The layout preference is a mockup for now; the live gallery always uses Columns.
func (s *State) Rows() int        { return (len(s.Order()) + Columns - 1) / Columns }
func TileWidth(width int) float64 { return math.Max(1, float64(width-32-(Columns-1)*Gap)/Columns) }
func (s *State) TileRect(index, width int) (x, y, w, h float64) {
	w = TileWidth(width)
	h = float64(s.RowHeight - 14)
	x = 16 + float64(index%Columns)*(w+Gap)
	y = float64(s.HeaderHeight()+(index/Columns)*s.RowHeight) - s.Scroll
	return
}
func (s *State) TileAt(x, y float64, width int) int {
	if y < float64(s.HeaderHeight()) {
		return -1
	}
	for i := range s.Order() {
		bx, by, bw, bh := s.TileRect(i, width)
		if x >= bx && x < bx+bw && y >= by && y < by+bh {
			return i
		}
	}
	return -1
}
