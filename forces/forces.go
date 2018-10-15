package forces

import (
	"../llog"
	"../structs"
	"fmt"
	"gopkg.in/cheggaaa/pb.v1"
	"math"
)

// forces_acting calculates the force inbetween the two given stars s1 and s2
// The function return the force
func forceActing(s1 structs.Star2D, s2 structs.Star2D) structs.Force {
	// Gravitational constant
	var G float64 = 6.674e-11

	// Distance between the stars
	var r21 = math.Sqrt(math.Pow(s2.C.X-s1.C.X, 2) + math.Pow(s2.C.Y-s1.C.Y, 2))

	// Unit vector pointing from s1 to s2
	rhat := structs.Force{s2.C.X - s1.C.X, s2.C.Y - s1.C.Y}

	// Calculate how strong the star is affected
	var FScalar = G * (s1.Mass * s2.Mass) / math.Pow(math.Abs(r21), 2)

	// Calculate the overall force by combining the scalar and the vector
	var Fx = FScalar * rhat.X
	var Fy = FScalar * rhat.Y

	// Pack the forces in a force structur
	F := structs.Force{Fx, Fy}

	return F
}

// forces calculates the forces acting in between a given star and all the other stars in a given array.
func forces(stars_arr []structs.Star2D, nr int) structs.Force {

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

// forcesThread calculates the forces acting on a given amount of stars in a given range for a given slice of stars
// as a go-routine
func forcesThread(starSlice []structs.Star2D, localRangeStar2Dt int, localRangeEnd int, channel chan structs.Star2D) {

	// iterate over the given range
	for index := localRangeStar2Dt; index < localRangeEnd; index++ {

		// Calculate the force acting inbetween the given star and all other stars
		var force = forces(starSlice, index)

		// create a new star
		newStar2D := structs.Star2D{
			structs.Coord{starSlice[index].C.X, starSlice[index].C.Y},
			structs.Force{force.X, force.Y},
			starSlice[index].Mass,
		}

		// push the new Star2D into the channel
		channel <- newStar2D
	}
}

// CalcAllForces calculates all the forces acting inbetween all the stars in the given starSlice slice and
// returns a "new" slice contaning the forces
func CalcAllForces(starSlice []structs.Star2D, threads int) []structs.Star2D {
	// create a channel for bundling the stars generaten in the go-routines
	channel := make(chan structs.Star2D, 1000)

	sliceLength := len(starSlice)

	// calculate the local range
	// Example: 100 stars with 4 threads = 25 stars / thread
	localRangeLen := sliceLength / threads

	// generate a new slice for storing the stars
	var newSlice []structs.Star2D

	llog.Good(fmt.Sprintf("Starting %d workers, each processing %d stars", threads, localRangeLen))

	// start n go threads
	for i := 0; i < threads; i++ {

		// define the local range
		localRangeStar2Dt := i * localRangeLen
		localRangeEnd := (i * localRangeLen) + localRangeLen

		// calculate the forces for all the stars in the given slice in the given range and return them using the
		// given channel
		go forcesThread(starSlice, localRangeStar2Dt, localRangeEnd, channel)
	}

	// Handle errors (10004 stars, but 1250 stars per thread, so 4 stars are not calculate and block the queue)
	if sliceLength > localRangeLen {

		// Calculate the amount of stars and their range
		remainingStar2Ds := sliceLength - (localRangeLen * threads)
		localRangeEnd := ((threads - 1) * localRangeLen) + localRangeLen

		// Run the Thread
		// go forcesThread(starSlice, localRangeEnd, localRangeEnd+remainingStar2Ds, channel)
		forcesThread(starSlice, localRangeEnd, localRangeEnd+remainingStar2Ds, channel)
	}

	// Initialize a new progress bar
	bar := pb.New(len(starSlice)).Prefix("Star2Ds: ")

	// iterate over the amount of stars
	for i := 0; i < sliceLength; i++ {

		// block until a star is finisehd
		var newStar2D structs.Star2D = <-channel

		// append the star from the channel to the newSlice for returning in the end
		newSlice = append(newSlice, newStar2D)

		// increment the progress bar and the counter
		bar.Increment()
	}

	bar.Finish()

	return newSlice
}

// Calculate the new positions of the stars using the
func NextTimestep(starSlice []structs.Star2D, deltat int) []structs.Star2D {
	// create a new slice for storing the "new" stars
	var newSlice []structs.Star2D

	// iterate over all the stars in the old slice
	for index := range starSlice {

		// calculate the new position
		newX := starSlice[index].C.X + starSlice[index].F.X*float64(deltat)
		newY := starSlice[index].C.Y + starSlice[index].F.Y*float64(deltat)

		// assemble the new star
		newStar2D := structs.Star2D{
			C:    structs.Coord{X: newX, Y: newY},
			F:    structs.Force{},
			Mass: starSlice[index].Mass,
		}

		// append the new star to the newSlice
		newSlice = append(newSlice, newStar2D)
	}

	return newSlice
}
