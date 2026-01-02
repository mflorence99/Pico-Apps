package main

import (
	"machine"
	"pico-apps/lib/displays"
	"pico-apps/lib/sensors"
	"pico-apps/lib/utils"
	"strconv"
	"time"

	"tinygo.org/x/tinyfont/freemono"
	"tinygo.org/x/tinyfont/freesans"
)

var fontData = &freemono.Regular9pt7b
var fontTitle = &freesans.Regular9pt7b

func main() {
	defer func() { utils.RecoverFromPanic(recover()) }()

	utils.WaitForSerial("😊 BME280 weather is ready")
	utils.BlinkLEDWhileAlive(machine.GPIO2, time.Millisecond*500)
	utils.BOOTSELOnButtonPress(machine.GPIO22)

	// 👇 we'll use this for the display
	machine.I2C0.Configure(
		machine.I2CConfig{
			SDA:       machine.GPIO16,
			SCL:       machine.GPIO17,
			Frequency: 400 * machine.KHz,
		},
	)

	// 👇 and this for the sensor
	machine.I2C1.Configure(
		machine.I2CConfig{
			SDA: machine.GPIO18,
			SCL: machine.GPIO19,
		},
	)

	// 🔥 who knows why?
	time.Sleep(time.Second)

	ssd1306 := displays.NewSSD1306(machine.I2C0)
	bme280 := sensors.NewBME280(machine.I2C1)

	for {
		ssd1306.Clear()

		w, h := ssd1306.Size()
		lh := h / 4

		ssd1306.CenterText(fontTitle, 0, lh*0, w, lh, "Weather", utils.Yellow)
		temp := strconv.FormatFloat(bme280.MustReadTemperature(), 'f', 2, 64)
		ssd1306.CenterText(fontData, 0, lh*1, w, lh, temp+"F", utils.Blue)
		pressure := strconv.FormatFloat(bme280.MustReadPressure(), 'f', 2, 64)
		ssd1306.CenterText(fontData, 0, lh*2, w, lh, pressure+"\"", utils.Blue)
		humidity := strconv.FormatFloat(bme280.MustReadHumidity(), 'f', 2, 64)
		ssd1306.CenterText(fontData, 0, lh*3, w, lh, humidity+"%", utils.Blue)

		ssd1306.Display()

		time.Sleep(2 * time.Second)
	}
}
