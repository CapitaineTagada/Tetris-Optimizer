package utils

func IsValid(tetrominos [][]string) bool {
	for _, tetromino := range tetrominos {
		if !isValidTetromino(tetromino) {
			return false
		}
	}
	return true
}

func isValidTetromino(tetromino []string) bool {
	tetrominoConnections := 0
	hashtagsCount := 0
	for indexVertical, line := range tetromino {
		for indexHorizontal, char := range line {
			if char != '#' && char != '.' {
				return false
			} else if char == '#' {
				hashtagsCount++
				hashtagConnections := 0
				if indexVertical > 0 && tetromino[indexVertical-1][indexHorizontal] == '#' {
					hashtagConnections++
				}
				if indexVertical < len(tetromino)-1 && tetromino[indexVertical+1][indexHorizontal] == '#' {
					hashtagConnections++
				}
				if indexHorizontal > 0 && tetromino[indexVertical][indexHorizontal-1] == '#' {
					hashtagConnections++
				}
				if indexHorizontal < len(line)-1 && tetromino[indexVertical][indexHorizontal+1] == '#' {
					hashtagConnections++
				}
				if hashtagConnections == 0 {
					return false
				}
				tetrominoConnections += hashtagConnections
			}
		}
	}
	return tetrominoConnections >= 6 && hashtagsCount == 4
}
