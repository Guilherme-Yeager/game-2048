package logic

func MoveUpGrid(score *int) {
	for {
		anyMovedInThisStep := false
		for c := range 4 {
			temp := [4]int{Grid[0][c], Grid[1][c], Grid[2][c], Grid[3][c]}
			if compress(&temp, score) {
				anyMovedInThisStep = true
				Grid[0][c], Grid[1][c], Grid[2][c], Grid[3][c] = temp[0], temp[1], temp[2], temp[3]
			}
		}
		if !anyMovedInThisStep {
			break
		}
		StepsGrid = append(StepsGrid, Grid)
	}
}

func MoveDownGrid(score *int) {
	for {
		anyMovedInThisStep := false
		for c := range 4 {
			temp := [4]int{Grid[3][c], Grid[2][c], Grid[1][c], Grid[0][c]}

			if compress(&temp, score) {
				anyMovedInThisStep = true
				Grid[3][c], Grid[2][c], Grid[1][c], Grid[0][c] = temp[0], temp[1], temp[2], temp[3]
			}
		}
		if !anyMovedInThisStep {
			break
		}
		StepsGrid = append(StepsGrid, Grid)
	}
}

func MoveRightGrid(score *int) {
	for {
		anyMovedInThisStep := false
		for r := range 4 {
			temp := [4]int{Grid[r][3], Grid[r][2], Grid[r][1], Grid[r][0]}

			if compress(&temp, score) {
				anyMovedInThisStep = true
				Grid[r][3], Grid[r][2], Grid[r][1], Grid[r][0] = temp[0], temp[1], temp[2], temp[3]
			}
		}
		if !anyMovedInThisStep {
			break
		}
		StepsGrid = append(StepsGrid, Grid)
	}
}

func MoveLeftGrid(score *int) {
	for {
		anyMovedInThisStep := false
		for r := range 4 {
			if compress(&Grid[r], score) {
				anyMovedInThisStep = true
			}
		}
		if !anyMovedInThisStep {
			break
		}
		StepsGrid = append(StepsGrid, Grid)
	}
}

func compress(line *[4]int, score *int) bool {
	moved := false
	var merged [4]bool
	for i := 1; i < 4; i++ {
		if line[i] == 0 {
			continue
		}
		if line[i-1] == 0 {
			line[i-1] = line[i]
			line[i] = 0
			moved = true
		} else if line[i-1] == line[i] && !merged[i-1] {
			line[i-1] *= 2
			line[i] = 0
			merged[i-1] = true
			*score += line[i-1]
			moved = true
		}
	}
	return moved
}
