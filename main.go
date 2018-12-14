package main

import (
	"encoding/json"
	"fmt"
	"git.darknebu.la/GalaxySimulator/Source/csv"
	"git.darknebu.la/GalaxySimulator/Source/structs"
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

	fmt.Printf("[ ] Utilizing %d threads ", config.Threads)
	fmt.Printf("for drawing %d Frames, ", config.Frames)
	fmt.Printf("each containing %d Stars.\n", config.RangeEnd)

	fmt.Printf("[ ] Getting previously existing Stars from %s ", config.LoadPath)
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

	// import existing stars from a csv
	// generate new stars in a homogeneous grid

	starsSlice = csv.GenerateHomogeneousGrid(starsSlice, -5e5, 5e5, 1e5)
	fmt.Printf("Amount of Stars: %d\n", len(starsSlice))
	//starsSlice = csv.Import(config.LoadPath, config.RangeStart, config.RangeEnd, starsSlice)

	fmt.Println("[+] Done loading the data.")

	// Iterate over all the frames
	for i := 0; i < config.Frames; i++ {
		fmt.Printf("[ ] Frame %d\n", i)

		// Create a new quadtree
		boundary := *structs.NewBoundingBox(structs.Vec2{0, 0}, 1e8)
		starsQuadtree := *structs.NewQuadtree(boundary)

		// Print all the elements in the stars Slice
		for _, element := range starsSlice {
			fmt.Println(element)
		}

		// Insert all the stars from the starsSlice into the Quadtree
		//starQuadtree := quadtree.InsertSlice(starsSlice)

		//starsSlice = forces.NextTimestep(starsSlice, 25*math.Pow(10, 4+7))
		//starsSlice = forces.CalcAllAccelerations(starsSlice, config.Threads)
		//var starsQuadtree quadtree.Quadtree = quadtree.CreateWithSlice(starsSlice)
		//quadtree.Print(&starsQuadtree)
		//quadtree.Draw(&starsQuadtree)
		//quadtree.DrawQuadtree(starsQuadtree)

		fmt.Println("[+] Done Calculating the forces acting.")

		// draw the galaxy
		//fmt.Println("[ ] Drawing the Stars")
		//outputName := fmt.Sprintf("out_%d.png", i+4)
		//draw.Slice(starsSlice, outputName)
		//fmt.Println("[+] Done drawing all the stars")
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
	err := jsonParser.Decode(&config)
	if err != nil {
		panic(err)
	}

	// Returning the config for further use
	return config
}
