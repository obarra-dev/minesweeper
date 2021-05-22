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

// Tile holds tile information in the board Game.
type Tile struct {
	State                StateTile
	Row                  int
	Column               int
	SurroundingMineCount int
	IsMine               bool
}

// Mine holds mine information in the board Game.
type Mine struct {
	Row    int
	Column int
}

//TODO must be private

// Game holds board Game information.
type Game struct {
	State      StateGame
	Board      [][]Tile
	Rows       int
	Columns    int
	MineAmount int
	FlagAmount int
}
