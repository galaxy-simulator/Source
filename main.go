package main

import (
	"fmt"
)

// force struct storing a force vector
type force struct{
	x, y	float64
}

// coordinate struct storing the position of stars
type coord struct{
	x, y	float64
}

// star struct storing information about the star
type star struct{
	c	coord
	f	force
	mass	float64
}

func main() {
	// stars_arr is a slice storing the stars
	stars_arr := []star{
		star{coord{1.0, 1.0}, force{0, 0}, 1000000},
		star{coord{3.0, 2.5}, force{0, 0}, 1000000},
		star{coord{1.0, 4.5}, force{0, 0}, 1000000},
	}

	fmt.Println(stars_arr)
}
