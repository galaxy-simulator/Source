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
	fmt.Println("Hello World")
}
