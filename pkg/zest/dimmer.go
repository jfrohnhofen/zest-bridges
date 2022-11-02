package zest

type Dimmer struct {
	channels []uint16
}

func NewDimmer(channels ...uint16) Dimmer {
	return Dimmer{channels}
}

func (d Dimmer) Channels() []uint16 {
	return d.channels
}

func (d Dimmer) Properties() map[string]uint16 {
	return map[string]uint16{
		"dimmer": 0,
	}
}

func (dimmer Dimmer) On() {
	for _, ch := range dimmer.channels {
		cue.setDmx(ch, 255)
	}
}

func (dimmer Dimmer) SetTo(val uint8) {
	for _, ch := range dimmer.channels {
		cue.setDmx(ch, val)
	}
}

func (dimmer Dimmer) Off() {
	for _, ch := range dimmer.channels {
		cue.setDmx(ch, 0)
	}
}

func (dimmer Dimmer) FadeFromTo(from, to uint8, dur uint32) {
	cue.animate(dur, func(dt uint32) {
		for _, ch := range dimmer.channels {
			cue.setDmx(ch, interpolate(dur, dt, from, to))
		}
	})
}

func (dimmer Dimmer) FadeTo(to uint8, dur uint32) {
	cue.animate(dur, func(dt uint32) {
		for _, ch := range dimmer.channels {
			cue.setDmx(ch, interpolate(dur, dt, cue.getDmx(ch), to))
		}
	})
}

func (dimmer Dimmer) FadeOff(dur uint32) {
	dimmer.FadeTo(0, dur)
}

func (dimmer Dimmer) FadeOn(dur uint32) {
	dimmer.FadeTo(255, dur)
}

func interpolate(dur, dt uint32, start, end uint8) uint8 {
	if dt > dur {
		dt = dur
	}
	return start + uint8(float64(dt)/float64(dur)*(float64(end)-float64(start)))
}
