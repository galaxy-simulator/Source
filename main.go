package main

import (
	"./forces"
	"./structs"
	"fmt"
)

func main() {
	// the slice starsSlice stores the star structures
	starsSlice := []structs.Star{
		{structs.Coord{X: 100, Y: 100}, structs.Force{0, 0}, 1000},
		{structs.Coord{X: 200, Y: 300}, structs.Force{0, 0}, 1000},
		{structs.Coord{X: 100, Y: 500}, structs.Force{0, 0}, 1000},
	}

	// Print the starsSlice
	fmt.Println(starsSlice)

	// Calculate the forces acting inbetween all the stars in the starsSlice slice and star nr 0 in the starsSlice slice
	fmt.Println(forces.CalcAllForces(starsSlice))
}
