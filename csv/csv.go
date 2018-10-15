package csv

import (
	"../structs"
	"encoding/csv"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// openStarCSV opens the file at the given path and reads its content. It then returns the content inform of a slice
// of slices
func openStarCSV(path string) [][]string {

	// Open the file located at the given path
	b, err := os.Open(path)

	// Handle errors
	if err != nil {
		log.Printf("openStarsCSV Panic! (cannot read file from %s)", path)
	}

	// Close the file afre reading it's content
	defer b.Close()

	// Parse the files conten usin a csv-reader
	lines, err := csv.NewReader(b).ReadAll()

	// Handle errors
	if err != nil {
		log.Println("openStarsCSV Panic! (cannot read the files content)")
	}

	return lines
}

// Import gets a file, a starting line, an ending line and  a struct. It then adds the content of the file to the struct
// For finding the length of the .csv, you can use the following command in linux:
// $ cat <csv> | wc -l
func Import(path string, start int, end int, slice []structs.Star) []structs.Star {
	lines := openStarCSV(path)

	// seed the random number generator
	rand.Seed(time.Now().UTC().UnixNano())

	// iterate over all the lines in the given range
	for linenr, line := range lines[start:end] {
		x, errx := strconv.ParseFloat(line[0], 64)
		y, erry := strconv.ParseFloat(line[1], 64)
		z, errz := strconv.ParseFloat(line[2], 64)

		// Handle errors
		if errx != nil {
			log.Printf("error reading value from csv in line nr. %d (%s)", linenr, errx)
		}
		if erry != nil {
			log.Printf("error reading value from csv in line nr. %d (%s)", linenr, erry)
		}
		if errz != nil {
			log.Printf("error reading value from csv in line nr. %d (%s)", linenr, errz)
		}

		// TODO: Export the code below to its own function
		var mass float64 = float64(10000 + rand.Intn(100000-10000))

		// Create a temporary star for assembling the star
		tempStar := structs.Star{
			structs.Coord{X: x, Y: y, Z: z},
			structs.Force{X: 0, Y: 0, Z: 0},
			mass,
		}

		// Add the Temporary star to the slice
		slice = append(slice, tempStar)

	}

	return slice
}
