package engine

import (
	"fmt"
	"log"
)

type StateTile int

const (
	StateTileCovered  StateTile = 1
	StateTileClear    StateTile = 2
	StateTileFlagged  StateTile = 3
	StateTileNumberd  StateTile = 4
	StateTileExploted StateTile = 5
)

type StateGame int

const (
	StateGameNew     StateGame = 1
	StateGameRunning StateGame = 2
	StateGameWon     StateGame = 3
	StateGameLost    StateGame = 4
)

type TypeMove int

const (
	TypeMoveClean          TypeMove = 1
	TypeMoveFlag           TypeMove = 2
	TypeMoveQuestion       TypeMove = 3
	TypeMoveRevertFlag     TypeMove = 4
	TypeMoveRevertQuestion TypeMove = 5
)

type Tile struct {
	State                StateTile
	Row                  int
	Column               int
	SurroundingMineCount int
	IsMine               bool
	ValueTest            int
}

type Mine struct {
	r      int
	c      int
	active bool
}

//TODO must be private to avoid invalid states
type Game struct {
	State      StateGame
	Board      [][]Tile
	Rows       int
	Columns    int
	MineAmount int
	FlagAmount int
}

func (g *Game) Play(r, c int, move TypeMove) Game {
	var game Game

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

func (g Game) isMovePlayed(r, c int, move TypeMove) bool {
	tile := g.Board[r][c]
	if tile.State == StateTileCovered {
		return false
	}

	if tile.State == StateTileNumberd || tile.State == StateTileClear || tile.State == StateTileExploted {
		return true
	}

	return tile.State == StateTileFlagged && move == TypeMoveFlag
}

func (g *Game) playOpenMove(r, c int) Game {
	tile := &g.Board[r][c]

	//game over, so show all tiles
	if tile.IsMine {
		log.Println("Game Over")
		tile.State = StateTileExploted
		g.State = StateGameLost
		return g.copyGame()
	}

	//it's no mine, so clear or show number
	if tile.SurroundingMineCount == 0 {
		log.Println("Tile was Cleaned")
		tile.State = StateTileClear
	} else {
		log.Println("Tile was Numbered")
		tile.State = StateTileNumberd
	}

	g.RevealEmptyAdjacentTiles(r, c)

	// game won, clear all tiles
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
func (g *Game) mark(r, c int) Game {
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

//TODO return points adjacent
func (g Game) RevealEmptyAdjacentTiles(r int, c int) {
	if g.Board[r][c].SurroundingMineCount == 0 {
		adjecentTiles := g.getAdjacentTiles(r, c)
		for i := 0; i < len(adjecentTiles); i++ {
			if adjecentTiles[i].IsMine != true &&
				(adjecentTiles[i].State == StateTileCovered || adjecentTiles[i].State == StateTileFlagged) {
				if adjecentTiles[i].SurroundingMineCount == 0 {
					g.Board[adjecentTiles[i].Row][adjecentTiles[i].Column].State = StateTileClear
					g.RevealEmptyAdjacentTiles(adjecentTiles[i].Row, adjecentTiles[i].Column)
				} else {
					g.Board[adjecentTiles[i].Row][adjecentTiles[i].Column].State = StateTileNumberd
				}
			}
		}
	}
}

func (g Game) getAdjacentTiles(f int, c int) []Tile {
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

func (g Game) copyGame() Game {
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

	return Game{g.State, board, g.Rows, g.Columns, g.MineAmount, g.FlagAmount}
}

//TODO no return matrix
func (g Game) buildGameWithVisibleTiles() Game {
	var board [][]Tile
	for i := 0; i < g.Rows; i++ {
		var column []Tile
		for j := 0; j < g.Columns; j++ {
			if board := g.Board[i][j]; !board.IsMine &&
				(board.State == StateTileClear || board.State == StateTileNumberd || board.State == StateTileFlagged) {
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

func (g Game) ShowBoard() {
	for i := 0; i < g.Rows; i++ {
		for j := 0; j < g.Columns; j++ {
			fmt.Print(g.Board[i][j], " ")
		}
		fmt.Println()
	}
}

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
