package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Board [][]bool

func main() {
	board := NewBoard(20, 20)
	board.Randomize(0.3)

	for i := 0; i < 100; i++ {
		board.Print()
		board.Next()
		time.Sleep(100 * time.Millisecond)
	}
}

func NewBoard(rows, cols int) Board {
	board := make(Board, rows)
	for i := range board {
		board[i] = make([]bool, cols)
	}
	return board
}

func (board Board) Randomize(prob float64) {
	rand.Seed(time.Now().UnixNano())
	for i := range board {
		for j := range board[i] {
			if rand.Float64() < prob {
				board[i][j] = true
			}
		}
	}
}

func (board Board) Next() {
	rows, cols := len(board), len(board[0])
	nextBoard := NewBoard(rows, cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			aliveNeighbors := board.countAliveNeighbors(i, j)
			if board[i][j] {
				if aliveNeighbors < 2 || aliveNeighbors > 3 {
					nextBoard[i][j] = false
				} else {
					nextBoard[i][j] = true
				}
			} else {
				if aliveNeighbors == 3 {
					nextBoard[i][j] = true
				} else {
					nextBoard[i][j] = false
				}
			}
		}
	}

	copy(board, nextBoard)
}

func (board Board) countAliveNeighbors(row, col int) int {
	rows, cols := len(board), len(board[0])
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			r := (row + i + rows) % rows
			c := (col + j + cols) % cols
			if board[r][c] {
				count++
			}
		}
	}
	return count
}

func (board Board) Print() {
	for i := range board {
		for j := range board[i] {
			if board[i][j] {
				fmt.Print(" X")
			} else {
				fmt.Print(" .")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
