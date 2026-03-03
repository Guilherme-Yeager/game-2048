package logic

func MoveUpGrid() {
	for col := range 4 {
		var merged [4]bool
		for row := 1; row < 4; row++ {
			if Grid[row][col] == 0 {
				continue
			}
			var curr int = row
			for curr > 0 && Grid[curr-1][col] == 0 {
				Grid[curr-1][col] = Grid[curr][col]
				Grid[curr][col] = 0
				curr--
			}
			if curr > 0 && Grid[curr-1][col] == Grid[curr][col] && !merged[curr-1] {
				Grid[curr-1][col] *= 2
				Grid[curr][col] = 0
				merged[curr-1] = true
			}
		}
	}
}
