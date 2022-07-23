package texture

import (
	"image/color"
)

// Pixelate provides pixelation for a field.
type Pixelate struct {
	Name       string
	Src        Field
	Resolution float64
}

// NewPixelate creates a new Pixelate with the specified resolution.
func NewPixelate(src Field, resolution float64) *Pixelate {
	return &Pixelate{"Pixelate", src, resolution}
}

// Eval2 implements the Field interface.
func (p *Pixelate) Eval2(x, y float64) float64 {
	x, y = pixelInd(x, y, p.Resolution)
	res := p.Src.Eval2(x, y)
	return res
}

// PixelateCF provides pixelation for a color field.
type PixelateCF struct {
	Name       string
	Src        ColorField
	Resolution float64
}

// NewPixelateCF creates a new Pixelate with the specified resolution.
func NewPixelateCF(src ColorField, resolution float64) *PixelateCF {
	return &PixelateCF{"PixelateCF", src, resolution}
}

// Eval2 implements the Field interface.
func (p *PixelateCF) Eval2(x, y float64) color.Color {
	x, y = pixelInd(x, y, p.Resolution)
	res := p.Src.Eval2(x, y)
	return res
}

func pixelInd(x, y, res float64) (float64, float64) {
	oneovrres, halfres := 1/res, res/2
	x *= oneovrres
	if x < 0 {
		x -= 1
	}
	ix := int(x)
	y *= oneovrres
	if y < 0 {
		y -= 1
	}
	iy := int(y)
	// Actual value returned is the value at the pixel's center
	return float64(ix)*res + halfres, float64(iy)*res + halfres
}
