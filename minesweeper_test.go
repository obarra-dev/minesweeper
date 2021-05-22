package minesweeper_test

import (
	"testing"

	"github.com/obarra-dev/minesweeper"
)

func Test_New(t *testing.T) {
	t.Run("board 3x3", func(t *testing.T) {
		game := minesweeper.New(3, 3, minesweeper.GenerateMines(0, 0, 0))

		if len(game.Board) != game.Rows || game.Rows != 3 {
			t.Errorf("expect %d got %d", 3, len(game.Board))
		}

		if len(game.Board[0]) != game.Columns  || game.Columns != 3 {
			t.Errorf("expect %d got %d", 3, len(game.Board[0]))
		}
	})
}

func Test_GenerateMines(t *testing.T) {
	t.Run("when board is big", func(t *testing.T) {
		mines := minesweeper.GenerateMines(200, 200, 200)
		expect := 200
		if len(mines) != expect {
			t.Errorf("expect %d got %d", expect, len(mines))
		}
	})

	t.Run("when board is small", func(t *testing.T) {
		mines := minesweeper.GenerateMines(1, 1, 1)
		expect := 1
		if len(mines) != expect {
			t.Errorf("expect %d got %d", expect, len(mines))
		}
	})

	t.Run("when board has no rows", func(t *testing.T) {
		mines := minesweeper.GenerateMines(0, 10, 3)
		expect := 0
		if len(mines) != expect {
			t.Errorf("expect %d got %d", expect, len(mines))
		}
	})

	t.Run("when board has no columns", func(t *testing.T) {
		mines := minesweeper.GenerateMines(10, 0, 3)
		expect := 0
		if len(mines) != expect {
			t.Errorf("expect %d got %d", expect, len(mines))
		}
	})

	t.Run("is random mines", func(t *testing.T) {
		mines := minesweeper.GenerateMines(17, 4, 8)
		if got := isRandomMines(mines, 17, 4); !got {
			t.Errorf("expect %t got %t", got, true)
		}
	})
}

func isRandomMines(mines []minesweeper.Mine, rows, columns int) bool {
	board := make([][]bool, rows)
	for r := 0; r < rows; r++ {
		board[r] = make([]bool, columns)
	}

	for _, m := range mines {
		if unique := board[m.R][m.C]; !unique {
			board[m.R][m.C] = true
		} else {
			return false
		}
	}

	return true
}

/**

func TestSetUpMines(t *testing.T) {
	minedPointTile := [][2]int{{0, 1}, {1, 1}, {1, 0}}
	game := minesweeper.New(3, 3, minedPointTile)
	game.ShowBoard()

	expected := [][]minesweeper.Tile{
		{{minesweeper.StateTileCovered, 0, 0, 3, false, 1}, minesweeper.Tile{minesweeper.StateTileCovered, 0, 1, 2, true, 2}, minesweeper.Tile{minesweeper.StateTileCovered, 0, 2, 2, false, 3}},
		{minesweeper.Tile{minesweeper.StateTileCovered, 1, 0, 2, true, 4}, minesweeper.Tile{minesweeper.StateTileCovered, 1, 1, 2, true, 5}, minesweeper.Tile{minesweeper.StateTileCovered, 1, 2, 2, false, 6}},
		{minesweeper.Tile{minesweeper.StateTileCovered, 2, 0, 2, false, 7}, minesweeper.Tile{minesweeper.StateTileCovered, 2, 1, 2, false, 8}, minesweeper.Tile{minesweeper.StateTileCovered, 2, 2, 1, false, 9}}}

	if !reflect.DeepEqual(expected, game.Board) {
		t.Error("Error", game.Board)
	}
}

func TestMarkPlayWhenRunning(t *testing.T) {
	//setup
	minedPointTile := [][2]int{{1, 1}}
	game := minesweeper.New(3, 3, minedPointTile)

	//execute
	gameCopy := game.Play(0, 0, minesweeper.TypeMoveClean)

	//assert
	if gameCopy.State != minesweeper.StateGameRunning {
		t.Error("Error", gameCopy)
	}
}

func TestMarkPlayWhenRunningAndShowNumber(t *testing.T) {
	//setup
	minedPointTile := [][2]int{{1, 1}}
	game := minesweeper.New(3, 3, minedPointTile)

	//execute
	gameCopy := game.Play(0, 0, minesweeper.TypeMoveClean)
	fmt.Println(gameCopy)

	//assert
	if gameCopy.State != minesweeper.StateGameRunning || len(gameCopy.Board) != 1 {
		t.Error("Error", gameCopy, len(gameCopy.Board))
	}
}

func TestMarkPlayWhenGameLost(t *testing.T) {
	//setup
	minedPointTile := [][2]int{{1, 1}}
	game := minesweeper.New(3, 3, minedPointTile)

	//execute
	gameCopy := game.Play(1, 1, minesweeper.TypeMoveClean)

	//assert
	if gameCopy.State != minesweeper.StateGameLost {
		t.Error("Error", gameCopy)
	}
}

func TestMarkPlayWhenGameWon(t *testing.T) {
	//setup
	minedPointTile := [][2]int{{1, 1}}
	game := minesweeper.New(3, 3, minedPointTile)

	//execute
	gameCopy := game.Play(0, 0, minesweeper.TypeMoveClean)
	//assert
	if gameCopy.State != minesweeper.StateGameRunning {
		t.Error("Error", gameCopy)
	}

	gameCopy = game.Play(0, 1, minesweeper.TypeMoveClean)
	//assert
	if gameCopy.State != minesweeper.StateGameRunning {
		t.Error("Error", gameCopy)
	}

	gameCopy = game.Play(0, 2, minesweeper.TypeMoveClean)
	//assert
	if gameCopy.State != minesweeper.StateGameRunning {
		t.Error("Error", gameCopy)
	}

	gameCopy = game.Play(1, 0, minesweeper.TypeMoveClean)
	//assert
	if gameCopy.State != minesweeper.StateGameRunning {
		t.Error("Error", gameCopy)
	}

	gameCopy = game.Play(1, 2, minesweeper.TypeMoveClean)
	//assert
	if gameCopy.State != minesweeper.StateGameRunning {
		t.Error("Error", gameCopy)
	}

	gameCopy = game.Play(2, 0, minesweeper.TypeMoveClean)
	//assert
	if gameCopy.State != minesweeper.StateGameRunning {
		t.Error("Error", gameCopy)
	}

	gameCopy = game.Play(2, 1, minesweeper.TypeMoveClean)
	//assert
	if gameCopy.State != minesweeper.StateGameRunning {
		t.Error("Error", gameCopy)
	}

	gameCopy = game.Play(2, 2, minesweeper.TypeMoveClean)
	//assert
	if gameCopy.State != minesweeper.StateGameWon {
		t.Error("Error", gameCopy)
	}
}

func TestRevealEmptyAdjacentTiles3x3(t *testing.T) {
	minedPointTile := [][2]int{{1, 1}}
	game := minesweeper.New(3, 3, minedPointTile)

	game.RevealEmptyAdjacentTiles(0, 0)
	game.ShowBoard()

	expected := [][]minesweeper.StateTile{
		{minesweeper.StateTileCovered, minesweeper.StateTileCovered, minesweeper.StateTileCovered},
		{minesweeper.StateTileCovered, minesweeper.StateTileCovered, minesweeper.StateTileCovered},
		{minesweeper.StateTileCovered, minesweeper.StateTileCovered, minesweeper.StateTileCovered}}

	if !reflect.DeepEqual(expected, game.GetStates()) {
		t.Error("Error", game.GetStates())
	}
}

func TestRevealEmptyAdjacentTiles3x8(t *testing.T) {
	minedPointTile := [][2]int{{1, 1}}
	game := minesweeper.New(3, 8, minedPointTile)

	game.RevealEmptyAdjacentTiles(0, 5)
	game.ShowBoard()

	expected := [][]minesweeper.StateTile{
		{minesweeper.StateTileCovered, minesweeper.StateTileCovered, minesweeper.StateTileNumbered, minesweeper.StateTileClear, minesweeper.StateTileClear, minesweeper.StateTileClear, minesweeper.StateTileClear, minesweeper.StateTileClear},
		{minesweeper.StateTileCovered, minesweeper.StateTileCovered, minesweeper.StateTileNumbered, minesweeper.StateTileClear, minesweeper.StateTileClear, minesweeper.StateTileClear, minesweeper.StateTileClear, minesweeper.StateTileClear},
		{minesweeper.StateTileCovered, minesweeper.StateTileCovered, minesweeper.StateTileNumbered, minesweeper.StateTileClear, minesweeper.StateTileClear, minesweeper.StateTileClear, minesweeper.StateTileClear, minesweeper.StateTileClear},
	}
	if !reflect.DeepEqual(expected, game.GetStates()) {
		t.Error("Error", game.GetStates())
	}
}

func TestPlayMoveWhenFlag(t *testing.T) {
	game := minesweeper.New(3, 3, [][2]int{})

	gameCopy := game.Play(1, 1, minesweeper.TypeMoveFlag)

	if gameCopy.State != minesweeper.StateGameRunning || game.FlagAmount != 1 || gameCopy.Board[0][0].State != minesweeper.StateTileFlagged {
		t.Error("Error", game.FlagAmount, gameCopy)
	}
}

func TestPlayMoveWhenRevertTheFlag(t *testing.T) {
	game := minesweeper.New(3, 3, [][2]int{})

	gameCopy := game.Play(1, 1, minesweeper.TypeMoveFlag)
	gameCopy = game.Play(1, 1, minesweeper.TypeMoveRevertFlag)

	if gameCopy.State != minesweeper.StateGameRunning || gameCopy.FlagAmount != 0 {
		t.Error("Error", gameCopy, game.FlagAmount)
	}
}
*/
