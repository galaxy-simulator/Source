package main

import (
	"math"
	"fmt"
	"io/ioutil"
)

// forces_acting calculates the force inbetween the two given stars s1 and s2
// The function return the force
func force_acting(s1 star, s2 star) force {

	// Gravitational constant
	var G float64 = 6.674 * math.Pow(10, -11)

	// Distance between the stars
	var r21 float64 = math.Sqrt(math.Pow(s2.c.x - s1.c.x, 2) + math.Pow(s2.c.y - s1.c.y, 2))

	// Unit vector pointing from s1 to s2
	rhat := force{s2.c.x - s1.c.x, s2.c.y - s1.c.y}

	// Calculate how strong the star is affected
	var F_scalar float64 = G * (s1.mass * s2.mass) / math.Pow(math.Abs(r21), 2)

	// Calculate the overall force by combining the scalar and the vector
	var Fx float64 =  F_scalar * rhat.x
	var Fy float64 =  F_scalar * rhat.y

	// Pack the forces in a force structur
	F := force{Fx, Fy}

	return F
}

// forces calculates the forces acting in between a given star and all the other stars in a given array.
func forces(stars_arr []star, nr int) {

	// Iterate over all the stars in the stars_arr
	for index, _ := range stars_arr {

		// If the current star is not the star itself
		if index != nr {

			// generate a new force and add it to the overall force of the star
			fa := force_acting(stars_arr[nr], stars_arr[index])
			stars_arr[nr].f.x += fa.x
			stars_arr[nr].f.y += fa.y
		}
	}
}

// calc_all_forces iterates over all the stars in a given array and calculate the forces acting on all stars
func calc_all_forces(arr []star) {

	// Iterate over all the stars
	for index, _ := range arr {

		// Calculate the force acting inbetween the given star and all other stars
		// This utilizes go-routines :D
		go forces(arr, index)
	}
}

// save_as_csv saves all the information stored in the given array to the given path.
// If no error occures, the return value is nil
func save_as_csv(stars_arr []star, path string) error {

	var filename string = path + ".txt"
	var content string

	// Iterate over all the stars in the given array and add the info the the content that will be saved
	for _, elem := range stars_arr {
		content += fmt.Sprintf("%f,%f,%f,%f,%f\n", elem.c.x, elem.c.y, elem.f.x, elem.f.y, elem.mass)
	}

	// Return the error and write the data to disc
	return ioutil.WriteFile(filename, []byte(content), 0600)
}
