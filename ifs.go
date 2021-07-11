package texture

import (
	g2d "github.com/jphsd/graphics2d"
)

type IFS struct {
	Src   Field
	Xfms  []*g2d.Aff3
	CFunc func(...float64) float64
	FFunc func(float64) float64
	Itr   int
}

// xfms are the inverses of the contractive affine transformations

func NewIFS(src Field, xfms []*g2d.Aff3, comb func(...float64) float64, itr int) *IFS {
	return &IFS{src, xfms, comb, nil, itr}
}

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

func Union(values ...float64) float64 {
	res := -1.0
	for _, v := range values {
		if v > res {
			res = v
		}
	}
	return res
}
