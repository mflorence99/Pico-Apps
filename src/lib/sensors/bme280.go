package sensors

import (
	"machine"

	"tinygo.org/x/drivers/bme280"
)

type BME280 struct {
	sensor bme280.Device
}

func NewBME280(i2c *machine.I2C) *BME280 {
	b := new(BME280)
	b.sensor = bme280.New(i2c)
	b.sensor.Configure()
	connected := b.sensor.Connected()
	if !connected {
		panic("BME280 not detected")
	}
	println("👍 BME280 detected")
	return b
}

func (b *BME280) MustReadTemperature() float64 {
	temp, err := b.sensor.ReadTemperature()
	if err != nil {
		panic(err.Error())
	}
	return (float64(temp) * 1.8 / 1000) + 32
}

func (b *BME280) MustReadPressure() float64 {
	pressure, err := b.sensor.ReadPressure()
	if err != nil {
		panic(err.Error())
	}
	return float64(pressure) * 0.02953 / 100000
}

func (b *BME280) MustReadHumidity() float64 {
	humidity, err := b.sensor.ReadHumidity()
	if err != nil {
		panic(err.Error())
	}
	return float64(humidity) / 100
}

// TODO 🔥 we include this for completeness but it is meaningless
//         as the sensor code ASSUMES the pressure at sea level
//         we are basically AT sea level and it always gives a wrong value

func (b *BME280) MustReadElevation() float64 {
	elevation, err := b.sensor.ReadAltitude()
	if err != nil {
		panic(err.Error())
	}
	return float64(elevation) * 3.28084
}
