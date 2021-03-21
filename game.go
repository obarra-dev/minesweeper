package minesweeper

import (
	"log"
)

// Play applies a user move.
func (g *game) Play(r, c int, move TypeMove) game {
	var game game

	if g.isMovePlayed(r, c, move) {
		g.State = StateGameRunning
		game = g.buildGameWithVisibleTiles()
	} else if move == TypeMoveClean {
		game = g.playOpenMove(r, c)
	} else if move != TypeMoveClean {
		game = g.mark(r, c)
	}

	return game
}

func (g game) isMovePlayed(r, c int, move TypeMove) bool {
	tile := g.Board[r][c]
	if tile.State == StateTileCovered {
		return false
	}

	if tile.State == StateTileNumbered || tile.State == StateTileClear || tile.State == StateTileExploited {
		return true
	}

	return tile.State == StateTileFlagged && move == TypeMoveFlag
}

func (g *game) playOpenMove(r, c int) game {
	tile := &g.Board[r][c]

	//game over, so show all tiles
	if tile.IsMine {
		log.Println("game Over")
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

	g.RevealEmptyAdjacentTiles(r, c)

	// game won, clear all tiles
	if g.isFlawlessVictory() {
		log.Println("Flawless Victory")
		g.State = StateGameWon
		return g.copyGame()
	}

	log.Println("The game is Running")
	g.State = StateGameRunning
	return g.buildGameWithVisibleTiles()
}

//TODO use Type Move Question
func (g *game) mark(r, c int) game {
	tile := &g.Board[r][c]

	if tile.State == StateTileCovered {
		log.Println("Flaging")
		tile.State = StateTileFlagged
		g.FlagAmount++
	} else if tile.State == StateTileFlagged {
		log.Println("Covering again")
		tile.State = StateTileCovered
		g.FlagAmount--
	}

	g.State = StateGameRunning
	return g.buildGameWithVisibleTiles()
}

func (g game) isFlawlessVictory() bool {
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

//TODO return points adjacent

// RevealEmptyAdjacentTiles makes visible  on the board all adjacent tiles from a point.
func (g game) RevealEmptyAdjacentTiles(r int, c int) {
	if g.Board[r][c].SurroundingMineCount == 0 {
		adjacentTiles := g.getAdjacentTiles(r, c)
		for i := 0; i < len(adjacentTiles); i++ {
			if adjacentTiles[i].IsMine != true &&
				(adjacentTiles[i].State == StateTileCovered || adjacentTiles[i].State == StateTileFlagged) {
				if adjacentTiles[i].SurroundingMineCount == 0 {
					g.Board[adjacentTiles[i].Row][adjacentTiles[i].Column].State = StateTileClear
					g.RevealEmptyAdjacentTiles(adjacentTiles[i].Row, adjacentTiles[i].Column)
				} else {
					g.Board[adjacentTiles[i].Row][adjacentTiles[i].Column].State = StateTileNumbered
				}
			}
		}
	}
}

func (g game) getAdjacentTiles(f int, c int) []Tile {
	minF := -1
	if f == 0 {
		minF = 0
	}

	minC := -1
	if c == 0 {
		minC = 0
	}

	maxF := 1
	if f == (g.Rows - 1) {
		maxF = 0
	}

	maxC := 1
	if c == (g.Columns - 1) {
		maxC = 0
	}

	var adjecentTiles []Tile
	for cc := minC; cc <= maxC; cc++ {
		for ff := minF; ff <= maxF; ff++ {
			if cc == 0 && ff == 0 {
				continue
			}

			var resultF = ff + f
			var resultC = cc + c

			adjecentTiles = append(adjecentTiles, g.Board[resultF][resultC])
		}
	}

	return adjecentTiles
}

func (g game) copyGame() game {
	board := make([][]Tile, g.Rows)

	for r := range board {
		board[r] = make([]Tile, g.Columns)
	}

	for i := 0; i < g.Rows; i++ {
		for j := 0; j < g.Columns; j++ {
			privateBoard := g.Board[i][j]
			board[i][j] = Tile{privateBoard.State, i, j, privateBoard.SurroundingMineCount, privateBoard.IsMine, -1}
		}
	}

	return game{g.State, board, g.Rows, g.Columns, g.MineAmount, g.FlagAmount}
}

//TODO no return matrix
func (g game) buildGameWithVisibleTiles() game {
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
	return game{g.State, board, g.Rows, g.Columns, g.MineAmount, g.FlagAmount}
}
