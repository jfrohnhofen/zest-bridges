package zest

const brightness = 255

type Stars struct {
	ch uint16
}

func NewStars(ch uint16) Stars {
	return Stars{ch}
}

func (s Stars) Channels() []uint16 {
	return []uint16{s.ch}
}

func (s Stars) Properties() map[string]uint16 {
	return map[string]uint16{
		"dimmer":  0,
		"r":       1,
		"g":       2,
		"b":       3,
		"flicker": 4,
		"depth":   5,
	}
}

func (s Stars) FadeOn(dur uint32) {
	cue.animate(dur, func(dt uint32) {
		cue.setDmx(s.ch, interpolate(dur, dt, 0, brightness))
	})
	cue.setDmx(s.ch+1, 255)
	cue.setDmx(s.ch+2, 255)
	cue.setDmx(s.ch+3, 255)
	cue.setDmx(s.ch+4, 200)
	cue.setDmx(s.ch+5, 255)
}

func (s Stars) FadeOff(dur uint32) {
	cue.animate(dur, func(dt uint32) {
		cue.setDmx(s.ch, interpolate(dur, dt, brightness, 0))
	})
}

func (s Stars) Off() {
	cue.setDmx(s.ch, 0)
}
