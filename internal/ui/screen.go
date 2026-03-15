package ui

import (
	"fmt"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const ScreenWidth int32 = 800
const ScreenHeight int32 = 600
const ScreenTilte string = "2048"

var Score int

func DrawScore() {
	scoreText := fmt.Sprintf("Score: %d", Score)
	rl.DrawText(scoreText, 10, 10, 24, rl.DarkGray)
}

func DrawGameOver(score int) bool {
	rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rl.Fade(rl.Black, 0.6))

	var winX int32 = (int32(rl.GetScreenWidth()) - 350) / 2
	var winY int32 = (int32(rl.GetScreenHeight()) - 250) / 2

	rl.DrawRectangle(winX+5, winY+5, 350, 250, rl.Fade(rl.Black, 0.3))
	rl.DrawRectangle(winX, winY, 350, 250, rl.RayWhite)
	rl.DrawRectangleLines(winX, winY, 350, 250, rl.DarkGray)
	rl.DrawText("GAME OVER", winX+60, winY+40, 40, rl.Maroon)

	scoreText := "Score: " + strconv.Itoa(score)
	rl.DrawText(scoreText, winX+78, winY+110, 20, rl.Gray)

	var btnX int32 = winX + 75
	var btnY int32 = winY + 160

	mousePos := rl.GetMousePosition()
	btnRect := rl.NewRectangle(float32(btnX), float32(btnY), 200, 50)

	btnColor := rl.Orange
	isHover := rl.CheckCollisionPointRec(mousePos, btnRect)

	if isHover {
		btnColor = rl.Gold
		rl.SetMouseCursor(rl.MouseCursorPointingHand)
	} else {
		rl.SetMouseCursor(rl.MouseCursorDefault)
	}

	rl.DrawRectangleRec(btnRect, btnColor)
	rl.DrawRectangleLinesEx(btnRect, 2, rl.DarkGray)
	rl.DrawText("PLAY AGAIN", btnX+40, btnY+15, 20, rl.White)

	return isHover && rl.IsMouseButtonPressed(rl.MouseLeftButton)
}
