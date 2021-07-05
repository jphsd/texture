package surface

import (
	gc "github.com/jphsd/graphics2d/color"
	"image/color"
)

// Light provides the At function to determine the color (RGB in [0,1]), unit direction, distance and power of a
// light at a location.
// If the direction is nil then the light is treated as an ambient one and any distance and power values ignored.
// If the distance is -ve then this is treated as a directional light at infinity.
// Otherwise the light is treated as a point light source with the power falling as the inverse square of the
// distance from the light.
type Light interface {
	Eval2(x, y float64) (*gc.FRGBA, []float64, float64, float64)
}

// Ambient describes an ambient light source.
type Ambient struct {
	Color *gc.FRGBA
}

// DefaultAmbient is a low gray light.
var DefaultAmbient = &Ambient{gc.NewFRGBA(color.RGBA{10, 10, 10, 255})}

// NewAmbient returns a new ambient light source.
func NewAmbient(col color.Color) *Ambient {
	return &Ambient{gc.NewFRGBA(col)}
}

// Eval2 implements the Eval2 function of the Light interface.
func (a *Ambient) Eval2(x, y float64) (*gc.FRGBA, []float64, float64, float64) {
	return a.Color, nil, -1, 0
}

// Directional describes a directional light source. The direction is from the surface to the light, normalized.
type Directional struct {
	Color     *gc.FRGBA
	Direction []float64
}

// NewDirectional returns a new directional light source.
func NewDirectional(col color.Color, dir []float64) *Directional {
	dir = Norm(dir)
	return &Directional{gc.NewFRGBA(col), dir}
}

// Eval2 implements the Eval2 function of the Light interface.
func (d *Directional) Eval2(x, y float64) (*gc.FRGBA, []float64, float64, float64) {
	return d.Color, d.Direction, -1, 0
}
