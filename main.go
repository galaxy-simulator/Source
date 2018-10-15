package main

import (
	"./csv"
	"./draw"
	"./forces"
	"./structs"
	"fmt"
	"git.darknebu.la/bit/logplus"
)

func main() {
	var threads int = 8
	var frames int = 1
	var stars int = 800

	// the slice starsSlice stores the star structures
	starsSlice := []structs.Star2D{}

	logplus.LogPositive("Opening the csv")
	starsSlice = csv.Import("data/U_ALL.csv", 0, stars, starsSlice)

	// Simulate the position of the stars after a specific time
	for i := 0; i < frames; i++ {

		// Print the iterator
		logplus.LogPositive("--- --- --- --- ---")
		logplus.LogPositive(fmt.Sprintf("Frames %d/%d", i, frames))

		// Calculate the position of stars after one timestep
		logplus.LogPositive("Calculate the new Star positions")
		starsSlice = forces.NextTimestep(starsSlice, 500000)

		// Calculate all the forces acting inside of the galaxy
		logplus.LogPositive("Calculate the acting forces")
		starsSlice = forces.CalcAllForces(starsSlice, threads)

		// Save the slice
		outputName := fmt.Sprintf("out_%d.png", i+1)
		logplus.LogPositive(fmt.Sprintf("draw the slice and save it to %s\n", outputName))
		draw.Slice(starsSlice, outputName)
	}
}
