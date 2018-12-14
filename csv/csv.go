package csv

import (
	"git.darknebu.la/GalaxySimulator/Source/file"
	"git.darknebu.la/GalaxySimulator/Source/structs"
	"git.darknebu.la/bit/logplus"
	"math/rand"
	"strconv"
)

// Import gets a file, a starting line, an ending line and  a struct. It then adds the content of the file to the struct
// For finding the length of the .csv, you can use the following command in linux:
// $ cat <csv> | wc -l
func Import(path string, start int, end int, slice []structs.Star2D) []structs.Star2D {
	f, _ := file.Open(path)
	lines, _ := f.ReadCSV()

	for linenr, line := range lines[start:end] {
		x, errx := strconv.ParseFloat(line[0], 64)
		y, erry := strconv.ParseFloat(line[1], 64)

		// Handle errors
		if errx != nil {
			logplus.LogErrorf("error reading value from csv in line nr. %d (%s)", linenr, errx)
		}
		if erry != nil {
			logplus.LogErrorf("error reading value from csv in line nr. %d (%s)", linenr, erry)
		}

		// Create a temporary star for assembling the star
		tempStar := structs.Star2D{
			C: structs.Vec2{X: x, Y: y},
			M: float64(rand.Intn(50000)),
		}

		// Add the Temporary star to the slice
		slice = append(slice, tempStar)

	}

	return slice
}

// Generate a homogeneous Grid
func GenerateHomogeneousGrid(slice []structs.Star2D, left int, right int, step int) []structs.Star2D {

	// Iterate over a grid
	for i := left; i < right; i += step {
		for j := left; j < right; j += step {

			// generate a new star with the coordinates
			tempStar := structs.Star2D{
				C: structs.Vec2{X: float64(i) + float64(rand.Intn(step)), Y: float64(j) + float64(rand.Intn(step))},
				M: float64(rand.Intn(500000)),
			}

			// add the star to the slice
			slice = append(slice, tempStar)
		}
	}

	// return the new slice containing a homogeneous grid of stars
	return slice
}
