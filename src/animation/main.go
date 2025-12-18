package main

import (
	"machine"
	"pico-apps/lib/utils"
	"time"

	"tinygo.org/x/drivers/ssd1306"
)

func show(display *ssd1306.Device, animation Animation, repeat int) {
	if repeat == 0 {
		repeat = 2
	}

	for x := 0; x < repeat; x++ {
		for i := 0; i < len(animation.Frames); i++ {
			time.Sleep(animation.Delay)
			err := display.SetBuffer(animation.Next())
			if err != nil {
				println(err)
			}
			display.Display()
		}
	}
}

func main() {
	utils.WaitForSerial()

	println("Animation is ready!")

	machine.I2C0.Configure(machine.I2CConfig{SDA: machine.GP0, SCL: machine.GP1, Frequency: 400 * machine.KHz})

	machine.I2C1.Configure(machine.I2CConfig{SDA: machine.GP2, SCL: machine.GP3, Frequency: 400 * machine.KHz})

	time.Sleep(time.Second * 1)

	d1 := ssd1306.NewI2C(machine.I2C0)
	d1.Configure(ssd1306.Config{Width: 128, Height: 64, Address: ssd1306.Address_128_32, VccState: ssd1306.SWITCHCAPVCC})
	d1.ClearDisplay()

	d2 := ssd1306.NewI2C(machine.I2C1)
	d2.Configure(ssd1306.Config{Width: 128, Height: 64, Address: ssd1306.Address_128_32, VccState: ssd1306.SWITCHCAPVCC})
	d2.ClearDisplay()

	for {

		show(d1, dragon, 1)
		show(d2, pepe, 1)
	}

}
