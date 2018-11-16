package main

import (
	"fmt"
	"git.darknebu.la/GalaxySimulator/Source/csv"
	"git.darknebu.la/GalaxySimulator/Source/draw"
	"git.darknebu.la/GalaxySimulator/Source/forces"
	"git.darknebu.la/GalaxySimulator/Source/structs"
	"math"
)

func main() {
	var threads int = 8
	var frames int = 1
	var rangeStart int = 0
	var rangeEnd int = 1
	var path string = "data/U_ALL.csv"

	// the slice starsSlice stores the star structures
	starsSlice := []structs.Star2D{}
	starsSlice = csv.Import(path, rangeStart, rangeEnd, starsSlice)

	fmt.Println("Done loading the data")

	// Simulate frames
	for i := 0; i < frames; i++ {
		fmt.Println("Calculating the frame")

		starsSlice = forces.NextTimestep(starsSlice, 25*math.Pow(10, 4+7))
		starsSlice = forces.CalcAllAccelerations(starsSlice, threads)

		fmt.Println("Done Calculating")

		// draw the galaxy
		outputName := fmt.Sprintf("out_%d.png", i+4)
		draw.Slice(starsSlice, outputName)
		fmt.Println("Done drawing all the stars")
	}
}
