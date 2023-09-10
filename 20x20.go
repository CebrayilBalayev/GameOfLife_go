package main

import (
	"image/color"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Board [][]bool

func NewRectangleWithColor(r, g, b uint8) *canvas.Rectangle {
	rect := canvas.NewRectangle(color.RGBA{r, g, b, 255})
	rect.SetMinSize(fyne.NewSize(20, 20))
	return rect
}

func GetRandomUint8() uint8 {
	return uint8(rand.Intn(256))
}

func main() {
	start := false
	timeToRefresh := time.Second
	minTimeToRefresh := time.Second - 900000000

	rand.Seed(time.Now().UnixNano())
	myApp := app.New()
	myWindow := myApp.NewWindow("Game of life")
	myWindow.Resize(fyne.NewSize(500, 500))

	n := 20
	rectContainer := container.NewGridWithColumns(n)
	board := NewBoard(n, n)
	board.Randomize(0.3)

	var Rectangles []*canvas.Rectangle
	for i := 0; i < (n); i++ {
		for j := 0; j < (n); j++ {
			var r, g, b uint8
			if board[i][j] {
				r, g, b = uint8(255), uint8(255), uint8(255)
			} else {
				r, g, b = uint8(0), uint8(0), uint8(0)
			}

			Rectangles = append(Rectangles, NewRectangleWithColor(r, g, b))
			rectContainer.Add(Rectangles[i*n+j])
		}
	}

	myLabel := widget.NewLabel(timeToRefresh.String())
	button1 := widget.NewButton(">", func() {
		board.Next()
		updateRectanglesColors(board, n, Rectangles)
		rectContainer.Refresh()
	})
	button2 := widget.NewButton("s/s", func() {
		start = !start
	})
	button3 := widget.NewButton("+", func() {
		timeToRefresh += minTimeToRefresh
		myLabel.SetText(timeToRefresh.String())
	})
	button4 := widget.NewButton("-", func() {
		if timeToRefresh != minTimeToRefresh {
			timeToRefresh -= minTimeToRefresh
			myLabel.SetText(timeToRefresh.String())
		}
	})

	buttons := container.NewVBox(button1, button2, button3, button4, myLabel)

	content := container.NewHBox(
		rectContainer,
		buttons,
	)

	go func() {
		for {
			time.Sleep(timeToRefresh)
			if start {
				board.Next()
				updateRectanglesColors(board, n, Rectangles)
				rectContainer.Refresh()
			}
		}
	}()

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func updateRectanglesColors(board Board, n int, rectangles []*canvas.Rectangle) {
	for i := 0; i < (n); i++ {
		for j := 0; j < (n); j++ {
			var r, g, b uint8
			if board[i][j] {
				r, g, b = uint8(255), uint8(255), uint8(255)
			} else {
				r, g, b = uint8(0), uint8(0), uint8(0)
			}
			rectangles[i*n+j].FillColor = color.RGBA{r, g, b, 255}
		}
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
