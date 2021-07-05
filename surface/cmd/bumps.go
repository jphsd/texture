// +build ignore

package main

import (
	col "image/color"

	gi "github.com/jphsd/graphics2d/image"
	"github.com/jphsd/graphics2d/util"
	"github.com/jphsd/texture"
	"github.com/jphsd/texture/color"
	"github.com/jphsd/texture/surface"
)

func main() {
	width, height := 400, 400

	nlf := &util.NLCircle2{}
	nm := texture.NewNormal(texture.NewNonLinear(60, 60, 0, nlf, 2), 20, 20, 1, 1)

	// directional light
	dlight := surface.NewDirectional(col.White, []float64{-1, -1, 1})

	// surface
	material := &myMaterial{
		color.NewFRGBA(col.Black),                  // Emissive
		color.NewFRGBA(col.White),                  // Ambient
		color.NewFRGBA(col.RGBA{0, 0xff, 0, 0xff}), // Diffuse
		color.NewFRGBA(col.RGBA{0xff, 0, 0, 0xff}), // Specular
		10, // Shininess
	}
	surf := &surface.Surface{surface.DefaultAmbient, []surface.Light{dlight}, material, nm}
	img := texture.NewRGBA(width, height, surf, 0, 0, 1, 1)
	gi.SaveImage(img, "bumps")
}

type myMaterial struct {
	Emissive, Ambient, Diffuse, Specular *color.FRGBA
	Shininess                            float64
}

func (m *myMaterial) Eval2(x, y float64) (*color.FRGBA, *color.FRGBA, *color.FRGBA, *color.FRGBA, float64) {
	return m.Emissive, m.Ambient, m.Diffuse, m.Specular, m.Shininess
}
