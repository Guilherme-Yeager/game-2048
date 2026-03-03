package logic

func MoveUpGrid() {
	for c := range 4 {
		var temp [4]int = [4]int{Grid[0][c], Grid[1][c], Grid[2][c], Grid[3][c]}
		compress(&temp)
		Grid[0][c], Grid[1][c], Grid[2][c], Grid[3][c] = temp[0], temp[1], temp[2], temp[3]
	}

}

func MoveLeftGrid() {
	for r := range 4 {
		compress(&Grid[r])
	}
}

func compress(line *[4]int) {
	var merged [4]bool
	for i := 1; i < 4; i++ {
		if line[i] == 0 {
			continue
		}
		var curr int = i
		for curr > 0 && line[curr-1] == 0 {
			line[curr-1] = line[curr]
			line[curr] = 0
			curr--
		}
		if curr > 0 && line[curr-1] == line[curr] && !merged[curr-1] {
			line[curr-1] *= 2
			line[curr] = 0
			merged[curr-1] = true
		}
	}
}
