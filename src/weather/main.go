package main

import (
	"machine"
	"pico-apps/lib/sensors"
	"pico-apps/lib/utils"
	"strconv"
	"time"
)

func main() {
	defer func() { utils.RecoverFromPanic(recover()) }()

	utils.WaitForSerial("😊 BME280 weather is ready")
	utils.BlinkLEDWhileAlive(machine.GPIO2, time.Millisecond*500)
	utils.BOOTSELOnButtonPress(machine.GPIO22)

	machine.I2C0.Configure(
		machine.I2CConfig{
			SDA: machine.GPIO20,
			SCL: machine.GPIO21,
		},
	)

	bme280 := sensors.NewBME280(machine.I2C0)

	for {
		println("Temperature:", strconv.FormatFloat(bme280.MustReadTemperature(), 'f', 2, 64), "°F")
		println("Pressure:", strconv.FormatFloat(bme280.MustReadPressure(), 'f', 2, 64), "inHg")
		println("Humidity:", strconv.FormatFloat(bme280.MustReadHumidity(), 'f', 2, 64), "%")

		time.Sleep(2 * time.Second)
	}
}
