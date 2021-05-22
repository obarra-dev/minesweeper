package minesweeper_test

import (
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
