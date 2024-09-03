package utils

func Solve(board [][]string, tetrominoes [][]string) [][]string {
	if Solvetetro(board, tetrominoes, 0, 'A') {
		return board
	}
	return nil
}

func Solvetetro(board [][]string, tetrominoes [][]string, index int, letter byte) bool {
	if index == len(tetrominoes) {
		return true
	}

	tetromino := tetrominoes[index]
	for y := range board {
		for x := range board[y] {
			if canPlace(board, tetromino, x, y) {
				placeTetromino(board, tetromino, x, y, letter)
				if Solvetetro(board, tetrominoes, index+1, letter+1) {
					return true
				}
				removeTetromino(board, tetromino, x, y)
			}
		}
	}
	return false
}

func canPlace(board [][]string, tetromino []string, x, y int) bool {
	for dy, row := range tetromino {
		for dx, char := range row {
			if char != '.' {
				if y+dy >= len(board) || x+dx >= len(board[0]) || board[y+dy][x+dx] != "." {
					return false
				}
			}
		}
	}
	return true
}

func placeTetromino(board [][]string, tetromino []string, x, y int, letter byte) {
	for dy, row := range tetromino {
		for dx, char := range row {
			if char != '.' {
				board[y+dy][x+dx] = string(letter)
			}
		}
	}
}

func removeTetromino(board [][]string, tetromino []string, x, y int) {
	for dy, row := range tetromino {
		for dx, char := range row {
			if char != '.' {
				board[y+dy][x+dx] = "."
			}
		}
	}
}
