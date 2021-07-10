package texture

type Displace struct {
	Src, SrcX, SrcY Field
	OffsX, OffsY    float64
	ScaleX, ScaleY  float64
}

func (d *Displace) Eval2(x, y float64) float64 {
	dvx := d.SrcX.Eval2(x+d.OffsX, y+d.OffsY) * d.ScaleX
	dvy := d.SrcY.Eval2(x+d.OffsX, y+d.OffsY) * d.ScaleY
	return d.Src.Eval2(dvx, dvy)
}

type Displace2 struct {
	Src            Field
	DistSrc        VectorField
	SelX, SelY     int
	OffsX, OffsY   float64
	ScaleX, ScaleY float64
}

func (d *Displace2) Eval2(x, y float64) float64 {
	dv := d.DistSrc.Eval2(x+d.OffsX, y+d.OffsY)
	dvx := dv[d.SelX] * d.ScaleX
	dvy := dv[d.SelY] * d.ScaleY
	return d.Src.Eval2(dvx, dvy)
}
