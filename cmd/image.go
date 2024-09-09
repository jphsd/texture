//go:build ignore

package main

import (
	"flag"
	g2d "github.com/jphsd/graphics2d"
	"github.com/jphsd/graphics2d/image"
	"github.com/jphsd/texture"
	"math"
	"math/rand"
)

func main() {
	// Read in image file indicated in command line
	flag.Parse()
	args := flag.Args()

	img, err := image.ReadImage(args[0])
	if err != nil {
		panic(err)
	}

	width, height := 800, 800
	iw, ih := img.Bounds().Dx(), img.Bounds().Dy()
	sx, sy := 2*float64(iw)/float64(width), 2*float64(ih)/float64(height)

	// Make new field from img
	f1 := texture.NewImage(img, texture.LinearInterp)
	f2 := texture.NewTilerCF(f1, []float64{float64(iw), float64(ih)})
	xfm := g2d.NewAff3()
	xfm.Scale(sx, sy)
	xfm.RotateAbout(rand.Float64()*math.Pi*2, float64(width)/2, float64(height)/2)
	f3 := texture.NewTransformCF(f2, xfm)

	out := texture.NewTextureRGBA(width, height, f3, 0, 0, 1, 1, false)
	image.SaveImage(out, "image")
	texture.SaveJSON(f3, "image")
}
