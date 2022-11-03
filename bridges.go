package main

import "bridges/pkg/zest"

func main() {
	//interactive()
	bridges()
}

func bridges() {
	// Fixtures
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
		kitchen(255, 4000)
	})
	show.AddCue("LX-1.14", false, "", func() {
		spot.FadeOff(2000)
		backlights.FadeOff(2000)
	})
	show.AddCue("LX-1.15", true, "", func() {
		// TODO: fade up 2-3 pools
	})
	show.AddCue("LX-1.16", false, "audio: on last (second / low) chord of music", func() {
		// TODO: fade out 2-3 pools
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
	})
	show.AddCue("LX-2.4", true, "audio: music is done", func() {
		// TODO: mover Marge's kitchen
	})
	show.AddCue("LX-2.5", true, "", func() {
		// TODO: mover: truck
	})
	show.AddCue("LX-2.6", false, "visual: hanging up phone", func() {
		// TODO: mover Marge's kitchen, quickly
	})

	// #3 - TEMPORARILY LOST
	show.AddCue("LX-3.1", false, "visual: Francesca sitting in truck", func() {
		kitchen(0, 3000)
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
		kitchen(255, 15000)
	})
	show.AddCue("LX-4.3", false, "visual: Robert enters kitchen", func() {
		// TODO: truck mover fades out
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
		kitchen(60, 5000)
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
		kitchen(255, 2000)
	})
	show.AddCue("LX-5.4", true, "", func() {
		// TODO: mover Marge's kitchen on: bigger 2s
		garden.FadeOn(5000)
		kitchen(60, 5000)
	})
	show.AddCue("LX-5.5", true, "", func() {
		// TODO: mover Marge's kitchen off
		garden.FadeOff(5000)
		kitchen(255, 5000)
	})

	// #6 - ANOTHER LIFE
	show.AddCue("LX-6.1", true, "", func() {
		video.Play(11077, 20078)
		spot.On()
		// TODO: movers: two random dots
	})
	show.AddCue("LX-6.2", true, "", func() {
		kitchen(128, 10000)
	})
	show.AddCue("LX-6.3", true, "", func() {
		spot.FadeOff(2000)
		kitchen(255, 8000)
		// TODO: movers fade out
	})
	show.AddCue("LX-6.4", true, "", func() {
		phone.FadeOn(2000)
	})
	show.AddCue("LX-6.6", false, "visual: Francesca hangs up the phone", func() {
		phone.Off()
	})
	show.AddCue("LX-6.7", true, "", func() {
		// TODO: mover truck (same as 2.5)
	})

	// #7 - WONDERING
	show.AddCue("LX-7.1", false, "visual: Robert leaves house", func() {
		stars.FadeOn(5000)
	})
	show.AddCue("LX-7.2", true, "", func() {
		kitchen(128, 5000)
	})
	show.AddCue("LX-7.3", false, "visual: Francesca leaves house", func() {
		kitchen(60, 5000)
		// TODO: porch to full
		// TODO: mover garden

	})
	show.AddCue("LX-7.4", true, "", func() {
		// TODO: mover truck slow fade off
		phone.FadeOn(2000)
	})
	show.AddCue("LX-7.5", false, "visual: Francesca enters house", func() {
		kitchen(128, 100)
	})
	show.AddCue("LX-7.6", false, "visual: Francesca hangs up", func() {
		phone.Off()
		kitchen(0, 5000)
	})

	// #8 - LOOK AT ME
	show.AddCue("LX-8.1", true, "", func() {
		video.Play(20079, 20580)
		stars.FadeOff(5000)
		kitchen(255, 25000)
	})
	show.AddCue("LX-8.2", true, "", func() {
		video.Play(20581, 20957)
		kitchen(0, 10000)
		centerWashUS.FadeOn(10000)
		centerWashDS.FadeOn(10000)
		rightWashUS.FadeOn(10000)
		rightWashDS.FadeOn(10000)
	})
	show.AddCue("LX-8.3", true, "", func() {
		video.Play(20958, 21259)
		centerWashUS.FadeOff(1000)
		centerWashDS.FadeOff(1000)
		rightWashUS.FadeOff(1000)
		rightWashDS.FadeOff(1000)
		// TODO: movers DSL
	})
	show.AddCue("LX-8.4", true, "", func() {
		video.Play(21160, 21411)
		// TODO: movers: DSL off / truck on
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
	})
	show.AddCue("LX-9.2", false, "visual: Robert letting Francesca go from hug", func() {
		video.Play(21412, 21613)
		// TODO: mover truck off
		// TODO: new spot on
	})
	show.AddCue("LX-9.3", true, "", func() {
		video.Play(21614, 21865)
		kitchen(255, 3000)
		// TODO: mover spot off
	})
	show.AddCue("LX-9.4", true, "", func() {
		phone.FadeOn(2000)
		// TODO: bar movers start fading on
		// TODO: add bar practicals
	})

	// #10 - SOMETHING FROM A DREAM
	show.AddCue("LX-10.1", true, "", func() {
		video.Play(21866, 22117)
		kitchen(60, 5000)
		// TODO: bar sign on
	})
	show.AddCue("LX-10.2", false, "", func() {
		spot.On()
	})
	show.AddCue("LX-10.3", true, "", func() {
		video.Play(22118, 22369)
		spot.Off()
		// TODO: bar movers off
		// TODO: bar sign off
	})
	show.AddCue("LX-10.4", true, "", func() {
		// TODO: movers pool on (same as 2.5)
	})
	show.AddCue("LX-10.5", false, "visual: Francesca hangs up", func() {
		// TODO: movers pool quick off
	})
	show.AddCue("LX-10.6", true, "start of music after voice on radio", func() {
		video.Play(22370, 29222)
		stars.FadeOn(1000)
		// TODO: mover pool SRC
		// TODO: on-air sign
	})

	// #11 - GET CLOSER / FALLING INTO YOU
	show.AddCue("LX-11.1", true, "", func() {
		kitchen(128, 3000)
	})
	show.AddCue("LX-11.2", false, "visual: Marge exiting", func() {
		// TODO: video
		// TODO: mover pool SRC off 2, 3, 4, 6
		// TODO: on-air off
	})
	show.AddCue("LX-11.3", true, "", func() {
		video.Play(29223, 29350)
		kitchen(0, 20000)
		stars.FadeOff(5000)
		// TODO: last mover off 1
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
	})
	show.AddCue("LX-12.2", false, "visual: State Fair Singer comes on stage", func() {
		video.Play(29351, 29402)
		kitchen(255, 100)
		centerWashUS.On()
		centerWashDS.On()
		rightWashUS.On()
		rightWashDS.SetTo(100)
		// TODO: movers blue white red
	})
	show.AddCue("LX-12.3", true, "", func() {
		video.Play(29403, 29654)
		kitchen(0, 5000)
		centerWashUS.FadeOff(5000)
		centerWashDS.FadeOff(5000)
		rightWashUS.FadeOff(5000)
		rightWashDS.FadeOff(5000)
		// TODO: USA movers off -> bed on
	})
	show.AddCue("LX-12.4", true, "", func() {
		kitchen(0, 10000)
		centerWashUS.FadeOn(10000)
		centerWashDS.FadeOn(10000)
		rightWashUS.FadeOn(10000)
		rightWashDS.FadeOn(10000)
	})
	show.AddCue("LX-12.5", false, "visual: Photo", func() {
		// TODO: mover photo
	})
	show.AddCue("LX-12.6", true, "", func() {
		video.Play(29655, 29856)
		centerWashUS.FadeOff(15000)
		centerWashDS.FadeOff(6000)
		rightWashUS.FadeOff(3000)
		rightWashDS.FadeOff(8000)
		kitchen(0, 10000)
		// TODO: extend bedroom with movers (move 2 + 3)
	})
	show.AddCue("LX-12.7", true, "", func() {
		// TODO: movers: pools off 5s
	})
	show.AddCue("LX-12.8", false, "visual: Francesca sits down on the bench", func() {
		garden.FadeOn(4000)
	})

	// #13 - ALMOST REAL
	show.AddCue("LX-13.1", true, "", func() {
		video.Play(29857, 30108)
		backlights.FadeTo(128, 3000)
		spot.On()
	})
	show.AddCue("LX-13.2", true, "", func() {
		// TODO mover: tight spot
	})
	show.AddCue("LX-13.3", true, "", func() {
		// TODO: other pool down center-stage
		// TODO: anther mover: door stage left
		// TODO: fade off "tight spot"
	})
	show.AddCue("LX-13.4", true, "", func() {
		// TODO: anther mover: stage right (restaurant) / half-bright (3 + 4)
		// TODO: "door" off
	})
	show.AddCue("LX-13.5", true, "", func() {
		// TODO: restaurant mover to full
		// TODO: mover on other door
	})
	show.AddCue("LX-13.6", true, "", func() {
		// TODO: club USR
	})
	show.AddCue("LX-13.7", true, "", func() {
		video.Play(30109, 32359)
	})
	show.AddCue("LX-13.8", true, "", func() {
		// TODO: club -> white pool (color: white, pulse 255)
	})
	show.AddCue("LX-13.9", true, "", func() {
		video.Play(32360, 33110)
	})
	show.AddCue("LX-13.10", true, "", func() {
		video.Play(33111, 33237)
		// TODO: fade off all pool (door on right is last to fade)
		// TODO: restaurant (3 + 4) up to full
	})
	show.AddCue("LX-13.11", false, "visual: as they walk off", func() {
		// TODO: restaurant off
	})
	show.AddCue("LX-13.12", true, "", func() {
		video.Play(33238, 33314)
		// TODO: "restaurant" on (warmer)
	})
	show.AddCue("LX-13.13", true, "", func() {
		video.Play(33315, 33391)
		// TODO: restaurant off
	})
	show.AddCue("LX-13.14", false, "visual: Francesca and Robert appearing", func() {
		// TODO: fade on bed
	})
	show.AddCue("LX-13.15", true, "", func() {
		phone.FadeOn(2000)
	})
	show.AddCue("LX-13.16", false, "visual: Francesca hangs up the phone", func() {
		phone.FadeOff(5000)
	})

	// #14 - BEORE AND AFTER YOU / A MILLION MILES
	show.AddCue("LX-14.1", true, "", func() {
		stars.FadeOn(5000)
	})
	show.AddCue("LX-14.2", true, "", func() {
		video.Play(33392, 33518)
		show.Wait()
		// TODO: movers off
		show.Wait()
		stars.FadeOff(5000)
	})
	show.AddCue("LX-14.3", true, "", func() {
		video.Play(33519, 33595)
		// TODO: bed on
	})
	show.AddCue("LX-14.4", true, "", func() {
		// TODO: bed off
		kitchen(60, 3000)
		spot.On()
	})
	show.AddCue("LX-14.5", true, "", func() {
		kitchen(255, 50)
		// TODO: truck on (same as 2.5)
		spot.Off()
		rightWashUS.FadeTo(128, 50)
		rightWashDS.FadeTo(128, 50)
	})
	show.AddCue("LX-14.6", true, "", func() {
		phone.FadeOn(2000)
	})
	show.AddCue("LX-14.7", true, "", func() {
		phone.Off()
		// TODO: small pool upstage right close to door - fade in over 3s
	})
	show.AddCue("LX-14.8", true, "", func() {
		// TODO: truck pool on
	})
	show.AddCue("LX-14.9", false, "visual", func() {
		// TODO: small pool upstage right off
		rightWashUS.On()
		rightWashDS.On()
	})
	show.AddCue("LX-14.10", true, "", func() {
		video.Play(33596, 33672)
		kitchen(0, 3000)
		rightWashUS.FadeOff(3000)
		rightWashDS.FadeOff(3000)
		show.Wait()
		// TODO: pool DSL
	})
	show.AddCue("LX-14.11", true, "", func() {
		phone.On()
	})

	// #15 - REWIND
	show.AddCue("LX-15.1", true, "", func() {
		// TODO: add mover to phone booth (purple)
	})
	show.AddCue("LX-15.2", false, "visual: Robert and Francesca hug", func() {
		video.Play(33673, 33699)
		// TODO: fade ice cream and truck to half
	})
	show.AddCue("LX-15.3", false, "visual: Robert freezes", func() {
		spot.On()
		phone.SetTo(60)
		// TODO: fade ice cream and truck to 1/4
		// TODO: set spot on phone to 1/4
	})
	show.AddCue("LX-15.4", false, "quickly!", func() {
		phone.On()
		// TODO: fade ice cream and truck to full
	})
	show.AddCue("LX-15.5", true, "", func() {
		phone.FadeOff(5000)
	})
	show.AddCue("LX-15.6", false, "visual: Robert exiting", func() {
		video.Play(33700, 33726)
		// TODO: fade out ice cream + truck to 5%
		// TODO: add stage-right pool 5%
	})
	show.AddCue("LX-15.7", true, "", func() {
		video.Play(33727, 33803)
		// TODO: fade out truck
		// TODO: fade up stage-right pool
	})

	// #16 - WHEN I'M GONE
	show.AddCue("LX-16.1", true, "", func() {
		spot.On()
		// TODO: mover heaven
		// TODO: fade off stage-right pool over 15s
	})
	show.AddCue("LX-16.2", true, "", func() {
		kitchen(255, 3000)
	})
	show.AddCue("LX-16.3", true, "", func() {
		kitchen(60, 3000)
	})
	show.AddCue("LX-16.4", true, "", func() {
		rightWashDS.FadeOn(3000)
		rightWashUS.FadeOn(3000)
	})
	show.AddCue("LX-16.5", true, "", func() {
		garden.FadeOn(3000)
		spot.FadeOff(2000)
	})
	show.AddCue("LX-16.6", true, "", func() {
		garden.FadeOff(3000)
		rightWashDS.FadeOff(3000)
		rightWashUS.FadeOff(3000)
		// TODO: mover pool
	})
	show.AddCue("LX-16.7", true, "", func() {
		graves.FadeOn(2000)
	})
	show.AddCue("LX-16.8", true, "", func() {
		kitchen(255, 15000)
	})
	show.AddCue("LX-16.9", false, "visual: Bud exiting", func() {
		graves.FadeOff(3000)
		spot.FadeOff(3000)
	})
	show.AddCue("LX-16.10", true, "", func() {
		video.Play(33804, 33930)
		kitchen(0, 5000)
		// TODO: center right pool
		// TODO: single mover spot
	})

	// #17 - IT ALL FADES AWAY
	show.AddCue("LX-17.1", false, "visual: Ginny walks off", func() {
		// TODO: single mover small spot off
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
		kitchen(255, 100)
		centerWashDS.On()
		centerWashUS.On()
		rightWashDS.On()
		rightWashUS.On()
	})
	show.AddCue("LX-19.2", true, "after second bow", func() {
		kitchen(255, 100)
		centerWashDS.Off()
		centerWashUS.Off()
		rightWashDS.Off()
		rightWashUS.Off()
	})
	show.AddCue("LX-19.3", true, "second run of bows", func() {
		kitchen(255, 100)
		centerWashDS.On()
		centerWashUS.On()
		rightWashDS.On()
		rightWashUS.On()
	})
	show.AddCue("LX-19.4", true, "after one bow", func() {
		kitchen(255, 100)
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

func kitchen(val uint8, dur uint32) {
	kitchenDS.FadeTo(val, dur)
	kitchenUS.FadeTo(val, dur)
	kitchenWall.FadeTo(val, dur)
	kitchenHangingLamp.FadeTo(val, dur)
	porch.FadeTo(val, dur)
}
