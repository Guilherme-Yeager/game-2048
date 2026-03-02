package engine

import (
	"github.com/Guilherme-Yeager/game-2048/internal/ui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Run() {
	rl.InitWindow(ui.ScreenWidth, ui.ScreenHeight, ui.ScreenTilte)
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		ui.DrawBoard()

		rl.EndDrawing()
	}
}
