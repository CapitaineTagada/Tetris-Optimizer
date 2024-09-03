package utils

// IsValid checks if all tetrominoes in the given input are valid.
func IsValid(tetrominos [][]string) bool {
	for _, tetromino := range tetrominos {
		// Validate each tetromino using the helper function isValidTetromino
		if !isValidTetromino(tetromino) {
			return false
		}
	}
	return true
}

// isValidTetromino checks if a single tetromino is valid.
func isValidTetromino(tetromino []string) bool {
	tetrominoConnections := 0
	hashCount := 0

	// Iterate over each line of the tetromino
	for indexVertical, line := range tetromino {
		// Iterate over each character in the line
		for indexHorizontal, char := range line {
			if char != '#' && char != '.' {
				return false
			} else if char == '#' {
				hashCount++

				hashConnections := 0

				// Check if there is a '#' above the current position
				if indexVertical > 0 && tetromino[indexVertical-1][indexHorizontal] == '#' {
					hashConnections++
				}
				// Check if there is a '#' below the current position
				if indexVertical < len(tetromino)-1 && tetromino[indexVertical+1][indexHorizontal] == '#' {
					hashConnections++
				}
				// Left
				if indexHorizontal > 0 && tetromino[indexVertical][indexHorizontal-1] == '#' {
					hashConnections++
				}
				// Right
				if indexHorizontal < len(line)-1 && tetromino[indexVertical][indexHorizontal+1] == '#' {
					hashConnections++
				}
				// If there are no connections for the current '#', the tetromino is invalid
				if hashConnections == 0 {
					return false
				}
				// Add the current character's connections to the total connections count
				tetrominoConnections += hashConnections
			}
		}
	}
	// Check if the tetromino has exactly 4 '#' characters and at least 6 connections
	return tetrominoConnections >= 6 && hashCount == 4
}
