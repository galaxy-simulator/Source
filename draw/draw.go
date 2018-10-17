package draw

import (
	"../structs"
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
	dc.InvertY()

	// Set the coordinate midpoint to the middle of the image
	dc.Translate(imageWidth/2, imageHeight/2)

	return dc
}

// saveImages saves the given context to a png at the given path
func saveImage(dc *gg.Context, path string) {
	dc.SavePNG(path)
}

// drawStar draws the given stars to the given context
func drawStar(dc *gg.Context, star structs.Star2D) {
	// User the value below to controll how big the stars are overall
	var StarScalingFactor float64 = 1

	// Default Star Size
	S := 2.0

	// Calculate the Stars Size according to its Mass (Maximal size = 5)
	if star.M < 5e4 {
		S = float64(math.Ceil(star.M/4e4) * StarScalingFactor)
	} else {
		S = 5.0
	}

	// draw the star / point
	dc.DrawPoint(star.C.X/50, star.C.Y/50, S)
	dc.Fill()
	dc.Stroke()
}

func drawVelocity(dc *gg.Context, star structs.Star2D) {
	// scaling factor for a better view of the velocity difference
	// Use this value to control how long the vectors are drawn
	var scalingFactor float64 = 25

	// Move the "cursor" to the start position of the vector
	dc.MoveTo(star.C.X/50, star.C.Y/50)

	// calculate the length of the vector
	vecLength := star.V.GetLength()

	// Use a sigmoid function to generate useful values for coloring the vectors according to their
	// strength
	var val = 1.0 / (1.0 + math.Exp(-vecLength*scalingFactor/2))

	// Set the color to a blue / red
	dc.SetRGB(val, 0, 1-val)

	// calculate the direction vector
	FUnit := (&star.V).Divide(vecLength)

	// set end-position of the vector line
	dc.LineTo(star.C.X/50+(FUnit.X*scalingFactor), star.C.Y/50+(FUnit.Y*scalingFactor))

	// set line width
	dc.SetLineWidth(3)

	// And finally: DRAW (stroke) the vector
	dc.Stroke()
}

// drawStars draws all the stars in the given slice to the given context
func drawStars(dc *gg.Context, slice []structs.Star2D) {
	// draw all the velocity in the given slice
	for _, star := range slice {
		drawVelocity(dc, star)
	}

	dc.SetRGB(1, 1, 1)

	// draw all the stars in the given slice
	for _, star := range slice {
		drawStar(dc, star)
	}
}

// Slice draws the stars and the forces acting on them and saves the result to the given path
func Slice(slice []structs.Star2D, path string) {

	// initialize the plot
	dc := initializePlot()

	// draw all the stars in the given slice
	drawStars(dc, slice)

	dc.SetRGB(1, 1, 1)

	// drawing the 4 big stars as bigger white dots
	//dc.DrawCircle(600, 600, 5)
	//dc.DrawCircle(-600, 600, 5)
	//dc.DrawCircle(-600, 0, 5)
	//dc.DrawCircle(600, -600, 5)

	dc.Fill()

	// save the plot to the given path
	saveImage(dc, path)
}
