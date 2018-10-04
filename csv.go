// collection of functions reading allready existing stars from a .csv

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// open_stars_csv opens the file defined using the path argument, reads its content and returns
// an array of arrays containing the coordinates of the stars
func open_stars_csv(path string) ([][]string) {

	// open the file in the given path
	b, err := os.Open(path)

	// handle errors
	if err != nil {
		fmt.Println("open_stars_csv panic")
		panic(err)
	}

	// close the file after reading it's content
	defer b.Close()

	// open the file using a csv reader and read it's content
	lines, err := csv.NewReader(b).ReadAll()

	// handle errors
	if err != nil {
		fmt.Println("read_csv panic")
		panic(err)
	}

	// return the content
	return lines
}

// add_csv_to_stars_arr adds the given stars to the given stars_arr
// the stars that should be added are expected to be stored in a[][]string
func add_csv_to_stars_arr(lines [][]string, stars_arr []star) {
	for _, line := range lines {

		// convert the coordinates to float64
		x, _ := strconv.ParseFloat(line[0], 64)
		y, _ := strconv.ParseFloat(line[1], 64)

		// create a temporary star for storing the information
		temp_star := star{
			coord{x, y},
			force{0, 0},
			1000000,
		}

		// add the temporary star to the stars_arr
		stars_arr = append(stars_arr, temp_star)
	}
}