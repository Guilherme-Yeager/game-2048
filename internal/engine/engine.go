package engine

import (
	"github.com/Guilherme-Yeager/game-2048/internal/logic"
	"github.com/Guilherme-Yeager/game-2048/internal/ui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var lastTime float64
var gameOver bool

const interval float32 = 0.065

func Run() {
	rl.InitWindow(ui.ScreenWidth, ui.ScreenHeight, ui.ScreenTilte)
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)
	logic.InitGrid()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		ui.DrawScore()
		handleDrawBoard()
		rl.EndDrawing()
	}
}

func handleDrawBoard() {
	timeCur := rl.GetTime()

	if len(logic.StepsGrid) > 0 {
		ui.DrawBoard(logic.StepsGrid[0])

		if timeCur-lastTime >= float64(interval) {
			logic.StepsGrid = logic.StepsGrid[1:]
			lastTime = timeCur
			if len(logic.StepsGrid) == 0 {
				ui.NewRow, ui.NewCol = logic.UpdateGrid()
				if !logic.CanMove() {
					gameOver = true
				}
			}
		}
	} else {
		ui.DrawBoard(logic.Grid)
		if !gameOver {
			handleInput()
		}
	}
	if gameOver {
		if ui.DrawGameOver(ui.Score) {
			logic.Grid = [4][4]int{}
			logic.InitGrid()
			ui.Score = 0
			ui.NewRow, ui.NewCol = -1, -1
			gameOver = false
			rl.SetMouseCursor(rl.MouseCursorDefault)
		}
	}
}

func handleInput() {
	if rl.IsKeyPressed(rl.KeyUp) {
		logic.MoveUpGrid(&ui.Score)
	}
	if rl.IsKeyPressed(rl.KeyDown) {
		logic.MoveDownGrid(&ui.Score)
	}
	if rl.IsKeyPressed(rl.KeyLeft) {
		logic.MoveLeftGrid(&ui.Score)
	}
	if rl.IsKeyPressed(rl.KeyRight) {
		logic.MoveRightGrid(&ui.Score)
	}
}
