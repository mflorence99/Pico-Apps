package displays

import (
	"image/color"
	"machine"

	"tinygo.org/x/drivers/ssd1306"
	"tinygo.org/x/tinyfont"
)

type SSD1306 struct {
	display *ssd1306.Device
}

func NewSSD1306(i2c *machine.I2C) *SSD1306 {
	d := new(SSD1306)
	d.display = ssd1306.NewI2C(i2c)
	d.display.Configure(
		ssd1306.Config{
			Width:    128,
			Height:   64,
			Address:  ssd1306.Address_128_32,
			VccState: ssd1306.SWITCHCAPVCC,
		},
	)
	d.display.Sleep(false)
	d.display.ClearBuffer()
	d.display.ClearDisplay()
	return d
}

func (d *SSD1306) CenterText(f tinyfont.Fonter, x, y, w, h int16, str string, c color.RGBA) {
	_, lw := tinyfont.LineWidth(f, str)
	lh := int16(tinyfont.GetGlyph(f, rune(str[0])).Info().Height)
	lx := max(w-int16(lw), 0) / 2
	// 👇 y coord for text is baseline
	baseLine := y + lh
	tinyfont.WriteLine(d.display, f, lx, baseLine, str, c)
}

func (d *SSD1306) Clear() {
	d.display.ClearBuffer()
}

func (d *SSD1306) Display() {
	d.display.Display()
}

func (d *SSD1306) Size() (int16, int16) {
	return d.display.Size()
}
