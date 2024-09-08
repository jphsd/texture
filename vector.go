package texture

import (
	"github.com/jphsd/texture/color"
	"math"
)

// VectorFields produces a vector field from a slice of fields
type VectorFields struct {
	Name string
	Srcs []Field
}

func NewVectorFields(srcs ...Field) *VectorFields {
	return &VectorFields{"VectorFields", srcs}
}

// Eval2 implements the VectorField interface.
func (v *VectorFields) Eval2(x, y float64) []float64 {
	res := make([]float64, len(v.Srcs))
	for i, f := range v.Srcs {
		res[i] = f.Eval2(x, y)
	}
	return res
}

// VectorColor uses the three values from a color field to populate it.
type VectorColor struct {
	Name string
	Src  ColorField
}

func NewVectorColor(src ColorField) *VectorColor {
	return &VectorColor{"VectorColor", src}
}

// Eval2 implements the VectorField interface.
func (v *VectorColor) Eval2(x, y float64) []float64 {
	c := color.NewFRGBA(v.Src.Eval2(x, y))
	res := make([]float64, 3)
	res[0] = c.R*2 - 1
	res[1] = c.G*2 - 1
	res[2] = c.B*2 - 1
	res[3] = c.A*2 - 1

	return res
}

// Normal provides a VectorField calculated from a Field using the finite difference method.
type Normal struct {
	Name     string
	Src      Field
	SDx, SDy float64
	Dx, Dy   float64
}

// NewNormal returns a new instance of Normal.
func NewNormal(src Field, sx, sy, dx, dy float64) *Normal {
	return &Normal{"Normal", src, sx / (2 * dx), sy / (2 * dy), dx, dy}
}

// Eval2 implements the VectorField interface.
func (n *Normal) Eval2(x, y float64) []float64 {
	dx := n.Src.Eval2(x-n.Dx, y) - n.Src.Eval2(x+n.Dx, y)
	dy := n.Src.Eval2(x, y-n.Dy) - n.Src.Eval2(x, y+n.Dy)
	dx *= n.SDx
	dy *= n.SDy
	div := 1 / math.Sqrt(dx*dx+dy*dy+1)
	return []float64{dx * div, dy * div, div}
}

// UnitVector provides a unit vector (i.e. magnitude = 1) version of a VectorField.
type UnitVector struct {
	Name string
	Src  VectorField
}

// Eval2 implements the VectorField interface.
func (u *UnitVector) Eval2(x, y float64) []float64 {
	v := u.Src.Eval2(x, y)
	var s float64
	for _, f := range v {
		s += f * f
	}
	s = 1 / math.Sqrt(s)
	res := make([]float64, len(v))
	for i, f := range v {
		res[i] = f * s
	}
	return res
}
