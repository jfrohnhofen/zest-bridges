package zest

const brightness = 255

type Stars struct {
	ch uint16
}

func NewStars(ch uint16) Stars {
	return Stars{ch}
}

func (s Stars) FadeOn(dur uint32) {
	cue.animate(dur, func(dt uint32) {
		cue.setDmx(s.ch, interpolate(dur, dt, 0, brightness))
	})
	cue.setDmx(s.ch+1, 255)
	cue.setDmx(s.ch+2, 255)
	cue.setDmx(s.ch+3, 255)
	cue.setDmx(s.ch+4, 100)
	cue.setDmx(s.ch+5, 100)
}

func (s Stars) FadeOff(dur uint32) {
	cue.animate(dur, func(dt uint32) {
		cue.setDmx(s.ch, interpolate(dur, dt, brightness, 0))
	})
}
