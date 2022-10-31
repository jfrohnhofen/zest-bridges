package main

import (
	"bridges/pkg/zest"
	"fmt"
)

func main() {
	// Fixtures
	house := zest.NewDimmer(1, 2, 3)

	show := zest.NewShow("/tty/usbserial", "192.168.1.1:9000")

	// Cues
	show.AddCue("LX-A.0", "house lights on", func() {
		house.On()
	})
	show.AddCue("LX-A.1", "house lights to half", func() {
		house.FadeTo(128, 5000)
	})
	show.AddCue("LX-A.2", "house lights off", func() {
		house.Off()
	})

	for i := 0; i < 100; i++ {
		show.AddCue(fmt.Sprintf("LX-B.%d", i), "", func() {})
	}

	// Break a leg
	show.Run()
}
