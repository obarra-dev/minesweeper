package minesweeper_test

import (
	"fmt"
	"github.com/obarra-dev/minesweeper"
	"testing"
)

func Test_New(t *testing.T) {
	t.Run("board 3x3", func(t *testing.T) {
		game := minesweeper.New(3, 3, minesweeper.GenerateMines(0, 0, 0))

		if len(game.Board) != game.Rows || game.Rows != 3 {
			t.Errorf("expect %d got %d", 3, len(game.Board))
		}

		if len(game.Board[0]) != game.Columns || game.Columns != 3 {
			t.Errorf("expect %d got %d", 3, len(game.Board[0]))
		}
	})
}

func Test_Play(t *testing.T) {
	t.Run("when clean into no mine should continue running", func(t *testing.T) {
		mines := []minesweeper.Mine{{R: 1, C: 1}}
		game := minesweeper.New(3, 3, mines)

		got := game.Play(0, 0, minesweeper.TypeMoveClean).State

		expect := minesweeper.StateGameRunning
		if got != expect {
			t.Errorf("expect %d got %d", expect, got)
		}
	})

	t.Run("when clean into mine should lose", func(t *testing.T) {
		mines := []minesweeper.Mine{{R: 1, C: 1}}
		game := minesweeper.New(3, 3, mines)

		got := game.Play(1, 1, minesweeper.TypeMoveClean).State

		expect := minesweeper.StateGameLost
		if got != expect {
			t.Errorf("expect %d got %d", expect, got)
		}
	})

	t.Run("when clean all ok should win", func(t *testing.T) {
		mines := []minesweeper.Mine{{R: 1, C: 1}}
		game := minesweeper.New(3, 3, mines)

		tests := []struct {
			r, c   int
			move   minesweeper.TypeMove
			expect minesweeper.StateGame
		}{
			{0, 0, minesweeper.TypeMoveClean, minesweeper.StateGameRunning},
			{0, 1, minesweeper.TypeMoveClean, minesweeper.StateGameRunning},
			{0, 2, minesweeper.TypeMoveClean, minesweeper.StateGameRunning},
			{1, 0, minesweeper.TypeMoveClean, minesweeper.StateGameRunning},
			{1, 2, minesweeper.TypeMoveClean, minesweeper.StateGameRunning},
			{2, 0, minesweeper.TypeMoveClean, minesweeper.StateGameRunning},
			{2, 1, minesweeper.TypeMoveClean, minesweeper.StateGameRunning},
			{2, 2, minesweeper.TypeMoveClean, minesweeper.StateGameWon},
		}
		for i, d := range tests {
			got := game.Play(d.r, d.c, d.move).State
			if got != d.expect {
				t.Errorf("Test[%d]: game.Play(%d,%d,%d) expect %d, got %d",
					i, d.r, d.c, d.move, d.expect, got)
			}
		}
	})
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

func showBoard(g minesweeper.Game) {
	for i := 0; i < g.Rows; i++ {
		for j := 0; j < g.Columns; j++ {
			fmt.Print(g.Board[i][j], " ")
		}
		fmt.Println()
	}
}

func GetStates(g minesweeper.Game) [][]minesweeper.StateTile {
	states := make([][]minesweeper.StateTile, g.Rows)

	for i := 0; i < g.Rows; i++ {
		states[i] = make([]minesweeper.StateTile, g.Columns)
		for j := 0; j < g.Columns; j++ {
			states[i][j] = g.Board[i][j].State
		}
	}
	return states
}
