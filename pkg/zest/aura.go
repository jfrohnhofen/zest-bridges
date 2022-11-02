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
	}
}

func (aura Aura) Open() {
	cue.setDmx(aura.ch, 255)
}

func (aura Aura) Close() {
	cue.setDmx(aura.ch, 0)
}

func (aura Aura) Dimmer(val uint8) {
	cue.setDmx(aura.ch+1, val)
}

func (aura Aura) Zoom(val uint8) {
	cue.setDmx(aura.ch+2, val)
}

func (aura Aura) Position(pan, tilt uint8) {
	cue.setDmx(aura.ch+3, pan)
	cue.setDmx(aura.ch+5, tilt)
}

func (aura Aura) Color(r, g, b uint8) {
	cue.setDmx(aura.ch+9, r)
	cue.setDmx(aura.ch+10, g)
	cue.setDmx(aura.ch+11, b)
}
