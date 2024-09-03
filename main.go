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

	// Convert the content into a [][]string format
	lines := strings.Split(string(content), "\n")
	var tetrominos [][]string
	var currentTetromino []string

	for _, line := range lines {
		if line == "" {
			if len(currentTetromino) > 0 {
				tetrominos = append(tetrominos, currentTetromino)
				currentTetromino = nil
			}
		} else {
			currentTetromino = append(currentTetromino, line)
		}
	}
	if len(currentTetromino) > 0 {
		tetrominos = append(tetrominos, currentTetromino)
	}

	if !utils.IsValid(tetrominos) {
		fmt.Println("ERROR")
		return
	}

	timeStarted := time.Now()
	trimmedTetrominos := utils.Trimmer(tetrominos)

	Size := int(math.Ceil(math.Sqrt(float64(len(trimmedTetrominos) * 4))))
	var finalboard [][]string
	for {
		board := utils.CreateBoard(Size)
		finalboard = utils.Solve(board, trimmedTetrominos)
		if finalboard != nil {
			break
		}
		Size++
	}

	utils.Print(finalboard)

	fmt.Println("\nSolving time:", time.Since(timeStarted))
}
