package utils

import (
	"machine"
	"time"
)

func RecoverFromPanic(r any) {
	println("Recovered from panic:", r)
	time.Sleep(time.Second * 1)
	machine.EnterBootloader()
}
