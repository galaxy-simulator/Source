package main

import (
	"./csv"
	"./draw"
	"./forces"
	"./structs"
)

func main() {
	// the slice starsSlice stores the star structures
	starsSlice := []structs.Star{
		{structs.Coord{X: 100, Y: 100}, structs.Force{0, 0}, 1000},
		{structs.Coord{X: 200, Y: 300}, structs.Force{0, 0}, 1000},
		{structs.Coord{X: 100, Y: 500}, structs.Force{0, 0}, 1000},
	}

	csv.Import("data/structure03.ita.uni-heidelberg.de_26635.csv", 0, 4000, starsSlice)

	forces.CalcAllForces(starsSlice)
	draw.Slice(starsSlice, "out.png")
}
