//go:build ignore

package main

import (
	"flag"
	g2dimg "github.com/jphsd/graphics2d/image"
	"github.com/jphsd/texture"
	"image"
	"os"

	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
	_ "image/jpeg"
	_ "image/png"
)

func main() {
	width, height := 1024, 1024

	// Read in Sample.jpg for texture.Sample use in random.go
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

	cf1 := texture.NewImage(img, texture.LinearInterp)
	f1 := texture.NewColorToGray(cf1)
	f2 := texture.NewCache(f1, 1, 10000)

	f3 := texture.NewEdge(f2, texture.Z4Support(1, 1))

	cf2 := texture.NewColorGray(f3)
	fimg := texture.NewRGBA(width, height, cf2, 0, 0, 1, 1)
	g2dimg.SaveImage(fimg, "morph")
}
