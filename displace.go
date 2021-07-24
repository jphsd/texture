package texture

type Displace struct {
	Src, SrcX, SrcY Field
	OffsX, OffsY    float64
	ScaleX, ScaleY  float64
	Indep           bool
}

func NewDisplace(in, dx, dy Field, scale float64) *Displace {
	return &Displace{in, dx, dy, 0, 0, scale, scale, false}
}

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

func NewDisplace2(in Field, disp VectorField, scale float64) *Displace2 {
	return &Displace2{in, disp, 0, 1, 0, 0, scale, scale, false}
}

type Displace2 struct {
	Src            Field
	DistSrc        VectorField
	SelX, SelY     int
	OffsX, OffsY   float64
	ScaleX, ScaleY float64
	Indep          bool
}

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
