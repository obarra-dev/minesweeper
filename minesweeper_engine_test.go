package minesweeper

import (
	"testing"
)

func Test_getAdjacentTiles(t *testing.T) {
	t.Run("1x1 ", func(t *testing.T) {
		minesweeper := New(1, 1, []Mine{})
		got := minesweeper.getAdjacentTiles(0, 0)
		expect := 0
		if len(got) != expect {
			t.Errorf("got %d expect %d", len(got), expect)
		}
	})

	t.Run("3x3 ", func(t *testing.T) {
		minesweeper := New(3, 3, []Mine{})

		t.Run("when tile is in the middle", func(t *testing.T) {
			got := minesweeper.getAdjacentTiles(1, 1)
			expect := 8
			if len(got) != expect {
				t.Errorf("got %d expect %d", len(got), expect)
			}
		})

		t.Run("when tile is in the upper right corner", func(t *testing.T) {
			got := minesweeper.getAdjacentTiles(0, 2)
			expect := 3
			if len(got) != expect {
				t.Errorf("got %d expect %d", len(got), expect)
			}
		})

		t.Run("when tile is in the lower right corner", func(t *testing.T) {
			got := minesweeper.getAdjacentTiles(2, 2)
			expect := 3
			if len(got) != expect {
				t.Errorf("got %d expect %d", len(got), expect)
			}
		})

		t.Run("when tile is in the upper left corner", func(t *testing.T) {
			got := minesweeper.getAdjacentTiles(0, 0)
			expect := 3
			if len(got) != expect {
				t.Errorf("got %d expect %d", len(got), expect)
			}
		})
	})

}
