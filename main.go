package main

import (
	"fmt"
	"os"
	"tetris-optimizer/src/utils"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: go run main.go path/to/<input file>")
		os.Exit(1)
	}

	// Read the file
	data, err := utils.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// Parse the data into a Tetromino
	Tetromino, err := utils.ParseTetromino(data)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	// Validate the Tetromino
	if !utils.IsValid(Tetromino) {
		fmt.Println("ERROR")
	}

	if utils.IsValid(Tetromino) {
		StartTime := time.Now()
		fmt.Println("Solving time: ", time.Since(StartTime))
	}

}
