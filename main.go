package main

import (
	"encoding/json"
	"fmt"
	"git.darknebu.la/GalaxySimulator/Source/csv"
	"git.darknebu.la/GalaxySimulator/Source/draw"
	"git.darknebu.la/GalaxySimulator/Source/forces"
	"git.darknebu.la/GalaxySimulator/Source/structs"
	"math"
	"os"
)

type Config struct {
	Threads    int    `json:"Threads"`
	Frames     int    `json:"Frames"`
	RangeStart int    `json:"RangeStart"`
	RangeEnd   int    `json:"RangeEnd"`
	LoadPath   string `json:"LoadPath"`
	OutPath    string `json:"OutPath"`
}

func main() {
	// Load the config
	var config Config = LoadConfiguration("config.json")

	fmt.Printf("[+] Utilizing %d threads ", config.Threads)
	fmt.Printf("for drawing %d Frames, ", config.Frames)
	fmt.Printf("each containing %d Stars.\n", config.RangeEnd)

	fmt.Printf("[+] Getting previously existing Stars from %s ", config.LoadPath)
	fmt.Printf("and writing the results to %s.\n", config.OutPath)

	// the slice starsSlice stores the star structures
	starsSlice := []structs.Star2D{
		structs.Star2D{
			C: structs.Vec2{
				X: -1e5,
				Y: 0,
			},
			V: structs.Vec2{
				X: 0,
				Y: 0,
			},
			M: 1e10,
		},
		structs.Star2D{
			C: structs.Vec2{
				X: 1e5,
				Y: 0,
			},
			V: structs.Vec2{
				X: 0,
				Y: 0,
			},
			M: 1e10,
		},
		structs.Star2D{
			C: structs.Vec2{
				X: 0,
				Y: 4e4,
			},
			V: structs.Vec2{
				X: 0,
				Y: 0,
			},
			M: 1e10,
		},
	}
	starsSlice = csv.Import(config.LoadPath, config.RangeStart, config.RangeEnd, starsSlice)

	fmt.Println("Done loading the data")

	// Simulate frames
	for i := 0; i < config.Frames; i++ {
		fmt.Println("Calculating the frame")

		starsSlice = forces.NextTimestep(starsSlice, 25*math.Pow(10, 4+7))
		starsSlice = forces.CalcAllAccelerations(starsSlice, config.Threads)

		fmt.Println("Done Calculating")

		// draw the galaxy
		outputName := fmt.Sprintf("out_%d.png", i+4)
		draw.Slice(starsSlice, outputName)
		fmt.Println("Done drawing all the stars")
	}
}

// LoadConfiguration loads a configuration file from a given path and returns a struct with
// the values that are defined inside of the configuration file.
func LoadConfiguration(file string) Config {

	// Define some config defaults
	var config = Config{
		Threads:    1,
		Frames:     1,
		RangeStart: 0,
		RangeEnd:   1,
		OutPath:    "",
	}

	// Reading the config file and closing when done
	configFile, _ := os.Open(file)
	defer configFile.Close()

	// Parsing the content and adding it to the config struct
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	// Returning the config for further use
	return config
}
