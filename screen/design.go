package screen

import (
	"image/color"

	"gioui.org/text"
)

func ButtonFontMain() (Font text.Font) {
	return text.Font{Variant: "Mono", Weight: text.Bold, Style: text.Italic}
}
func ButtonBackgroundMain() (Background color.NRGBA) {
	return color.NRGBA{A: 0xff, R: 0x3C, G: 0x3C, B: 0x3C}
}
