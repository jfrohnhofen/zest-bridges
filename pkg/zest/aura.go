package zest

type Aura struct {
	ch uint16
}

func NewAura(ch uint16) Aura {
	return Aura{ch}
}

func (aura Aura) Channels() []uint16 {
	return []uint16{aura.ch}
}

func (aura Aura) Properties() map[string]uint16 {
	return map[string]uint16{
		"shutter": 0,
		"dimmer":  1,
		"zoom":    2,
		"pan":     3,
		"tilt":    5,
		"red":     9,
		"green":   10,
		"blue":    11,
		"temp":    13,
	}
}

func (aura Aura) Prepare(shutter, zoom, pan, tilt, temp, r, g, b uint8) {
	cue.setDmx(aura.ch, shutter)
	cue.setDmx(aura.ch+2, zoom)
	cue.setDmx(aura.ch+3, pan)
	cue.setDmx(aura.ch+5, tilt)
	cue.setDmx(aura.ch+9, r)
	cue.setDmx(aura.ch+10, g)
	cue.setDmx(aura.ch+11, b)
	cue.setDmx(aura.ch+13, temp)
}

func (aura Aura) FadeTo(to uint8, dur uint32) {
	cue.animate(dur, func(dt uint32) {
		cue.setDmx(aura.ch+1, interpolate(dur, dt, cue.getDmx(aura.ch+1), to))
	})
}

func (aura Aura) On() {
	cue.setDmx(aura.ch+1, 255)
}

func (aura Aura) FadeAll(shutter, zoom, pan, tilt, temp, r, g, b uint8, dur uint32) {
	cue.animate(dur, func(dt uint32) {
		cue.setDmx(aura.ch, interpolate(dur, dt, cue.getDmx(aura.ch), shutter))
		cue.setDmx(aura.ch+2, interpolate(dur, dt, cue.getDmx(aura.ch), zoom))
		cue.setDmx(aura.ch+3, interpolate(dur, dt, cue.getDmx(aura.ch), pan))
		cue.setDmx(aura.ch+5, interpolate(dur, dt, cue.getDmx(aura.ch), tilt))
		cue.setDmx(aura.ch+9, interpolate(dur, dt, cue.getDmx(aura.ch), r))
		cue.setDmx(aura.ch+10, interpolate(dur, dt, cue.getDmx(aura.ch), g))
		cue.setDmx(aura.ch+11, interpolate(dur, dt, cue.getDmx(aura.ch), b))
		cue.setDmx(aura.ch+13, interpolate(dur, dt, cue.getDmx(aura.ch), temp))
	})
}
