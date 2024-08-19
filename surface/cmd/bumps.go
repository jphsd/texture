//go:build ignore
// +build ignore

package main

import (
	"fmt"
	col "image/color"

	g2d "github.com/jphsd/graphics2d"
	gi "github.com/jphsd/graphics2d/image"
	"github.com/jphsd/texture"
	"github.com/jphsd/texture/color"
	"github.com/jphsd/texture/surface"
)

func main() {
	width, height := 800, 800

	nlf := texture.NewNLCircle1()
	wf := texture.NewNLWave([]float64{60}, []*texture.NonLinear{nlf}, false, true)
	rg := texture.NewRadialGradient(wf)
	xfm := g2d.Translate(-60, -60)
	xg := texture.NewTransform(rg, xfm)
	tf := texture.NewTiler(xg, []float64{120, 120})
	nm := texture.NewNormal(tf, 20, 20, 1, 1)

	// directional light
	dlight := surface.NewDirectional(col.White, []float64{-1, -1, 1})

	// surface
	material := &myMaterial{
		color.NewFRGBA(col.Black),                  // Emissive
		color.NewFRGBA(col.White),                  // Ambient
		color.NewFRGBA(col.RGBA{0, 0xff, 0, 0xff}), // Diffuse
		color.NewFRGBA(col.RGBA{0xff, 0, 0, 0xff}), // Specular
		10, // Shininess
		0,  // Roughness
	}
	surf := &surface.Surface{surface.DefaultAmbient, []surface.Light{dlight}, material, nm}
	rvals := []float64{0, 0.01, 0.05, 0.1, 0.5, 1}
	for i, r := range rvals {
		material.Roughness = r
		img := texture.NewRGBA(width, height, surf, 0, 0, 1, 1)
		gi.SaveImage(img, fmt.Sprintf("bumps%d", i))
	}
}

type myMaterial struct {
	Emissive, Ambient, Diffuse, Specular color.FRGBA
	Shininess                            float64
	Roughness                            float64
}

func (m *myMaterial) Eval2(x, y float64) (color.FRGBA, color.FRGBA, color.FRGBA, color.FRGBA, float64, float64) {
	return m.Emissive, m.Ambient, m.Diffuse, m.Specular, m.Shininess, m.Roughness
}
