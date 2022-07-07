package texture

import (
	"image/color"
)

// Component composes a value, vector and color field.
type Component struct {
	Name   string
	Value  Field
	Vector VectorField
	Color  ColorField
}

// NewComponent takes a source field and returns a component instance. The three colors control
// the color values at t = 0, 0.5, 1 in the supplied nonlinear color field created from the
// nonlinear and lerp functions. The vector field is produced by calculating the normals of the
// source field, scaled by bscale.
func NewComponent(src Field,
	c1, c2, c3 color.Color,
	lerp LerpType,
	bscale float64) *Component {
	normals := NewNormal(src, bscale, bscale, 1, 1)
	cc := NewColorConv(src, c1, c3, []color.Color{c2}, []float64{0.5}, lerp)
	return &Component{"Component", src, normals, cc}
}
