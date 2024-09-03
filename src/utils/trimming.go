package utils

// Here the function that trim the result

func Trimmer(tetrominos [][]string) [][]string {
	if len(tetrominos) == 0 {
		return [][]string{} // Return an empty slice if input is empty
	}

	trimmedTetrominos := make([][]string, len(tetrominos))

	for i, tetromino := range tetrominos {
		if len(tetromino) == 0 {
			continue // Skip empty tetrominos
		}

		// Find the boundaries of the tetromino
		minX, minY, maxX, maxY := len(tetromino[0]), len(tetromino), 0, 0

		for y, row := range tetromino {
			for x, cell := range row {
				if cell == '#' {
					minX = min(minX, x)
					minY = min(minY, y)
					maxX = max(maxX, x)
					maxY = max(maxY, y)
				}
			}
		}

		// Trim the tetromino
		trimmed := make([]string, maxY-minY+1)
		for y := minY; y <= maxY; y++ {
			trimmed[y-minY] = tetromino[y][minX : maxX+1]
		}

		trimmedTetrominos[i] = trimmed
	}

	return trimmedTetrominos
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
