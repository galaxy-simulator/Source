package draw

import (
	"../structs"
	"fmt"
	"github.com/fogleman/gg"
	"math"
)

// initializePlot generates a new plot and returns the plot context
func initializePlot() *gg.Context {
	// Define the image size
	const imageWidth = 8192
	const imageHeight = 8192

	// Initialize the new context
	dc := gg.NewContext(imageWidth, imageHeight)

	// Set the background black
	dc.SetRGB(0, 0, 0)
	dc.Clear()

	// Invert the Y axis (positive values are on the top and right)
	// dc.InvertY()

	// Set the coordinate midpoint to the middle of the image
	dc.Translate(imageWidth/2, imageHeight/2)

	return dc
}

// saveImages saves the given context to a png at the given path
func saveImage(dc *gg.Context, path string) {
	dc.SavePNG(path)
}

// drawStar draws the given stars to the given context
func drawStar(dc *gg.Context, star structs.Star) {
	dc.DrawPoint(star.C.X, star.C.Y, 2)
	dc.Fill()
	dc.Stroke()
}

// vectorLength calculates the length of the given vector
func vectorLength(force structs.Force) float64 {
	return math.Sqrt(math.Pow(force.X, 2) + math.Pow(force.Y, 2))
}

func drawForce(dc *gg.Context, star structs.Star) {
	// controll the length of the vector
	var scalingFactor float64 = 40

	// Move the "cursor" to the start position of the vector
	dc.MoveTo(star.C.X, star.C.Y)

	// calculate the length of the vector
	vecLength := vectorLength(star.F)

	// Use a sigmoid function to generate useful values for coloring the vectors according to their
	// strength
	var val = 1.0 / (1.0 + math.Exp(-vecLength))

	// Set the color to a blue / red
	dc.SetRGB(val, 0, 1-val)

	// trace the Vector
	FxUnit := star.F.X / math.Abs(vecLength)
	FyUnit := star.F.Y / math.Abs(vecLength)
	dc.LineTo(star.C.X+(FxUnit*scalingFactor), star.C.Y+(FyUnit*scalingFactor))

	// css
	dc.SetLineWidth(2)

	// And finally: DRAW (stroke) the vector
	dc.Stroke()
}

// drawStars draws all the stars in the given slice to the given context
func drawStars(dc *gg.Context, slice []structs.Star) {
	// draw all the forces in the given slice
	for _, star := range slice {
		drawForce(dc, star)
	}

	dc.SetRGB(1, 1, 1)

	// draw all the stars in the given slice
	for _, star := range slice {
		drawStar(dc, star)
	}
}

// Slice draws the stars and the forces acting on them and saves the result to the given path
func Slice(slice []structs.Star, path string) {

	fmt.Printf("%-60s", "Plot init")
	// initialize the plot
	dc := initializePlot()
	fmt.Printf("Done\n")

	fmt.Printf("%-60s", "Drawing the Stars")
	// draw all the stars in the given slice
	drawStars(dc, slice)
	fmt.Printf("Done\n")

	fmt.Printf("%-60s", "Saving image")
	// save the plot to the given path
	saveImage(dc, path)
	fmt.Printf("Done\n")
}
