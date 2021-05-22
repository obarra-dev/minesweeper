package minesweeper

// StateTile is a enum, it's represents the state of the tile.
type StateTile int

// All possible states of the tile.
const (
	StateTileCovered StateTile = iota
	StateTileClear
	StateTileFlagged
	StateTileNumbered
	StateTileExploited
)

// StateGame is a enum, it's represents the state of the Game.
type StateGame int

// All possible states of the Game.
const (
	StateGameNew StateGame = iota
	StateGameRunning
	StateGameWon
	StateGameLost
)

// TypeMove is a enum, it's represents the type of user move.
type TypeMove int

// All possible types of the user move.
const (
	TypeMoveClean TypeMove = iota
	TypeMoveFlag
	TypeMoveRevertFlag
	TypeMoveQuestion
	TypeMoveRevertQuestion
)

// Mine holds mine information in the board Game.
type Mine struct {
	Row    int
	Column int
}

//TODO must be private? if it's public, it can be set with invalid values
type (
	// Game holds board Game information.
	Game struct {
		State      StateGame
		Board      [][]Tile
		Rows       int
		Columns    int
		MineAmount int
	}

	// Tile holds tile information in the board Game.
	Tile struct {
		State                StateTile
		Row                  int
		Column               int
		SurroundingMineCount int
		IsMine               bool
	}
)
