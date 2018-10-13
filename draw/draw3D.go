package draw

import (
	"../llog"
	"../structs"
	"github.com/fogleman/ln/ln"
)

func drawStar3D(scene ln.Scene, star structs.Star) ln.Scene {
	starSize := 0.1
	oneCorner := ln.Vector{star.C.X - starSize, star.C.Y - starSize, star.C.Z - starSize}
	otherCorner := ln.Vector{star.C.X + starSize, star.C.Y + starSize, star.C.Z + starSize}

	scene.Add(ln.NewCube(oneCorner, otherCorner))

	return scene
}

func drawStars3D(scene ln.Scene, slice []structs.Star) ln.Scene {
	for _, star := range slice {
		scene = drawStar3D(scene, star)
	}

	return scene
}

func Slice3D(slice []structs.Star, path string) {
	// create a scene and add a single cube
	scene := ln.Scene{}

	llog.Good("Drawing the Stars")
	scene = drawStars3D(scene, slice)
	llog.Good("Done Drawing the Stars")

	// scene.Add(ln.NewCube(ln.Vector{-1, -1, -1}, ln.Vector{1, 1, 1}))

	// define camera parameters
	eye := ln.Vector{4, 3, 2}    // camera position
	center := ln.Vector{0, 0, 0} // camera looks at
	up := ln.Vector{0, 0, 1}     // up direction

	// define rendering parameters
	width := 1024.0  // rendered width
	height := 1024.0 // rendered height
	fovy := 50.0     // vertical field of view, degrees
	znear := 0.1     // near z plane
	zfar := 10.0     // far z plane
	step := 0.01     // how finely to chop the paths for visibility testing

	llog.Good("Configuring path")
	// compute 2D paths that depict the 3D scene
	paths := scene.Render(eye, center, up, width, height, fovy, znear, zfar, step)
	llog.Good("Done")

	// render the paths in an image
	llog.Good("Writing to png")
	paths.WriteToPNG(path, width, height)
}
