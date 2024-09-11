package texture_test

import (
	"fmt"
	"github.com/jphsd/graphics2d"
	"github.com/jphsd/graphics2d/image"
	"github.com/jphsd/texture"
	"image/color"
)

func Example_flames() {
	f := texture.NewPerlin(12345)

	fx := texture.NewPerlin(12346)

	fy := texture.NewPerlin(12347)

	f2 := texture.NewDisplace(f, fx, fy, 1)

	xfm := graphics2d.Scale(1, 0.3)
	f3 := texture.NewTransform(f2, xfm)

	img := texture.NewTextureGray16(600, 600, f3, 0, 0, .015, .015, false)
	// Colorize it
	c1, c2 := color.RGBA{0x19, 0, 0, 0xff}, color.RGBA{0xff, 0xff, 0x85, 0xff}
	stops := []int{
		64,
		128,
		160,
		250,
	}
	colors := []color.Color{
		color.RGBA{0x76, 0, 0, 0xff},
		color.RGBA{0xff, 0, 0, 0xff},
		color.RGBA{0xff, 0x7c, 0, 0xff},
		color.RGBA{0xff, 0xff, 0x7a, 0xff},
	}
	img2 := image.NewColorizer(img, c1, c2, stops, colors, false)
	image.SaveImage(img2, "Example_flames")
	fmt.Printf("Generated Example_flames")
	// Output: Generated Example_flames
}
