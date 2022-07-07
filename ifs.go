package texture

import (
	g2d "github.com/jphsd/graphics2d"
)

// IFS represents a collection of affine transforms comprising an iterated function system.
type IFS struct {
	Name string
	Dom  []float64
	Xfms []*g2d.Aff3 // Inverses of the IFS contractive affine transformations
	Itr  int
}

// NewIFS returns a new instance of IFS. Note that the number of sub evaluations required is the number of
// transforms to the power of the number of iterations per evaluation.
func NewIFS(dom []float64, xfms []*g2d.Aff3, itr int) *IFS {
	invxfms := make([]*g2d.Aff3, len(xfms))
	for i, xfm := range xfms {
		invxfms[i], _ = xfm.InverseOf()
	}
	return &IFS{"IFS", dom, invxfms, itr}
}

// Eval2 implements the Field interface.
func (f *IFS) Eval2(x, y float64) float64 {
	q := make([][]float64, 0, len(f.Xfms))
	q = append(q, []float64{x, y})
	for i := 0; i < f.Itr; i++ {
		nq := make([][]float64, 0, len(q))
		for _, pt := range q {
			for _, xfm := range f.Xfms {
				// Assumes the transforms are injective into dom
				np := xfm.Apply(pt)[0]
				if np[0] > 0 && np[0] < f.Dom[0] && np[1] > 0 && np[1] < f.Dom[1] {
					nq = append(nq, np)
				}
			}
		}
		if len(nq) == 0 {
			return -1
		}
		q = nq
	}
	return 1
}

// IFS represents a collection of affine transforms comprising an iterated function system.
type IFSCombiner struct {
	Name string
	Src1 Field
	Src2 Field
	Dom  []float64
	Xfms []*g2d.Aff3 // Inverses of the IFS contractive affine transformations
	Itr  int
}

// NewIFSCombiner returns a new instance of IFS. Note that the number of sub evaluations required is the number of
// transforms to the power of the number of iterations per evaluation.
func NewIFSCombiner(src1, src2 Field, dom []float64, xfms []*g2d.Aff3, itr int) *IFSCombiner {
	invxfms := make([]*g2d.Aff3, len(xfms))
	for i, xfm := range xfms {
		invxfms[i], _ = xfm.InverseOf()
	}
	return &IFSCombiner{"IFSCombiner", src1, src2, dom, invxfms, itr}
}

// Eval2 implements the Field interface.
func (f *IFSCombiner) Eval2(x, y float64) float64 {
	q := make([][]float64, 0, len(f.Xfms))
	q = append(q, []float64{x, y})
	for i := 0; i < f.Itr; i++ {
		nq := make([][]float64, 0, len(q))
		for _, pt := range q {
			for _, xfm := range f.Xfms {
				// Assumes the transforms are injective into dom
				np := xfm.Apply(pt)[0]
				if np[0] > 0 && np[0] < f.Dom[0] && np[1] > 0 && np[1] < f.Dom[1] {
					nq = append(nq, np)
				}
			}
		}
		if len(nq) == 0 {
			return f.Src2.Eval2(x, y)
		}
		q = nq
	}
	return f.Src1.Eval2(x, y)
}
