package minesweeper

// StateTile is a enum, it's represents the state of the tile.
type StateTile int

// All possible states of the tile.
const (
	StateTileCovered   StateTile = 1
	StateTileClear     StateTile = 2
	StateTileFlagged   StateTile = 3
	StateTileNumbered  StateTile = 4
	StateTileExploited StateTile = 5
)

// StateGame is a enum, it's represents the state of the Game.
type StateGame int

// All possible states of the Game.
const (
	StateGameNew     StateGame = 1
	StateGameRunning StateGame = 2
	StateGameWon     StateGame = 3
	StateGameLost    StateGame = 4
)

// TypeMove is a enum, it's represents the type of user move.
type TypeMove int

// All possible types of the user move.
const (
	TypeMoveClean          TypeMove = 1
	TypeMoveFlag           TypeMove = 2
	TypeMoveQuestion       TypeMove = 3
	TypeMoveRevertFlag     TypeMove = 4
	TypeMoveRevertQuestion TypeMove = 5
)

// Tile holds tile information in the board Game.
type Tile struct {
	State                StateTile
	Row                  int
	Column               int
	SurroundingMineCount int
	IsMine               bool
	ValueTest            int
}

// Mine holds mine information in the board Game.
type Mine struct {
	r      int
	c      int
	active bool
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
