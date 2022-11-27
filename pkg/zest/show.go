package zest

import (
	"image"
	"log"
	"net"
	"sync"
	"time"

	"github.com/marcusolsson/tui-go"
	"github.com/ziutek/ftdi"
)

const frameDuration = 40

var cue *Cue

type Frame struct {
	dmx      []uint8
	videoCmd string
}

func (f Frame) clone() Frame {
	o := Frame{}
	o.dmx = make([]uint8, 512)
	copy(o.dmx, f.dmx)
	return o
}

type Cue struct {
	name      string
	called    bool
	descr     string
	frames    []Frame
	time      uint32
	animateDt uint32
	maxTime   uint32
}

func (cue *Cue) getDmx(ch uint16) uint8 {
	frameIdx := (cue.time + frameDuration - 1) / frameDuration
	for len(cue.frames) <= int(frameIdx) {
		newFrame := cue.frames[len(cue.frames)-1].clone()
		cue.frames = append(cue.frames, newFrame)
	}
	return cue.frames[frameIdx].dmx[ch-1]
}

func (cue *Cue) setDmx(ch uint16, val uint8) {
	frameIdx := (cue.time + cue.animateDt + frameDuration) / frameDuration
	for len(cue.frames) <= int(frameIdx) {
		newFrame := cue.frames[len(cue.frames)-1].clone()
		cue.frames = append(cue.frames, newFrame)
	}
	for i := frameIdx; i < uint32(len(cue.frames)); i++ {
		cue.frames[i].dmx[ch-1] = val
	}
}

func (cue *Cue) setVideo(cmd string) {
	frameIdx := (cue.time + cue.animateDt + frameDuration) / frameDuration
	for len(cue.frames) <= int(frameIdx) {
		newFrame := cue.frames[len(cue.frames)-1].clone()
		cue.frames = append(cue.frames, newFrame)
	}
	cue.frames[frameIdx].videoCmd = cmd
}

func (cue *Cue) animate(dur uint32, fn func(dt uint32)) {
	for dt := (cue.time + frameDuration - 1) / frameDuration * frameDuration; dt < cue.time+dur+frameDuration-1; dt += frameDuration {
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
		log.Fatal("negative wait time")
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
	cues []Cue
}

func (show *Show) AddCue(name string, called bool, descr string, fn func()) {
	prevFrame := Frame{make([]uint8, 512), ""}
	if len(show.cues) > 0 {
		prevCue := show.cues[len(show.cues)-1]
		prevFrame = prevCue.frames[len(prevCue.frames)-1]
	}
	cue = &Cue{name, called, descr, []Frame{prevFrame}, 0, 0, 0}
	fn()
	show.cues = append(show.cues, *cue)
	cue = nil
}

var dmx *ftdi.Device

func (show Show) Run(first bool) {
	numCues := len(show.cues)
	selectedCue := 0
	activeCue := -1
	scrollPos := 0

	cueList := tui.NewTable(0, 0)
	for _, cue := range show.cues {
		called := tui.NewLabel("not called")
		if cue.called {
			called = tui.NewLabel("called")
		}
		cueList.AppendRow(tui.NewLabel(" "), tui.NewLabel(cue.name), called, tui.NewLabel(cue.descr))
	}
	cueList.SetColumnStretch(0, 2)
	cueList.SetColumnStretch(1, 5)
	cueList.SetColumnStretch(2, 8)
	cueList.SetColumnStretch(3, 80)
	cueList.Select(selectedCue)

	scrollArea := tui.NewScrollArea(cueList)
	scrollArea.Scroll(0, 0)

	ui, err := tui.New(scrollArea)
	if err != nil {
		log.Fatal(err)
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

	//ui.SetKeybinding("Esc", func() {
	//	ui.Quit()
	//	log.Fatal("Done")
	//})
	ui.SetKeybinding("Down", down)
	ui.SetKeybinding("Up", up)
	ui.SetKeybinding("Enter", next)
	ui.SetKeybinding(" ", next)

	if first {
		go func() {
			outputCue(show.cues[0])
		}()
	} else {
		go func() {
			if err := ui.Run(); err != nil {
				log.Fatal(err)
			}
		}()
	}
	go releaseToken()

	runDmxLoop()
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
	for i, frame := range cue.frames[1:] {
		if !stillHasToken() {
			return
		}
		time.Sleep((time.Duration(i) * frameDuration * time.Millisecond) - time.Since(start))

		mu.Lock()
		dmxFrame = frame.dmx[:]
		mu.Unlock()

		if frame.videoCmd != "" {
			conn, err := net.Dial("udp", "192.168.1.1:9000")
			if err == nil {
				conn.Write([]byte(frame.videoCmd))
			}
		}
	}

	releaseToken()
}

var (
	mu       = sync.Mutex{}
	dmxFrame = make([]byte, 512)
)

func runDmxLoop() {
	dmx, err := ftdi.OpenFirst(0x0403, 0x6001, ftdi.ChannelAny)
	if err != nil {
		log.Fatal(err)
	}

	if err := dmx.SetBaudrate(250000); err != nil {
		log.Fatal(err)
	}
	if err := dmx.SetLineProperties(ftdi.DataBits8, ftdi.StopBits2, ftdi.ParityNone); err != nil {
		log.Fatal(err)
	}
	if err := dmx.SetFlowControl(ftdi.FlowCtrlDisable); err != nil {
		log.Fatal(err)
	}

	if err := dmx.PurgeBuffers(); err != nil {
		log.Fatal(err)
	}

	if err := dmx.SetRTS(0); err != nil {
		log.Fatal(err)
	}

	for {
		dmx.SetLineProperties2(ftdi.DataBits8, ftdi.StopBits2, ftdi.ParityNone, ftdi.BreakOn)
		dmx.SetLineProperties2(ftdi.DataBits8, ftdi.StopBits2, ftdi.ParityNone, ftdi.BreakOff)

		data := make([]byte, 513)
		mu.Lock()
		copy(data[1:], dmxFrame[:])
		mu.Unlock()
		dmx.Write(data)
		time.Sleep(33 * time.Millisecond)
	}
}
