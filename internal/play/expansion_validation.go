package play

import (
	"errors"
	"math"
)

func (c *Collection) validateExpansion() error {
	bad := func() error { return errors.New("invalid expanded game save") }
	bounded := func(v, lo, hi float64) bool { return !math.IsNaN(v) && !math.IsInf(v, 0) && v >= lo && v <= hi }
	if c.Ready[DriftID] {
		d := &c.Drift
		if d.Level < 1 || d.Level > 3 || d.Lives < 0 || d.Lives > 3 || d.Clock < 0 || d.Clock >= 18 || d.Shield < 0 || d.Shield > 120 || d.Hits < 0 || len(d.Rocks) > 20 || len(d.Shots) > 6 || !bounded(d.X, 0, 1) || !bounded(d.Y, 0, 1) || !bounded(d.VX, -.36, .36) || !bounded(d.VY, -.36, .36) || !bounded(d.Angle, -math.Pi, math.Pi) {
			return bad()
		}
		for _, r := range d.Rocks {
			if r.Size < 0 || r.Size > 2 || !bounded(r.X, 0, 1) || !bounded(r.Y, 0, 1) || !bounded(r.VX, -1, 1) || !bounded(r.VY, -1, 1) {
				return bad()
			}
		}
		for _, s := range d.Shots {
			if s.Life < 1 || s.Life > 75 || !bounded(s.X, 0, 1) || !bounded(s.Y, 0, 1) || !bounded(s.VX, -1, 1) || !bounded(s.VY, -1, 1) {
				return bad()
			}
		}
	}
	if c.Ready[StackID] {
		s := &c.Stack
		if s.Piece < 0 || s.Piece > 6 || s.Next < 0 || s.Next > 6 || s.Rotation < 0 || s.Rotation > 3 || s.X < -3 || s.X >= StackW || s.Y < 0 || s.Y >= StackH || s.Pace < 0 || s.Pace > 2 || s.Clock < 0 || s.Clock >= 90 || s.Lines < 0 || s.Lines > 43 || s.Score < 0 || s.BagIndex < 0 || s.BagIndex > 7 {
			return bad()
		}
		seen := [7]bool{}
		for _, v := range s.Bag {
			if v < 0 || v > 6 || seen[v] {
				return bad()
			}
			seen[v] = true
		}
		for _, v := range s.Cells {
			if v < 0 || v > 7 {
				return bad()
			}
		}
		if !s.Over && !s.Won && !s.Fits(s.X, s.Y, s.Rotation) {
			return bad()
		}
	}
	if c.Ready[DeductionID] {
		d := &c.Deduction
		for i, v := range d.Solution {
			if v < 1 || v > 9 || !sudokuAllowed(&d.Solution, i, v) || d.Clues[i] < 0 || d.Clues[i] > 9 || d.Clues[i] != 0 && d.Clues[i] != v || d.Cells[i] < 0 || d.Cells[i] > 9 || d.Clues[i] != 0 && d.Cells[i] != d.Clues[i] || d.Previous[i] < 0 || d.Previous[i] > 9 || d.CanUndo && d.Clues[i] != 0 && d.Previous[i] != d.Clues[i] {
				return bad()
			}
		}
		d.Won = d.Cells == d.Solution
	}
	if c.Ready[ClusterID] {
		v := &c.Cluster
		if v.Color < 1 || v.Color > 4 || v.Next < 1 || v.Next > 4 || v.Score < 0 || v.Shots < 0 || !bounded(v.Aim, -1.1, 1.1) || !bounded(v.X, 0, 1) || !bounded(v.Y, 0, 1) || !bounded(v.VX, -1, 1) || !bounded(v.VY, -1, 1) {
			return bad()
		}
		for _, b := range v.Cells {
			if b < 0 || b > 4 {
				return bad()
			}
		}
	}
	if c.Ready[WaveID] {
		w := &c.Wave
		if w.Level < 1 || w.Level > 3 || w.Lives < 0 || w.Lives > 3 || w.Clock < 0 || w.Clock >= 180 || w.Shield < 0 || w.Shield > 120 || w.Hits < 0 || w.Direction != 1 && w.Direction != -1 || !bounded(w.X, .035, .30) || !bounded(w.Y, .1, 1) || !bounded(w.Player, .04, .96) || len(w.Shots) > 30 {
			return bad()
		}
		for _, s := range w.Shots {
			if !bounded(s.X, 0, 1) || !bounded(s.Y, 0, 1) {
				return bad()
			}
		}
	}
	if c.Ready[CrossingID] {
		v := &c.Crossing
		if v.Level < 1 || v.Level > 3 || v.Lives < 0 || v.Lives > 3 || v.Y < 0 || v.Y > 10 || !bounded(v.X, -1, 10) || v.Shield < 0 || v.Shield > 45 {
			return bad()
		}
		for _, p := range v.Phases {
			if !bounded(p, 0, 3) {
				return bad()
			}
		}
	}
	if c.Ready[ShiftID] {
		s := &c.Shift
		boards := [][16]int{s.Cells, s.Initial}
		if s.CanUndo {
			boards = append(boards, s.Previous)
		}
		for _, board := range boards {
			seen := [16]bool{}
			for _, v := range board {
				if v < 0 || v > 15 || seen[v] {
					return bad()
				}
				seen[v] = true
			}
		}
		if s.Moves < 0 || s.PreviousMoves < 0 {
			return bad()
		}
		s.Won = s.Solved()
	}
	return nil
}
