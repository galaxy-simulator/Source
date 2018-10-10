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
	var threads int = 8

	// the slice starsSlice stores the star structures
	starsSlice := []structs.Star{
		{structs.Coord{X: 30000, Y: 30000}, structs.Force{0, 0}, 500000000},
		{structs.Coord{X: -30000, Y: 30000}, structs.Force{0, 0}, 500000000},
		{structs.Coord{X: -30000, Y: 0}, structs.Force{0, 0}, 500000000},
		{structs.Coord{X: 30000, Y: -30000}, structs.Force{0, 0}, 500000000},
	}

	llog.Good("Opening the csv")
	starsSlice = csv.Import("data/U_ALL.csv", 0, 50000, starsSlice)

	llog.Good("Calculate the acting forces")
	starsSlice = forces.CalcAllForces(starsSlice, threads)

	path := "out_2.png"

	llog.Good(fmt.Sprintf("draw the slice and save it to %s", path))
	draw.Slice(starsSlice, path)

}
