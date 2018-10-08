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
	}

	fmt.Println(starsSlice)
	fmt.Println(forces.ForceActing(starsSlice[0], starsSlice[1]))
}
