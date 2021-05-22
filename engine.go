package minesweeper

import (
	"log"
)

// New creates a new board Game instance.
func New(rows, columns int, mines []Mine) *Game {
	board := make([][]Tile, rows)

	for r := 0; r < rows; r++ {
		board[r] = make([]Tile, columns)
		for c := 0; c < columns; c++ {
			board[r][c] = Tile{StateTileCovered, r, c, 0, false}
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

func (g Game) setUpMines(mines []Mine) {
	for _, mine := range mines {
		g.Board[mine.R][mine.C].IsMine = true
		adjacentTiles := g.getAdjacentTiles(mine.R, mine.C)
		for i := 0; i < len(adjacentTiles); i++ {
			g.Board[adjacentTiles[i].Row][adjacentTiles[i].Column].SurroundingMineCount++
		}
	}
}

func (g Game) isMovePlayed(r, c int, move TypeMove) bool {
	tile := g.Board[r][c]

	switch tile.State {
	case StateTileCovered:
		return false
	case StateTileNumbered, StateTileClear, StateTileExploited:
		return true
	}

	return tile.State == StateTileFlagged && move == TypeMoveFlag
}

func (g *Game) playMoveClean(r, c int) Game {
	tile := &g.Board[r][c]

	//Game over, so show all tiles
	if tile.IsMine {
		log.Println("Game Over")
		tile.State = StateTileExploited
		g.State = StateGameLost
		return g.copyGame()
	}

	//it's no mine, so clear or show number
	if tile.SurroundingMineCount == 0 {
		log.Println("Tile was Cleaned")
		tile.State = StateTileClear
	} else {
		log.Println("Tile was Numbered")
		tile.State = StateTileNumbered
	}

	g.revealEmptyAdjacentTiles(r, c)

	// Game won, clear all tiles
	if g.isFlawlessVictory() {
		log.Println("Flawless Victory")
		g.State = StateGameWon
		return g.copyGame()
	}

	log.Println("The Game is Running")
	g.State = StateGameRunning
	return g.buildGameWithVisibleTiles()
}

//TODO use Type Move Question
func (g *Game) playMoveFlag(r, c int) Game {
	tile := &g.Board[r][c]

	if tile.State == StateTileCovered {
		log.Println("Flagging")
		tile.State = StateTileFlagged
		g.FlagAmount++
	} else if tile.State == StateTileFlagged {
		log.Println("Covering")
		tile.State = StateTileCovered
		g.FlagAmount--
	}

	g.State = StateGameRunning
	return g.buildGameWithVisibleTiles()
}

func (g Game) isFlawlessVictory() bool {
	for i := 0; i < g.Rows; i++ {
		for j := 0; j < g.Columns; j++ {
			if board := g.Board[i][j]; !board.IsMine &&
				(board.State == StateTileCovered || board.State == StateTileFlagged) {
				return false
			}
		}
	}

	return true
}

// RevealEmptyAdjacentTiles makes visible all adjacent tiles from a point on the board.
func (g Game) revealEmptyAdjacentTiles(r, c int) {
	if g.Board[r][c].SurroundingMineCount == 0 {
		adjacentTiles := g.getAdjacentTiles(r, c)
		for _, tile := range adjacentTiles {
			if !tile.IsMine && (tile.State == StateTileCovered || tile.State == StateTileFlagged) {
				if tile.SurroundingMineCount == 0 {
					g.Board[tile.Row][tile.Column].State = StateTileClear
					g.revealEmptyAdjacentTiles(tile.Row, tile.Column)
				} else {
					g.Board[tile.Row][tile.Column].State = StateTileNumbered
				}
			}
		}
	}
}

func (g Game) getAdjacentTiles(r, c int) []Tile {
	minR := -1
	if r == 0 {
		minR = 0
	}

	minC := -1
	if c == 0 {
		minC = 0
	}

	maxR := 1
	if r == (g.Rows - 1) {
		maxR = 0
	}

	maxC := 1
	if c == (g.Columns - 1) {
		maxC = 0
	}

	var tiles []Tile
	for cc := minC; cc <= maxC; cc++ {
		for rr := minR; rr <= maxR; rr++ {
			if cc == 0 && rr == 0 {
				continue
			}

			var adjacentR = rr + r
			var adjacentC = cc + c

			tiles = append(tiles, g.Board[adjacentR][adjacentC])
		}
	}

	return tiles
}

func (g Game) copyGame() Game {
	board := make([][]Tile, g.Rows)

	for r := range board {
		board[r] = make([]Tile, g.Columns)
	}

	for i := 0; i < g.Rows; i++ {
		for j := 0; j < g.Columns; j++ {
			privateBoard := g.Board[i][j]
			board[i][j] = Tile{privateBoard.State, i, j, privateBoard.SurroundingMineCount, privateBoard.IsMine}
		}
	}

	return Game{g.State, board, g.Rows, g.Columns, g.MineAmount, g.FlagAmount}
}

//TODO no return matrix?
func (g Game) buildGameWithVisibleTiles() Game {
	var board [][]Tile
	for i := 0; i < g.Rows; i++ {
		var column []Tile
		for j := 0; j < g.Columns; j++ {
			if board := g.Board[i][j]; !board.IsMine &&
				(board.State == StateTileClear || board.State == StateTileNumbered || board.State == StateTileFlagged) {
				column = append(column, g.Board[i][j])
			}
		}
		if column != nil && len(column) > 0 {
			board = append(board, column)
		}
	}

	if board == nil {
		board = [][]Tile{}
	}
	return Game{g.State, board, g.Rows, g.Columns, g.MineAmount, g.FlagAmount}
}
