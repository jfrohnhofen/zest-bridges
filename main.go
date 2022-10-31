package main

import (
	"bridges/pkg/zest"
)

func main() {
	// Fixtures
	house := zest.NewDimmer(501, 502, 503, 504, 506, 507, 508, 509)
	backlights := zest.NewDimmer(45, 48)
	//kitchenHangingLamp := zest.NewDimmer(28)
	//barHangingLamps := zest.NewDimmer(10)
	spot := zest.NewDimmer(17)
	stars := zest.NewStars(400)

	video := zest.NewVideo("/mnt/sdcard/Render_09-10-22_new.led")

	show := zest.Show{}

	show.AddCue("testing", "", func() {
		stars.FadeOn(1000)
		show.WaitFor(4000)
		stars.FadeOff(1000)
		spot.On()
	})

	show.Run()

	return

	// BEGINNING
	show.AddCue("LX-A.0", "house lights on", func() {
		house.On()
		video.Play(0, 0)
	})
	show.AddCue("LX-A.1", "house lights to half", func() {
		house.FadeTo(128, 2000)
	})
	show.AddCue("LX-A.2", "house lights off", func() {
		house.FadeOff(2000)
	})

	// #1 - TO BUILD A HOME
	show.AddCue("LX-1.1", "", func() {
		video.Play(1, 251)
		backlights.FadeTo(255, 6000)
	})
	show.AddCue("LX-1.2", "", func() {
		video.Play(252, 401)
	})
	show.AddCue("LX-1.3", "", func() {
		video.Play(402, 501)
	})
	show.AddCue("LX-1.4", "", func() {
		video.Play(502, 512)
	})
	show.AddCue("LX-1.5", "", func() {
		video.Play(513, 524)
	})
	show.AddCue("LX-1.6", "", func() {
		video.Play(525, 536)
	})
	show.AddCue("LX-1.7", "", func() {
		video.Play(537, 548)
	})
	show.AddCue("LX-1.8", "", func() {
		video.Play(549, 560)
	})
	show.AddCue("LX-1.9", "", func() {
		video.Play(561, 572)
	})
	show.AddCue("LX-1.10", "", func() {
		video.Play(573, 701)
	})
	show.AddCue("LX-1.11", "", func() {
		video.Play(702, 1075)
	})
	show.AddCue("LX-1.12", "", func() {

	})
	show.AddCue("LX-1.13", "", func() {

	})
	show.AddCue("LX-1.14", "", func() {

	})
	show.AddCue("LX-1.15", "", func() {

	})
	show.AddCue("LX-1.16", "", func() {
	})

	// #2 - HOME BEFORE YOU KNOW IT
	show.AddCue("LX-2.1", "", func() {
	})
	show.AddCue("LX-2.2", "", func() {
	})
	show.AddCue("LX-2.3", "", func() {
	})
	show.AddCue("LX-2.4", "", func() {
	})
	show.AddCue("LX-2.5", "", func() {
	})

	// #3 - TEMPORARILY LOST
	show.AddCue("LX-3.1", "", func() {
	})
	show.AddCue("LX-3.2", "", func() {
		video.Play(1076, 7076)
	})
	show.AddCue("LX-3.3", "", func() {
		video.Play(7077, 7252)
	})
	show.AddCue("LX-3.4", "", func() {
	})

	// #4 - WHAT DO YOU CALL A MAN LIKE THAT
	show.AddCue("LX-4.1", "", func() {
		video.Play(7253, 10899)
	})
	show.AddCue("LX-4.2", "", func() {
		video.Play(10900, 11076)
	})
	show.AddCue("LX-4.3", "", func() {
	})
	show.AddCue("LX-4.4", "", func() {
	})
	show.AddCue("LX-4.5", "", func() {
	})

	// #5 - YOU'RE NEVER ALONE
	show.AddCue("LX-5.1", "", func() {
	})
	show.AddCue("LX-5.2", "", func() {
	})
	show.AddCue("LX-5.4", "", func() {
	})
	show.AddCue("LX-5.5", "", func() {
	})
	show.AddCue("LX-5.6", "", func() {
	})

	// #6 - ANOTHER LIFE
	show.AddCue("LX-6.1", "", func() {
		video.Play(11077, 20078)
	})
	show.AddCue("LX-6.2", "", func() {
	})
	show.AddCue("LX-6.4", "", func() {
	})
	show.AddCue("LX-6.5", "", func() {
	})
	show.AddCue("LX-6.6", "", func() {
	})
	show.AddCue("LX-6.7", "", func() {
	})

	// #7 - WONDERING
	show.AddCue("LX-7.1", "", func() {
	})
	show.AddCue("LX-7.2", "", func() {
	})
	show.AddCue("LX-7.3", "", func() {
	})
	show.AddCue("LX-7.4", "", func() {
	})
	show.AddCue("LX-7.5", "", func() {
	})
	show.AddCue("LX-7.6", "", func() {
	})

	// #8 - LOOK AT ME
	show.AddCue("LX-8.1", "", func() {
		video.Play(20079, 20580)
	})
	show.AddCue("LX-8.3", "", func() {
		video.Play(20581, 20957)
	})
	show.AddCue("LX-8.4", "", func() {
		video.Play(20958, 21259)
	})
	show.AddCue("LX-8.5", "", func() {
		video.Play(21160, 21411)
	})

	// #9 - THE WORLD INSIDE A FRAME
	show.AddCue("LX-9.1", "", func() {
	})
	show.AddCue("LX-9.2", "", func() {
		video.Play(21412, 21613)
	})
	show.AddCue("LX-9.3", "", func() {
		video.Play(21614, 21865)
	})
	show.AddCue("LX-9.4", "", func() {
	})

	// #10 - SOMETHING FROM A DREAM
	show.AddCue("LX-10.1", "", func() {
		video.Play(21866, 22117)
	})
	show.AddCue("LX-10.2", "", func() {
	})
	show.AddCue("LX-10.3", "", func() {
		video.Play(22118, 22369)
	})
	show.AddCue("LX-10.5", "", func() {
	})
	show.AddCue("LX-10.6", "", func() {
	})
	show.AddCue("LX-10.7", "", func() {
	})

	// #11 - GET CLOSER / FALLING INTO YOU
	show.AddCue("LX-11.1", "", func() {
		video.Play(22370, 29222)
	})
	show.AddCue("LX-11.2", "", func() {
	})
	show.AddCue("LX-11.3", "", func() {
	})
	show.AddCue("LX-11.4", "", func() {
		video.Play(29223, 29350)
	})

	// INTERMISSION
	show.AddCue("LX-B.1", "", func() {
		house.FadeOn(4000)
	})
	show.AddCue("LX-B.2", "", func() {
		house.FadeOff(4000)
	})

	// #12 - STATE ROAD 21 / THE REAL WORLD / WHO WE ARE
	show.AddCue("LX-12.1", "", func() {
		video.Play(29351, 29402)
	})
	show.AddCue("LX-12.2", "", func() {
		video.Play(29403, 29654)
	})
	show.AddCue("LX-12.3", "", func() {
	})
	show.AddCue("LX-12.4", "", func() {
		video.Play(29655, 29856)
	})
	show.AddCue("LX-12.5", "", func() {
	})
	show.AddCue("LX-12.6", "", func() {
	})

	// #13 - ALMOST REAL
	show.AddCue("LX-13.1", "", func() {
		video.Play(29857, 30108)
	})
	show.AddCue("LX-13.2", "", func() {
	})
	show.AddCue("LX-13.3", "", func() {
	})
	show.AddCue("LX-13.4", "", func() {
	})
	show.AddCue("LX-13.5", "", func() {
		video.Play(30109, 32359)
	})
	show.AddCue("LX-13.7", "", func() {
	})
	show.AddCue("LX-13.10", "", func() {
	})
	show.AddCue("LX-13.11", "", func() {
		video.Play(32360, 33110)
	})
	show.AddCue("LX-13.12", "", func() {
		video.Play(33111, 33237)
	})
	show.AddCue("LX-13.13", "", func() {
	})
	show.AddCue("LX-13.14", "", func() {
	})
	show.AddCue("LX-13.15", "", func() {
		video.Play(33238, 33314)
	})
	show.AddCue("LX-13.16", "", func() {
		video.Play(33315, 33391)
	})
	show.AddCue("LX-13.17", "", func() {
	})
	show.AddCue("LX-13.18", "", func() {
	})
	show.AddCue("LX-13.19", "", func() {
	})

	// #14 - BEORE AND AFTER YOU / A MILLION MILES
	show.AddCue("LX-14.1", "", func() {
	})
	show.AddCue("LX-14.2", "", func() {
		video.Play(33392, 33518)
	})
	show.AddCue("LX-14.3", "", func() {
	})
	show.AddCue("LX-14.4", "", func() {
		video.Play(33519, 33595)
	})
	show.AddCue("LX-14.5", "", func() {
	})
	show.AddCue("LX-14.6", "", func() {
	})
	show.AddCue("LX-14.7", "", func() {
	})
	show.AddCue("LX-14.8", "", func() {
	})
	show.AddCue("LX-14.9", "", func() {
	})
	show.AddCue("LX-14.10", "", func() {
		video.Play(33596, 33672)
	})
	show.AddCue("LX-14.11", "", func() {
	})
	show.AddCue("LX-14.12", "", func() {
	})

	// #15 - REWIND
	show.AddCue("LX-15.1", "", func() {
	})
	show.AddCue("LX-15.2", "", func() {
		video.Play(33673, 33699)
	})
	show.AddCue("LX-15.3", "", func() {
	})
	show.AddCue("LX-15.4", "", func() {
		video.Play(33700, 33726)
	})
	show.AddCue("LX-15.5", "", func() {
	})
	show.AddCue("LX-15.6", "", func() {
	})
	show.AddCue("LX-15.7", "", func() {
		video.Play(33727, 33803)
	})

	// #16 - WHEN I'M GONE
	show.AddCue("LX-16.1", "", func() {
	})
	show.AddCue("LX-16.2", "", func() {
	})
	show.AddCue("LX-16.3", "", func() {
	})
	show.AddCue("LX-16.4", "", func() {
	})
	show.AddCue("LX-16.6", "", func() {
	})
	show.AddCue("LX-16.7", "", func() {
	})
	show.AddCue("LX-16.9", "", func() {
	})
	show.AddCue("LX-16.10", "", func() {
	})
	show.AddCue("LX-16.11", "", func() {
	})
	show.AddCue("LX-16.12", "", func() {
		video.Play(33804, 33930)
	})

	// #17 - IT ALL FADES AWAY
	show.AddCue("LX-17.1", "", func() {
	})
	show.AddCue("LX-17.2", "", func() {
		video.Play(33931, 33957)
	})
	show.AddCue("LX-17.3", "", func() {
		video.Play(33958, 33984)
	})
	show.AddCue("LX-17.4", "", func() {
		video.Play(33985, 34011)
	})
	show.AddCue("LX-17.5", "", func() {
		video.Play(34012, 34038)
	})
	show.AddCue("LX-17.6", "", func() {
		video.Play(34039, 34065)
	})
	show.AddCue("LX-17.7", "", func() {
		video.Play(34066, 34092)
	})
	show.AddCue("LX-17.8", "", func() {
		video.Play(34093, 34119)
	})
	show.AddCue("LX-17.9", "", func() {
		video.Play(34120, 34146)
	})
	show.AddCue("LX-17.10", "", func() {
		video.Play(34147, 34173)
	})
	show.AddCue("LX-17.11", "", func() {
	})
	show.AddCue("LX-17.12", "", func() {
	})
	show.AddCue("LX-17.13", "", func() {
		video.Play(34174, 34300)
	})

	// #18 - ALWAYS BETTER
	show.AddCue("LX-18.1", "", func() {
	})
	show.AddCue("LX-18.2", "", func() {
		video.Play(34301, 34427)
	})
	show.AddCue("LX-18.4", "", func() {
	})
	show.AddCue("LX-18.5", "", func() {
		video.Play(34428, 34504)
	})
	show.AddCue("LX-18.6", "", func() {
	})

	// #19 - BOWS / EXIT
	show.AddCue("LX-19.1", "", func() {
		video.Play(34505, 34556)
	})
	show.AddCue("LX-19.2", "", func() {
	})
	show.AddCue("LX-19.3", "", func() {
	})
	show.AddCue("LX-19.4", "", func() {
		video.Play(34557, 34632)
	})

	// Break a leg
	show.Run()
}
