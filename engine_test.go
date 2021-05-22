package minesweeper_test

import (
	"fmt"
	"github.com/obarra-dev/minesweeper"
	"reflect"
	"testing"
)

func Test_New(t *testing.T) {
	t.Run("board 3x3", func(t *testing.T) {
		mines := []minesweeper.Mine{
			{R: 0, C: 1},
			{R: 1, C: 1},
			{R: 1, C: 0},
		}
		game := minesweeper.New(3, 3, mines)

		if len(game.Board) != game.Rows || game.Rows != 3 {
			t.Errorf("expect %d got %d", 3, len(game.Board))
		}

		if len(game.Board[0]) != game.Columns || game.Columns != 3 {
			t.Errorf("expect %d got %d", 3, len(game.Board[0]))
		}

		expect := [][]minesweeper.Tile{
			{{minesweeper.StateTileCovered, 0, 0, 3, false}, {minesweeper.StateTileCovered, 0, 1, 2, true}, {minesweeper.StateTileCovered, 0, 2, 2, false}},
			{{minesweeper.StateTileCovered, 1, 0, 2, true}, {minesweeper.StateTileCovered, 1, 1, 2, true}, {minesweeper.StateTileCovered, 1, 2, 2, false}},
			{{minesweeper.StateTileCovered, 2, 0, 2, false}, {minesweeper.StateTileCovered, 2, 1, 2, false}, {minesweeper.StateTileCovered, 2, 2, 1, false}}}
		if !reflect.DeepEqual(expect, game.Board) {
			t.Errorf("expect %+v got %+v", expect, game.Board)
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

		caseTests := []struct {
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
		for i, ct := range caseTests {
			got := game.Play(ct.r, ct.c, ct.move).State
			if got != ct.expect {
				t.Errorf("Test[%d]: game.Play(%d,%d,%d) expect %d, got %d",
					i, ct.r, ct.c, ct.move, ct.expect, got)
			}
		}
	})

	t.Run("when flag and revert flag", func(t *testing.T) {
		type expect struct {
			flagAmount int
			gameState  minesweeper.StateGame
			tileState  minesweeper.StateTile
		}

		caseTests := []struct {
			r, c   int
			move   minesweeper.TypeMove
			expect expect
		}{
			{1, 1, minesweeper.TypeMoveFlag, expect{1, minesweeper.StateGameRunning, minesweeper.StateTileFlagged}},
			{1, 1, minesweeper.TypeMoveRevertFlag, expect{0, minesweeper.StateGameRunning, minesweeper.StateTileFlagged}},
		}

		game := minesweeper.New(3, 3, []minesweeper.Mine{})
		for i, ct := range caseTests {
			gameCopy := game.Play(ct.r, ct.c, ct.move)
			got := expect{
				flagAmount: gameCopy.FlagAmount,
				gameState:  gameCopy.State,
				tileState:  minesweeper.StateTileFlagged,
			}
			if got != ct.expect {
				t.Errorf("Test[%d]: game.Play(%d,%d,%d) expect %d, got %d",
					i, ct.r, ct.c, ct.move, ct.expect, got)
			}
		}
	})
}

func showBoard(g minesweeper.Game) {
	for i := 0; i < g.Rows; i++ {
		for j := 0; j < g.Columns; j++ {
			fmt.Print(g.Board[i][j], " ")
		}
		fmt.Println()
	}
}
