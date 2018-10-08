package forces

import (
	"../structs"
	"math"
)

// forces_acting calculates the force inbetween the two given stars s1 and s2
// The function return the force
func ForceActing(s1 structs.Star, s2 structs.Star) structs.Force {
	// Gravitational constant
	var G = 6.674 * math.Pow(10, -11)

	// Distance between the stars
	var r21 = math.Sqrt(math.Pow(s2.C.X-s1.C.X, 2) + math.Pow(s2.C.Y-s1.C.Y, 2))

	// Unit vector pointing from s1 to s2
	rhat := structs.Force{s2.C.X - s1.C.X, s2.C.Y - s1.C.Y}

	// Calculate how strong the star is affected
	var F_scalar = G * (s1.Mass * s2.Mass) / math.Pow(math.Abs(r21), 2)

	// Calculate the overall force by combining the scalar and the vector
	var Fx = F_scalar * rhat.X
	var Fy = F_scalar * rhat.Y

	// Pack the forces in a force structur
	F := structs.Force{Fx, Fy}

	return F
}
