//go:build ignore

package main

import (
	"fmt"

	g2d "github.com/jphsd/graphics2d"
	col "github.com/jphsd/graphics2d/color"
	gi "github.com/jphsd/graphics2d/image"
	"github.com/jphsd/texture"
	"github.com/jphsd/texture/color"
	"github.com/jphsd/texture/surface"
)

func main() {
	width, height := 800, 800

	// Bump normals
	nlf := texture.NewNLCircle1()
	wf := texture.NewNLWave([]float64{60}, []*texture.NonLinear{nlf}, false, true)
	rg := texture.NewRadialGradient(wf)
	xfm := g2d.Translate(-60, -60)
	xg := texture.NewTransform(rg, xfm)
	tf := texture.NewTiler(xg, []float64{120, 120})
	nm := texture.NewNormal(tf, 20, 20, 1, 1)
	//nm := texture.DefaultNormal

	// directional lights
	lights := []surface.Directional{
		surface.NewDirectional(col.White, []float64{1, 1, 1}),
	}

	black := color.NewFRGBA(col.Black)
	white := color.NewFRGBA(col.White)

	// surface material
	material := &myMaterial{
		black,                     // Emissive
		white,                     // Ambient reflection
		color.NewFRGBA(col.Green), // Diffuse reflection
		color.NewFRGBA(col.Red),   // Specular reflection
		5,                         // Shininess
		0,                         // Roughness [0,1]
	}

	// surface
	surfAmb := surface.DefaultAmbient
	surf := &surface.Surface{surfAmb, lights, material, nm}

	// range of roughnesses
	rvals := []float64{0, 0.01, 0.05, 0.1, 0.3, 0.5, 0.7, 0.9, 1}
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
