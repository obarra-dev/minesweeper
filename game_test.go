package minesweeper

// TODO use minesweeper_test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGenerateMines(t *testing.T) {
	t.Run("generate mine when is big", func(t *testing.T) {
		mines := GenerateMines(200, 200, 200)
		expect := 200
		if len(mines) != expect {
			t.Errorf("got %d expect %d", expect, len(mines))
		}
	})

	t.Run("generate mine when is matrix is small", func(t *testing.T) {
		mines := GenerateMines(1, 1, 1)
		expect := 1
		if len(mines) != expect {
			t.Errorf("got %d expect %d", expect, len(mines))
		}
	})

	t.Run("generate mine when is matrix has no rows", func(t *testing.T) {
		mines := GenerateMines(0, 10, 3)
		expect := 0
		if len(mines) != expect {
			t.Errorf("got %d expect %d", expect, len(mines))
		}
	})

	t.Run("generate mine when is matrix has no columns", func(t *testing.T) {
		mines := GenerateMines(10, 0, 3)
		expect := 0
		if len(mines) != expect {
			t.Errorf("got %d expect %d", expect, len(mines))
		}
	})
}

func TestSetUpMines(t *testing.T) {
	minedPointTile := [][2]int{{0, 1}, {1, 1}, {1, 0}}
	game := New(3, 3, minedPointTile)
	game.ShowBoard()

	expected := [][]Tile{
		{Tile{StateTileCovered, 0, 0, 3, false, 1}, Tile{StateTileCovered, 0, 1, 2, true, 2}, Tile{StateTileCovered, 0, 2, 2, false, 3}},
		{Tile{StateTileCovered, 1, 0, 2, true, 4}, Tile{StateTileCovered, 1, 1, 2, true, 5}, Tile{StateTileCovered, 1, 2, 2, false, 6}},
		{Tile{StateTileCovered, 2, 0, 2, false, 7}, Tile{StateTileCovered, 2, 1, 2, false, 8}, Tile{StateTileCovered, 2, 2, 1, false, 9}}}

	if !reflect.DeepEqual(expected, game.Board) {
		t.Error("Error", game.Board)
	}
}

func TestMarkPlayWhenRunning(t *testing.T) {
	//setup
	minedPointTile := [][2]int{{1, 1}}
	game := New(3, 3, minedPointTile)

	//execute
	gameCopy := game.Play(0, 0, TypeMoveClean)

	//assert
	if gameCopy.State != StateGameRunning {
		t.Error("Error", gameCopy)
	}
}

func TestMarkPlayWhenRunningAndShowNumber(t *testing.T) {
	//setup
	minedPointTile := [][2]int{{1, 1}}
	game := New(3, 3, minedPointTile)

	//execute
	gameCopy := game.Play(0, 0, TypeMoveClean)
	fmt.Println(gameCopy)

	//assert
	if gameCopy.State != StateGameRunning || len(gameCopy.Board) != 1 {
		t.Error("Error", gameCopy, len(gameCopy.Board))
	}
}

func TestMarkPlayWhenGameLost(t *testing.T) {
	//setup
	minedPointTile := [][2]int{{1, 1}}
	game := New(3, 3, minedPointTile)

	//execute
	gameCopy := game.Play(1, 1, TypeMoveClean)

	//assert
	if gameCopy.State != StateGameLost {
		t.Error("Error", gameCopy)
	}
}

func TestMarkPlayWhenGameWon(t *testing.T) {
	//setup
	minedPointTile := [][2]int{{1, 1}}
	game := New(3, 3, minedPointTile)

	//execute
	gameCopy := game.Play(0, 0, TypeMoveClean)
	//assert
	if gameCopy.State != StateGameRunning {
		t.Error("Error", gameCopy)
	}

	gameCopy = game.Play(0, 1, TypeMoveClean)
	//assert
	if gameCopy.State != StateGameRunning {
		t.Error("Error", gameCopy)
	}

	gameCopy = game.Play(0, 2, TypeMoveClean)
	//assert
	if gameCopy.State != StateGameRunning {
		t.Error("Error", gameCopy)
	}

	gameCopy = game.Play(1, 0, TypeMoveClean)
	//assert
	if gameCopy.State != StateGameRunning {
		t.Error("Error", gameCopy)
	}

	gameCopy = game.Play(1, 2, TypeMoveClean)
	//assert
	if gameCopy.State != StateGameRunning {
		t.Error("Error", gameCopy)
	}

	gameCopy = game.Play(2, 0, TypeMoveClean)
	//assert
	if gameCopy.State != StateGameRunning {
		t.Error("Error", gameCopy)
	}

	gameCopy = game.Play(2, 1, TypeMoveClean)
	//assert
	if gameCopy.State != StateGameRunning {
		t.Error("Error", gameCopy)
	}

	gameCopy = game.Play(2, 2, TypeMoveClean)
	//assert
	if gameCopy.State != StateGameWon {
		t.Error("Error", gameCopy)
	}
}

func TestRevealEmptyAdjacentTiles3x3(t *testing.T) {
	minedPointTile := [][2]int{{1, 1}}
	game := New(3, 3, minedPointTile)

	game.RevealEmptyAdjacentTiles(0, 0)
	game.ShowBoard()

	expected := [][]StateTile{
		{StateTileCovered, StateTileCovered, StateTileCovered},
		{StateTileCovered, StateTileCovered, StateTileCovered},
		{StateTileCovered, StateTileCovered, StateTileCovered}}

	if !reflect.DeepEqual(expected, game.GetStates()) {
		t.Error("Error", game.GetStates())
	}
}

func TestRevealEmptyAdjacentTiles3x8(t *testing.T) {
	minedPointTile := [][2]int{{1, 1}}
	game := New(3, 8, minedPointTile)

	game.RevealEmptyAdjacentTiles(0, 5)
	game.ShowBoard()

	expected := [][]StateTile{
		{StateTileCovered, StateTileCovered, StateTileNumbered, StateTileClear, StateTileClear, StateTileClear, StateTileClear, StateTileClear},
		{StateTileCovered, StateTileCovered, StateTileNumbered, StateTileClear, StateTileClear, StateTileClear, StateTileClear, StateTileClear},
		{StateTileCovered, StateTileCovered, StateTileNumbered, StateTileClear, StateTileClear, StateTileClear, StateTileClear, StateTileClear},
	}
	if !reflect.DeepEqual(expected, game.GetStates()) {
		t.Error("Error", game.GetStates())
	}
}

func TestPlayMoveWhenFlag(t *testing.T) {
	game := New(3, 3, [][2]int{})

	gameCopy := game.Play(1, 1, TypeMoveFlag)

	if gameCopy.State != StateGameRunning || game.FlagAmount != 1 || gameCopy.Board[0][0].State != StateTileFlagged {
		t.Error("Error", game.FlagAmount, gameCopy)
	}
}

func TestPlayMoveWhenRevertTheFlag(t *testing.T) {
	game := New(3, 3, [][2]int{})

	gameCopy := game.Play(1, 1, TypeMoveFlag)
	gameCopy = game.Play(1, 1, TypeMoveRevertFlag)

	if gameCopy.State != StateGameRunning || gameCopy.FlagAmount != 0 {
		t.Error("Error", gameCopy, game.FlagAmount)
	}
}

func TestGetAdjacentTilesShouldBe8(t *testing.T) {
	game := New(3, 3, [][2]int{})
	result := game.getAdjacentTiles(1, 1)

	for d := 0; d < len(result); d++ {
		fmt.Println(result[d].ValueTest)
	}

	if len(result) != 8 {
		t.Error("Error", len(result))
	}
}

func TestGetAdjacentTilesShouldBe3(t *testing.T) {
	game := New(3, 3, [][2]int{})
	result := game.getAdjacentTiles(2, 2)
	if len(result) != 3 {
		t.Error("Error", len(result))
	}
}

func TestGetAdjacentTilesShouldBe5(t *testing.T) {
	game := New(3, 3, [][2]int{})
	result := game.getAdjacentTiles(0, 1)
	if len(result) != 5 {
		t.Error("Error", len(result))
	}
}

func TestGetAdjacentTilesShouldBe0(t *testing.T) {
	game := New(1, 1, [][2]int{})
	result := game.getAdjacentTiles(0, 0)
	if len(result) != 0 {
		t.Error("Error", len(result))
	}
}

func TestBuildGame(t *testing.T) {
	game := New(3, 3, [][2]int{})
	game.ShowBoard()

	if len(game.Board) != 3 || game.Rows != 3 &&
		len(game.Board[0]) != 3 || game.Columns != 3 {
		t.Error("Error", len(game.Board), len(game.Board[0]))
	}
}
