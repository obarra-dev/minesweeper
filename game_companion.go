package engine

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func BuildNewGame(rows, columns int, minedPointTiles [][2]int) *Game {
	board := make([][]Tile, rows)
	for r := range board {
		board[r] = make([]Tile, columns)
	}

	cont := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			cont++
			board[r][c] = Tile{StateTileCovered, r, c, 0, false, cont}
		}
	}

	game := &Game{StateGameNew, board, rows, columns, len(minedPointTiles), 0}
	game.setUpMines(minedPointTiles)

	return game
}

func (g Game) setUpMines(minedPointTiles [][2]int) {
	for _, mine := range minedPointTiles {
		r := mine[0]
		c := mine[1]

		g.Board[r][c].IsMine = true

		adjacentTiles := g.getAdjacentTiles(r, c)
		for i := 0; i < len(adjacentTiles); i++ {
			g.Board[adjacentTiles[i].Row][adjacentTiles[i].Column].SurroundingMineCount++
		}
	}
}

func GenerateMinedPoints(maxRowIncluded, maxColumnIncluded, amountPoints int) [][2]int {
	tileMinePoints := make([][2]int, amountPoints)

	setPoints := make(map[string]bool)

	for len(setPoints) < amountPoints {
		concatenated := fmt.Sprint(rand.Intn(maxRowIncluded), "-", rand.Intn(maxColumnIncluded))
		setPoints[concatenated] = true
	}

	i := 0
	for key := range setPoints {
		point := strings.Split(key, "-")
		tileMinePoints[i][0], _ = strconv.Atoi(point[0])
		tileMinePoints[i][1], _ = strconv.Atoi(point[1])
		i++
	}

	return tileMinePoints
}
