package main

import (
	"bridges/pkg/zest"
)

func interactive() {
	zest.Interactive(map[string]zest.Light{
		"working":            workingLights,
		"house":              house,
		"back":               backlights,
		"kitchenUS":          kitchenUS,
		"kitchenDS":          kitchenDS,
		"kitchenWall":        kitchenWall,
		"kitchenHangingLamp": kitchenHangingLamp,
		"porch":              porch,
		"mover1":             mover1,
		"mover2":             mover2,
		"mover3":             mover3,
		"mover4":             mover4,
		"mover5":             mover5,
		"mover6":             mover6,

		"band":   band,
		"garden": garden,
		"phone":  phone,

		"centerWashUS": centerWashUS,
		"centerWashDS": centerWashDS,
		"rightWashUS":  centerWashUS,
		"rightWashDS":  centerWashDS,

		"barHangingLamps": barHangingLamps,
		"graves":          graves,
		"stars":           stars,
		"spot":            spot,
	})
}

var (
	workingLights = zest.NewDimmer(11, 46)
	house         = zest.NewDimmer(501, 502, 503, 504, 506, 507, 508, 509, 510)

	backlights = zest.NewDimmer(45, 48)

	kitchenUS          = zest.NewDimmer(38)
	kitchenDS          = zest.NewDimmer(21, 13, 34)
	kitchenWall        = zest.NewDimmer(20)
	kitchenHangingLamp = zest.NewDimmer(28)
	porch              = zest.NewDimmer(37, 47, 19)

	mover1 = zest.NewAura(300)
	mover2 = zest.NewAura(315)
	mover3 = zest.NewAura(330)
	mover4 = zest.NewAura(345)
	mover5 = zest.NewAura(360)
	mover6 = zest.NewAura(375)

	band   = zest.NewDimmer(25)
	garden = zest.NewDimmer(33, 36)
	phone  = zest.NewDimmer(15, 14)

	centerWashUS = zest.NewDimmer(41)
	centerWashDS = zest.NewDimmer(1, 12)
	rightWashUS  = zest.NewDimmer(39)
	rightWashDS  = zest.NewDimmer(2, 3)

	barHangingLamps = zest.NewDimmer(10)
	graves          = zest.NewDimmer(24)
	stars           = zest.NewStars(400)
	spot            = zest.NewDimmer(17)
	
	onAirSign = zest.NewDimmer(22)

	video = zest.NewVideo("/mnt/sdcard/Render_04-11-22.led")
)
