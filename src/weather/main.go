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

var (
	fontData  = &freemono.Regular9pt7b
	fontTitle = &freesans.Regular9pt7b
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

	for _, led := range []machine.Pin{ledRed, ledYellow, ledGreen} {
		led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	}

	utils.WaitForSerial("😊 BME280 weather is ready")
	utils.BlinkLEDWhileAlive(ledGreen, time.Millisecond*500)
	utils.BOOTSELOnButtonPress(button)

	machine.I2C1.Configure(
		machine.I2CConfig{
			SDA:       sda,
			SCL:       scl,
			Frequency: 400 * machine.KHz,
		},
	)

	// 🔥 who knows why?
	time.Sleep(time.Second)

	hd44780 := displays.NewHD44780(machine.I2C1, 0x27)
	ssd1306 := displays.NewSSD1306(machine.I2C1)
	bme280 := sensors.NewBME280(machine.I2C1)

	for {
		ssd1306.Clear()

		w, h := ssd1306.Size()
		lh := h / 4

		ssd1306.CenterText(fontTitle, 0, lh*0, w, lh, "Weather", utils.Yellow)

		raw := bme280.MustReadTemperature()
		ledRed.Set(raw > 78.5) // 👈 just a hack test
		temperature := strconv.FormatFloat(raw, 'f', 2, 64)
		ssd1306.CenterText(fontData, 0, lh*1, w, lh, temperature+"F", utils.Blue)

		raw = bme280.MustReadPressure()
		pressure := strconv.FormatFloat(raw, 'f', 2, 64)
		ssd1306.CenterText(fontData, 0, lh*2, w, lh, pressure+"\"", utils.Blue)

		raw = bme280.MustReadHumidity()
		humidity := strconv.FormatFloat(raw, 'f', 2, 64)
		ssd1306.CenterText(fontData, 0, lh*3, w, lh, humidity+"%", utils.Blue)

		ssd1306.Display()

		hd44780.Clear()
		hd44780.Display("Temp...." + temperature + "F\n" + pressure + "\"   " + humidity + "%")

		time.Sleep(2 * time.Second)
	}
}
