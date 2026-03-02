package ui

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func GetDefaultBoardColor() rl.Color {
	return rl.NewColor(173, 173, 171, 255)
}

func GetInternalColorRec(value int) rl.Color {
	if value == 0 {
		return rl.NewColor(205, 193, 180, 255)
	}
	var factor uint8 = uint8(math.Log2(float64(value)) * 15)
	var diff int = 255 - int(factor)
	var red uint8 = uint8(255)
	var green uint8
	var blue uint8

	if diff > 255 {
		green = 255
		blue = 255
	} else {
		green = uint8(255 - factor)
		blue = uint8(255 - factor)
	}

	return rl.NewColor(red, green, blue, 255)
}
