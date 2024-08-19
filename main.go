package main

import (
	// Import the package used in the main function//+
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./tetromino_solver <file_path>")
		return
	}

	tetrominoes, err := ParseFile(os.Args[1])
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	solution := findSmallestSquare(tetrominoes)
	if solution != nil {
		for _, row := range solution {
			fmt.Println(string(row))
		}
	} else {
		fmt.Println("ERROR")
	}
}
