//go:build debug

package utils

import (
	"machine"
	"time"
)

// 🟧 Wait for the serial port to be opened for debugging

func WaitForSerial(msg string) {
	for !machine.Serial.DTR() {
		time.Sleep(100 * time.Millisecond)
	}
	println(msg)
}
