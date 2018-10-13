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
	var frames int = 1

	// the slice starsSlice stores the star structures
	starsSlice := []structs.Star{}

	llog.Good("Opening the csv")
	starsSlice = csv.Import("data/U_ALL.csv", 0, 2500, starsSlice)

	// Simulate the position of the stars after a specific time
	for i := 0; i < frames; i++ {
		llog.Great("--- --- --- --- ---")
		llog.Great(fmt.Sprintf("Frames %d/%d", i, frames))

		llog.Good("Calculate the new Star positions")
		starsSlice = forces.NextTimestep(starsSlice, 500000)

		llog.Good("Calculate the acting forces")
		starsSlice = forces.CalcAllForces(starsSlice, threads)

		outputName := fmt.Sprintf("out_%d.png", i+1)
		llog.Good(fmt.Sprintf("draw the slice and save it to %s\n", outputName))
		draw.Slice(starsSlice, outputName)
	}
}
