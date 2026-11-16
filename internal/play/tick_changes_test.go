package play

import (
	"reflect"
	"testing"
)

func TestEchoTickReportsEveryStateChange(t *testing.T) {
	for _, phase := range []int{EchoShow, EchoInput, EchoSuccess, EchoFailed, EchoWon} {
		for _, flash := range []int{0, 1, 12} {
			e := NewEcho(1)
			e.Phase, e.Flash = phase, flash
			for tick := 0; tick < 120; tick++ {
				before := e
				before.Sequence = append([]int(nil), e.Sequence...)
				changed := e.Tick()
				if changed == reflect.DeepEqual(before, e) {
					t.Fatalf("incorrect change report: phase=%d flash=%d tick=%d", phase, flash, tick)
				}
			}
		}
	}
}

func TestClusterTickReportsMotionLandingAndIdle(t *testing.T) {
	c := NewCluster(1)
	if c.Tick() {
		t.Fatal("idle tick reported change")
	}
	c.Fire()
	for tick := 0; tick < 300 && c.Flying; tick++ {
		before := c
		if !c.Tick() || c == before {
			t.Fatal("moving/landing tick did not report change")
		}
	}
	if c.Flying {
		t.Fatal("shot did not land")
	}
	before := c
	if c.Tick() || c != before {
		t.Fatal("stationary tick changed state")
	}
}
