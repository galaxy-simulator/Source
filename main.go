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
	// var starsSlice []structs.Star

	// llog.Good("Opening the csv")
	// starsSlice = csv.Import("data/structure03.ita.uni-heidelberg.de_26635.csv", 0, 2000, starsSlice)
	starsSlice = csv.Import("data/U_ALL.csv", 0, 50000, starsSlice)
	// fmt.Printf("Done\n")

	// llog.Good("Calculate the acting forces")
	starsSlice = forces.CalcAllForces(starsSlice, threads)
	// starsSlice = forces.CalcAllForcesOld(starsSlice)

	// for _, e := range starsSlice {
	// 	fmt.Printf("%f %f | %f %f | %f\n", e.C.X, e.C.Y, e.F.X, e.F.Y, e.Mass)
	// }

	path := "out_2.png"

	llog.Good(fmt.Sprintf("draw the slice and save it to %s", path))
	draw.Slice(starsSlice, path)

}
