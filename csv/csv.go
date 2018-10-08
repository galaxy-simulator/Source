package csv

import (
	"../structs"
	"encoding/csv"
	"log"
	"os"
	"strconv"
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

	for linenr, line := range lines[start:end] {
		x, errx := strconv.ParseFloat(line[0], 64)
		y, erry := strconv.ParseFloat(line[1], 64)

		// Handle errors
		if errx != nil {
			log.Printf("error reading value from csv in line nr. %d (%s)", linenr, errx)
		}
		if errx != nil {
			log.Printf("error reading value from csv in line nr. %d (%s)", linenr, erry)
		}

		// Create a temporary star for assembling the star
		tempStar := structs.Star{
			structs.Coord{x, y},
			structs.Force{0, 0},
			1000000,
		}

		// Add the Temporary star to the slice
		slice = append(slice, tempStar)

	}

	return slice
}
