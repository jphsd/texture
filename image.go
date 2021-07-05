package texture

import (
	g2dcol "github.com/jphsd/graphics2d/color"
	"image"
	"image/color"
	"image/draw"
	"math"
)

// Image holds the data to support a continuous bicubic interpolation over an image.
type Image struct {
	Image        *image.NRGBA
	Max          []float64
	LastX, LastY int
	HSL          bool
}

// NewImage sets up a new field with the supplied image. The image is converted to a {0, 0}
// offset image.
func NewImage(img image.Image) *Image {
	r := img.Bounds()
	w, h := r.Dx(), r.Dy()
	nr := image.Rect(0, 0, w, h)
	gimg := image.NewNRGBA(nr)
	draw.Draw(gimg, nr, img, r.Min, draw.Src)
	return &Image{gimg, []float64{float64(w), float64(h)}, w - 1, h - 1, false}
}

const (
	epsilon = 0.0000001
)

// Eval2 returns the values [-1,1] of the field at x, y.
func (f *Image) Eval2(x, y float64) []float64 {
	if x < 0 || x >= f.Max[0] || y < 0 || y >= f.Max[1] {
		return []float64{0, 0, 0, 1}
	}
	ix, iy := int(math.Floor(x+epsilon)), int(math.Floor(y+epsilon))
	rx, ry := x-float64(ix), y-float64(iy)
	p := f.getValues(ix, iy)
	v := BiCubic(rx, ry, p)
	// Scale from [0,1] to [-1,1]
	return []float64{v[0]*2 - 1, v[1]*2 - 1, v[2]*2 - 1, v[3]*2 - 1}
}

// Get 4x4 patch
func (f *Image) getValues(x, y int) [][][]float64 {
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
func (f *Image) getValue(x, y int) []float64 {
	var col color.Color
	if x < 0 {
		if y < 0 {
			col = f.Image.At(0, 0)
		} else if y > f.LastY {
			col = f.Image.At(0, f.LastY)
		} else {
			col = f.Image.At(0, y)
		}
	} else if x > f.LastX {
		if y < 0 {
			col = f.Image.At(f.LastX, 0)
		} else if y > f.LastY {
			col = f.Image.At(f.LastX, f.LastY)
		} else {
			col = f.Image.At(f.LastX, y)
		}
	} else if y < 0 {
		col = f.Image.At(x, 0)
	} else if y > f.LastY {
		col = f.Image.At(x, f.LastY)
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
	r := uint32(c.R)
	r = r<<8 | r
	rv := float64(r)
	rv /= 0xffff
	g := uint32(c.G)
	g = g<<8 | g
	gv := float64(g)
	gv /= 0xffff
	b := uint32(c.B)
	b = b<<8 | b
	bv := float64(b)
	bv /= 0xffff
	a := uint32(c.A)
	a = a<<8 | a
	av := float64(a)
	av /= 0xffff
	// Scale to [-1,1]
	return []float64{rv, gv, bv, av}
}

// Cubic calculates the value of f(t) for t in range [0,1] given the values of t at -1, 0, 1, 2 in p[]
// fitted to a cubic polynomial: f(t) = at^3 + bt^2 + ct + d. Clamped because it over/undershoots.
func Cubic(t float64, p []float64) float64 {
	v := p[1] + 0.5*t*(p[2]-p[0]+t*(2.0*p[0]-5.0*p[1]+4.0*p[2]-p[3]+t*(3.0*(p[1]-p[2])+p[3]-p[0])))
	if v < 0 {
		v = 0
	} else if v > 1 {
		v = 1
	}
	return v
}

// BiCubic uses Cubic to calculate the value of f(u,v) for u,v in range [0,1].
func BiCubic(u, v float64, p [][][]float64) []float64 {
	res := make([]float64, 4)
	for i := 0; i < 4; i++ {
		np := make([]float64, 4)
		np[0] = Cubic(v, p[i][0])
		np[1] = Cubic(v, p[i][1])
		np[2] = Cubic(v, p[i][2])
		np[3] = Cubic(v, p[i][3])
		res[i] = Cubic(u, np)
	}
	return res
}

// NewRGBA renders the texture into a new RGBA image.
func NewRGBA(width, height int, src ColorField, ox, oy, dx, dy float64) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	y := oy
	for r := 0; r < height; r++ {
		x := ox
		for c := 0; c < width; c++ {
			v := src.Eval2(x, y)
			img.Set(c, r, v)
			x += dx
		}
		y += dy
	}

	return img
}
