package utils

import (
	"bufio"
	"errors"
	"os"
)

type Tetromino struct {
	shape  [][]rune
	width  int
	height int
}

func NewTetromino(shape [][]rune) (*Tetromino, error) {
	if len(shape) != 4 || len(shape[0]) != 4 {
		return nil, errors.New("invalid tetromino shape")
	}
	minX, minY, maxX, maxY := 4, 4, 0, 0
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			if shape[y][x] == '#' {
				if x < minX {
					minX = x
				}
				if y < minY {
					minY = y
				}
				if x > maxX {
					maxX = x
				}
				if y > maxY {
					maxY = y
				}
			}
		}
	}
	width := maxX - minX + 1
	height := maxY - minY + 1
	croppedShape := make([][]rune, height)
	for i := range croppedShape {
		croppedShape[i] = shape[minY+i][minX : maxX+1]
	}
	return &Tetromino{croppedShape, width, height}, nil
}

func parseFile(path string) ([]*Tetromino, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tetrominoes []*Tetromino
	scanner := bufio.NewScanner(file)
	var shape [][]rune
	lineCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) != 4 && len(line) != 0 {
			return nil, errors.New("ERROR")
		}

		if len(line) == 4 {
			shape = append(shape, []rune(line))
			lineCount++
		}

		if lineCount == 4 {
			tetromino, err := NewTetromino(shape)
			if err != nil {
				return nil, err
			}
			tetrominoes = append(tetrominoes, tetromino)
			shape = [][]rune{}
			lineCount = 0
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return tetrominoes, nil
}

func canPlace(grid [][]rune, t *Tetromino, x, y int) bool {
	for i := 0; i < t.height; i++ {
		for j := 0; j < t.width; j++ {
			if t.shape[i][j] == '#' && grid[y+i][x+j] != '.' {
				return false
			}
		}
	}
	return true
}

func placeTetromino(grid [][]rune, t *Tetromino, x, y int, ch rune) {
	for i := 0; i < t.height; i++ {
		for j := 0; j < t.width; j++ {
			if t.shape[i][j] == '#' {
				grid[y+i][x+j] = ch
			}
		}
	}
}

func removeTetromino(grid [][]rune, t *Tetromino, x, y int) {
	for i := 0; i < t.height; i++ {
		for j := 0; j < t.width; j++ {
			if t.shape[i][j] == '#' {
				grid[y+i][x+j] = '.'
			}
		}
	}
}

func solve(tetrominoes []*Tetromino, gridSize int) [][]rune {
	grid := make([][]rune, gridSize)
	for i := range grid {
		grid[i] = make([]rune, gridSize)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	var backtrack func(int) bool
	backtrack = func(index int) bool {
		if index == len(tetrominoes) {
			return true
		}
		t := tetrominoes[index]
		for y := 0; y <= gridSize-t.height; y++ {
			for x := 0; x <= gridSize-t.width; x++ {
				if canPlace(grid, t, x, y) {
					placeTetromino(grid, t, x, y, rune('A'+index))
					if backtrack(index + 1) {
						return true
					}
					removeTetromino(grid, t, x, y)
				}
			}
		}
		return false
	}

	if backtrack(0) {
		return grid
	}
	return nil
}

func findSmallestSquare(tetrominoes []*Tetromino) [][]rune {
	size := 2
	for size*size < len(tetrominoes)*4 {
		size++
	}

	for {
		solution := solve(tetrominoes, size)
		if solution != nil {
			return solution
		}
		size++
	}
}
