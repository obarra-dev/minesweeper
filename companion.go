package minesweeper

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

// New creates a new board Game instance.
func New(rows, columns int, mines [][2]int) *Game {
	board := make([][]Tile, rows)
	for r := range board {
		board[r] = make([]Tile, columns)
	}

	cont := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			cont++
			board[r][c] = Tile{StateTileCovered, r, c, 0, false, cont}
		}
	}

	game := &Game{StateGameNew, board, rows, columns, len(mines), 0}
	game.setUpMines(mines)

	return game
}

// GenerateMinedPoints generates mines with random points.
func GenerateMinedPoints(maxRowIncluded, maxColumnIncluded, amountPoints int) [][2]int {
	tileMinePoints := make([][2]int, amountPoints)
	setPoints := make(map[string]bool)
	for len(setPoints) < amountPoints {
		concatenated := fmt.Sprint(rand.Intn(maxRowIncluded), "-", rand.Intn(maxColumnIncluded))
		setPoints[concatenated] = true
	}

	i := 0
	for key := range setPoints {
		point := strings.Split(key, "-")
		tileMinePoints[i][0], _ = strconv.Atoi(point[0])
		tileMinePoints[i][1], _ = strconv.Atoi(point[1])
		i++
	}

	return tileMinePoints
}

// ShowBoard shows all tile information of the Game board.
func (g Game) ShowBoard() {
	for i := 0; i < g.Rows; i++ {
		for j := 0; j < g.Columns; j++ {
			fmt.Print(g.Board[i][j], " ")
		}
		fmt.Println()
	}
}

// GetStates returns state tiles of the board Game.
func (g Game) GetStates() [][]StateTile {
	states := make([][]StateTile, g.Rows)

	for i := 0; i < g.Rows; i++ {
		states[i] = make([]StateTile, g.Columns)
		for j := 0; j < g.Columns; j++ {
			states[i][j] = g.Board[i][j].State
		}
	}
	return states
}
