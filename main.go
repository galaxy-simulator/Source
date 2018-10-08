package main

import (
	"./csv"
	"./draw"
	"./forces"
	"./llog"
	"./structs"
	"fmt"
)

func AddToChannel(starsChan chan structs.Star, slice []structs.Star) {

	for index, _ := range slice {
		starsChan <- slice[index]
	}

}

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

	// Define a channel
	starsChannel := make(chan structs.Star, len(starsSlice))

	llog.Good("Opening the csv")
	starsSlice = csv.Import("data/structure03.ita.uni-heidelberg.de_26635.csv", 0, 2000, starsSlice)
	go AddToChannel(starsChannel, starsSlice)
	fmt.Printf("Done\n")

	llog.Good("Calculate the acting forces")
	forces.CalcAllForces(starsSlice)

	llog.Good("Drawing the slice")
	draw.Slice(starsSlice, "out_50000.png")
	fmt.Printf("Done\n")
}
