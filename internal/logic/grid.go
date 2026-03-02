package logic

import "math/rand"

var Grid [4][4]int

func InitGrid() {
	Grid = [4][4]int{}
	var positionedNumbers int8 = 2
	for {
		var col int = rand.Intn(4)
		var row int = rand.Intn(4)
		if Grid[row][col] != 0 {
			continue
		} else {
			positionedNumbers--
		}

		var randomNumber int = rand.Intn(100)
		if randomNumber < 90 {
			Grid[col][row] = 2
		} else {
			Grid[col][row] = 4
		}

		if positionedNumbers == 0 {
			break
		}
	}

}
