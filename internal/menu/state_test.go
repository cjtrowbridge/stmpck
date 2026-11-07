package menu

import (
	"slices"
	"testing"
)

func TestFreshSessionDefaults(t *testing.T) {
	s := New()
	if s.Theme != DefaultIntensity || !s.Autoplay() || s.Sound || s.Tab != Personal {
		t.Fatal("expected middle, autoplay, silent, personal", s)
	}
	s.Theme = 2
	s.Cycle(-1)
	s.Sound = true
	fresh := New()
	if fresh.Theme != DefaultIntensity || !fresh.Autoplay() || fresh.Sound {
		t.Fatal("session settings leaked into a new launch")
	}
}

func TestNavigationPreservesPosition(t *testing.T) {
	s := New()
	s.Move(4, 400)
	if s.Focus != 4 || s.Scroll <= 0 {
		t.Fatal("last card must be visible", s)
	}
	before := s.Scroll
	s.Open(4)
	s.Back()
	if s.Scroll != before || s.Selected != -1 {
		t.Fatal("return lost position", s)
	}
	s.Move(-20, 400)
	if s.Focus != 0 || s.Scroll != 0 {
		t.Fatal("top boundary", s)
	}
	s.Scroll = 99999
	s.Clamp(2000)
	if s.Scroll != 0 {
		t.Fatal("large viewport must not scroll", s)
	}
}

func TestGesturesAndIntensityBounds(t *testing.T) {
	for _, tt := range []struct {
		dx, dy float64
		want   string
	}{{2, 3, "tap"}, {80, 12, "horizontal"}, {20, 80, "vertical"}, {30, 30, "vertical"}} {
		if got := Gesture(tt.dx, tt.dy); got != tt.want {
			t.Fatalf("%+v: %s", tt, got)
		}
	}
	s := New()
	s.Cycle(-1)
	s.Cycle(-10)
	if s.Theme != 0 || s.Autoplay() {
		t.Fatal("left must clamp with autoplay off", s)
	}
	s.Cycle(-1)
	if s.Theme != 0 {
		t.Fatal("left wrapped")
	}
	s.Cycle(99)
	if s.Theme != RainbowFull {
		t.Fatal("right did not clamp")
	}
	s.Cycle(1)
	if s.Theme != RainbowFull {
		t.Fatal("right wrapped")
	}
	s.Open(-1)
	s.Open(Count)
	if s.Selected != -1 {
		t.Fatal("invalid selection accepted")
	}
}

func TestPersonalSortingAndPreviewSelection(t *testing.T) {
	s := New()
	s.Open(3)
	s.Back()
	if s.Plays != ([Count]uint64{}) {
		t.Fatal("preview counted as play")
	}
	s.RecordPlay(3)
	s.RecordPlay(3)
	s.RecordPlay(1)
	want := []int{3, 1, 0, 2, 4}
	if got := s.Order(); !slices.Equal(got, want) {
		t.Fatalf("personal: %v", got)
	}
	s.Open(0)
	if s.Selected != 3 {
		t.Fatal("selected row must open its game, not its position")
	}
}

func TestWorldwideAndDiscovery(t *testing.T) {
	s := New()
	s.RecordPlay(3)
	s.SetTab(Worldwide)
	if s.HasGlobalRanking || !slices.Equal(s.Order(), []int{1, 0, 4, 3, 2}) {
		t.Fatal("placeholder order", s.Order())
	}
	if s.SetGlobalRanking([]int{0, 1, 1, 3, 4}) || s.SetGlobalRanking([]int{0}) {
		t.Fatal("invalid ranking")
	}
	if !s.SetGlobalRanking([]int{4, 3, 2, 1, 0}) {
		t.Fatal("valid ranking")
	}
	s.SetTab(Discover)
	if !slices.Equal(s.Order(), []int{4, 2, 1, 0}) {
		t.Fatal("discovery must filter played and preserve global ranking", s.Order())
	}
	s.RecordPlay(4)
	if !slices.Equal(s.Order(), []int{2, 1, 0}) {
		t.Fatal("newly played game remains discoverable")
	}
	for id := 0; id < Count; id++ {
		s.RecordPlay(id)
	}
	s.Scroll = 100
	s.Focus = 4
	s.Clamp(600)
	s.Move(1, 600)
	s.Open(0)
	if len(s.Order()) != 0 || s.Scroll != 0 || s.Focus != 0 || s.Selected != -1 {
		t.Fatal("empty discovery navigation", s)
	}
}
func TestResponsiveRows(t *testing.T) {
	s := New()
	s.RowHeight = 440
	s.Move(4, 800)
	if s.Scroll != float64(2*440+Header+Footer-800) {
		t.Fatal("row scroll bounds", s.Scroll)
	}
	s.SetTab(Discover)
	for i := 0; i < 4; i++ {
		s.RecordPlay(i)
	}
	s.Clamp(800)
	if s.Scroll != 0 {
		t.Fatal("filtered catalog should fit")
	}
}

func TestIntensityAutoplaySequence(t *testing.T) {
	s := New()
	for _, step := range []struct {
		delta, level int
		autoplay     bool
	}{
		{-1, Still, false}, {-1, Wireframe, false}, {-1, Wireframe, false},
		{1, Still, false}, {1, DefaultIntensity, true}, {1, Vivid, true}, {1, Breathe, true}, {1, RainbowSoft, true}, {1, RainbowRich, true}, {1, RainbowFull, true}, {1, RainbowFull, true},
	} {
		s.Cycle(step.delta)
		if s.Theme != step.level || s.Autoplay() != step.autoplay {
			t.Fatalf("step %+v: %+v", step, s)
		}
	}
}

func TestBackGoesUpOneLevel(t *testing.T) {
	s := New()
	s.Scroll = 123
	s.Open(2)
	s.Settings = true
	s.Back()
	if s.Settings || s.Selected != 2 {
		t.Fatal("Back must close settings before leaving preview", s)
	}
	s.Back()
	if s.Selected != -1 || s.Scroll != 123 {
		t.Fatal("Back must return to the same catalog position", s)
	}
	s.Back()
	if s.Selected != -1 || s.Scroll != 123 {
		t.Fatal("Back at home must stay home", s)
	}
	s.Settings = true
	s.Back()
	if s.Settings || s.Selected != -1 {
		t.Fatal("home settings must close to home", s)
	}
}

func TestMosaicHitTargets(t *testing.T) {
	s := New()
	width := 450
	s.RowHeight = int(TileWidth(width)) + 14
	if s.Rows() != 2 {
		t.Fatal("five games should occupy two rows")
	}
	for i := 0; i < Count; i++ {
		x, y, w, h := s.TileRect(i, width)
		if s.TileAt(x+w/2, y+h/2, width) != i {
			t.Fatalf("wrong tile %d", i)
		}
	}
	x, y, w, h := s.TileRect(5, width)
	if s.TileAt(x+w/2, y+h/2, width) != -1 {
		t.Fatal("empty last cell opens a game")
	}
	x, y, w, _ = s.TileRect(0, width)
	if s.TileAt(x+w+Gap/2, y+10, width) != -1 {
		t.Fatal("gutter opens a game")
	}
	s.Scroll = 40
	x, y, w, h = s.TileRect(3, width)
	if s.TileAt(x+w/2, y+h/2, width) != 3 {
		t.Fatal("scrolled hit test")
	}
}
