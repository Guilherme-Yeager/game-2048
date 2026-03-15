package logic

import (
	"math/rand"
)

var Grid [4][4]int
var StepsGrid [][4][4]int

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

func UpdateGrid() (int32, int32) {
	if !hasCellEmpty() {
		return -1, -1
	}

	value := 2
	if rand.Intn(100) < 10 {
		value = 4
	}

	for {
		row, col := rand.Intn(4), rand.Intn(4)
		if Grid[row][col] == 0 {
			Grid[row][col] = value
			return int32(row), int32(col)
		}
	}
}

func CanMove() bool {
	if hasCellEmpty() {
		return true
	}
	for r := range 4 {
		for c := range 4 {
			if c < 3 && Grid[r][c] == Grid[r][c+1] {
				return true
			}
			if r < 3 && Grid[r][c] == Grid[r+1][c] {
				return true
			}
		}
	}
	return false
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
