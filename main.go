package main

import (
	"./csv"
	"./draw"
	"./forces"
	"./llog"
	"./structs"
	"fmt"
)

func main() {
	// the slice starsSlice stores the star structures
	starsSlice := []structs.Star{
		{structs.Coord{X: 100, Y: 100}, structs.Force{0, 0}, 1000},
		{structs.Coord{X: 200, Y: 300}, structs.Force{0, 0}, 1000},
		{structs.Coord{X: 100, Y: 500}, structs.Force{0, 0}, 1000},
		{structs.Coord{X: -200000.0, Y: 8000.0}, structs.Force{X: 0, Y: 0}, 8000000000},
		{structs.Coord{X: 200000.0, Y: 0.0}, structs.Force{X: 0, Y: 0}, 8000000000},
		{structs.Coord{X: 0.0, Y: 200000}, structs.Force{X: 0, Y: 0}, 8000000000},
		{structs.Coord{X: 20000.0, Y: 80000}, structs.Force{X: 0, Y: 0}, 4000000},
	}

	llog.Good("Opening the csv")
	starsSlice = csv.Import("data/structure03.ita.uni-heidelberg.de_26635.csv", 0, 2000, starsSlice)
	// starsSlice = csv.Import("data/U_ALL.csv", 0, 2000, starsSlice)
	fmt.Printf("Done\n")

	llog.Good("Calculate the acting forces")
	forces.CalcAllForces(starsSlice)

	// Filename without suffix and extension
	path := "out"

	llog.Good(fmt.Sprintf("draw the slice and save it to %s", path))
	draw.Slice(starsSlice, path)

}
