package utils

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func IsValid(Tetromino [][]string) bool {
	for _, r := range Tetromino {
		TetrominoConnect := 0
		HashCount := 0
		for CheckVertical, line := range r {
			for CheckHorizontal, char := range line {
				if char != '#' && char != '.' {
					return false
				} else if char == '#' {
					HashConnect := 0
					HashCount++
					if CheckVertical > 0 && r[CheckVertical-1][CheckHorizontal] == '#' {
						HashConnect++
					}
					if CheckVertical < len(r)-1 && r[CheckVertical+1][CheckHorizontal] == '#' {
						HashConnect++
					}
					if CheckHorizontal > 0 && r[CheckVertical][CheckHorizontal-1] == '#' {
						HashConnect++
					}
					if CheckHorizontal < len(line)-1 && r[CheckVertical][CheckHorizontal+1] == '#' {
						HashConnect++
					}
					if HashConnect == 0 {
						return false
					} else {
						TetrominoConnect += HashConnect
					}
				}
			}
		}
	}
	return true
}

// ReadFile reads the contents of a file specified by the path and returns it as a byte slice.
func ReadFile(path string) ([]byte, error) {
	// Open the file
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("ERROR: %v", err)
	}
	defer file.Close() // Ensure the file is closed after reading

	// Read the file contents
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("ERROR: %v", err)
	}
	return data, nil
}

// parseTetromino converts the file data into a Tetromino ([][]string)
func ParseTetromino(data []byte) ([][]string, error) {
	lines := strings.Split(string(data), "\n")
	var Tetromino [][]string

	for _, line := range lines {
		if line == "" {
			continue
		}
		row := strings.Split(line, "")
		Tetromino = append(Tetromino, row)
	}

	// Ensure Tetromino is 4x4
	if len(Tetromino) != 4 {
		fmt.Println("ERROR")
	}
	for _, row := range Tetromino {
		if len(row) != 4 {
			fmt.Println("ERROR")
		}
	}

	return Tetromino, nil
}

//Work on the solving function
