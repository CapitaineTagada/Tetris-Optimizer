package utils

// Solve the board by placing all tetrominoes. It returns the solved board if successful, otherwise returns nil.
func Solve(board [][]string, tetrominoes [][]string) [][]string {
	if SolveTetris(board, tetrominoes, 0, 'A') {
		return board
	}
	return nil
}

// Tries to place tetrominoes on the board. It returns true if all tetrominoes are placed successfully.
func SolveTetris(board [][]string, tetrominoes [][]string, index int, letter byte) bool {
	// Base case: If all tetrominoes are placed, return true
	if index == len(tetrominoes) {
		return true
	}

	// Get the current tetromino to be placed
	tetromino := tetrominoes[index]

	// Attempt to place the current tetromino at every possible position on the board
	for y := range board {
		for x := range board[y] {
			// Check if the tetromino can be placed at position
			if CanPlace(board, tetromino, x, y) {
				// Place the tetromino on the board
				Placing(board, tetromino, x, y, letter)
				// Recursively try to place the next tetromino
				if SolveTetris(board, tetrominoes, index+1, letter+1) {
					return true // If successful, return true
				}
				// If placing the next tetromino fails, remove the current one and try the next position
				Remove(board, tetromino, x, y)
			}
		}
	}
	return false // false if the current tetromino cannot be placed at any position
}

// checks if a tetromino can be placed at position (x, y) on the board.
func CanPlace(board [][]string, tetromino []string, x, y int) bool {
	for dy, row := range tetromino {
		for dx, char := range row {
			// If the current cell is not empty
			if char != '.' {
				// Check if the cell goes out of the board boundaries or overlaps with an already occupied cell
				if y+dy >= len(board) || x+dx >= len(board[0]) || board[y+dy][x+dx] != "." {
					return false
				}
			}
		}
	}
	return true // if no conflict, return true
}

// places a tetromino on the board at position (x, y) using the specified letter
func Placing(board [][]string, tetromino []string, x, y int, letter byte) {
	// Iterate over each cell of the tetromino
	for dy, row := range tetromino {
		for dx, char := range row {
			// if the current cell is not empty
			if char != '.' {
				// Place the letter at the corresponding position on the board
				board[y+dy][x+dx] = string(letter)
			}
		}
	}
}

// removes a tetromino from the board at position (x, y), restoring the cells to empty ('.')
func Remove(board [][]string, tetromino []string, x, y int) {
	for dy, row := range tetromino {
		for dx, char := range row {
			if char != '.' {
				// Reset the corresponding cell on the board to empty ('.')
				board[y+dy][x+dx] = "."
			}
		}
	}
}
