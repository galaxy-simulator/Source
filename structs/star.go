package structs

import "C"

type Star2D struct {
	C Vec2    // coordinates of the star
	V Vec2    // velocity    of the star
	M float64 // mass        of the star
}

func (s *Star2D) Copy() Star2D {
	return Star2D{s.C.Copy(), s.V.Copy(), s.M}
}

// accelerate the star with the acceleration a for the time t
func (s *Star2D) AccelerateVelocity(a Vec2, t float64) {
	s.V = s.V.Add(a.Multiply(t))
}

// move the star with it's velocity for the time t
func (s *Star2D) Move(t float64) {
	s.C = s.C.Add(s.V.Multiply(t))
}

// accelerate and move the star with it's velocity and the acceleration a for the time t
func (s *Star2D) Accelerate(a Vec2, t float64) {
	s.AccelerateVelocity(a, t)
	s.Move(t)
}
