package structs

// Force struct soring a force vector
type Force struct {
	X, Y float64
}

// Coord struct storing coordinates
type Coord struct {
	X, Y float64
}

// Star struct storing all necessary star information
type Star struct {
	C    Coord
	F    Force
	Mass float64
}
