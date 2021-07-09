package texture

import (
	g2dcol "github.com/jphsd/graphics2d/color"
	"github.com/jphsd/graphics2d/util"
	"image/color"
	"math"
)

type Color struct {
	Src1, Src2, Src3, Src4 Field
	HSL                    bool
}

func (c *Color) Eval2(x, y float64) color.Color {
	if c.Src2 == nil && c.Src3 == nil {
		v := (c.Src1.Eval2(x, y) + 1) / 2
		v *= 0xffff
		return &color.Gray16{uint16(v)}
	}
	a := 1.0
	if c.Src4 != nil {
		a = (c.Src4.Eval2(x, y) + 1) / 2
	}
	v1, v2, v3 := (c.Src1.Eval2(x, y)+1)/2, (c.Src2.Eval2(x, y)+1)/2, (c.Src3.Eval2(x, y)+1)/2
	if c.HSL {
		return &g2dcol.HSL{v1, v2, v3, a}
	}
	return &color.NRGBA{uint8(v1 * 0xff), uint8(v2 * 0xff), uint8(v3 * 0xff), uint8(a * 0xff)}
}

type ColorVector struct {
	Src VectorField
	HSL bool
}

func (cv *ColorVector) Eval2(x, y float64) color.Color {
	v := cv.Src.Eval2(x, y)
	a := 1.0
	if cv.HSL {
		return &g2dcol.HSL{v[0], v[1], v[2], a}
	}
	return &color.NRGBA{uint8(v[0] * 0xff), uint8(v[1] * 0xff), uint8(v[2] * 0xff), uint8(a * 0xff)}
}

type ColorConv struct {
	Src  Field
	Conv *ColorNL
}

func (c *ColorConv) Eval2(x, y float64) color.Color {
	return c.Conv.ColorNLerp((c.Src.Eval2(x, y) + 1) / 2)
}

type ColorBlend struct {
	Src1, Src2 ColorField
	Src3       Field
	Lerp       func(float64, color.Color, color.Color) color.Color
}

func (c *ColorBlend) Eval2(x, y float64) color.Color {
	return c.Lerp((c.Src3.Eval2(x, y)+1)/2, c.Src1.Eval2(x, y), c.Src2.Eval2(x, y))
}

type ColorSubstitute struct {
	Src1, Src2 ColorField
	Src3       Field
	Start, End float64
}

func (c *ColorSubstitute) Eval2(x, y float64) color.Color {
	v := c.Src3.Eval2(x, y)
	if v < c.Start || v > c.End {
		return c.Src1.Eval2(x, y)
	}
	return c.Src2.Eval2(x, y)
}

// Color lerps; t [0,1] => color.Color

type ColorNL struct {
	NL     util.NonLinear
	Lerp   func(float64, color.Color, color.Color) color.Color
	Colors []color.Color
	TVals  []float64
}

func NewColorNL(start, end color.Color, colors []color.Color, tvals []float64, nl util.NonLinear, lerp func(float64, color.Color, color.Color) color.Color) *ColorNL {
	if tvals == nil {
		tvals = []float64{0, 1}
		colors = []color.Color{start, end}
	} else {
		nt := []float64{0}
		nt = append(nt, tvals...)
		nt = append(nt, 1)
		tvals = nt
		nc := []color.Color{start}
		nc = append(nc, colors...)
		nc = append(nc, end)
		colors = nc
	}
	n := len(tvals)
	if n > len(colors) {
		n = len(colors)
	}
	nt := make([]float64, n)
	for i, v := range tvals {
		nt[i] = nl.Transform(v)
	}
	return &ColorNL{nl, lerp, colors, nt}
}

// ColorNLerp returns the color at t given the nonlinear function and color/t points.
func (c *ColorNL) ColorNLerp(t float64) color.Color {
	t = c.NL.Transform(t)
	var i int
	for i = 1; i < len(c.TVals) && t > c.TVals[i]; i++ {
	}
	if i == len(c.TVals) {
		return c.Colors[i-1]
	}
	nt := (t - c.TVals[i-1]) / (c.TVals[i] - c.TVals[i-1])
	return c.Lerp(nt, c.Colors[i-1], c.Colors[i])
}

// Color Lerps

// ColorRGBALerp calculates the color value at t [0,1] given a start and end color in RGB space.
func ColorRGBALerp(t float64, start, end color.Color) color.Color {
	rs, gs, bs, as := start.RGBA() // uint32 [0,0xffff]
	re, ge, be, ae := end.RGBA()
	rt := uint32(math.Floor((1-t)*float64(rs) + t*float64(re) + 0.5))
	gt := uint32(math.Floor((1-t)*float64(gs) + t*float64(ge) + 0.5))
	bt := uint32(math.Floor((1-t)*float64(bs) + t*float64(be) + 0.5))
	at := uint32(math.Floor((1-t)*float64(as) + t*float64(ae) + 0.5))
	rt >>= 8 // uint32 [0,0xff]
	gt >>= 8
	bt >>= 8
	at >>= 8
	return &color.RGBA{uint8(rt), uint8(gt), uint8(bt), uint8(at)}
}

// ColorHSLLerp calculates the color value at t [0,1] given a start and end color in HSL space.
func ColorHSLLerp(t float64, start, end color.Color) color.Color {
	cs, ce := g2dcol.NewHSL(start), g2dcol.NewHSL(end)
	ht := (1-t)*cs.H + t*ce.H // Will never cross 1:0
	st := (1-t)*cs.S + t*ce.S
	lt := (1-t)*cs.L + t*ce.L
	at := (1-t)*cs.A + t*ce.A
	return &g2dcol.HSL{ht, st, lt, at}
}

// ColorHSLLerpS calculates the color value at t [0,1] given a start and end color in HSL space.
// Differs from ColorHSLLerp in that the shortest path for hue is taken.
func ColorHSLLerpS(t float64, start, end color.Color) color.Color {
	cs, ce := g2dcol.NewHSL(start), g2dcol.NewHSL(end)
	hd := ce.H - cs.H
	ht := 0.0
	// Handle hue being circular
	if hd > 0.5 || hd < -0.5 {
		h := ce.H - 1
		ht = (1-t)*cs.H + t*h
		if ht < 0 {
			ht += 1
		}
	} else {
		ht = (1-t)*cs.H + t*ce.H
	}
	st := (1-t)*cs.S + t*ce.S
	lt := (1-t)*cs.L + t*ce.L
	at := (1-t)*cs.A + t*ce.A
	return &g2dcol.HSL{ht, st, lt, at}
}
