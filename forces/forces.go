package forces

import (
	"../structs"
	"fmt"
	"git.darknebu.la/bit/logplus"
	"gopkg.in/cheggaaa/pb.v1"
	"math"
)

// forces_acting calculates the force inbetween the two given stars s1 and s2
// The function return the force
func accelerationActing(s1 structs.Star2D, s2 structs.Star2D) structs.Vec2 {

	// Gravitational constant
	var G = 6.674E-11

	// the vector from star s1 to star s2
	var r12 = s2.C.Subtract(s1.C)

	// the distance between the stars
	var deltaR = r12.GetLength()

	// the scalar acceleration of star s1 by star s2
	scalarA := G * (s2.M) / math.Pow(deltaR, 2)

	// calculate the direction vector of the vector from s1 to s2
	directionVectorA := r12.GetDirVector()

	// the vector acceleration of scalarA
	A := directionVectorA.Multiply(scalarA)

	return A
}

// accelerations calculates the acceleration acting in between a given star and all the other stars in a given array.
func accelerations(stars_arr []structs.Star2D, nr int) structs.Vec2 {

	var a = /*acceleration*/ structs.Vec2{}
	// Iterate over all the stars in the stars_arr
	for index := range stars_arr {

		// If the current star is not the star itself
		if index != nr {

			// calculate the acceleration and add it to the overall acceleration of the star
			aa := accelerationActing(stars_arr[nr], stars_arr[index])
			a = a.Add(aa)

		}
	}

	return a
}

// accelerationThread calculates the acceleration acting on a given amount of stars in a given range for a given slice of stars
// as a go-routine
func accelerationThread(starSlice []structs.Star2D, localRangeStart int, localRangeEnd int, channel chan structs.Star2D) {

	// iterate over the given range
	for index := localRangeStart; index < localRangeEnd; index++ {

		// Calculate the acceleration acting inbetween the given star and all other stars
		var a = accelerations(starSlice, index)

		// create a new star
		newStar := starSlice[index].Copy()
		newStar.AccelerateVelocity(a, 1)

		// push the new Star into the channel
		channel <- newStar
	}
}

// CalcAllAccelerations calculates all the accelerations acting in between all the stars in the given starSlice slice and
// returns a "new" slice containing the stars with their new velocities
func CalcAllAccelerations(starSlice []structs.Star2D, threads int) []structs.Star2D {
	// create a channel for bundling the stars generated in the go-routines
	channel := make(chan structs.Star2D, 1000)

	sliceLength := len(starSlice)

	// calculate the local range
	// Example: 100 stars with 4 threads = 25 stars / thread
	localRangeLen := sliceLength / threads

	// generate a new slice for storing the stars
	var newSlice []structs.Star2D

	logplus.LogNeutral(fmt.Sprintf("Starting %d workers, each processing %d stars", threads, localRangeLen))

	// start n go threads
	for i := 0; i < threads; i++ {

		// define the local range
		localRangeStart := i * localRangeLen
		localRangeEnd := (i * localRangeLen) + localRangeLen

		// fmt.Printf("starting worker nr. %d, processing %d stars\n", i, localRangeEnd-localRangeStart)

		// calculate the accelerations for all the stars in the given slice in the given range and return them using the
		// given channel
		go accelerationThread(starSlice, localRangeStart, localRangeEnd, channel)
	}

	// Handle errors (10004 stars, but 1250 stars per thread, so 4 stars are not calculate and block the queue)
	if sliceLength > localRangeLen {

		// Calculate the amount of stars and their range
		remainingStars := sliceLength - (localRangeLen * threads)
		localRangeEnd := ((threads - 1) * localRangeLen) + localRangeLen

		// Run the Thread
		// go accelerationThread(starSlice, localRangeEnd, localRangeEnd+remainingStars, channel)
		accelerationThread(starSlice, localRangeEnd, localRangeEnd+remainingStars, channel)
	}

	// Initialize a new progress bar
	bar := pb.New(len(starSlice)).Prefix("Stars: ")

	bar.Start()

	// iterate over the amount of stars
	for i := 0; i < sliceLength; i++ {

		// block until a star is finisehd
		var newStar = <-channel

		// append the star from the channel to the newSlice for returning in the end
		newSlice = append(newSlice, newStar)

		// increment the progress bar and the counter
		bar.Increment()
	}

	bar.Finish()

	return newSlice
}

// Calculate the new positions of the stars using the
func NextTimestep(starSlice []structs.Star2D, deltat float64) []structs.Star2D {
	// create a new slice for storing the "new" stars
	var newSlice []structs.Star2D

	// iterate over all the stars in the old slice
	for index := range starSlice {

		// move the star with it's velocity for time deltat
		newStar := starSlice[index].Copy()
		newStar.Move(deltat)

		// append the new star to the newSlice
		newSlice = append(newSlice, newStar)
	}

	return newSlice
}
