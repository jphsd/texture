package surface

import (
	"github.com/jphsd/texture/color"
	col "image/color"
)

// Material provides the At function to determine the emissive light, various reflectances, shininess and
// roughness at a location. Reflectances are ordered as ambient, diffuse and specular.
type Material interface {
	Eval2(x, y float64) (color.FRGBA, color.FRGBA, color.FRGBA, color.FRGBA, float64, float64)
}

type defaultMaterial struct {
	Emissive, Ambient, Diffuse, Specular color.FRGBA
}

// DefaultMaterial describes a material with 0 emissivity, full white ambient and directional, and no specular
// components.
var DefaultMaterial = defaultMaterial{
	color.FRGBA{},             // Emissive
	color.NewFRGBA(col.White), // Ambient
	color.NewFRGBA(col.White), // Diffuse
	color.FRGBA{}}             // Specular

// Eval2 implements the Material interface.
func (d defaultMaterial) Eval2(x, y float64) (color.FRGBA, color.FRGBA, color.FRGBA, color.FRGBA, float64, float64) {
	return d.Emissive, d.Ambient, d.Diffuse, d.Specular, 0, 0
}
