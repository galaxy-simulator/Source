package csv

import (
	"../file"
	"../structs"
	"git.darknebu.la/bit/logplus"
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
			logplus.LogFError("error reading value from csv in line nr. %d (%s)", linenr, errx)
		}
		if erry != nil {
			logplus.LogFError("error reading value from csv in line nr. %d (%s)", linenr, erry)
		}

		// Create a temporary star for assembling the star
		tempStar := structs.Star2D{
			C: structs.Vec2{X: x, Y: y},
			M: 50000,
		}

		// Add the Temporary star to the slice
		slice = append(slice, tempStar)

	}

	return slice
}
