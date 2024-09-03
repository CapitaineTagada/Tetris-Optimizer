package utils

import (
	"math"
)

// Tetromino represents a single tetromino piece
type Tetromino struct {
	Shape  []string
	Letter string
}

// func ParseTetrominoes(content string) []Tetromino {
// 	blocks := strings.Split(content, "\n\n")
// 	tetrominoes := make([]Tetromino, 0)
// 	for i, block := range blocks {
// 		lines := strings.Split(strings.TrimSpace(block), "\n")
// 		shape := make([][]string, len(lines))
// 		for j, line := range lines {
// 			shape[j] = strings.Split(line, "")
// 		}
// 		tetrominoes = append(tetrominoes, Tetromino{
// 			Shape:  shape,
// 			Letter: string(rune('A' + i)),
// 		})
// 	}
// 	return tetrominoes
// }

func MakeBoard(size int) [][]string {
	board := make([][]string, size)
	for i := range board {
		board[i] = make([]string, size)
		for j := range board[i] {
			board[i][j] = "."
		}
	}
	return board
}

func backtrack(board [][]string, tetrominoes []Tetromino, index int) bool {
	if index == len(tetrominoes) {
		return true
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if CanPlace(board, tetrominoes[index], i, j) {
				Place(board, tetrominoes[index], i, j)
				if backtrack(board, tetrominoes, index+1) {
					return true
				}
				Remove(board, tetrominoes[index], i, j)
			}
		}
	}
	return false
}

func SolveTetris(tetrominoes []Tetromino) [][]string {
	size := int(math.Ceil(math.Sqrt(float64(len(tetrominoes) * 4))))
	for {
		board := MakeBoard(size)
		if backtrack(board, tetrominoes, 0) {
			return board
		}
		size++
	}
}

func CanPlace(board [][]string, tetromino Tetromino, row, col int) bool {
	for i := 0; i < len(tetromino.Shape); i++ {
		for j := 0; j < len(tetromino.Shape[i]); j++ {
			if tetromino.Shape[i][j] == '#' {
				newRow, newCol := row+i, col+j
				if newRow >= len(board) || newCol >= len(board) || board[newRow][newCol] != "." {
					return false
				}
			}
		}
	}
	return true
}

func Place(board [][]string, tetromino Tetromino, row, col int) {
	for i := 0; i < len(tetromino.Shape); i++ {
		for j := 0; j < len(tetromino.Shape[i]); j++ {
			if tetromino.Shape[i][j] == '#' {
				board[row+i][col+j] = tetromino.Letter
			}
		}
	}
}

func Remove(board [][]string, tetromino Tetromino, row, col int) {
	for i := 0; i < len(tetromino.Shape); i++ {
		for j := 0; j < len(tetromino.Shape[i]); j++ {
			if tetromino.Shape[i][j] == '#' {
				board[row+i][col+j] = "."
			}
		}
	}
}
