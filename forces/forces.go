package forces

import (
	"../structs"
	"fmt"
	"gopkg.in/cheggaaa/pb.v1"
	"math"
)

// forces_acting calculates the force inbetween the two given stars s1 and s2
// The function return the force
func forceActing(s1 structs.Star, s2 structs.Star) structs.Force {
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

// forces calculates the forces acting in between a given star and all the other stars in a given array.
func forces(stars_arr []structs.Star, nr int) structs.Force {

	var force structs.Force

	// Iterate over all the stars in the stars_arr
	for index := range stars_arr {

		// If the current star is not the star itself
		if index != nr {

			// generate a new force and add it to the overall force of the star
			fa := forceActing(stars_arr[nr], stars_arr[index])
			stars_arr[nr].F.X += fa.X
			stars_arr[nr].F.Y += fa.Y

			force.X += fa.X
			force.Y += fa.Y
		}
	}

	return force
}

// CalcAllForces calculates all the forces acting inbetween all the stars in the given starSlice slice and
// returns a "new" slice contaning the forces
func CalcAllForces(starSlice []structs.Star) []structs.Star {

	// Define a new slice in which the stars (and the forces acting on them) should be saved
	var new_slice []structs.Star

	fmt.Printf("\n")

	// Define a progres-bar
	bar := pb.StartNew(len(starSlice)).Prefix("Stars done: ")
	bar.SetWidth(80)

	// Iterate over all the stars in the original slice
	for index := range starSlice {

		// Increment the progress-bar
		bar.Increment()

		// Calculate the force acting inbetween the given star and all other stars
		// This utilizes go-routines :D
		var force = forces(starSlice, index)

		// create a new star
		current_star := structs.Star{
			structs.Coord{starSlice[index].C.X, starSlice[index].C.Y},
			structs.Force{force.X, force.Y},
			starSlice[index].Mass,
		}

		// append the new star to the new slice
		new_slice = append(new_slice, current_star)

	}

	// Print a newline after the progressbar
	bar.FinishPrint("")

	return new_slice
}
