package texture

import (
	"image"
	"image/color"
	"image/draw"
	"math"
)

// Image holds the data to support a continuous bicubic interpolation over an image.
type Image struct {
	Image         *image.Gray16
	Max           []float64
	Width, Height int
	FFunc         func(float64) float64
}

// NewImage sets up a new field with the supplied image. The image is converted to a {0, 0}
// offset Gray16 image.
func NewImage(img image.Image) *Image {
	r := img.Bounds()
	w, h := r.Dx(), r.Dy()
	nr := image.Rect(0, 0, w, h)
	gimg := image.NewGray16(nr)
	draw.Draw(gimg, nr, img, r.Min, draw.Src)
	return &Image{gimg, []float64{float64(w), float64(h)}, w - 1, h - 1, nil}
}

// Eval2 returns the value [-1,1] of the field at x, y.
func (f *Image) Eval2(x, y float64) float64 {
	if x < 0 || x >= f.Max[0] || y < 0 || y >= f.Max[1] {
		return 0
	}
	nx, ny := math.Floor(x), math.Floor(y)
	rx, ry := x-nx, y-ny
	ix, iy := int(nx), int(ny)
	p := f.getValues(ix, iy)
	v := 0.0
	if f.FFunc == nil {
		v = f.BiCubic(rx, ry, p)
	} else {
		v = f.FFunc(f.BiCubic(rx, ry, p))
	}
	return v*2 - 1
}

// Get 4x4 patch
func (f *Image) getValues(x, y int) [][]float64 {
	res := make([][]float64, 4)
	for r, i := y-1, 0; r < y+3; r++ {
		res[i] = make([]float64, 4)
		for c, j := x-1, 0; c < x+3; c++ {
			res[i][j] = f.getValue(c, r)
			j++
		}
		i++
	}
	return res
}

// Get converted gray value and handle edges
func (f *Image) getValue(x, y int) float64 {
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
	gc, _ := col.(color.Gray16)
	v := float64(gc.Y)
	v /= 0xffff
	return v
}

// Cubic calculates the value of f(t) for t in range [0,1] given the values of t at -1, 0, 1, 2 in p[]
// fitted to a cubic polynomial: f(t) = at^3 + bt^2 + ct + d. Clamped because it over/undershoots.
func (f *Image) Cubic(t float64, p []float64) float64 {
	v := p[1] + 0.5*t*(p[2]-p[0]+t*(2.0*p[0]-5.0*p[1]+4.0*p[2]-p[3]+t*(3.0*(p[1]-p[2])+p[3]-p[0])))
	if v < 0 {
		v = 0
	} else if v > 1 {
		v = 1
	}
	return v
}

// BiCubic uses Cubic to calculate the value of f(u,v) for u,v in range [0,1].
func (f *Image) BiCubic(u, v float64, p [][]float64) float64 {
	np := make([]float64, 4)
	np[0] = f.Cubic(v, p[0])
	np[1] = f.Cubic(v, p[1])
	np[2] = f.Cubic(v, p[2])
	np[3] = f.Cubic(v, p[3])
	return f.Cubic(u, np)
}
