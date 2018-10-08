package structs

type Force struct {
	X, Y float64
}

type Coord struct {
	X, Y float64
}

type Star struct {
	C    Coord
	F    Force
	Mass float64
}
