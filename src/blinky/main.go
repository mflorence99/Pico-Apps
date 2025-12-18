package main

import (
	"machine"
	"pico-apps/lib/utils"
	"time"
)

func main() {
	utils.WaitForSerial()

	led := machine.GP15
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	println("Blinky is ready!")

	for {
		led.Low()
		time.Sleep(time.Millisecond * 500)

		led.High()
		time.Sleep(time.Millisecond * 500)
	}
}
