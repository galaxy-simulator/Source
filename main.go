package main

import (
	"fmt"
	"git.darknebu.la/GalaxySimulator/Source/csv"
	"git.darknebu.la/GalaxySimulator/Source/draw"
	// "git.darknebu.la/GalaxySimulator/Source/forces"
	"git.darknebu.la/GalaxySimulator/Source/structs"
	"git.darknebu.la/bit/logplus"
	// "math"
	"os"
)

func main() {
	// Define a logging level for logplus
	logplus.SetLogLevel(logplus.LevelAll)

	// var threads int = 8
	var frames int = 1
	var rangeStart int = 0

	// Error handling (panic if there enouth arguments are provided)
	if len(os.Args) < 2 {
		panic("It seems like you forgot to supply a number of stars!")
	}
	rangeEnd, _ := os.Args[1]

	// the slice starsSlice stores the star structures
	starsSlice := []structs.Star2D{}

	// Import data from a csv
	logplus.LogNeutralf("Opening the csv")
	starsSlice = csv.Import("data/U_ALL.csv", rangeStart, rangeEnd, starsSlice)

	// Simulate frames
	for i := 0; i < frames; i++ {
		logplus.LogPositivef("--- --- --- --- ---")
		// logplus.LogPositive(fmt.Sprintf("Frames %d/%d", i, frames))	logplus.LogPositive("Done drawing the quadtree")

		// logplus.LogNeutral("Calculate the new Star positions")
		// starsSlice = forces.NextTimestep(starsSlice, 25*math.Pow(10, 4+7))

		// logplus.LogNeutral("Calculate the acting accelerations")
		// starsSlice = forces.CalcAllAccelerations(starsSlice, threads)

		// draw the galaxy
		outputName := fmt.Sprintf("out_%d.png", i+4)
		logplus.LogNeutralf(fmt.Sprintf("draw the slice and save it to %s", outputName))
		draw.Slice(starsSlice, outputName)
	}
}
