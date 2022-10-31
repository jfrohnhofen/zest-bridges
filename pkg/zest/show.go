package zest

import (
	"image"
	"time"

	"github.com/marcusolsson/tui-go"
)

const frameDuration = 1000

var cue *Cue

type Frame struct {
	dmx [10]uint8
}

type Cue struct {
	name      string
	descr     string
	frames    []Frame
	time      uint32
	animateDt uint32
	maxTime   uint32
}

func (cue *Cue) getDmx(ch uint16) uint8 {
	frameIdx := (cue.time + frameDuration - 1) / frameDuration
	for len(cue.frames) <= int(frameIdx) {
		cue.frames = append(cue.frames, cue.frames[len(cue.frames)-1])
	}
	return cue.frames[frameIdx].dmx[ch]
}

func (cue *Cue) setDmx(ch uint16, val uint8) {
	frameIdx := (cue.time + cue.animateDt + frameDuration) / frameDuration
	for len(cue.frames) <= int(frameIdx) {
		cue.frames = append(cue.frames, cue.frames[len(cue.frames)-1])
	}
	cue.frames[frameIdx].dmx[ch] = val
}

func (cue *Cue) animate(dur uint32, fn func(dt uint32)) {
	for dt := (cue.time + frameDuration - 1) / frameDuration; dt < cue.time+dur+frameDuration-1; dt += frameDuration {
		cue.animateDt = dt - cue.time
		fn(cue.animateDt)
	}

	if cue.time+dur > cue.maxTime {
		cue.maxTime = cue.time + dur
	}
	cue.animateDt = 0
}

func (Show) WaitFor(dur uint32) {
	cue.time += dur
	if cue.time > cue.maxTime {
		cue.maxTime = cue.time
	}
}

func (Show) WaitUntil(time uint32) {
	if cue.time > time {
		panic("negative wait time")
	}
	cue.time = time
	if cue.time > cue.maxTime {
		cue.maxTime = cue.time
	}
}

func (Show) Wait() {
	cue.time = cue.maxTime
}

type Show struct {
	dmxDevice   string
	videoDevice string
	cues        []Cue
}

func NewShow(dmxDevice, videoDevice string) Show {
	return Show{dmxDevice, videoDevice, nil}
}

func (show *Show) AddCue(name string, descr string, fn func()) {
	prevFrame := Frame{}
	if len(show.cues) > 0 {
		prevCue := show.cues[len(show.cues)-1]
		prevFrame = prevCue.frames[len(prevCue.frames)-1]
	}
	cue = &Cue{name, descr, []Frame{prevFrame}, 0, 0, 0}
	fn()
	show.cues = append(show.cues, *cue)
	cue = nil
}

func (show Show) Run() {
	numCues := len(show.cues)
	selectedCue := 0
	activeCue := -1
	scrollPos := 0

	cueList := tui.NewTable(0, 0)
	for _, cue := range show.cues {
		cueList.AppendRow(tui.NewLabel(" "), tui.NewLabel(cue.name), tui.NewLabel(cue.descr))
	}
	cueList.Select(selectedCue)

	scrollArea := tui.NewScrollArea(cueList)
	scrollArea.Scroll(0, 0)

	ui, err := tui.New(scrollArea)
	if err != nil {
		panic(err)
	}

	down := func() {
		if selectedCue < numCues-1 {
			selectedCue++
			cueList.Select(selectedCue)
		}
		if scrollPos+scrollArea.Size().Y-selectedCue < 5 {
			oldScrollPos := scrollPos
			scrollPos += 10
			if scrollPos > numCues-scrollArea.Size().Y+1 {
				scrollPos = numCues - scrollArea.Size().Y + 1
			}
			scrollArea.Scroll(0, scrollPos-oldScrollPos)
		}
	}

	up := func() {
		if selectedCue > 0 {
			selectedCue--
			cueList.Select(selectedCue)
		}
		if selectedCue-scrollPos < 5 {
			oldScrollPos := scrollPos
			scrollPos -= 10
			if scrollPos < 0 {
				scrollPos = 0
			}
			scrollArea.Scroll(0, scrollPos-oldScrollPos)
		}
	}

	next := func() {
		if activeCue > -1 {
			cueList.Grid.SetCell(image.Point{0, activeCue}, tui.NewLabel(""))
		}
		activeCue = selectedCue
		go outputCue(show.cues[activeCue])
		cueList.Grid.SetCell(image.Point{0, activeCue}, tui.NewLabel("  *"))
		down()
	}

	ui.SetKeybinding("Esc", func() { ui.Quit() })
	ui.SetKeybinding("Down", down)
	ui.SetKeybinding("Up", up)
	ui.SetKeybinding("Enter", next)
	ui.SetKeybinding(" ", next)

	go releaseToken()

	if err := ui.Run(); err != nil {
		panic(err)
	}
}

var token = make(chan bool)

func claimToken() {
	<-token
}

func stillHasToken() bool {
	select {
	case token <- true:
		return false
	default:
		return true
	}
}

func releaseToken() {
	token <- true
}

func outputCue(cue Cue) {
	claimToken()

	start := time.Now()
	for i := range cue.frames[1:] {
		if !stillHasToken() {
			break
		}
		time.Sleep((time.Duration(i) * frameDuration * time.Millisecond) - time.Since(start))
	}

	releaseToken()
}