package texture

import (
	"math"
)

// FixedNormal provides a fixed VectorField.
type FixedNormal struct {
	Val []float64
}

// DefaultNormal describes the unit normal point straight up from the XY plane.
var DefaultNormal = &FixedNormal{[]float64{0, 0, 1}}

// Eval2 implements
func (n *FixedNormal) Eval2(x, y float64) []float64 {
	return n.Val
}

// Normal provides a VectorField calculated from a Field using the finite difference method.
type Normal struct {
	Source   Field
	SDx, SDy float64
	Dx, Dy   float64
}

// NewNormal returns a new instance of Normal.
func NewNormal(src Field, sx, sy, dx, dy float64) *Normal {
	return &Normal{src, sx / (2 * dx), sy / (2 * dy), dx, dy}
}

// Eval2 implements the VectorField interface.
func (n *Normal) Eval2(x, y float64) []float64 {
	dx := n.Source.Eval2(x-n.Dx, y) - n.Source.Eval2(x+n.Dx, y)
	dy := n.Source.Eval2(x, y-n.Dy) - n.Source.Eval2(x, y+n.Dy)
	dx *= n.SDx
	dy *= n.SDy
	div := 1 / math.Sqrt(dx*dx+dy*dy+1)
	return []float64{dx * div, dy * div, div}
}

// UnitNormal provides a normalized version of a VectorField.
type UnitNormal struct {
	Source VectorField
}

// Eval2 implements the VectorField interface.
func (u *UnitNormal) Eval2(x, y float64) []float64 {
	v := u.Source.Eval2(x, y)
	s := v[0]*v[0] + v[1]*v[1] + v[2]*v[2]
	s = 1 / math.Sqrt(s)
	return []float64{v[0] * s, v[1] * s, v[2] * s}
}

const oneOverPi = 1 / math.Pi

// Direction converts a VectorField to a Field based on the vector's direction.
type Direction struct {
	Source VectorField
	FFunc  func(float64) float64
}

// Eval2 implements the Field interface.
func (d *Direction) Eval2(x, y float64) float64 {
	normal := d.Source.Eval2(x, y)
	theta := math.Atan2(normal[1], normal[0])
	if d.FFunc == nil {
		return theta * oneOverPi
	}
	return d.FFunc(theta * oneOverPi)
}

// Magnitude converts a normalized VectorField to a Field based on the vector's magnitude.
type Magnitude struct {
	Source VectorField
	FFunc  func(float64) float64
}

// Eval2 implements the Field interface. Always >= 0
func (m *Magnitude) Eval2(x, y float64) float64 {
	normal := m.Source.Eval2(x, y)
	if m.FFunc == nil {
		return 1 - normal[2]
	}
	return m.FFunc(1 - normal[2])
}

// Select converts a VectorField to a field by selecting one of its components.
type Select struct {
	Src   VectorField
	Chan  int
	FFunc func(float64) float64
}

// Eval2 implements the Field interface.
func (s *Select) Eval2(x, y float64) float64 {
	v := s.Src.Eval2(x, y)[s.Chan]
	if s.FFunc == nil {
		return v
	}
	return s.FFunc(v)
}

// VectorCombine converts a VectorField to a Field using a combiner function.
type VectorCombine struct {
	Source VectorField
	CFunc  func(...float64) float64
	FFunc  func(float64) float64
}

// Eval2 implements the Field interface.
func (vb *VectorCombine) Eval2(x, y float64) float64 {
	v := vb.Source.Eval2(x, y)
	res := vb.CFunc(v...)
	if vb.FFunc == nil {
		return res
	}
	return vb.FFunc(res)
}
