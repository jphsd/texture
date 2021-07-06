package texture

import (
	"github.com/jphsd/graphics2d/util"
	"image/color"
)

type Component struct {
	Value  Field
	Vector VectorField
	Color  ColorField
}

func NewComponent(src Field,
	c1, c2, c3 color.Color,
	nl util.NonLinear,
	lerp func(float64, color.Color, color.Color) color.Color, bscale float64) *Component {
	normals := NewNormal(src, bscale, bscale, 1, 1)
	cnl := NewColorNL(c1, c3, []color.Color{c2}, []float64{0.5}, nl, lerp)
	return &Component{src, normals, &ColorConv{src, cnl}}
}
