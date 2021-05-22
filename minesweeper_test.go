package minesweeper

import (
	"reflect"
	"testing"
)

func Test_getAdjacentTiles(t *testing.T) {
	t.Run("1x1 ", func(t *testing.T) {
		minesweeper := New(1, 1, []Mine{})
		got := minesweeper.getAdjacentTiles(0, 0)
		expect := 0
		if len(got) != expect {
			t.Errorf("expect %d got %d", expect, len(got))
		}
	})

	t.Run("3x3 ", func(t *testing.T) {
		minesweeper := New(3, 3, []Mine{})

		t.Run("when tile is in the middle", func(t *testing.T) {
			got := minesweeper.getAdjacentTiles(1, 1)
			expect := 8
			if len(got) != expect {
				t.Errorf("expect %d got %d", expect, len(got))
			}
		})

		t.Run("when tile is in the upper right corner", func(t *testing.T) {
			got := minesweeper.getAdjacentTiles(0, 2)
			expect := 3
			if len(got) != expect {
				t.Errorf("expect %d got %d", expect, len(got))
			}
		})

		t.Run("when tile is in the lower right corner", func(t *testing.T) {
			got := minesweeper.getAdjacentTiles(2, 2)
			expect := 3
			if len(got) != expect {
				t.Errorf("expect %d got %d", expect, len(got))
			}
		})

		t.Run("when tile is in the upper left corner", func(t *testing.T) {
			got := minesweeper.getAdjacentTiles(0, 0)
			expect := 3
			if len(got) != expect {
				t.Errorf("expect %d got %d", expect, len(got))
			}
		})
	})
}

func Test_revealEmptyAdjacentTiles(t *testing.T) {
	getStateTiles := func(g Game) [][]StateTile {
		states := make([][]StateTile, g.Rows)

		for i := 0; i < g.Rows; i++ {
			states[i] = make([]StateTile, g.Columns)
			for j := 0; j < g.Columns; j++ {
				states[i][j] = g.Board[i][j].State
			}
		}
		return states
	}

	t.Run("3x3", func(t *testing.T) {
		mines := []Mine{{Row: 1, Column: 1}}
		game := New(3, 3, mines)

		game.revealEmptyAdjacentTiles(0, 0)

		expect := [][]StateTile{
			{StateTileCovered, StateTileCovered, StateTileCovered},
			{StateTileCovered, StateTileCovered, StateTileCovered},
			{StateTileCovered, StateTileCovered, StateTileCovered}}

		got := getStateTiles(*game)
		if !reflect.DeepEqual(expect, got) {
			t.Errorf("expect %+v got %+v", expect, game.Board)
		}
	})

	t.Run("3x8", func(t *testing.T) {
		mines := []Mine{{Row: 1, Column: 1}}
		game := New(3, 8, mines)

		game.revealEmptyAdjacentTiles(0, 5)

		expect := [][]StateTile{
			{StateTileCovered, StateTileCovered, StateTileNumbered, StateTileClear, StateTileClear, StateTileClear, StateTileClear, StateTileClear},
			{StateTileCovered, StateTileCovered, StateTileNumbered, StateTileClear, StateTileClear, StateTileClear, StateTileClear, StateTileClear},
			{StateTileCovered, StateTileCovered, StateTileNumbered, StateTileClear, StateTileClear, StateTileClear, StateTileClear, StateTileClear},
		}

		got := getStateTiles(*game)
		if !reflect.DeepEqual(expect, got) {
			t.Errorf("expect %+v got %+v", expect, game.Board)
		}
	})
}
