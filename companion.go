package minesweeper

import (
	"math/rand"
)

// GenerateMines generates random mines for given length board.
func GenerateMines(rows, columns, amountMines int) []Mine {
	if rows <= 0 || columns <= 0 {
		return []Mine{}
	}

	type point struct {
		r int
		c int
	}

	generateRandomPoints := func() map[point]bool {
		points := make(map[point]bool)
		for len(points) < amountMines {
			p := point{
				r: rand.Intn(rows),
				c: rand.Intn(columns),
			}
			points[p] = true
		}
		return points
	}

	var mines []Mine
	for p := range generateRandomPoints() {
		mines = append(mines, Mine{
			Row:    p.r,
			Column: p.c,
		})
	}

	return mines
}
