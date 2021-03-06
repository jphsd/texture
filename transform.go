package texture

import (
	g2d "github.com/jphsd/graphics2d"
	"image/color"
)

// Transform applies an affine transform to the values passed into the Eval2 function.
type Transform struct {
	Name string
	Src  Field
	Xfm  *g2d.Aff3
}

func NewTransform(src Field, xfm *g2d.Aff3) *Transform {
	return &Transform{"Transform", src, xfm}
}

// Eval2 implements the Field interface.
func (t *Transform) Eval2(x, y float64) float64 {
	pt := []float64{x, y}
	pts := t.Xfm.Apply(pt)
	return t.Src.Eval2(pts[0][0], pts[0][1])
}

// TransformVF applies an affine transform to the values passed into the Eval2 function.
type TransformVF struct {
	Name string
	Src  VectorField
	Xfm  *g2d.Aff3
}

func NewTransformVF(src VectorField, xfm *g2d.Aff3) *TransformVF {
	return &TransformVF{"TransformVF", src, xfm}
}

// Eval2 implements the VectorField interface.
func (t *TransformVF) Eval2(x, y float64) []float64 {
	pt := []float64{x, y}
	pts := t.Xfm.Apply(pt)
	return t.Src.Eval2(pts[0][0], pts[0][1])
}

// TransformCF applies an affine transform to the values passed into the Eval2 function.
type TransformCF struct {
	Name string
	Src  ColorField
	Xfm  *g2d.Aff3
}

func NewTransformCF(src ColorField, xfm *g2d.Aff3) *TransformCF {
	return &TransformCF{"TransformCF", src, xfm}
}

// Eval2 implements the ColorField interface.
func (t *TransformCF) Eval2(x, y float64) color.Color {
	pt := []float64{x, y}
	pts := t.Xfm.Apply(pt)
	return t.Src.Eval2(pts[0][0], pts[0][1])
}
