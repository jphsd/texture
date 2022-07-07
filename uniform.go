package texture

import "image/color"

type Uniform struct {
	Name  string
	Value float64
}

func NewUniform(v float64) *Uniform {
	return &Uniform{"Uniform", v}
}

func (f *Uniform) Eval2(x, y float64) float64 {
	return f.Value
}

type UniformCF struct {
	Name  string
	Value color.Color
}

func NewUniformCF(v color.Color) *UniformCF {
	return &UniformCF{"UniformCF", v}
}

func (f *UniformCF) Eval2(x, y float64) color.Color {
	return f.Value
}

type UniformVF struct {
	Name  string
	Value []float64
}

func NewUniformVF(v []float64) *UniformVF {
	return &UniformVF{"UniformVF", v}
}

func (f *UniformVF) Eval2(x, y float64) []float64 {
	return f.Value
}

// DefaultNormal describes the unit normal point straight up from the XY plane.
var DefaultNormal = &UniformVF{"UniformVF", []float64{0, 0, 1}}
