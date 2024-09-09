package texture

import (
	"image/color"
	"math"
)

type ColorToGray struct {
	Name string
	Src  ColorField
}

func NewColorToGray(src ColorField) *ColorToGray {
	return &ColorToGray{"ColorToGray", src}
}

func (c *ColorToGray) Eval2(x, y float64) float64 {
	col := c.Src.Eval2(x, y)
	g := color.Gray16Model.Convert(col).(color.Gray16)
	v := float64(g.Y)
	v /= 0xffff
	return v*2 - 1
}

type ColorSelect struct {
	Name string
	Src  ColorField
	Chan int
}

func NewColorSelect(src ColorField, ch int) *ColorSelect {
	return &ColorSelect{"ColorSelect", src, ch}
}

func (c *ColorSelect) Eval2(x, y float64) float64 {
	r, g, b, a := c.Src.Eval2(x, y).RGBA()
	var col uint32
	switch c.Chan {
	default:
		fallthrough
	case 0: // Red
		col = r
	case 1: // Green
		col = g
	case 2: // Blue
		col = b
	case 3: // Alpha
		col = a
	}
	v := float64(col)
	v /= 0xffff
	return v*2 - 1
}

const oneOverPi = 1 / math.Pi

// Direction converts a VectorField to a Field based on the vector's direction in the XY plane.
type Direction struct {
	Name string
	Src  VectorField
}

func NewDirection(src VectorField) *Direction {
	return &Direction{"Direction", src}
}

// Eval2 implements the Field interface.
func (d *Direction) Eval2(x, y float64) float64 {
	vec := d.Src.Eval2(x, y)
	theta := math.Atan2(vec[1], vec[0])
	return theta * oneOverPi
}

// Magnitude converts a VectorField to a Field based on the vector's magnitude.
type Magnitude struct {
	Name  string
	Src   VectorField
	Scale float64
}

func NewMagnitude(src VectorField, scale float64) *Magnitude {
	return &Magnitude{"Magnitude", src, scale}
}

// Eval2 implements the Field interface. Always >= 0
func (m *Magnitude) Eval2(x, y float64) float64 {
	v := m.Src.Eval2(x, y)
	var s float64
	for _, f := range v {
		s += f * f
	}
	r := math.Sqrt(s) * m.Scale
	return clamp(r)
}

// Select converts a VectorField to a field by selecting one of its components.
type Select struct {
	Name  string
	Src   VectorField
	Chan  int
	Scale float64
}

func NewSelect(src VectorField, ch int, scale float64) *Select {
	return &Select{"Select", src, ch, scale}
}

// Eval2 implements the Field interface.
func (s *Select) Eval2(x, y float64) float64 {
	v := s.Src.Eval2(x, y)[s.Chan] * s.Scale
	return clamp(v)
}

// Weighted converts a VectorField to a field by selecting one of its components.
type Weighted struct {
	Name    string
	Src     VectorField
	Weights []float64
}

func NewWeighted(src VectorField, w []float64) *Weighted {
	return &Weighted{"Weighted", src, w}
}

// Eval2 implements the Field interface.
func (w *Weighted) Eval2(x, y float64) float64 {
	v := w.Src.Eval2(x, y)
	var s float64
	n, j := len(v), len(w.Weights)
	if j < n {
		n = j // Implicitly set additional weights to 0
	}
	for i := 0; i < n; i++ {
		s += v[i] * w.Weights[i]
	}
	return clamp(s)
}
