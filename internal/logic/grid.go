package logic

import (
	"math/rand"
)

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
			Grid[row][col] = 2
		} else {
			Grid[row][col] = 4
		}

		if positionedNumbers == 0 {
			break
		}
	}
}

func UpdateGrid() {
	if hasCellEmpty() {
		for {
			var col int = rand.Intn(4)
			var row int = rand.Intn(4)
			if Grid[row][col] != 0 {
				continue
			} else {
				Grid[row][col] = 2
				break
			}

		}
	}
}

func hasCellEmpty() bool {
	for r := range 4 {
		for c := range 4 {
			if Grid[r][c] == 0 {
				return true
			}
		}
	}
	return false
}
