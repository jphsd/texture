package texture

import "github.com/jphsd/graphics2d"

// Displace allows a source field to be evaluated at locations determined by an offset and scaling of
// the input x, y coordinates taken from other sources. If Indep is true, then the mapped x, y
// is independent of the original x, y location.
type Displace struct {
	Name            string
	Src, SrcX, SrcY Field
	Xfm             *graphics2d.Aff3
	Indep           bool
}

// NewDisplace creates a new Displace instance using the same transform for both x and y displacements.
func NewDisplace(in, dx, dy Field, scale float64) *Displace {
	xfm := graphics2d.NewAff3()
	xfm.Scale(scale, scale)
	return &Displace{"Displace", in, dx, dy, xfm, false}
}

// Eval2 implements the Field interface.
func (d *Displace) Eval2(x, y float64) float64 {
	vp := []float64{d.SrcX.Eval2(x, y), d.SrcY.Eval2(x, y)}
	p := d.Xfm.Apply(vp)[0]
	if !d.Indep {
		return d.Src.Eval2(x+p[0], y+p[1])
	}
	return d.Src.Eval2(p[0], p[1])
}

// Displace2 allows a source field to be evaluated at locations determined by axis independent transforms of
// the input x, y coordinates taken from other sources. If Indep is true, then the mapped x, y
// is independent of the original x, y location.
type Displace2 struct {
	Name            string
	Src, SrcX, SrcY Field
	XfmX, XfmY      *graphics2d.Aff3
	Indep           bool
}

// NewDisplace2 creates a new Displace2 instance using the same source for both x and y displacements.
func NewDisplace2(in, dx, dy Field, xfmx, xfmy *graphics2d.Aff3) *Displace2 {
	return &Displace2{"Displace2", in, dx, dy, xfmx, xfmy, false}
}

// Eval2 implements the Field interface.
func (d *Displace2) Eval2(x, y float64) float64 {
	px := d.XfmX.Apply([]float64{d.SrcX.Eval2(x, y), 0})[0]
	py := d.XfmY.Apply([]float64{0, d.SrcY.Eval2(x, y), 0})[0]
	if !d.Indep {
		return d.Src.Eval2(x+px[0], y+py[1])
	}
	return d.Src.Eval2(px[0], py[1])
}
