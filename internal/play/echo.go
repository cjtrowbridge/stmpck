package play

const (
	EchoShow = iota
	EchoInput
	EchoSuccess
	EchoFailed
	EchoWon
)

type Echo struct {
	Sequence                        []int
	RNG                             RNG
	Phase, Clock, Input, Flash, Pad int
}

func NewEcho(seed uint64) Echo {
	e := Echo{RNG: RNG(seed), Pad: -1}
	e.Sequence = []int{e.RNG.Next(4)}
	return e
}
func (e *Echo) Cue() int {
	if e.Phase == EchoShow && e.Clock/45 < len(e.Sequence) && e.Clock%45 < 30 {
		return e.Sequence[e.Clock/45]
	}
	if e.Flash > 0 {
		return e.Pad
	}
	return -1
}
func (e *Echo) Tick() bool {
	changed := e.Flash > 0
	if e.Flash > 0 {
		e.Flash--
	}
	switch e.Phase {
	case EchoShow:
		changed = true
		e.Clock++
		if e.Clock >= len(e.Sequence)*45 {
			e.Phase = EchoInput
			e.Clock = 0
		}
	case EchoSuccess:
		changed = true
		e.Clock++
		if e.Clock >= 60 {
			if len(e.Sequence) >= 32 {
				e.Phase = EchoWon
			} else {
				e.Sequence = append(e.Sequence, e.RNG.Next(4))
				e.Retry()
			}
		}
	}
	return changed
}
func (e *Echo) Press(pad int) bool {
	if e.Phase != EchoInput || pad < 0 || pad > 3 {
		return false
	}
	e.Pad = pad
	e.Flash = 12
	if e.Sequence[e.Input] != pad {
		e.Phase = EchoFailed
		return true
	}
	e.Input++
	if e.Input == len(e.Sequence) {
		e.Phase = EchoSuccess
		e.Clock = 0
	}
	return true
}
func (e *Echo) Retry() { e.Phase = EchoShow; e.Clock = 0; e.Input = 0; e.Flash = 0; e.Pad = -1 }
