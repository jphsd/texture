package texture

import g2d "github.com/jphsd/graphics2d"

// Displace allows a source field to be evaluated locations determining by an offset and scaling of
// the input x, y coordinates taken from other sources. If Indep is true, then the mapped x, y
// is independent of the original x, y location.
type Displace struct {
	Src, SrcX, SrcY Field
	Xfm             *g2d.Aff3
	Indep           bool
}

// NewDisplace creates a new Dispalce instance using the same source for both x and y displacements.
func NewDisplace(in, dx, dy Field, scale float64) *Displace {
	xfm := g2d.NewAff3()
	xfm.Scale(scale, scale)
	return &Displace{in, dx, dy, xfm, false}
}

// Eval2 implements the Field interface.
func (d *Displace) Eval2(x, y float64) float64 {
	vp := []float64{d.SrcX.Eval2(x, y), d.SrcY.Eval2(x, y)}
	p := d.Xfm.Apply(vp)
	if !d.Indep {
		return d.Src.Eval2(x+p[0][0], y+p[0][1])
	}
	return d.Src.Eval2(p[0][0], p[0][1])
}

// Displace2 is similar to Displace but utilizes a vector field in place of the two value fields.
// SelX and SelY determine which vector component is used to distort the source location.
type Displace2 struct {
	Src        Field
	DistSrc    VectorField
	SelX, SelY int
	Xfm        *g2d.Aff3
	Indep      bool
}

// NewDisplace2 creates a new Displace2 instance.
func NewDisplace2(in Field, disp VectorField, scale float64) *Displace2 {
	xfm := g2d.NewAff3()
	xfm.Scale(scale, scale)
	return &Displace2{in, disp, 0, 1, xfm, false}
}

// Eval2 implements the Field interface.
func (d *Displace2) Eval2(x, y float64) float64 {
	dv := d.DistSrc.Eval2(x, y)
	vp := []float64{dv[d.SelX], dv[d.SelY]}
	p := d.Xfm.Apply(vp)
	if !d.Indep {
		return d.Src.Eval2(x+p[0][0], y+p[0][1])
	}
	return d.Src.Eval2(p[0][0], p[0][1])
}
