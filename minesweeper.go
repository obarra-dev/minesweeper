package minesweeper

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
)

// New creates a new board Game instance.
func New(rows, columns int, mines []Mine) *Game {
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

// Play applies a user move.
func (g *Game) Play(r, c int, move TypeMove) Game {
	var game Game
	if g.isMovePlayed(r, c, move) {
		game.State = StateGameRunning
		return g.buildGameWithVisibleTiles()
	}

	switch move {
	case TypeMoveClean:
		game = g.playMoveClean(r, c)
	case TypeMoveFlag, TypeMoveRevertFlag:
		game = g.playMoveFlag(r, c)
	default:
		log.Println("invalid type move")
	}

	return game
}

// GenerateMines generates mines with random mines for given matrix.
func GenerateMines(rows, columns, amountMines int) []Mine {
	if rows <= 0 || columns <= 0 {
		return []Mine{}
	}

	const (
		r         = 0
		c         = 1
		separator = "-"
	)

	generateRandomPoints := func() map[string]bool {
		setOfPoints := make(map[string]bool)
		for len(setOfPoints) < amountMines {
			point := fmt.Sprint(rand.Intn(rows), separator, rand.Intn(columns))
			setOfPoints[point] = true
		}
		return setOfPoints
	}

	var mines []Mine
	for key := range generateRandomPoints() {
		point := strings.Split(key, separator)
		row, _ := strconv.Atoi(point[r])
		column, _ := strconv.Atoi(point[c])
		mines = append(mines, Mine{
			r: row,
			c: column,
		})
	}

	return mines
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
