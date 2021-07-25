package texture

import (
	"github.com/jphsd/graphics2d/util"
	"image/color"
)

// Component composes a value, vector and color field.
type Component struct {
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
	nl util.NonLinear,
	lerp func(float64, color.Color, color.Color) color.Color,
	bscale float64) *Component {
	normals := NewNormal(src, bscale, bscale, 1, 1)
	cnl := NewColorNL(c1, c3, []color.Color{c2}, []float64{0.5}, nl, lerp)
	return &Component{src, normals, &ColorConv{src, cnl}}
}
