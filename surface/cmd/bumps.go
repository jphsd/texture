// +build ignore

package main

import (
	"image/color"

	gc "github.com/jphsd/graphics2d/color"
	gi "github.com/jphsd/graphics2d/image"
	"github.com/jphsd/graphics2d/util"
	text "github.com/jphsd/texture"
	surf "github.com/jphsd/texture/surface"
)

func main() {
	width, height := 400, 400

	nlf := &util.NLCircle2{}
	nm := text.NewNormal(text.NewNonLinear(60, 60, 0, nlf, 2), 20, 20, 1, 1)

	// directional light
	dlight := surf.NewDirectional(color.White, []float64{-1, -1, 1})

	// surface
	material := &myMaterial{
		gc.NewFRGBA(color.Black),                  // Emissive
		gc.NewFRGBA(color.White),                  // Ambient
		gc.NewFRGBA(color.RGBA{0, 0xff, 0, 0xff}), // Diffuse
		gc.NewFRGBA(color.RGBA{0xff, 0, 0, 0xff}), // Specular
		10, // Shininess
	}
	surface := &surf.Surface{surf.DefaultAmbient, []surf.Light{dlight}, material, nm}
	img := text.NewRGBA(width, height, surface, 0, 0, 1, 1)
	gi.SaveImage(img, "bumps")
}

type myMaterial struct {
	Emissive, Ambient, Diffuse, Specular *gc.FRGBA
	Shininess                            float64
}

func (m *myMaterial) Eval2(x, y float64) (*gc.FRGBA, *gc.FRGBA, *gc.FRGBA, *gc.FRGBA, float64) {
	return m.Emissive, m.Ambient, m.Diffuse, m.Specular, m.Shininess
}
