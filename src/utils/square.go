package utils

import "fmt"

// Here the functions to manage the board : print the result and set the max perimeter of board

// CreateBoard creates a board for placing tetrominos.
func CreateBoard(size int) [][]string {
	board := make([][]string, size)
	for i := range board {
		board[i] = make([]string, size)
		for j := range board[i] {
			board[i][j] = "."
		}
	}
	return board
}

// Print prints the final square.
func Print(board [][]string) {
	for _, row := range board {
		for _, cell := range row {
			fmt.Print(cell)
		}
		fmt.Println()
	}
	fmt.Println()
}
