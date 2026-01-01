package utils

import (
	"machine"
	"time"
)

func BlinkLEDWhileAlive(led machine.Pin, rate time.Duration) {
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	go func() {
		for {
			led.Set(!led.Get())
			time.Sleep(rate)
		}
	}()
}
