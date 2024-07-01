package texture

// Distort applies a distortion to a source field based on the values it contains.
// Based on Musgrave's VLNoise3() in Ch 8 of Texturing and Modeling
type Distort struct {
	Name           string
	Src            Field
	OffsX, OffsY   float64
	NOffsX, NOffsY float64
	Distortion     float64
}

// NewDistort creates a new Distort instance.
func NewDistort(src Field, dist float64) *Distort {
	return &Distort{"Distort", src, 0.5, 0.5, 0, 3.333, dist}
}

// Eval2 implements the Field interface.
func (d *Distort) Eval2(x, y float64) float64 {
	x += d.OffsX
	y += d.OffsY
	nx, ny := d.Src.Eval2(x+d.NOffsX, y+d.NOffsX), d.Src.Eval2(x+d.NOffsY, y+d.NOffsY)
	nx *= d.Distortion
	ny *= d.Distortion
	v := d.Src.Eval2(nx, ny)
	return v
}
