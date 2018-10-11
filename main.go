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
	var path1 string = "out_0.png"
	var frames int = 250

	// the slice starsSlice stores the star structures
	starsSlice := []structs.Star{
		{structs.Coord{X: 30000, Y: 30000}, structs.Force{X: 0, Y: 0}, 500000000},
		{structs.Coord{X: -30000, Y: 30000}, structs.Force{X: 0, Y: 0}, 500000000},
		{structs.Coord{X: -30000, Y: 0}, structs.Force{X: 0, Y: 0}, 500000000},
		{structs.Coord{X: 30000, Y: -30000}, structs.Force{X: 0, Y: 0}, 500000000},
	}

	llog.Good("Opening the csv")
	starsSlice = csv.Import("data/U_ALL.csv", 0, 25000, starsSlice)

	// Step 1
	llog.Good("Calculate the acting forces")
	starsSlice = forces.CalcAllForces(starsSlice, threads)

	llog.Good(fmt.Sprintf("draw the slice and save it to %s\n", path1))
	draw.Slice(starsSlice, path1)

	// Step 2
	// Simulate the position of the stars after a specific time
	for i := 0; i < frames; i++ {
		llog.Great("--- --- --- --- ---")
		llog.Great(fmt.Sprintf("Frames %d/%d", i, frames))

		llog.Good("Calculate the new Star positions")
		starsSlice = forces.NextTimestep(starsSlice, 250000)

		llog.Good("Calculate the acting forces")
		starsSlice = forces.CalcAllForces(starsSlice, threads)

		outputName := fmt.Sprintf("out_%d.png", i+1)
		llog.Good(fmt.Sprintf("draw the slice and save it to %s\n", outputName))
		draw.Slice(starsSlice, outputName)
	}
}
