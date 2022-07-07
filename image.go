package texture

import (
	tcol "github.com/jphsd/texture/color"
	"image"
	"image/color"
)

// Image holds the data to support a continuous bicubic interpolation over an image.
type Image struct {
	Name  string
	image image.Image
	MinX  int
	LastX int
	MinY  int
	LastY int
}

// NewImage sets up a new field with the supplied image. The image is converted to a {0, 0}
// offset NRGBA image.
func NewImage(img image.Image) *Image {
	rect := img.Bounds()
	return &Image{"Image", img, rect.Min.X, rect.Max.X - 1, rect.Min.Y, rect.Max.Y - 1}
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
	c := BiCubic(rx, ry, p)
	return c
}

// Get 4x4 patch
func (f *Image) getValues(x, y int) [][]*tcol.FRGBA {
	res := make([][]*tcol.FRGBA, 4)
	for r, i := y-1, 0; r < y+3; r++ {
		res[i] = make([]*tcol.FRGBA, 4)
		for c, j := x-1, 0; c < x+3; c++ {
			res[i][j] = f.getValue(c, r)
			j++
		}
		i++
	}
	return res
}

// Get converted values as FRGBA, and handle edges
func (f *Image) getValue(x, y int) *tcol.FRGBA {
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

	fc, _ := tcol.FRGBAModel.Convert(col).(*tcol.FRGBA)
	return fc
}

// BiCubic uses Cubic to calculate the value of f(u,v) for u,v in range [0,1).
func BiCubic(u, v float64, p [][]*tcol.FRGBA) *tcol.FRGBA {
	row := make([]float64, 4)
	col := make([]float64, 4)

	// R
	for j := 0; j < 4; j++ {
		for i := 0; i < 4; i++ {
			col[i] = p[i][j].R
		}
		row[j] = bcclamp(Cubic(v, col))
	}
	r := bcclamp(Cubic(u, row))

	// G
	for j := 0; j < 4; j++ {
		for i := 0; i < 4; i++ {
			col[i] = p[i][j].G
		}
		row[j] = bcclamp(Cubic(v, col))
	}
	g := bcclamp(Cubic(u, row))

	// B
	for j := 0; j < 4; j++ {
		for i := 0; i < 4; i++ {
			col[i] = p[i][j].B
		}
		row[j] = bcclamp(Cubic(v, col))
	}
	b := bcclamp(Cubic(u, row))

	// A
	for j := 0; j < 4; j++ {
		for i := 0; i < 4; i++ {
			col[i] = p[i][j].A
		}
		row[j] = bcclamp(Cubic(v, col))
	}
	a := bcclamp(Cubic(u, row))

	return &tcol.FRGBA{r, g, b, a}
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
