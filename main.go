package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"tetris-optimizer/src/utils"
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

	// Split the content into individual tetrominoes
	tetrominoStrings := strings.Split(strings.TrimSpace(string(content)), "\n\n")
	var tetrominoes []utils.Tetromino

	for i, tetrominoString := range tetrominoStrings {
		lines := strings.Split(strings.TrimSpace(tetrominoString), "\n")
		if !utils.IsValid(lines) {
			fmt.Println("ERROR")
			return
		}
		tetrominoes = append(tetrominoes, utils.Tetromino{
			Shape:  lines,
			Letter: string(rune('A' + i)),
		})
	}

	solution := utils.SolveTetris(tetrominoes)

	// Print the solution
	for _, row := range solution {
		fmt.Println(strings.Join(row, ""))
	}
}
