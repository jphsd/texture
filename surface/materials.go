package surface

import (
	"github.com/jphsd/texture/color"
	col "image/color"
)

// Material provides the At function to determine the emissive light, various reflectances and shininess at a location.
// Reflectances are ordered as ambient, diffuse and specular.
type Material interface {
	Eval2(x, y float64) (*color.FRGBA, *color.FRGBA, *color.FRGBA, *color.FRGBA, float64)
}

type defaultMaterial struct {
	Ambient, Diffuse *color.FRGBA
}

// DefaultMaterial describes a material with 0 emissivity, full white ambient and directional, and no specular
// components.
var DefaultMaterial = &defaultMaterial{color.NewFRGBA(col.White), color.NewFRGBA(col.White)}

// Eval2 implements the Eval2 function in the Material interface.
func (d *defaultMaterial) Eval2(x, y float64) (*color.FRGBA, *color.FRGBA, *color.FRGBA, *color.FRGBA, float64) {
	return nil, d.Ambient, d.Diffuse, nil, 0
}
