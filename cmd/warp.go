//go:build ignore

package main

import (
	g2d "github.com/jphsd/graphics2d"
	"github.com/jphsd/graphics2d/image"
	"github.com/jphsd/texture"
)

// Demonstrate texture warps
func main() {
	width, height := 1000, 1000

	f1 := texture.NewSquares(25)
	cf := texture.NewColorGray(f1)

	//rwf := IdentityWF{}
	//rwf := texture.NewRadialWF([]float64{500, 500}, 1,  1)
	//rwf := texture.NewRadialNLWF([]float64{500, 500}, texture.NewNLExponential(3), 300)
	//rwf := texture.NewRadialNLWF([]float64{500, 500}, texture.NewNLLogarithmic(3), 300)
	//rwf := texture.NewPinchXWF([]float64{500, 500}, 0.3, 0.002, 2)
	//rwf := texture.NewSwirlWF([]float64{500, 500}, -0.05)
	//rwf := texture.NewDrainWF([]float64{500, 500}, 3.1412, 250)
	//rwf := texture.NewRippleXWF(100, 20, 12.5)
	rwf := texture.NewRadialRippleWF([]float64{500, 500}, 100, 10, 0)
	//rwf := texture.NewRadialWiggleWF([]float64{500, 500}, 100, 0.1, 0)

	wf := texture.NewWarpCF(cf, rwf)

	img := texture.NewRGBA(width, height, wf, 0, 0, 1, 1)

	image.SaveImage(img, "warp")
}

type IdentityWF struct{}

func (iw IdentityWF) Eval(x, y float64) (float64, float64) {
	return x, y
}
