//go:build ignore

package main

import (
	"flag"
	"image"
	"math"
	"math/rand"
	"os"
	"time"

	g2d "github.com/jphsd/graphics2d"
	gi "github.com/jphsd/graphics2d/image"
	"github.com/jphsd/texture"

	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
	_ "image/jpeg"
	_ "image/png"
)

func main() {
	// Read in image file indicated in command line
	flag.Parse()
	args := flag.Args()
	f, err := os.Open(args[0])
	if err != nil {
		panic(err)
	}
	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	_ = f.Close()

	rand.Seed(int64(time.Now().Nanosecond()))

	width, height := 800, 800
	iw, ih := img.Bounds().Dx(), img.Bounds().Dy()
	sx, sy := 2*float64(iw)/float64(width), 2*float64(ih)/float64(height)

	// Make new field from img
	f1 := texture.NewImage(img)
	f2 := texture.NewTilerCF(f1, []float64{float64(iw), float64(ih)})
	xfm := g2d.NewAff3()
	xfm.Scale(sx, sy)
	xfm.RotateAbout(rand.Float64()*math.Pi*2, float64(width)/2, float64(height)/2)
	f3 := texture.NewTransformCF(f2, xfm)

	out := texture.NewRGBA(width, height, f3, 0, 0, 1, 1)
	gi.SaveImage(out, "image")
	texture.SaveJSON(f3, "image")
}
