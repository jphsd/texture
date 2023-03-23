package texture

import (
	tcol "github.com/jphsd/texture/color"
	"image"
	"image/color"
)

type Interp int

const (
	NearestInterp Interp = iota
	LinearInterp
	CubicInterp
	P3Interp
	P5Interp
)

// Image holds the data to support a continuous bicubic interpolation over an image.
type Image struct {
	Name   string
	image  image.Image
	MinX   int
	LastX  int
	MinY   int
	LastY  int
	Func   Interp
	interp func(float64, []float64) float64
}

// NewImage sets up a new field with the supplied image. The image is converted to a {0, 0}
// offset NRGBA image.
func NewImage(img image.Image, interp Interp) *Image {
	rect := img.Bounds()
	var f func(float64, []float64) float64
	switch interp {
	default:
		fallthrough
	case NearestInterp:
		f = Nearest
	case LinearInterp:
		f = Linear
	case CubicInterp:
		f = Cubic
	case P3Interp:
		f = P3
	case P5Interp:
		f = P5
	}
	return &Image{"Image", img, rect.Min.X, rect.Max.X - 1, rect.Min.Y, rect.Max.Y - 1, interp, f}
}

// Eval2 implements the ColorField interface.
func (f *Image) Eval2(x, y float64) color.Color {
	// Image.At is defined over the entire plane.
	if _, ok := f.image.(*image.Uniform); ok {
		return f.image.At(0, 0)
	}

	ix, iy := int(x), int(y)
	if ix < f.MinX || ix > f.LastX || iy < f.MinY || iy > f.LastY {
		return color.Black
	}
	rx, ry := x-float64(ix), y-float64(iy)
	p := f.getValues(ix, iy)
	c := f.biPatch(rx, ry, p)
	return c
}

// Get 4x4 patch
func (f *Image) getValues(x, y int) [][]tcol.FRGBA {
	res := make([][]tcol.FRGBA, 4)
	for r, i := y-1, 0; r < y+3; r++ {
		res[i] = make([]tcol.FRGBA, 4)
		for c, j := x-1, 0; c < x+3; c++ {
			res[i][j] = f.getValue(c, r)
			j++
		}
		i++
	}
	return res
}

// Get converted values as FRGBA, and handle edges
func (f *Image) getValue(x, y int) tcol.FRGBA {
	var col color.Color
	if x < f.MinX {
		if y < f.MinY {
			col = f.image.At(f.MinX, f.MinY)
		} else if y > f.LastY {
			col = f.image.At(f.MinX, f.LastY)
		} else {
			col = f.image.At(f.MinX, y)
		}
	} else if x > f.LastX {
		if y < f.MinY {
			col = f.image.At(f.LastX, f.MinY)
		} else if y > f.LastY {
			col = f.image.At(f.LastX, f.LastY)
		} else {
			col = f.image.At(f.LastX, y)
		}
	} else if y < f.MinY {
		col = f.image.At(x, f.MinY)
	} else if y > f.LastY {
		col = f.image.At(x, f.LastY)
	} else {
		col = f.image.At(x, y)
	}

	fc, _ := tcol.FRGBAModel.Convert(col).(tcol.FRGBA)
	return fc
}

// biPatch uses interp to calculate the value of f(u,v) for u,v in range [0,1).
func (f *Image) biPatch(u, v float64, p [][]tcol.FRGBA) tcol.FRGBA {
	row := make([]float64, 4)
	col := make([]float64, 4)

	// R
	for j := 0; j < 4; j++ {
		for i := 0; i < 4; i++ {
			col[i] = p[i][j].R
		}
		row[j] = bcclamp(f.interp(v, col))
	}
	r := bcclamp(f.interp(u, row))

	// G
	for j := 0; j < 4; j++ {
		for i := 0; i < 4; i++ {
			col[i] = p[i][j].G
		}
		row[j] = bcclamp(Cubic(v, col))
	}
	g := bcclamp(f.interp(u, row))

	// B
	for j := 0; j < 4; j++ {
		for i := 0; i < 4; i++ {
			col[i] = p[i][j].B
		}
		row[j] = bcclamp(f.interp(v, col))
	}
	b := bcclamp(f.interp(u, row))

	// A
	for j := 0; j < 4; j++ {
		for i := 0; i < 4; i++ {
			col[i] = p[i][j].A
		}
		row[j] = bcclamp(f.interp(v, col))
	}
	a := bcclamp(f.interp(u, row))

	return tcol.FRGBA{r, g, b, a}
}

func bcclamp(v float64) float64 {
	if v < 0 {
		return 0
	}
	if v > 1 {
		return 1
	}
	return v
}

// Cubic calculates the value of f(t) for t in range [0,1] given the values of t at -1, 0, 1, 2 in p[]
// fitted to a cubic polynomial: f(t) = at^3 + bt^2 + ct + d. Clamped because it over/undershoots.
// (From graphics2d/util/nlerp.go and https://www.paulinternet.nl/?page=bicubic)
func Cubic(t float64, p []float64) float64 {
	v := p[1] + 0.5*t*(p[2]-p[0]+t*(2.0*p[0]-5.0*p[1]+4.0*p[2]-p[3]+t*(3.0*(p[1]-p[2])+p[3]-p[0])))
	return v
}

// Linear calculates the value of f(t) for t in range [0,1] given the values of t at -1, 0, 1, 2 in p[]
// using linear interpolation.
func Linear(t float64, p []float64) float64 {
	return (1-t)*p[1] + t*p[2]
}

// Nearest calculates the value of f(t) for t in range [0,1] given the values of t at -1, 0, 1, 2 in p[]
// using the closest value to t.
func Nearest(t float64, p []float64) float64 {
	if t < 0.5 {
		return p[1]
	}
	return p[2]
}

// P3 calculates the value of f(t) for t in range [0,1] given the values of t at -1, 0, 1, 2 in p[]
// uses a cubic s-curve
func P3(t float64, p []float64) float64 {
	// first derivative is 0 at t = 0 or 1
	t = t * t * (3 - 2*t)
	t = t * t * t * (t*(t*6-15) + 10)
	return Linear(t, p)
}

// P5 calculates the value of f(t) for t in range [0,1] given the values of t at -1, 0, 1, 2 in p[]
// uses a quintic s-curve
func P5(t float64, p []float64) float64 {
	// first and second derivatives are 0 at t = 0 or 1
	t = t * t * t * (t*(t*6-15) + 10)
	return Linear(t, p)
}
