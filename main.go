package main

import (
	"./csv"
	"./draw"
	"./forces"
	"./structs"
	"fmt"
	"git.darknebu.la/bit/logplus"
	"math"
)

func main() {
	logplus.SetLogLevel(logplus.LevelAll)
	var threads int = 8
	var frames int = 1
	var rangeStart int = 0
	var rangeEnd int = 5000

	// the slice starsSlice stores the star structures
	starsSlice := []structs.Star2D{
		{C: structs.Vec2{}, M: 5E8},
	}

	logplus.LogNeutral("Opening the csv")
	starsSlice = csv.Import("data/U_ALL.csv", rangeStart, rangeEnd, starsSlice)

	// Simulate frames
	for i := 0; i < frames; i++ {
		logplus.LogPositive("--- --- --- --- ---")
		logplus.LogPositive(fmt.Sprintf("Frames %d/%d", i, frames))

		logplus.LogNeutral("Calculate the new Star positions")
		starsSlice = forces.NextTimestep(starsSlice, 25*math.Pow(10, 4+7))

		logplus.LogNeutral("Calculate the acting accelerations")
		starsSlice = forces.CalcAllAccelerations(starsSlice, threads)

		outputName := fmt.Sprintf("out_%d.png", i+1)
		logplus.LogNeutral(fmt.Sprintf("draw the slice and save it to %s\n", outputName))
		draw.Slice(starsSlice, outputName)
	}
}
