package structs

/*
	The structs package defines structs that are used to store information that is used for processing
	star related data.
*/

// Force struct soring a force vector
type Force struct {
	X, Y float64
}

// Coord struct storing coordinates
type Coord struct {
	X, Y float64
}

// Star struct storing all necessary star information
type Star2D struct {
	C    Coord
	F    Force
	Mass float64
}
