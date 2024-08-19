package main

import (
	"fmt"
	"os"
	"tetris-optimizer/src/utils" // Import the package used in the main function
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./tetromino_solver <file_path>")
		return
	}

	tetrominoes, err := utils.ParseFile(os.Args[1])
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	solution := utils.FindSmallestSquare(tetrominoes)
	if solution != nil {
		for _, row := range solution {
			fmt.Println(string(row))
		}
	} else {
		fmt.Println("ERROR")
	}
}
