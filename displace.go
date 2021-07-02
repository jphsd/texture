package texture

type Displace struct {
	Src, SrcX, SrcY Field
	ScaleX, ScaleY  float64
}

func (d *Displace) Eval2(x, y float64) float64 {
	dvx := d.SrcX.Eval2(x, y) * d.ScaleX
	dvy := d.SrcY.Eval2(x, y) * d.ScaleY
	return d.Src.Eval2(dvx, dvy)
}
