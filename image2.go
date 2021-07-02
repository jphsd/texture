package texture

import (
	g2dcol "github.com/jphsd/graphics2d/color"
	"image"
	"image/color"
	"image/draw"
	"math"
)

// Image2 holds the data to support a continuous bicubic interpolation over an image.
type Image2 struct {
	Image         *image.NRGBA
	Max           []float64
	Width, Height int
	HSL           bool
}

// NewImage2 sets up a new field with the supplied image. The image is converted to a {0, 0}
// offset image.
func NewImage2(img image.Image) *Image2 {
	r := img.Bounds()
	w, h := r.Dx(), r.Dy()
	nr := image.Rect(0, 0, w, h)
	gimg := image.NewNRGBA(nr)
	draw.Draw(gimg, nr, img, r.Min, draw.Src)
	return &Image2{gimg, []float64{float64(w), float64(h)}, w - 1, h - 1, false}
}

// Eval2 returns the value [0,1] of the field at x, y.
func (f *Image2) Eval2(x, y float64) []float64 {
	if x < 0 || x >= f.Max[0] || y < 0 || y >= f.Max[1] {
		return []float64{0, 0, 0, 1}
	}
	nx, ny := math.Floor(x), math.Floor(y)
	rx, ry := x-nx, y-ny
	ix, iy := int(nx), int(ny)
	p := f.getValues(ix, iy)
	v := f.BiCubic(rx, ry, p)
	return v
}

// Get 4x4 patch
func (f *Image2) getValues(x, y int) [][][]float64 {
	res := make([][][]float64, 4)
	for r, i := y-1, 0; r < y+3; r++ {
		res[i] = make([][]float64, 4)
		for c, j := x-1, 0; c < x+3; c++ {
			res[i][j] = f.getValue(c, r)
			j++
		}
		i++
	}
	return res
}

// Get converted values and handle edges
func (f *Image2) getValue(x, y int) []float64 {
	var col color.Color
	if x < 0 {
		if y < 0 {
			col = f.Image.At(0, 0)
		} else if y > f.Height {
			col = f.Image.At(0, f.Height)
		} else {
			col = f.Image.At(0, y)
		}
	} else if x > f.Width {
		if y < 0 {
			col = f.Image.At(f.Width, 0)
		} else if y > f.Height {
			col = f.Image.At(f.Width, f.Height)
		} else {
			col = f.Image.At(f.Width, y)
		}
	} else if y < 0 {
		col = f.Image.At(x, 0)
	} else if y > f.Height {
		col = f.Image.At(x, f.Height)
	} else {
		col = f.Image.At(x, y)
	}
	if f.HSL {
		// HSLA
		hsl := g2dcol.NewHSL(col)
		return []float64{hsl.H, hsl.S, hsl.L, hsl.A}
	}
	// NRGBA
	c, _ := col.(color.NRGBA)
	rv := float64(c.R | (c.R << 8))
	rv /= 0xffff
	gv := float64(c.G | (c.G << 8))
	gv /= 0xffff
	bv := float64(c.B | (c.B << 8))
	bv /= 0xffff
	av := float64(c.A | (c.A << 8))
	av /= 0xffff
	return []float64{rv, gv, bv, av}
}

// Cubic calculates the value of f(t) for t in range [0,1] given the values of t at -1, 0, 1, 2 in p[]
// fitted to a cubic polynomial: f(t) = at^3 + bt^2 + ct + d. Clamped because it over/undershoots.
func (f *Image2) Cubic(t float64, p []float64) float64 {
	v := p[1] + 0.5*t*(p[2]-p[0]+t*(2.0*p[0]-5.0*p[1]+4.0*p[2]-p[3]+t*(3.0*(p[1]-p[2])+p[3]-p[0])))
	if v < 0 {
		v = 0
	} else if v > 1 {
		v = 1
	}
	return v
}

// BiCubic uses Cubic to calculate the value of f(u,v) for u,v in range [0,1].
func (f *Image2) BiCubic(u, v float64, p [][][]float64) []float64 {
	res := make([]float64, 4)
	for i := 0; i < 4; i++ {
		np := make([]float64, 4)
		np[0] = f.Cubic(v, p[i][0])
		np[1] = f.Cubic(v, p[i][1])
		np[2] = f.Cubic(v, p[i][2])
		np[3] = f.Cubic(v, p[i][3])
		res[i] = f.Cubic(u, np)
	}
	return res
}
