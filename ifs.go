package texture

import (
	g2d "github.com/jphsd/graphics2d"
)

// IFS represents a collection of affine transforms comprising an iterated function system. The combiner function
// is usually the union of the sub evaluations but other combiners can be used, for example the ones used in
// Fractal.
type IFS struct {
	Src   Field
	Xfms  []*g2d.Aff3 // Inverses of the IFS contractive affine transformations
	CFunc func(...float64) float64
	FFunc func(float64) float64
	Itr   int
}

// NewIFS returns a new instance of IFS. Note that the number of sub evaluations required is the number of
// transforms to the power of the number of iterations per evaluation.
func NewIFS(src Field, xfms []*g2d.Aff3, comb func(...float64) float64, itr int) *IFS {
	invxfms := make([]*g2d.Aff3, len(xfms))
	for i, xfm := range xfms {
		invxfms[i], _ = xfm.InverseOf()
	}
	return &IFS{src, invxfms, comb, nil, itr}
}

// Eval2 implements the Field interface.
func (f *IFS) Eval2(x, y float64) float64 {
	q := make([][]float64, 0, len(f.Xfms))
	q = append(q, []float64{x, y})
	for i := 0; i < f.Itr; i++ {
		nq := make([][]float64, 0, len(q))
		for _, pt := range q {
			for _, xfm := range f.Xfms {
				nq = append(nq, xfm.Apply(pt)[0])
			}
		}
		q = nq
	}

	nv := make([]float64, len(q))
	for i, pt := range q {
		nv[i] = f.Src.Eval2(pt[0], pt[1])
	}

	v := clamp(f.CFunc(nv...))

	if f.FFunc == nil {
		return v
	}
	return f.FFunc(v)
}

// Union returns the largest of the input values.
func Union(values ...float64) float64 {
	res := -1.0
	for _, v := range values {
		if v > res {
			res = v
		}
	}
	return res
}
