package texture

import "github.com/jphsd/graphics2d"

// Region is used to specify a rectangular region rotated by theta.
type Region struct {
	Width  float64
	Height float64
	Xfm    *graphics2d.Aff3
}

// NewRegion returns a Region instance based on the origin, extant and rotation of the inputs.
func NewRegion(origin, extent []float64, theta float64) *Region {
	xfm := graphics2d.NewAff3()
	xfm.Rotate(-theta)
	xfm.Translate(-origin[0], -origin[1])
	return &Region{extent[0], extent[1], xfm}
}

// InRegion returns true if the point is within the region.
func (r *Region) InRegion(x, y float64) bool {
	pts := r.Xfm.Apply([]float64{x, y})
	nx, ny := pts[0][0], pts[0][1]
	return !(nx < 0 || nx > r.Width || ny < 0 || ny > r.Height)
}

// Box is used to specify a region within a field. The value of In is returned if within the region bounds. The
// value of Out otherwise.
type Box struct {
	*Region
	In, Out float64
}

// Eval2 implements the Field interface.
func (b *Box) Eval2(x, y float64) float64 {
	if b.InRegion(x, y) {
		return b.In
	}
	return b.Out
}

// Window is used to specify a region for which the source field value is returned. The
// value of Out is returned otherwise.
type Window struct {
	*Region
	Src Field
	Out float64
}

// Eval2 implements the Field interface.
func (w *Window) Eval2(x, y float64) float64 {
	if w.InRegion(x, y) {
		return w.Src.Eval2(x, y)
	}
	return w.Out
}

// Window2 is used to specify a region for which the first source field value is returned. The
// second source field value is returned otherwise.
type Window2 struct {
	*Region
	Src1, Src2 Field
}

// Eval2 implements the Field interface.
func (w *Window2) Eval2(x, y float64) float64 {
	if w.InRegion(x, y) {
		return w.Src1.Eval2(x, y)
	}
	return w.Src2.Eval2(x, y)
}
