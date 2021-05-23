package minesweeper_test

import (
	"testing"

	"github.com/obarra-dev/minesweeper"
)

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
		if unique := board[m.Row][m.Column]; !unique {
			board[m.Row][m.Column] = true
		} else {
			return false
		}
	}

	return true
}
