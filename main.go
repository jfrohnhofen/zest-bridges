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
		"moverUSL":           moverUSL,
		"moverUSR":           moverUSR,
		"moverDSLL":          moverDSLL,
		"moverDSL":           moverDSL,
		"moverDSR":           moverDSR,
		"moverDSRR":          moverDSRR,

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
	house         = zest.NewDimmer(501, 502, 503, 504, 506, 507, 508, 509)

	backlights = zest.NewDimmer(45, 48)

	kitchenUS          = zest.NewDimmer(38)
	kitchenDS          = zest.NewDimmer(21, 13, 34)
	kitchenWall        = zest.NewDimmer(20)
	kitchenHangingLamp = zest.NewDimmer(28)
	porch              = zest.NewDimmer(37, 47, 19)

	moverUSL  = zest.NewAura(360)
	moverUSR  = zest.NewAura(375)
	moverDSLL = zest.NewAura(300)
	moverDSL  = zest.NewAura(315)
	moverDSR  = zest.NewAura(330)
	moverDSRR = zest.NewAura(345)

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

	video = zest.NewVideo("/mnt/sdcard/Render_09-10-22_new.led")
)

func foo() {
	workingLights.Off()
	house.Off()
	backlights.Off()

	kitchenUS.Off()
	kitchenDS.Off()
	kitchenWall.Off()
	kitchenHangingLamp.Off()
	porch.Off()

	moverUSL.Close()
	moverUSR.Close()
	moverDSLL.Close()
	moverDSL.Close()
	moverDSR.Close()
	moverDSRR.Close()

	band.Off()
	garden.Off()
	phone.Off()

	centerWashUS.Off()
	centerWashDS.Off()
	rightWashUS.Off()
	rightWashDS.Off()

	barHangingLamps.Off()
	graves.Off()
	spot.Off()
	stars.Off()

}
