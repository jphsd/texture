package texture

import (
	"github.com/jphsd/graphics2d/util"
	"math"
)

// Generator is used to construct a one dimensional pattern with a fixed wavelength, center and phase, rotated
// by theta. It has an optional filter. The generator function takes a value in the range [0,1] and returns a
// value in [-1,1].
type Generator struct {
	Lambda       float64 // [1,...)
	Center       float64 // [0,1]
	Phase        float64 // [0,1]
	GFunc        func(float64) float64
	FFunc        func(float64) float64
	CosTh, SinTh float64
}

// NewGenerator constructs a new Generator instance with wavelength lambda and rotation theta using the
// supplied generator function.
func NewGenerator(lambda, theta float64, f func(float64) float64) *Generator {
	if lambda < 1 {
		lambda = 1
	}
	// Snap to quad
	ct := math.Cos(theta)
	if closeTo(0, ct) {
		ct = 0
	} else if closeTo(1, ct) {
		ct = 1
	} else if closeTo(-1, ct) {
		ct = -1
	}
	st := math.Sin(theta)
	if closeTo(0, st) {
		st = 0
	} else if closeTo(1, st) {
		st = 1
	} else if closeTo(-1, st) {
		st = -1
	}
	return &Generator{lambda, 0.5, 0, f, nil, ct, st}
}

// Eval2 implements the Field interface.
func (g *Generator) Eval2(x, y float64) float64 {
	v := x*g.CosTh + y*g.SinTh
	if g.FFunc == nil {
		return g.GFunc(g.VtoT(v))
	}
	return g.FFunc(g.GFunc(g.VtoT(v)))
}

// VtoT converts a value in (-inf,inf) to [0,1] based on the generator's orientation, lambda and phase values.
func (g *Generator) VtoT(v float64) float64 {
	// Vs Div and Floor ...
	for v < 0 {
		v += g.Lambda
	}
	for v > g.Lambda {
		v -= g.Lambda
	}
	t := v/g.Lambda + g.Phase
	if t > 1 {
		t -= 1
	}
	if t <= g.Center {
		return t * 0.5 / g.Center
	}
	return 0.5*(t-g.Center)/(1-g.Center) + 0.5
}

// FlatGF is a flat generator function and returns a fixed value.
type FlatGF struct {
	Val float64
}

// Flat returns a fixed value regardless of t.
func (fgf *FlatGF) Flat(t float64) float64 {
	return fgf.Val
}

// Sin returns a sine wave (offset by -90 degrees)
func Sin(t float64) float64 {
	t *= 2 * math.Pi
	t -= math.Pi / 2
	return math.Sin(t)
}

// Square returns a square wave.
func Square(t float64) float64 {
	if t > 0.5 {
		return 1
	}
	return -1
}

// Triangle returns a triangle wave.
func Triangle(t float64) float64 {
	v := 0.0
	if t < 0.5 {
		v = t * 2
	} else {
		v = (1 - t) * 2
	}
	return v*2 - 1
}

// Saw returns a saw wave.
func Saw(t float64) float64 {
	return t*2 - 1
}

// NLGF captures a non-linear function.
type NLGF struct {
	NL util.NonLinear
}

// NL1 uses the NLGF function twice (once up and once down) to make a wave form.
func (g *NLGF) NL1(t float64) float64 {
	v := 0.0
	if t < 0.5 {
		v = g.NL.Transform(2 * t)
	} else {
		v = g.NL.Transform(2 * (1 - t))
	}
	return v*2 - 1
}

// NL2GF captures two non-linear functions, one for up and the other for down.
type NL2GF struct {
	NLU, NLD util.NonLinear
}

// NL2 uses the NL2GF functions to make a wave form.
func (g *NL2GF) NL2(t float64) float64 {
	v := 0.0
	if t < 0.5 {
		v = g.NLU.Transform(2 * t)
	} else {
		v = g.NLD.Transform(2 * (1 - t))
	}
	return v*2 - 1
}

// Noise1DGF captures a scaled Perlin noise instance.
type Noise1DGF struct {
	Scale        float64
	Noise        *Perlin
	OffsX, OffsY float64
}

// NewNoise1DGF returns a new Noise1DGF instance.
func NewNoise1DGF(scale float64) *Noise1DGF {
	return &Noise1DGF{scale, NewPerlin(), 0, 0}
}

// Noise1D returns a wave derived from a Perlin noise function.
func (n *Noise1DGF) Noise1D(t float64) float64 {
	return n.Noise.Eval2(t*n.Scale+n.OffsX, n.OffsY)
}

func closeTo(a, b float64) bool {
	d := a - b
	if d < 0 {
		d = -d
	}

	return d < 0.000001
}
