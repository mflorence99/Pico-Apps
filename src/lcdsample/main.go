package main

import (
	"fmt"
	"machine"
	"pico-apps/lib/utils"
	"time"

	"tinygo.org/x/drivers/hd44780i2c"
)

func main() {
	utils.WaitForSerial("LCD Sample is ready!")

	machine.I2C0.Configure(
		machine.I2CConfig{
			SDA:       machine.GP0,
			SCL:       machine.GP1,
			Frequency: 400 * machine.KHz,
		},
	)

	lcd := hd44780i2c.New(machine.I2C0, 0x27)

	lcd.Configure(hd44780i2c.Config{
		Width:       16, // required
		Height:      2,  // required
		CursorOn:    true,
		CursorBlink: true,
	})

	for ix := 0; ; ix++ {
		lcd.ClearDisplay()
		lcd.Print([]byte(fmt.Sprintf(" Hello, Mark!\n  LCD Test #%d", ix%1000)))
		time.Sleep(time.Millisecond * 100)
	}
}
