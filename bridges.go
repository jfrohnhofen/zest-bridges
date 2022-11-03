package main

import "bridges/pkg/zest"

func main() {
	show := zest.Show{}

	// BEGINNING
	show.AddCue("LX-A.0", false, "house lights on", func() {
		house.On()
		video.Play(0, 0)
	})
	show.AddCue("LX-A.1", true, "house lights to half", func() {
		house.FadeTo(128, 2000)
	})
	show.AddCue("LX-A.2", true, "house lights off", func() {
		house.FadeOff(2000)
	})

	// #1 - TO BUILD A HOME
	show.AddCue("LX-1.1", false, "", func() {
		video.Play(1, 251)
		backlights.FadeTo(255, 4000)
	})
	show.AddCue("LX-1.2", false, "", func() {
		video.Play(252, 401)
		spot.FadeOn(2000)
	})
	show.AddCue("LX-1.3", false, "", func() {
		video.Play(402, 501)
	})
	show.AddCue("LX-1.4", false, "", func() {
		video.Play(502, 512)
	})
	show.AddCue("LX-1.5", false, "", func() {
		video.Play(513, 524)
	})
	show.AddCue("LX-1.6", false, "", func() {
		video.Play(525, 536)
	})
	show.AddCue("LX-1.7", false, "", func() {
		video.Play(537, 548)
	})
	show.AddCue("LX-1.8", false, "", func() {
		video.Play(549, 560)
	})
	show.AddCue("LX-1.9", false, "", func() {
		video.Play(561, 572)
	})
	show.AddCue("LX-1.10", false, "", func() {
		video.Play(573, 701)
	})
	show.AddCue("LX-1.11", false, "", func() {
		video.Play(702, 1075)
	})
	show.AddCue("LX-1.12", false, "", func() {
		porch.FadeOn(4000)
	})
	show.AddCue("LX-1.13", false, "", func() {
		kitchen(255, 4000, true)
	})
	show.AddCue("LX-1.14", false, "", func() {
		spot.FadeOff(2000)
		backlights.FadeOff(2000)
		mover1.Prepare(255, 255, 22, 200, 0, 255, 255, 255)
		mover2.Prepare(255, 191, 21, 165, 0, 255, 255, 255)
		mover3.Prepare(255, 255, 127, 70, 0, 255, 255, 255)
		mover4.Prepare(255, 190, 46, 168, 0, 255, 255, 255)
		mover6.Prepare(255, 202, 108, 117, 0, 255, 255, 255)
	})
	show.AddCue("LX-1.15", true, "", func() {
		mover1.FadeTo(255, 3000)
		mover2.FadeTo(255, 3000)
		mover3.FadeTo(255, 3000)
		mover4.FadeTo(255, 3000)
		mover6.FadeTo(255, 3000)
	})
	show.AddCue("LX-1.16", false, "audio: on last (second / low) chord of music", func() {
		mover1.FadeTo(0, 3000)
		mover2.FadeTo(0, 3000)
		mover3.FadeTo(0, 3000)
		mover4.FadeTo(0, 3000)
		mover6.FadeTo(0, 3000)
	})

	// #2 - HOME BEFORE YOU KNOW IT
	show.AddCue("LX-2.1", true, "audio: when music starts", func() {
		garden.FadeOn(2000)
		centerWashUS.FadeOn(2000)
		centerWashDS.FadeOn(2000)
		rightWashUS.FadeOn(2000)
		rightWashDS.FadeOn(2000)
	})
	show.AddCue("LX-2.2", false, "visual: Francesca enters kitchen", func() {
		garden.FadeOff(2000)
	})
	show.AddCue("LX-2.3", false, "visual: Bud exiting", func() {
		centerWashUS.FadeOff(2000)
		centerWashDS.FadeOff(2000)
		rightWashUS.FadeOff(2000)
		rightWashDS.FadeOff(2000)

		// Marge's kitchen
		mover1.Prepare(255, 179, 23, 203, 0, 255, 255, 255)
		mover4.Prepare(255, 163, 46, 179, 0, 255, 255, 255)
		mover6.Prepare(255, 149, 148, 117, 0, 255, 255, 255)
	})
	show.AddCue("LX-2.4", true, "audio: music is done", func() {
		mover1.FadeTo(255, 2000)
		mover4.FadeTo(255, 2000)
		mover6.FadeTo(255, 2000)

		// Truck
		mover2.Prepare(255, 162, 21, 166, 0, 255, 255, 255)
		mover3.Prepare(255, 162, 127, 90, 0, 255, 255, 255)
	})
	show.AddCue("LX-2.5", true, "", func() {
		mover2.FadeTo(255, 2000)
		mover3.FadeTo(255, 2000)
	})
	show.AddCue("LX-2.6", false, "visual: hanging up phone", func() {
		mover1.FadeTo(0, 500)
		mover4.FadeTo(0, 500)
		mover6.FadeTo(0, 500)
	})

	// #3 - TEMPORARILY LOST
	show.AddCue("LX-3.1", false, "visual: Francesca sitting in truck", func() {
		kitchen(0, 3000, true)
	})
	show.AddCue("LX-3.2", false, "visual: Robert reverses and looks forward again", func() {
		video.Play(1076, 7076)
	})
	show.AddCue("LX-3.3", true, "", func() {
		video.Play(7077, 7252)
		centerWashUS.FadeOn(10000)
		centerWashDS.FadeOn(10000)
		rightWashUS.FadeOn(10000)
		rightWashDS.FadeOn(10000)
	})

	// #4 - WHAT DO YOU CALL A MAN LIKE THAT
	show.AddCue("LX-4.1", false, "visual: Francesca sits in the truck", func() {
		video.Play(7253, 10899)
		centerWashUS.FadeOff(10000)
		centerWashDS.FadeOff(10000)
		rightWashUS.FadeOff(10000)
		rightWashDS.FadeOff(10000)
	})
	show.AddCue("LX-4.2", true, "", func() {
		video.Play(10900, 11076)
		kitchen(255, 15000, true)
	})
	show.AddCue("LX-4.3", false, "visual: Robert enters kitchen", func() {
		// Truck
		mover2.FadeTo(0, 3000)
		mover3.FadeTo(0, 3000)
	})
	show.AddCue("LX-4.4", true, "", func() {
		phone.FadeOn(5000)
	})
	show.AddCue("LX-4.5", true, "", func() {
		phone.Off()
	})

	// #5 - YOU'RE NEVER ALONE
	show.AddCue("LX-5.1", true, "", func() {
		centerWashUS.FadeOn(5000)
		centerWashDS.FadeOn(5000)
		rightWashUS.FadeOn(5000)
		rightWashDS.FadeOn(5000)
		kitchen(60, 5000, true)
	})
	show.AddCue("LX-5.2", true, "", func() {
		spot.On()
		centerWashUS.FadeOff(3000)
		show.Wait()
		centerWashDS.FadeOff(3000)
		show.Wait()
		rightWashUS.FadeOff(3000)
		show.Wait()
		rightWashDS.FadeOff(3000)
	})
	show.AddCue("LX-5.3", false, "visual: Bud leaving stage", func() {
		spot.Off()
		kitchen(255, 2000, true)
		// Marge's kitchen
		mover1.Prepare(255, 171, 23, 203, 0, 255, 255, 255)
		mover4.Prepare(255, 141, 46, 179, 0, 255, 255, 255)
		mover6.Prepare(255, 87, 148, 117, 0, 255, 255, 255)
	})
	show.AddCue("LX-5.4", true, "", func() {
		mover1.FadeTo(255, 5000)
		mover4.FadeTo(255, 5000)
		mover6.FadeTo(255, 5000)
		garden.FadeOn(5000)
		kitchen(60, 5000, true)
	})
	show.AddCue("LX-5.5", true, "", func() {
		mover1.FadeTo(0, 5000)
		mover4.FadeTo(0, 5000)
		mover6.FadeTo(0, 5000)
		garden.FadeOff(5000)
		kitchen(255, 5000, true)
		mover3.Prepare(255, 208, 131, 77, 0, 255, 255, 255)
		mover6.Prepare(255, 175, 145, 117, 0, 255, 255, 255)
	})

	// #6 - ANOTHER LIFE
	show.AddCue("LX-6.1", true, "", func() {
		video.Play(11077, 20078)
		spot.On()
		mover3.FadeTo(255, 500)
		mover6.FadeTo(255, 500)
	})
	show.AddCue("LX-6.2", true, "", func() {
		kitchen(128, 10000, true)
	})
	show.AddCue("LX-6.3", true, "", func() {
		spot.FadeOff(2000)
		kitchen(255, 8000, true)
		mover3.FadeTo(0, 2000)
		mover6.FadeTo(0, 2000)
	})
	show.AddCue("LX-6.4", true, "", func() {
		phone.FadeOn(2000)
	})
	show.AddCue("LX-6.6", false, "visual: Francesca hangs up the phone", func() {
		phone.Off()
		mover2.Prepare(255, 162, 21, 166, 86, 255, 255, 255)
		mover3.Prepare(255, 162, 127, 90, 86, 255, 255, 255)
		mover4.Prepare(255, 163, 49, 169, 86, 255, 255, 255)
		mover6.Prepare(255, 77, 72, 117, 86, 255, 255, 255)
	})
	show.AddCue("LX-6.7", true, "", func() {
		mover2.FadeTo(255, 2000)
		mover3.FadeTo(255, 2000)
		mover4.FadeTo(255, 2000)
		mover6.FadeTo(255, 2000)
	})

	// #7 - WONDERING
	show.AddCue("LX-7.1", false, "visual: Robert leaves house", func() {
		stars.FadeOn(5000)
	})
	show.AddCue("LX-7.2", true, "", func() {
		kitchen(128, 5000, true)
		mover1.Prepare(255, 146, 46, 155, 86, 255, 255, 255)
	})
	show.AddCue("LX-7.3", false, "visual: Francesca leaves house", func() {
		kitchen(60, 5000, false)
		porch.FadeTo(255, 5000)
		mover1.FadeTo(190, 5000)
	})
	show.AddCue("LX-7.4", true, "", func() {
		phone.FadeOn(2000)
		mover2.FadeTo(0, 5000)
		mover3.FadeTo(0, 5000)
		mover4.FadeTo(0, 5000)
		mover6.FadeTo(0, 5000)
	})
	show.AddCue("LX-7.5", false, "visual: Francesca enters house", func() {
		kitchen(128, 0, true)
	})
	show.AddCue("LX-7.6", false, "visual: Francesca hangs up", func() {
		phone.Off()
		kitchen(0, 5000, true)
		mover1.FadeTo(0, 5000)
	})

	// #8 - LOOK AT ME
	show.AddCue("LX-8.1", true, "", func() {
		video.Play(20079, 20580)
		stars.FadeOff(5000)
		kitchen(255, 25000, true)
	})
	show.AddCue("LX-8.2", true, "", func() {
		video.Play(20581, 20957)
		kitchen(0, 10000, true)
		centerWashUS.FadeOn(10000)
		centerWashDS.FadeOn(10000)
		rightWashUS.FadeOn(10000)
		rightWashDS.FadeOn(10000)
		mover1.Prepare(255, 109, 33, 169, 190, 255, 255, 255)
		mover5.Prepare(255, 97, 78, 110, 190, 255, 255, 255)
	})
	show.AddCue("LX-8.3", true, "", func() {
		video.Play(20958, 21259)
		centerWashUS.FadeOff(1000)
		centerWashDS.FadeOff(1000)
		rightWashUS.FadeOff(1000)
		rightWashDS.FadeOff(1000)
		mover1.FadeTo(255, 1000)
		mover5.FadeTo(255, 1000)

		// Truck
		mover2.Prepare(255, 162, 21, 166, 0, 255, 255, 255)
		mover3.Prepare(255, 162, 127, 90, 0, 255, 255, 255)
	})
	show.AddCue("LX-8.4", true, "", func() {
		video.Play(21160, 21411)
		mover1.FadeTo(0, 10000)
		mover5.FadeTo(0, 10000)

		mover2.FadeTo(255, 10000)
		mover3.FadeTo(255, 10000)
		centerWashUS.FadeOn(10000)
		centerWashDS.FadeOn(10000)
		rightWashUS.FadeOn(10000)
		rightWashDS.FadeOn(10000)
	})

	// #9 - THE WORLD INSIDE A FRAME
	show.AddCue("LX-9.1", false, "visual: Francesca and Robert sitting on bench", func() {
		centerWashUS.FadeOff(4000)
		centerWashDS.FadeOff(4000)
		rightWashUS.FadeOff(4000)
		rightWashDS.FadeOff(4000)
		mover1.Prepare(255, 176, 19, 198, 0, 255, 243, 156)
		mover4.Prepare(255, 163, 48, 179, 0, 255, 243, 156)
		mover6.Prepare(255, 77, 161, 117, 0, 255, 243, 156)
	})
	show.AddCue("LX-9.2", false, "visual: Robert letting Francesca go from hug", func() {
		video.Play(21412, 21613)
		// Truck off
		mover2.FadeTo(0, 8000)
		mover3.FadeTo(0, 8000)

		mover1.FadeTo(255, 8000)
		mover4.FadeTo(255, 8000)
		mover6.FadeTo(255, 8000)
	})
	show.AddCue("LX-9.3", true, "", func() {
		video.Play(21614, 21865)
		kitchen(255, 3000, true)
		mover1.FadeTo(0, 3000)
		mover4.FadeTo(0, 3000)
		mover6.FadeTo(0, 3000)
		show.Wait()
		// bar
		mover1.Prepare(255, 211, 25, 190, 0, 255, 255, 255)
		mover2.Prepare(255, 255, 24, 202, 0, 255, 255, 255)
		mover3.Prepare(255, 162, 127, 72, 0, 255, 255, 255)
		mover4.Prepare(255, 138, 45, 175, 0, 255, 255, 255)
		mover6.Prepare(255, 77, 161, 117, 0, 255, 255, 255)
	})
	show.AddCue("LX-9.4", true, "", func() {
		phone.FadeOn(2000)
		barHangingLamps.FadeOn(10000)
		// bar
		mover1.FadeTo(255, 10000)
		mover2.FadeTo(165, 10000)
		mover3.FadeTo(255, 10000)
		mover4.FadeTo(255, 10000)
		mover6.FadeTo(117, 10000)
	})

	// #10 - SOMETHING FROM A DREAM
	show.AddCue("LX-10.1", true, "", func() {
		video.Play(21866, 22117)
		kitchen(60, 5000, true)
		// TODO: bar sign on
	})
	show.AddCue("LX-10.2", false, "", func() {
		spot.On()
	})
	show.AddCue("LX-10.3", true, "", func() {
		video.Play(22118, 22369)
		spot.Off()
		barHangingLamps.Off()
		// TODO: bar sign off
		mover1.FadeTo(0, 500)
		mover2.FadeTo(0, 500)
		mover3.FadeTo(0, 500)
		mover4.FadeTo(0, 500)
		mover6.FadeTo(0, 500)
		show.Wait()
		// Truck
		mover2.Prepare(255, 162, 21, 166, 0, 255, 255, 255)
		mover3.Prepare(255, 162, 127, 90, 0, 255, 255, 255)
	})
	show.AddCue("LX-10.4", true, "", func() {
		mover2.FadeTo(255, 2000)
		mover3.FadeTo(255, 2000)
	})
	show.AddCue("LX-10.5", false, "visual: Francesca hangs up", func() {
		mover2.FadeTo(0, 500)
		mover3.FadeTo(0, 500)
		show.Wait()

		mover1.Prepare(255, 127, 37, 173, 219, 255, 255, 255)
		mover2.Prepare(255, 204, 24, 192, 0, 255, 255, 255)
		mover3.Prepare(255, 175, 122, 82, 0, 255, 255, 255)
		mover4.Prepare(255, 189, 53, 175, 0, 255, 255, 255)
		mover6.Prepare(255, 127, 128, 107, 0, 255, 255, 255)
	})
	show.AddCue("LX-10.6", true, "start of music after voice on radio", func() {
		video.Play(22370, 29222)
		stars.FadeOn(1000)
		// TODO: on-air sign
		mover1.FadeTo(255, 1000)
		mover2.FadeTo(255, 1000)
		mover3.FadeTo(255, 1000)
		mover4.FadeTo(255, 1000)
		mover6.FadeTo(255, 1000)
	})

	// #11 - GET CLOSER / FALLING INTO YOU
	show.AddCue("LX-11.1", true, "", func() {
		kitchen(128, 3000, true)
	})
	show.AddCue("LX-11.2", false, "visual: Marge exiting", func() {
		// TODO: video?
		// TODO: on-air off
		mover2.FadeTo(0, 2000)
		mover3.FadeTo(0, 2000)
		mover4.FadeTo(0, 2000)
		mover6.FadeTo(0, 2000)
	})
	show.AddCue("LX-11.3", true, "", func() {
		video.Play(29223, 29350)
		kitchen(0, 20000, true)
		stars.FadeOff(5000)
		mover1.FadeTo(0, 5000)
	})

	// INTERMISSION
	show.AddCue("LX-B.1", true, "", func() {
		house.FadeOn(5000)
	})
	show.AddCue("LX-B.2", true, "", func() {
		workingLights.On()
	})
	show.AddCue("LX-B.3", true, "", func() {
		workingLights.Off()
	})
	show.AddCue("LX-B.4", true, "", func() {
		house.FadeTo(128, 2000)
	})
	show.AddCue("LX-B.5", true, "", func() {
		house.FadeOff(2000)
	})

	// #12 - STATE ROAD 21 / THE REAL WORLD / WHO WE ARE
	show.AddCue("LX-12.1", true, "", func() {
		band.FadeOn(1000)
		// USA
		mover1.Prepare(255, 190, 37, 22, 0, 0, 0, 255)
		mover2.Prepare(255, 190, 36, 15, 0, 255, 0, 0)
		mover3.Prepare(255, 190, 137, 235, 0, 0, 0, 255)
		mover4.Prepare(255, 190, 49, 18, 0, 255, 0, 0)
	})
	show.AddCue("LX-12.2", false, "visual: State Fair Singer comes on stage", func() {
		video.Play(29351, 29402)
		kitchen(255, 0, true)
		centerWashUS.On()
		centerWashDS.On()
		rightWashUS.On()
		rightWashDS.SetTo(100)
		// USA
		mover1.FadeTo(255, 500)
		mover2.FadeTo(255, 500)
		mover3.FadeTo(255, 500)
		mover4.FadeTo(255, 500)
	})
	show.AddCue("LX-12.3", true, "", func() {
		video.Play(29403, 29654)
		kitchen(0, 5000, true)
		centerWashUS.FadeOff(5000)
		centerWashDS.FadeOff(5000)
		rightWashUS.FadeOff(5000)
		rightWashDS.FadeOff(5000)
		// USA -> bed
		mover1.FadeAll(255, 146, 16, 195, 200, 255, 255, 255, 2000)
		mover2.FadeAll(255, 129, 20, 184, 200, 255, 255, 255, 2000)
		mover3.FadeAll(255, 134, 114, 80, 200, 255, 255, 255, 2000)
		mover4.FadeAll(255, 155, 43, 166, 200, 255, 255, 255, 2000)
	})
	show.AddCue("LX-12.4", true, "", func() {
		kitchen(0, 10000, true)
		centerWashUS.FadeOn(10000)
		centerWashDS.FadeOn(10000)
		rightWashUS.FadeOn(10000)
		rightWashDS.FadeOn(10000)
		// Photo
		mover5.Prepare(255, 19, 155, 144, 0, 255, 255, 255)
	})
	show.AddCue("LX-12.5", false, "visual: Photo", func() {
		mover5.On()
		show.WaitFor(200)
		mover5.FadeTo(0, 800)
	})
	show.AddCue("LX-12.6", true, "", func() {
		video.Play(29655, 29856)
		centerWashUS.FadeOff(15000)
		centerWashDS.FadeOff(6000)
		rightWashUS.FadeOff(3000)
		rightWashDS.FadeOff(8000)
		kitchen(0, 10000, true)

		// extend bedroom
		mover2.FadeAll(255, 129, 22, 175, 200, 255, 255, 255, 2000)
		mover3.FadeAll(255, 134, 124, 83, 200, 255, 255, 255, 2000)
	})
	show.AddCue("LX-12.7", true, "", func() {
		mover1.FadeTo(0, 5000)
		mover2.FadeTo(0, 5000)
		mover3.FadeTo(0, 5000)
		mover4.FadeTo(0, 5000)
	})
	show.AddCue("LX-12.8", false, "visual: Francesca sits down on the bench", func() {
		garden.FadeOn(4000)
	})

	// #13 - ALMOST REAL
	show.AddCue("LX-13.1", true, "", func() {
		video.Play(29857, 30108)
		backlights.FadeTo(128, 3000)
		spot.On()
		mover3.Prepare(255, 164, 119, 80, 0, 255, 255, 255)
		mover4.Prepare(255, 175, 50, 175, 0, 255, 255, 255)
	})
	show.AddCue("LX-13.2", true, "", func() {
		mover3.FadeTo(255, 3000)
		mover4.FadeTo(255, 3000)
		mover1.Prepare(255, 131, 22, 185, 0, 255, 255, 255)
		mover2.Prepare(255, 221, 36, 198, 0, 255, 255, 255)
		mover6.Prepare(255, 153, 184, 89, 0, 255, 255, 255)
	})
	show.AddCue("LX-13.3", true, "", func() {
		mover1.FadeTo(255, 3000)
		mover2.FadeTo(255, 3000)
		mover6.FadeTo(255, 3000)
		mover3.FadeTo(0, 3000)
		mover4.FadeTo(0, 3000)
		show.Wait()
		mover3.Prepare(255, 142, 116, 80, 0, 255, 255, 255)
		mover4.Prepare(255, 142, 116, 80, 0, 255, 255, 255)
	})
	show.AddCue("LX-13.4", true, "", func() {
		mover3.FadeTo(128, 2000)
		mover4.FadeTo(128, 2000)
		// TODO: "door" off
		mover2.Prepare(255, 221, 26, 200, 0, 255, 255, 255)
	})
	show.AddCue("LX-13.5", true, "", func() {
		mover3.FadeTo(255, 2000)
		mover4.FadeTo(255, 2000)
		// other door
		mover2.FadeTo(255, 2000)

	})
	show.AddCue("LX-13.6", true, "", func() {
		mover2.FadeTo(0, 1000)
		show.Wait()
		// Club
		mover2.Prepare(216, 255, 46, 176, 0, 255, 0, 0)
		mover5.Prepare(222, 200, 205, 144, 0, 0, 255, 0)
		show.WaitFor(1000)
		mover2.FadeTo(255, 1000)
		mover5.FadeTo(255, 1000)
	})
	show.AddCue("LX-13.7", true, "", func() {
		video.Play(30109, 32359)
	})
	show.AddCue("LX-13.8", true, "", func() {
		// Club -> white pools
		mover2.FadeAll(255, 255, 46, 172, 0, 255, 255, 255, 3000)
		mover5.FadeAll(255, 200, 205, 144, 0, 255, 255, 255, 3000)
	})
	show.AddCue("LX-13.9", true, "", func() {
		video.Play(32360, 33110)
	})
	show.AddCue("LX-13.10", true, "", func() {
		video.Play(33111, 33237)

		// fade off all pool (door on right is last to fade)
		mover1.FadeTo(0, 2000)
		mover2.FadeTo(0, 5000)
		mover5.FadeTo(0, 7000)
		mover6.FadeTo(0, 8000)

		// Restaurant up to full
		mover3.FadeTo(255, 2000)
		mover4.FadeTo(255, 2000)
	})
	show.AddCue("LX-13.11", false, "visual: as they walk off", func() {
		mover3.FadeTo(0, 2000)
		mover4.FadeTo(0, 2000)
	})
	show.AddCue("LX-13.12", true, "", func() {
		video.Play(33238, 33314)
		mover3.Prepare(255, 142, 116, 80, 200, 255, 255, 255)
		mover4.Prepare(255, 142, 116, 80, 200, 255, 255, 255)
		mover3.FadeTo(255, 2000)
		mover4.FadeTo(255, 2000)
	})
	show.AddCue("LX-13.13", true, "", func() {
		video.Play(33315, 33391)
		mover3.FadeTo(0, 2000)
		mover4.FadeTo(0, 2000)
		show.Wait()
		mover2.Prepare(255, 100, 130, 80, 0, 255, 255, 255)
		mover3.Prepare(255, 100, 45, 175, 0, 255, 255, 255)
	})
	show.AddCue("LX-13.14", false, "visual: Francesca and Robert appearing", func() {
		mover2.FadeTo(255, 3000)
		mover3.FadeTo(255, 3000)
	})
	show.AddCue("LX-13.15", true, "", func() {
		phone.FadeOn(2000)
	})
	show.AddCue("LX-13.16", false, "visual: Francesca hangs up the phone", func() {
		phone.FadeOff(5000)
		mover1.Prepare(255, 75, 112, 73, 0, 255, 0, 255)
		mover4.Prepare(255, 89, 147, 77, 0, 0, 0, 255)
	})

	// #14 - BEORE AND AFTER YOU / A MILLION MILES
	show.AddCue("LX-14.1", true, "", func() {
		stars.FadeOn(5000)
		mover1.FadeTo(255, 5000)
		mover2.FadeTo(182, 5000)
		mover3.FadeTo(182, 5000)
		mover4.FadeTo(255, 5000)
	})
	show.AddCue("LX-14.2", true, "", func() {
		video.Play(33392, 33518)
		show.WaitFor(5000)
		mover1.FadeTo(0, 5000)
		mover2.FadeTo(0, 5000)
		mover3.FadeTo(0, 5000)
		mover4.FadeTo(0, 5000)
		show.Wait()
		stars.FadeOff(5000)
		// Bed
		mover1.Prepare(255, 75, 112, 73, 228, 255, 255, 255)
		mover2.Prepare(255, 100, 130, 80, 228, 255, 255, 255)
		mover3.Prepare(255, 100, 45, 175, 228, 255, 255, 255)
		mover4.Prepare(255, 89, 147, 77, 228, 255, 255, 255)
	})
	show.AddCue("LX-14.3", true, "", func() {
		video.Play(33519, 33595)
		// Bed
		mover1.FadeTo(255, 3000)
		mover2.FadeTo(255, 3000)
		mover3.FadeTo(255, 3000)
		mover4.FadeTo(255, 3000)
	})
	show.AddCue("LX-14.4", true, "", func() {
		mover1.FadeTo(0, 3000)
		mover2.FadeTo(0, 3000)
		mover3.FadeTo(0, 3000)
		mover4.FadeTo(0, 3000)
		kitchen(60, 3000, true)
		spot.On()
		show.Wait()
		mover2.Prepare(255, 121, 26, 171, 0, 255, 255, 255)
		mover3.Prepare(255, 121, 137, 81, 0, 255, 255, 255)
	})
	show.AddCue("LX-14.5", true, "", func() {
		kitchen(255, 0, true)
		mover2.FadeTo(255, 500)
		mover3.FadeTo(255, 500)
		spot.Off()
		rightWashUS.FadeTo(128, 50)
		rightWashDS.FadeTo(128, 50)
	})
	show.AddCue("LX-14.6", true, "", func() {
		phone.FadeOn(2000)
		mover4.Prepare(255, 168, 132, 82, 0, 255, 255, 255)
	})
	show.AddCue("LX-14.7", true, "", func() {
		phone.Off()
		mover2.FadeTo(128, 3000)
		mover3.FadeTo(128, 3000)
		mover4.FadeTo(255, 3000)
	})
	show.AddCue("LX-14.8", true, "", func() {
		mover2.FadeTo(255, 2000)
		mover3.FadeTo(255, 2000)
		mover4.FadeTo(0, 2000)
	})
	show.AddCue("LX-14.9", false, "visual", func() {
		rightWashUS.On()
		rightWashDS.On()
		mover5.Prepare(255, 132, 103, 119, 188, 255, 255, 255)
	})
	show.AddCue("LX-14.10", true, "", func() {
		video.Play(33596, 33672)
		kitchen(0, 3000, true)
		rightWashUS.FadeOff(3000)
		rightWashDS.FadeOff(3000)
		show.Wait()
		mover5.FadeTo(255, 3000)
	})
	show.AddCue("LX-14.11", true, "", func() {
		phone.On()
		mover4.Prepare(255, 147, 120, 92, 0, 255, 127, 255)
	})

	// #15 - REWIND
	show.AddCue("LX-15.1", true, "", func() {
		// phone booth (purple)
		mover4.FadeTo(255, 3000)
		mover1.Prepare(255, 212, 109, 50, 0, 255, 127, 255)
		mover6.Prepare(255, 106, 93, 133, 0, 255, 127, 255)
	})
	show.AddCue("LX-15.2", false, "visual: Robert and Francesca hug", func() {
		video.Play(33673, 33699)
		mover1.FadeTo(255, 1000)
		mover6.FadeTo(255, 1000)
	})
	show.AddCue("LX-15.3", false, "visual: Robert freezes", func() {
		spot.On()
		phone.SetTo(60)
		mover1.FadeAll(255, 255, 109, 50, 49, 255, 255, 255, 2000)
		mover2.FadeTo(60, 2000)
		mover3.FadeTo(60, 2000)
		mover4.FadeTo(60, 2000)
		mover5.FadeTo(60, 2000)
		mover6.FadeTo(60, 2000)
	})
	show.AddCue("LX-15.4", false, "quickly!", func() {
		phone.On()
		mover1.FadeTo(0, 2000)
		mover2.FadeTo(255, 2000)
		mover3.FadeTo(255, 2000)
		mover4.FadeTo(0, 2000)
		mover5.FadeTo(255, 2000)
		mover6.FadeTo(0, 2000)
	})
	show.AddCue("LX-15.5", true, "", func() {
		phone.FadeOff(5000)
		mover3.Prepare(255, 169, 144, 76, 188, 255, 255, 255)
		mover6.Prepare(255, 164, 39, 113, 170, 255, 255, 255)
	})
	show.AddCue("LX-15.6", false, "visual: Robert exiting", func() {
		video.Play(33700, 33726)
		mover3.FadeTo(25, 3000)
		mover6.FadeTo(9, 3000)
	})
	show.AddCue("LX-15.7", true, "", func() {
		video.Play(33727, 33803)
		mover3.FadeTo(255, 3000)
		mover6.FadeTo(91, 3000)
		mover6.FadeAll(255, 164, 39, 113, 170, 255, 255, 255, 3000)
		mover4.Prepare(255, 191, 116, 104, 188, 255, 255, 255)
	})

	// #16 - WHEN I'M GONE
	show.AddCue("LX-16.1", true, "", func() {
		spot.On()
		mover3.FadeTo(0, 15000)
		mover6.FadeTo(0, 15000)
		// Heaven
		mover4.FadeTo(255, 15000)
	})
	show.AddCue("LX-16.2", true, "", func() {
		kitchen(255, 3000, true)
	})
	show.AddCue("LX-16.3", true, "", func() {
		kitchen(60, 3000, true)
	})
	show.AddCue("LX-16.4", true, "", func() {
		rightWashDS.FadeOn(3000)
		rightWashUS.FadeOn(3000)
	})
	show.AddCue("LX-16.5", true, "", func() {
		garden.FadeOn(3000)
		spot.FadeOff(2000)
		mover1.Prepare(255, 110, 115, 77, 49, 255, 255, 255)
		mover2.Prepare(255, 110, 119, 82, 49, 255, 255, 255)
		mover3.Prepare(255, 110, 123, 78, 49, 255, 255, 255)
		mover5.Prepare(200, 142, 91, 83, 188, 255, 255, 255)
		mover6.Prepare(200, 142, 0, 97, 188, 255, 255, 255)
	})
	show.AddCue("LX-16.6", true, "", func() {
		garden.FadeOff(3000)
		rightWashDS.FadeOff(3000)
		rightWashUS.FadeOff(3000)
		mover1.FadeTo(255, 3000)
		mover2.FadeTo(255, 3000)
		mover3.FadeTo(255, 3000)
		mover4.FadeAll(255, 155, 118, 104, 188, 255, 255, 255, 3000)
		mover5.FadeTo(200, 3000)
		mover6.FadeTo(200, 3000)
	})
	show.AddCue("LX-16.7", true, "", func() {
		graves.FadeOn(2000)
	})
	show.AddCue("LX-16.8", true, "", func() {
		kitchen(255, 15000, true)
	})
	show.AddCue("LX-16.9", false, "visual: Bud exiting", func() {
		graves.FadeOff(3000)
		spot.FadeOff(3000)
	})
	show.AddCue("LX-16.10", true, "", func() {
		video.Play(33804, 33930)
		kitchen(0, 5000, true)
		// TODO: ???
	})

	// #17 - IT ALL FADES AWAY
	show.AddCue("LX-17.1", false, "visual: Ginny walks off", func() {
		// TODO: ??
	})
	show.AddCue("LX-17.2", false, "", func() {
		video.Play(33931, 33957)
	})
	show.AddCue("LX-17.3", false, "", func() {
		video.Play(33958, 33984)
	})
	show.AddCue("LX-17.4", false, "", func() {
		video.Play(33985, 34011)
	})
	show.AddCue("LX-17.5", false, "", func() {
		video.Play(34012, 34038)
	})
	show.AddCue("LX-17.6", false, "", func() {
		video.Play(34039, 34065)
	})
	show.AddCue("LX-17.7", false, "", func() {
		video.Play(34066, 34092)
	})
	show.AddCue("LX-17.8", false, "", func() {
		video.Play(34093, 34119)
	})
	show.AddCue("LX-17.9", false, "", func() {
		video.Play(34120, 34146)
	})
	show.AddCue("LX-17.10", true, "", func() {
		video.Play(34147, 34173)
		spot.On()
		backlights.FadeOn(2000)
		// TODO: far down stage center very small
		// TODO fade one desk mover after 10s
		// TODO fade other desk mover after 10s
	})
	show.AddCue("LX-17.11", false, "visual: as Robert stands up walks away", func() {
		video.Play(34174, 34300)
		spot.FadeOff(5000)
	})

	// #18 - ALWAYS BETTER
	show.AddCue("LX-18.1", false, "visual: Francesca walks on and takes letter", func() {
		spot.FadeOn(1000)
	})
	show.AddCue("LX-18.2", true, "", func() {
		video.Play(34301, 34427)
	})
	show.AddCue("LX-18.3", true, "", func() {
		backlights.FadeOff(10000)
	})
	show.AddCue("LX-18.4", true, "", func() {
		video.Play(34428, 34504)
		// TODO: spot slowly fading off
	})
	show.AddCue("LX-18.5", true, "", func() {
		stars.FadeOn(3000)
	})

	// #19 - BOWS / EXIT
	show.AddCue("LX-19.1", true, "audio: strings", func() {
		video.Play(34505, 34556)
		kitchen(255, 0, true)
		centerWashDS.On()
		centerWashUS.On()
		rightWashDS.On()
		rightWashUS.On()
	})
	show.AddCue("LX-19.2", true, "after second bow", func() {
		kitchen(255, 0, true)
		centerWashDS.Off()
		centerWashUS.Off()
		rightWashDS.Off()
		rightWashUS.Off()
	})
	show.AddCue("LX-19.3", true, "second run of bows", func() {
		kitchen(255, 0, true)
		centerWashDS.On()
		centerWashUS.On()
		rightWashDS.On()
		rightWashUS.On()
	})
	show.AddCue("LX-19.4", true, "after one bow", func() {
		kitchen(255, 0, true)
		centerWashDS.Off()
		centerWashUS.Off()
		rightWashDS.Off()
		rightWashUS.Off()
	})
	show.AddCue("LX-19.5", true, "house to full", func() {
		house.FadeOn(5000)
	})
	show.AddCue("LX-19.6", true, "end of music: video off", func() {
		video.Play(34557, 34632)
	})

	// Break a leg
	show.Run(false)
}

func kitchen(val uint8, dur uint32, withPorch bool) {
	kitchenDS.FadeTo(val, dur)
	kitchenUS.FadeTo(val, dur)
	kitchenWall.FadeTo(val, dur)
	kitchenHangingLamp.FadeTo(val, dur)
	if withPorch {
		porch.FadeTo(val, dur)
	}
}
