//go:build ignore

package main

import (
	"flag"
	"github.com/jphsd/graphics2d/image"
	"github.com/jphsd/texture"
)

func main() {
	width, height := 1024, 1024

	// Read in Sample.jpg for texture.Sample use in random.go
	flag.Parse()
	args := flag.Args()
	img, err := image.ReadImage(args[0])
	if err != nil {
		panic(err)
	}

	cf1 := texture.NewImage(img, texture.LinearInterp)
	f1 := texture.NewColorToGray(cf1)
	f2 := texture.NewCache(f1, 1, 10000)

	f3 := texture.NewEdge(f2, texture.Z4Support(1, 1))

	fimg := texture.NewTextureGray16(width, height, f3, 0, 0, 1, 1, false)
	image.SaveImage(fimg, "morph")
}
