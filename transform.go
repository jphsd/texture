package texture

import (
	g2d "github.com/jphsd/graphics2d"
	"image/color"
)

type Transform struct {
	Src Field
	Xfm *g2d.Aff3
}

func (t *Transform) Eval2(x, y float64) float64 {
	pt := []float64{x, y}
	pts := t.Xfm.Apply(pt)
	return t.Src.Eval2(pts[0][0], pts[0][1])
}

type TransformVF struct {
	Src VectorField
	Xfm *g2d.Aff3
}

func (t *TransformVF) Eval2(x, y float64) []float64 {
	pt := []float64{x, y}
	pts := t.Xfm.Apply(pt)
	return t.Src.Eval2(pts[0][0], pts[0][1])
}

type TransformCF struct {
	Src ColorField
	Xfm *g2d.Aff3
}

func (t *TransformCF) Eval2(x, y float64) color.Color {
	pt := []float64{x, y}
	pts := t.Xfm.Apply(pt)
	return t.Src.Eval2(pts[0][0], pts[0][1])
}
