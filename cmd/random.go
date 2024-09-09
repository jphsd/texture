//go:build ignore

package main

import (
	"fmt"
	"github.com/jphsd/graphics2d/image"
	"github.com/jphsd/texture"
	"github.com/jphsd/texture/random"
)

func main() {
	width, height := 800, 800

	// Read in Sample.jpg for texture.Sample use in random.go
	img, err := image.ReadImage("Sample.jpg")
	if err != nil {
		panic(err)
	}
	random.Sample = img

	cnt := 0
	for cnt < 100 {
		name := fmt.Sprintf("%06d", cnt)
		cf := random.MakeColorField(6, 0)
		img := texture.NewTextureRGBA(width, height, cf, 0, 0, 1, 1, false)
		image.SaveImage(img, name)
		texture.SaveJSON(cf, name)
		cnt++
	}
}
