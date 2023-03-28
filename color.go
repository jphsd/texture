package texture

import (
	g2dcol "github.com/jphsd/graphics2d/color"
	tcol "github.com/jphsd/texture/color"
	"image/color"
	"math"
)

// ColorGray contains the field to use in the color evaluation and produces a grayscale color.
type ColorGray struct {
	Name string
	Src  Field
}

func NewColorGray(src Field) *ColorGray {
	return &ColorGray{"ColorGray", src}
}

// Eval2 implements the ColorField interface.
func (c *ColorGray) Eval2(x, y float64) color.Color {
	t := (c.Src.Eval2(x, y) + 1) / 2
	return color.Gray16{uint16(t * 0xffff)}
}

// ColorSinCos contains the field used in the color evaluation. The color produced depends
// on the mode and color space.
type ColorSinCos struct {
	Name string
	Src  Field
	Mode int
	HSL  bool
}

func NewColorSinCos(src Field, mode int, hsl bool) *ColorSinCos {
	return &ColorSinCos{"ColorSinCos", src, mode, hsl}
}

// Eval2 implements the ColorField interface.
func (c *ColorSinCos) Eval2(x, y float64) color.Color {
	t := c.Src.Eval2(x, y) * math.Pi
	var v1, v2, v3 float64
	switch c.Mode {
	default:
		v1, v2, v3 = (1+math.Sin(t))/2, (1+math.Cos(t))/2, 0.5
	case 1:
		v1, v2, v3 = (1+math.Sin(t))/2, 0.5, (1+math.Cos(t))/2
	case 2:
		v1, v2, v3 = (1+math.Cos(t))/2, (1+math.Sin(t))/2, 0.5
	case 3:
		v1, v2, v3 = (1+math.Cos(t))/2, 0.5, (1+math.Sin(t))/2
	case 4:
		v1, v2, v3 = 0.5, (1+math.Sin(t))/2, (1+math.Cos(t))/2
	case 5:
		v1, v2, v3 = 0.5, (1+math.Cos(t))/2, (1+math.Sin(t))/2
	}
	if c.HSL {
		return g2dcol.HSL{v1, v2, v3, 1}
	}
	return tcol.FRGBA{v1, v2, v3, 1}
}

// ColorConv combines the field with a color interpolator.
type ColorConv struct {
	Name   string
	Src    Field
	Colors []color.Color
	TVals  []float64
	Lerp   LerpType
}

func NewColorConv(src Field, start, end color.Color, cols []color.Color, tvals []float64, lerp LerpType) *ColorConv {
	if tvals == nil {
		tvals = []float64{0, 1}
		cols = []color.Color{start, end}
	} else {
		cl, tl := len(cols), len(tvals)
		if cl > tl {
			cols = cols[:tl]
		} else if tl > cl {
			tvals = tvals[:cl]
		}
		// Rewrite cols and tvals with start and end values
		nt := []float64{0}
		nt = append(nt, tvals...)
		nt = append(nt, 1)
		tvals = nt
		nc := []color.Color{start}
		nc = append(nc, cols...)
		nc = append(nc, end)
		cols = nc
	}
	return &ColorConv{"ColorConv", src, cols, tvals, lerp}
}

// Eval2 implements the ColorField interface.
func (c *ColorConv) Eval2(x, y float64) color.Color {
	t := (c.Src.Eval2(x, y) + 1) / 2

	var i int
	for i = 1; i < len(c.TVals) && t > c.TVals[i]; i++ {
	}
	if i == len(c.TVals) {
		return c.Colors[i-1]
	}
	nt := (t - c.TVals[i-1]) / (c.TVals[i] - c.TVals[i-1])
	c1, c2 := c.Colors[i-1], c.Colors[i]

	switch c.Lerp {
	default:
		fallthrough
	case LerpRGBA:
		return g2dcol.ColorRGBALerp(nt, c1, c2)
	case LerpHSL:
		return g2dcol.ColorHSLLerp(nt, c1, c2)
	case LerpHSLs:
		return g2dcol.ColorHSLLerpS(nt, c1, c2)
	}
}

// ColorFields uses the four field sources to form {R, G, B, A} or {H, S, L, A}.
type ColorFields struct {
	Name string
	Src1 Field
	Src2 Field
	Src3 Field
	Src4 Field
	HSL  bool
}

func NewColorFields(src1, src2, src3, src4 Field, hsl bool) *ColorFields {
	return &ColorFields{"ColorFields", src1, src2, src3, src4, hsl}
}

// Eval2 implements the ColorField interface.
func (c *ColorFields) Eval2(x, y float64) color.Color {
	a := 1.0
	if c.Src4 != nil {
		a = (c.Src4.Eval2(x, y) + 1) / 2
	}
	v1, v2, v3 := (c.Src1.Eval2(x, y)+1)/2, (c.Src2.Eval2(x, y)+1)/2, (c.Src3.Eval2(x, y)+1)/2
	if c.HSL {
		return g2dcol.HSL{v1, v2, v3, a}
	}
	return color.NRGBA{uint8(v1 * 0xff), uint8(v2 * 0xff), uint8(v3 * 0xff), uint8(a * 0xff)}
}

// ColorVector uses the three values from a vector field to populate either R, G, B or H, S, L.
// Alpha is set to opaque.
type ColorVector struct {
	Name string
	Src  VectorField
	HSL  bool
}

func NewColorVector(src VectorField, hsl bool) *ColorVector {
	return &ColorVector{"ColorVector", src, hsl}
}

// Eval2 implements the ColorField interface.
func (c *ColorVector) Eval2(x, y float64) color.Color {
	v := c.Src.Eval2(x, y)
	v1, v2, v3 := (v[0]+1)/2, (v[1]+1)/2, (v[2]+1)/2
	a := 1.0
	if c.HSL {
		return g2dcol.HSL{v1, v2, v3, a}
	}
	return color.NRGBA{uint8(v1 * 0xff), uint8(v2 * 0xff), uint8(v3 * 0xff), uint8(a * 0xff)}
}

// LerpType defines the type of lerp -
type LerpType int

// Constants for lerp types.
const (
	LerpRGBA LerpType = iota
	LerpHSL
	LerpHSLs
)

// ColorBlend takes colors from the two color field sources and blends them based on the value
// from the third source using the specified color lerp.
type ColorBlend struct {
	Name string
	Src1 ColorField
	Src2 ColorField
	Src3 Field
	Lerp LerpType
}

func NewColorBlend(src1, src2 ColorField, src3 Field, lerp LerpType) *ColorBlend {
	return &ColorBlend{"ColorrBlend", src1, src2, src3, lerp}
}

// Eval2 implements the ColorField interface.
func (c *ColorBlend) Eval2(x, y float64) color.Color {
	switch c.Lerp {
	default:
		fallthrough
	case LerpRGBA:
		return g2dcol.ColorRGBALerp((c.Src3.Eval2(x, y)+1)/2, c.Src1.Eval2(x, y), c.Src2.Eval2(x, y))
	case LerpHSL:
		return g2dcol.ColorHSLLerp((c.Src3.Eval2(x, y)+1)/2, c.Src1.Eval2(x, y), c.Src2.Eval2(x, y))
	case LerpHSLs:
		return g2dcol.ColorHSLLerpS((c.Src3.Eval2(x, y)+1)/2, c.Src1.Eval2(x, y), c.Src2.Eval2(x, y))
	}
}

// ColorSubstitute returns a color from either Src1 or Src2 depending on if the value
// from the third source is between the supplied start and end values.
type ColorSubstitute struct {
	Name string
	Src1 ColorField
	Src2 ColorField
	Src3 Field
	A, B float64
}

func NewColorSubstitute(src1, src2 ColorField, src3 Field, a, b float64) *ColorSubstitute {
	return &ColorSubstitute{"ColorSubstitute", src1, src2, src3, a, b}
}

// Eval2 implements the ColorField interface.
func (c *ColorSubstitute) Eval2(x, y float64) color.Color {
	v := c.Src3.Eval2(x, y)
	if v < c.A || v > c.B {
		return c.Src1.Eval2(x, y)
	}
	return c.Src2.Eval2(x, y)
}

// Color lerps; t [0,1] => color.Color

// ColorRGBABiLerp calculates the color value at u, v [0,1] given colors at [0,0], [1,0], [0,1], [1,1] in RGB space.
func ColorRGBABiLerp(u, v float64, colors []color.Color) color.RGBA {
	c1 := g2dcol.ColorRGBALerp(v, colors[0], colors[2])
	c2 := g2dcol.ColorRGBALerp(v, colors[1], colors[3])
	return g2dcol.ColorRGBALerp(u, c1, c2)
}
