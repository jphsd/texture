//go:build ignore

package main

import (
	"flag"
	"github.com/jphsd/graphics2d/image"
	"github.com/jphsd/texture"
)

func main() {
	vv := flag.Float64("v", 0, "threshold")
	flag.Parse()
	v := *vv
	name := flag.Args()[0]

	img, err := image.ReadImage(name)
	if err != nil {
		panic(err)
	}

	width, height := img.Bounds().Dx(), img.Bounds().Dy()

	cf := texture.NewImage(img, texture.LinearInterp)
	f2 := texture.NewColorToGray(cf)
	vf := texture.NewThresholdFilter(f2, v, -1, 1)
	img2 := texture.NewTextureGray16(width, height, vf, 0, 0, 1, 1, false)

	image.SaveImage(img2, "thresh")
}
