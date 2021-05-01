package minesweeper
/**
import (
	"fmt"
	"testing"
)

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
**/