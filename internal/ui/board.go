package ui

import rl "github.com/gen2brain/raylib-go/raylib"

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

func DrawBoard() {
	b.posX = (int32(rl.GetScreenWidth()) - b.width) / 2
	b.posY = (int32(rl.GetScreenHeight()) - b.height) / 2
	rl.DrawRectangle(b.posX, b.posY, b.width, b.height, GetDefaultBoardColor())
	rl.DrawRectangleLines(b.posX, b.posY, b.width, b.height, rl.Black)
}
