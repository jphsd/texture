package texture

import (
	"github.com/jphsd/datastruct"
)

// BlinnField implements the ideas from Blinn's 1982 paper, A Generalization of Algebraic Surface Drawing.
// In Blinn's paper, D() is the distance squared, F() is the exponential function and the values of A and B
// are determined from the desired metaball radius and blobiness.
type BlinnField struct {
	Name   string
	Points [][]float64
	A      []float64
	B      []float64
	D      func([]float64, []float64) float64
	F      func(float64) float64
	Scale  float64
	Offset float64
}

func NewBlinnField(points [][]float64, a, b []float64,
	d func([]float64, []float64) float64, f func(float64) float64,
	scale, offset float64) *BlinnField {
	return &BlinnField{"BlinnField", points, a, b, d, f, scale, offset}
}

// Eval2 implements the Field interface.
func (pf *BlinnField) Eval2(x, y float64) float64 {
	loc := []float64{x, y}
	var sum float64
	for i, pi := range pf.Points {
		d := pf.D(loc, pi)
		if i < len(pf.A) {
			d *= pf.A[i]
		}
		v := pf.F(d)
		if i < len(pf.B) {
			v *= pf.B[i]
		}
		sum += v
	}
	return clamp(sum*pf.Scale + pf.Offset)
}

// WorleyField implements the ideas from Worley's 1996 paper, A Cellular Texture Basis Function.
// In Worley's paper, the D() is the Euclidean distance function, F() is the identity function, all A
// are 1 and B determines the weights based on the distance from the point being queried.
type WorleyField struct {
	Name   string
	Points [][]float64
	A      []float64
	B      []float64
	F      func(float64) float64
	Scale  float64
	Offset float64
	kdtree *datastruct.KDTree
	np     int
}

func NewWorleyField(points [][]float64, a, b []float64,
	d func([]float64, []float64) float64, f func(float64) float64,
	scale, offset float64) *WorleyField {
	kdtree := datastruct.NewKDTree(2, points...)
	kdtree.Dist = d
	np := len(b)
	if np > len(points) {
		np = len(points)
	}
	return &WorleyField{"WorleyField", points, a, b, f, scale, offset, kdtree, np}
}

// Eval2 implements the Field interface.
func (pf *WorleyField) Eval2(x, y float64) float64 {
	loc := []float64{x, y}

	_, ds, _ := pf.kdtree.KNN(loc, pf.np)
	var sum float64
	for i, d := range ds {
		if i < len(pf.A) {
			d *= pf.A[i]
		}
		sum += pf.F(d) * pf.B[i]
	}
	return clamp(sum*pf.Scale + pf.Offset)
}
