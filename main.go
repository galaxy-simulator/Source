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
	var path1 string = "out_0.png"
	var frames int = 1

	// the slice starsSlice stores the star structures
	starsSlice := []structs.Star2D{
		//		{C: structs.Vec2{X:  30000,  Y: 30000},M: 5E8},
		//		{C: structs.Vec2{X: -30000,  Y: 30000},M: 5E8},
		//		{C: structs.Vec2{X: -30000           },M: 5E8},
		//		{C: structs.Vec2{X:  30000, Y: -30000},M: 5E8},
		{C: structs.Vec2{}, M: 5E8},
	}

	logplus.LogNeutral("Opening the csv")
	starsSlice = csv.Import("data/U_ALL.csv", 0, 25000, starsSlice)

	// Simulate the position of the stars after a specific time
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
