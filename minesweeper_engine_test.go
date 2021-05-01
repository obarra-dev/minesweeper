package minesweeper

import (
	"testing"
)

func TestGetAdjacentTilesWhenIsMatrix3x3(t *testing.T) {
	minesweeper := New(3, 3, []Mine{})
	t.Run("getAdjacentTiles when the point is in the middle", func(t *testing.T) {
		got := minesweeper.getAdjacentTiles(1, 1)
		expect := 8
		if len(got) != expect {
			t.Errorf("got %d expect %d", len(got), expect)
		}
	})

	t.Run("getAdjacentTiles when the point is in the upper right corner", func(t *testing.T) {
		got := minesweeper.getAdjacentTiles(0, 2)
		expect := 3
		if len(got) != expect {
			t.Errorf("got %d expect %d", len(got), expect)
		}
	})

	t.Run("getAdjacentTiles when the point is in the lower right corner", func(t *testing.T) {
		got := minesweeper.getAdjacentTiles(2, 2)
		expect := 3
		if len(got) != expect {
			t.Errorf("got %d expect %d", len(got), expect)
		}
	})

	t.Run("getAdjacentTiles when the point is in the upper left corner", func(t *testing.T) {
		got := minesweeper.getAdjacentTiles(0, 0)
		expect := 3
		if len(got) != expect {
			t.Errorf("got %d expect %d", len(got), expect)
		}
	})

	t.Run("getAdjacentTiles when the point is in the lower left corner", func(t *testing.T) {
		got := minesweeper.getAdjacentTiles(2, 0)
		expect := 3
		if len(got) != expect {
			t.Errorf("got %d expect %d", len(got), expect)
		}
	})
}

func TestGetAdjacentTilesShouldBe0(t *testing.T) {
	game := New(1, 1, []Mine{})
	result := game.getAdjacentTiles(0, 0)
	if len(result) != 0 {
		t.Error("Error", len(result))
	}
}
