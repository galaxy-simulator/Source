package main

import (
	"./structs"
	"fmt"
)

func main() {
	// the slice starsSlice stores the star structures
	starsSlice := []structs.Star{
		{structs.Coord{X: 100, Y: 100}, structs.Force{0, 0}, 1000},
	}

	fmt.Println(starsSlice)
}
