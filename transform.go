package texture

import g2d "github.com/jphsd/graphics2d"

type Transform struct {
	Src Field
	Xfm *g2d.Aff3
}

func (t *Transform) Eval2(x, y float64) float64 {
	pt := []float64{x, y}
	pts := t.Xfm.Apply(pt)
	return t.Src.Eval2(pts[0][0], pts[0][1])
}
