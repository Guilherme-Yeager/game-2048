package ui

import (
	"math"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type board struct {
	posX   int32
	posY   int32
	width  int32
	height int32
}

var b = board{
	width:  500,
	height: 500,
}

const internalSizeRec int32 = 120
const padding int32 = 4

var NewRow int32 = -1
var NewCol int32 = -1

func DrawBoard(grid [4][4]int) {
	b.posX = (int32(rl.GetScreenWidth()) - b.width) / 2
	b.posY = (int32(rl.GetScreenHeight()) - b.height) / 2

	rl.DrawRectangle(b.posX, b.posY, b.width, b.height, GetDefaultBoardColor())
	rl.DrawRectangleLines(b.posX, b.posY, b.width, b.height, rl.Black)

	for r := range int32(4) {
		for c := range int32(4) {
			var posX int32 = b.posX + (c * (internalSizeRec + padding)) + padding
			var posY int32 = b.posY + (r * (internalSizeRec + padding)) + padding

			rl.DrawRectangle(posX, posY, internalSizeRec, internalSizeRec, GetInternalColorRec(grid[r][c]))

			var color rl.Color
			if r == NewRow && c == NewCol {
				color = rl.Green
			} else {
				color = rl.Black
			}

			rl.DrawRectangleLines(posX, posY, internalSizeRec, internalSizeRec, color)
			drawValueBoard(grid[r][c], posX, posY)
		}
	}
}

func drawValueBoard(value int, posX int32, posY int32) {
	if value == 0 {
		return
	}
	var text string = strconv.Itoa(value)
	var fontSize int32 = int32(float64(internalSizeRec) / math.Max(1.5, float64(int32(len(text)))*0.65))
	var textWidth int32 = rl.MeasureText(text, fontSize)

	posX = posX + (internalSizeRec / 2) - (textWidth / 2)
	posY = posY + (internalSizeRec / 2) - (fontSize / 2)

	rl.DrawText(text, posX, posY, fontSize, rl.Black)
}
