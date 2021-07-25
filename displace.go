package texture

// Displace allows a source field to be evaluated locations determing by an offset and scaling of
// the input x, y coordinates taken from other sources. If Indep is true, then the mapped x, y
// is independent of the original x, y location.
type Displace struct {
	Src, SrcX, SrcY Field
	OffsX, OffsY    float64
	ScaleX, ScaleY  float64
	Indep           bool
}

// NewDisplace creates a new Dispalce instance using the same source for both x and y displacements.
func NewDisplace(in, dx, dy Field, scale float64) *Displace {
	return &Displace{in, dx, dy, 0, 0, scale, scale, false}
}

// Eval2 implements the Field interface.
func (d *Displace) Eval2(x, y float64) float64 {
	if !d.Indep {
		dvx := x + d.SrcX.Eval2(x+d.OffsX, y+d.OffsY)*d.ScaleX
		dvy := y + d.SrcY.Eval2(x+d.OffsX, y+d.OffsY)*d.ScaleY
		return d.Src.Eval2(dvx, dvy)
	}
	dvx := d.SrcX.Eval2(x+d.OffsX, y+d.OffsY) * d.ScaleX
	dvy := d.SrcY.Eval2(x+d.OffsX, y+d.OffsY) * d.ScaleY
	return d.Src.Eval2(dvx, dvy)
}

// Displace2 is similar to Displace but utilizes a vector field in place of the two value fields.
// SelX and SelY determine which vector component is used to distort the source location.
type Displace2 struct {
	Src            Field
	DistSrc        VectorField
	SelX, SelY     int
	OffsX, OffsY   float64
	ScaleX, ScaleY float64
	Indep          bool
}

// NewDisplace2 creates a new Displace2 instance.
func NewDisplace2(in Field, disp VectorField, scale float64) *Displace2 {
	return &Displace2{in, disp, 0, 1, 0, 0, scale, scale, false}
}

// Eval2 implements the Field interface.
func (d *Displace2) Eval2(x, y float64) float64 {
	if !d.Indep {
		dv := d.DistSrc.Eval2(x+d.OffsX, y+d.OffsY)
		dvx := x + dv[d.SelX]*d.ScaleX
		dvy := y + dv[d.SelY]*d.ScaleY
		return d.Src.Eval2(dvx, dvy)
	}
	dv := d.DistSrc.Eval2(x+d.OffsX, y+d.OffsY)
	dvx := dv[d.SelX] * d.ScaleX
	dvy := dv[d.SelY] * d.ScaleY
	return d.Src.Eval2(dvx, dvy)
}
