package main

import (
	"fmt"
	"machine"
	"pico-apps/lib/displays"
	"pico-apps/lib/utils"
	"time"
)

var (
	ledRed    = machine.GPIO0
	ledYellow = machine.GPIO1
	ledGreen  = machine.GPIO2
	button    = machine.GPIO22
	sda       = machine.GPIO18
	scl       = machine.GPIO19
)

func main() {
	defer func() { utils.RecoverFromPanic(recover()) }()

	utils.WaitForSerial("😊 LCD is ready!")
	utils.BlinkLEDWhileAlive(ledRed, time.Millisecond*500)
	utils.BOOTSELOnButtonPress(button)

	machine.I2C1.Configure(
		machine.I2CConfig{
			SDA:       machine.GPIO18,
			SCL:       machine.GPIO19,
			Frequency: 400 * machine.KHz,
		},
	)

	hd44780 := displays.NewHD44780(machine.I2C1, 0x27)

	for ix := 0; ; ix++ {
		hd44780.Clear()
		hd44780.Display(fmt.Sprintf(" Hello, Mark!\n  LCD Test #%d", ix%1000))
		time.Sleep(time.Millisecond * 100)
	}
}
