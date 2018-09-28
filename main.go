package main

import (
	"fmt"
	"encoding/csv"
	"os"
	"strconv"
)

// force struct storing a force vector
type force struct{
	x, y	float64
}

// coordinate struct storing the position of stars
type coord struct{
	x, y	float64
}

// star struct storing information about the star
type star struct{
	c	coord
	f	force
	mass	float64
}

func main() {
	// stars_arr is a slice storing the stars
	stars_arr := []star{
		star{coord{1.0, 1.0}, force{0, 0}, 1000000},
		star{coord{3.0, 2.5}, force{0, 0}, 1000000},
		star{coord{1.0, 4.5}, force{0, 0}, 1000000},
	}

	// Open all the files from data/* and use their content as stars
	b, err := os.Open("data/structure03.ita.uni-heidelberg.de_26635.csv")

	// Error handling for opening files
	if err != nil {
		panic(err)
	}

	// Close the file after reading it's content
	defer b.Close()

	// Read the lines using a csv-reader
	lines, err := csv.NewReader(b).ReadAll()

	// CSV-Reader error handler
	if err != nil {
		panic(err)
	}

	for i := 0; i < 3; i++ {
		fmt.Println(i)

		// Iterate over all the lines in the given files, convert the data to float64 and
		// append the data to the stars_arr
		for _, line := range lines {

			// convert to float64
			x, _ := strconv.ParseFloat(line[0], 64)
			y, _ := strconv.ParseFloat(line[0], 64)

			// Define a temporary star that gets appended to the stars_arr
			temp_star := star{
				coord{x,y},
				force{0,0},
				1000000,
			}

			// Append the stars
			stars_arr = append(stars_arr, temp_star)
		}

	}

	// calculate the forces acting

	calc_all_forces(stars_arr)

	print_arr(stars_arr)

}
