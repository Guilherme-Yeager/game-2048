package engine

import (
	logic "github.com/Guilherme-Yeager/game-2048/internal/LOGIC"
	"github.com/Guilherme-Yeager/game-2048/internal/ui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Run() {
	rl.InitWindow(ui.ScreenWidth, ui.ScreenHeight, ui.ScreenTilte)
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	logic.InitGrid()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		ui.DrawInitiBoard(logic.Grid)

		rl.EndDrawing()
	}
}
