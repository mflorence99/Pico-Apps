package displays

import (
	"machine"

	"tinygo.org/x/drivers/hd44780i2c"
)

type HD44780 struct {
	display hd44780i2c.Device
}

func NewHD44780(i2c *machine.I2C, addr uint8) *HD44780 {
	d := new(HD44780)
	d.display = hd44780i2c.New(i2c, addr)
	d.display.Configure(
		hd44780i2c.Config{
			Width:       16,
			Height:      2,
			CursorOn:    false,
			CursorBlink: false,
		},
	)
	d.display.ClearDisplay()
	return d
}

func (d *HD44780) Clear() {
	d.display.ClearDisplay()
}

func (d *HD44780) Display(str string) {
	d.display.Print([]byte(str))
}
