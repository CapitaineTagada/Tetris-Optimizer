package utils

import "fmt"

// create a square board of the given size with all cells filled with '.'
func CreateSquare(size int) [][]string {
	board := make([][]string, size)

	// Loop through each row of the board
	for i := range board {
		// Create a slice for each row with the same size as the board
		board[i] = make([]string, size)
		// Initialize all cells in the current row with '.'
		for j := range board[i] {
			board[i][j] = "."
		}
	}
	return board
}

// prints the board to the console
func PrintSquare(board [][]string) {
	// Iterate through each row of the board
	for _, row := range board {
		// iterate through each cell then print them
		for _, cell := range row {
			fmt.Print(cell)
		}
		fmt.Println()
	}
}
