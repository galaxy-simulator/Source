package draw

import (
	"../structs"
	"github.com/fogleman/gg"
)

// initializePlot generates a new plot and returns the plot context
func initializePlot() *gg.Context {

	// Define the image size
	const image_width = 8192
	const image_height = 8192

	// Initialize the new context
	dc := gg.NewContext(image_width, image_height)

	// Set the background black
	dc.SetRGB(0, 0, 0)
	dc.Clear()

	// Set the coordinate midpoint to the middle of the image
	dc.Translate(image_width/2, image_height/2)

	return dc
}

func Slice(slice []structs.Star, path string) {
	dc := initializePlot()
	dc.SetRGB(1, 1, 1)
	dc.DrawCircle(0, 0, 100)
	dc.Stroke()
	dc.SavePNG(path)
}
