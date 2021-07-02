package texture

type Distort struct {
	Src            Field
	OffsX, OffsY   float64
	NOffsX, NOffsY float64
	Distortion     float64
}

func NewDistort(src Field, dist float64) *Distort {
	return &Distort{src, 0.5, 0.5, 0, 3.333, dist}
}

func (d *Distort) Eval2(x, y float64) float64 {
	x += d.OffsX
	y += d.OffsY
	nx, ny := d.Src.Eval2(x+d.NOffsX, y+d.NOffsX), d.Src.Eval2(x+d.NOffsY, y+d.NOffsY)
	nx *= d.Distortion
	ny *= d.Distortion
	return d.Src.Eval2(nx, ny)
}
