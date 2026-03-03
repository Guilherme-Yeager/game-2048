package engine

import (
	"github.com/Guilherme-Yeager/game-2048/internal/logic"
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
		handleInput()

		rl.EndDrawing()
	}
}

func handleInput() {
	if rl.IsKeyPressed(rl.KeyUp) {
		logic.MoveUpGrid()
		logic.UpdateGrid()
	}
	if rl.IsKeyPressed(rl.KeyDown) {
		logic.UpdateGrid()
	}

	if rl.IsKeyPressed(rl.KeyLeft) {
		logic.MoveLeftGrid()
		logic.UpdateGrid()
	}

	if rl.IsKeyPressed(rl.KeyRight) {
		logic.UpdateGrid()
	}
}
