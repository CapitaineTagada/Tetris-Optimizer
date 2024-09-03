package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"
	"tetris-optimizer/src/utils"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go path/to/<input file>")
		return
	}

	filePath := os.Args[1]

	// Read the content of the file
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert the content into a [][]strings (representing tetrominoes)
	lines := strings.Split(string(content), "\n")
	var tetrominos [][]string
	var currentTetromino []string

	// Process the input file
	for _, line := range lines {
		if line == "" {
			// If an empty line is encountered, save the current tetromino and reset for the next one
			if len(currentTetromino) > 0 {
				tetrominos = append(tetrominos, currentTetromino)
				currentTetromino = nil
			}
			// Append line to the current tetromino
		} else {
			currentTetromino = append(currentTetromino, line)
		}
	}

	// Append the last tetromino
	if len(currentTetromino) > 0 {
		tetrominos = append(tetrominos, currentTetromino)
	}

	if !utils.IsValid(tetrominos) {
		fmt.Println("ERROR")
		return
	}
	// Set a timer to know solving time
	timeStarted := time.Now()

	// Trim excess spaces for uniformity
	trimmedTetrominos := utils.Trimmer(tetrominos)

	// Calculate the initial size of the board
	Size := int(math.Ceil(math.Sqrt(float64(len(trimmedTetrominos) * 4))))
	var finalboard [][]string
	for {
		board := utils.CreateSquare(Size)
		// Try to solve the board with the trimmed tetrominoes
		finalboard = utils.Solve(board, trimmedTetrominos)
		// If a valid solution is found, break out of the loop
		if finalboard != nil {
			break
		}
		// If no solution is found, increase the board size and try again
		Size++
	}

	utils.PrintSquare(finalboard)

	// end of timer
	fmt.Println("\nSolving time:", time.Since(timeStarted))
}
