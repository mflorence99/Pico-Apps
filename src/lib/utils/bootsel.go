package utils

import (
	"machine"
)

func BOOTSELOnButtonPress(button machine.Pin) {
	button.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	button.SetInterrupt(machine.PinRising, func(p machine.Pin) {
		machine.EnterBootloader()
	})
}
