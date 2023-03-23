package color

import "image/color"

// FRGBA represents a non-premultiplied RGBA tuple in floating point [0,1] such that the tuple
// can be added to, scaled and multiplied.
type FRGBA struct {
	R, G, B, A float64
}

// NewFRGBA returns a new FRGB using the supplied color.
func NewFRGBA(col color.Color) FRGBA {
	r, g, b, a := col.RGBA()
	if a == 0 {
		return FRGBA{0, 0, 0, 0}
	}

	// r, g, and b are premultiplied - undo
	fr, fg, fb, fa := float64(r)/0xffff, float64(g)/0xffff, float64(b)/0xffff, float64(a)/0xffff
	return FRGBA{fr / fa, fg / fa, fb / fa, fa}
}

// RGBA implements the RGBA function from the Color interface.
func (c FRGBA) RGBA() (uint32, uint32, uint32, uint32) {
	// r, g, b premultiplied
	r, g, b, a := c.R*c.A*0xffff, c.G*c.A*0xffff, c.B*c.A*0xffff, c.A*0xffff
	return uint32(r), uint32(g), uint32(b), uint32(a)
}

// FRGBAModel for conversion of a color to FRGBA.
var FRGBAModel color.Model = color.ModelFunc(frgbaModel)

func frgbaModel(col color.Color) color.Color {
	frgba, ok := col.(FRGBA)
	if !ok {
		frgba = NewFRGBA(col)
	}
	return frgba
}

// Add returns the addition of c1 to the color (with clamping).
func (c FRGBA) Add(c1 FRGBA) FRGBA {
	r, g, b := c.R+c1.R, c.G+c1.G, c.B+c1.B
	if r < 0 {
		r = 0
	} else if r > 1 {
		r = 1
	}
	if g < 0 {
		g = 0
	} else if g > 1 {
		g = 1
	}
	if b < 0 {
		b = 0
	} else if b > 1 {
		b = 1
	}
	return FRGBA{r, g, b, c.A}
}

// Prod returns the product of c1 with the color.
func (c FRGBA) Prod(c1 FRGBA) FRGBA {
	return FRGBA{c.R * c1.R, c.G * c1.G, c.B * c1.B, c.A}
}

// Scale returns the color scaled by the value (except alpha).
func (c FRGBA) Scale(v float64) FRGBA {
	return FRGBA{c.R * v, c.G * v, c.B * v, c.A}
}

// IsBlack returns true if the color is black.
func (c FRGBA) IsBlack() bool {
	return c.R < 0.0001 && c.G < 0.0001 && c.B < 0.0001
}
