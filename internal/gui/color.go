package gui

import (
	"fmt"
	"image/color"
    "strings"
)

func ColorToHex(c color.Color) string {
	r, g, b, a := c.RGBA()
	return fmt.Sprintf("#%02X%02X%02X%02X", uint8(r>>8), uint8(g>>8), uint8(b>>8), uint8(a>>8))
}

func HexToColor(hex string) (color.Color, error) {
	var r, g, b, a uint8 = 0, 0, 0, 255 // default alpha = 255

	if len(hex) == 7 { // #RRGGBB
		_, err := fmt.Sscanf(hex, "#%02X%02X%02X", &r, &g, &b)
		return color.RGBA{r, g, b, a}, err
	}
	if len(hex) == 9 { // #RRGGBBAA
		_, err := fmt.Sscanf(hex, "#%02X%02X%02X%02X", &r, &g, &b, &a)
		return color.RGBA{r, g, b, a}, err
	}
	return nil, fmt.Errorf("invalid hex color: %s", hex)
}

func HexToNRGBA(hex string) (color.NRGBA, error) {
	var r, g, b, a uint8 = 0, 0, 0, 255 // default alpha = 255
	hex = strings.TrimPrefix(hex, "#")

	switch len(hex) {
	case 6:
		_, err := fmt.Sscanf(hex, "%02X%02X%02X", &r, &g, &b)
		return color.NRGBA{R: r, G: g, B: b, A: a}, err
	case 8:
		_, err := fmt.Sscanf(hex, "%02X%02X%02X%02X", &r, &g, &b, &a)
		return color.NRGBA{R: r, G: g, B: b, A: a}, err
	default:
		return color.NRGBA{}, fmt.Errorf("invalid hex color: %s", hex)
	}
}
