package utils

// removes empty rows and columns around each tetromino,
func Trimmer(tetrominos [][]string) [][]string {
	if len(tetrominos) == 0 {
		return [][]string{}
	}

	// Prepare a slice to store trimmed tetrominoes
	trimmedTetrominos := make([][]string, len(tetrominos))

	// Iterate over each tetromino to trim it
	for i, tetromino := range tetrominos {
		if len(tetromino) == 0 {
			continue
		}

		// Initialize boundaries for trimming (minimum and maximum x and y coordinates)
		minX, minY, maxX, maxY := len(tetromino[0]), len(tetromino), 0, 0

		// Find the boundaries of the '#' characters
		for y, row := range tetromino {
			for x, cell := range row {
				if cell == '#' {
					// Update the minimum and maximum coordinates for trimming
					minX = min(minX, x)
					minY = min(minY, y)
					maxX = max(maxX, x)
					maxY = max(maxY, y)
				}
			}
		}

		// Create a new trimmed tetromino with only the necessary rows and columns
		trimmed := make([]string, maxY-minY+1)
		for y := minY; y <= maxY; y++ {
			// Extract the part of the row that contains the '#' characters
			trimmed[y-minY] = tetromino[y][minX : maxX+1]
		}

		// Store the trimmed tetromino
		trimmedTetrominos[i] = trimmed
	}

	return trimmedTetrominos
}

// min returns the smaller of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// max returns the larger of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
