package texture

import (
	"github.com/jphsd/graphics2d/util"
	"math"
)

type Generator struct {
	Lambda       float64 // [1,...)
	Center       float64 // [0,1]
	Phase        float64 // [0,1]
	GFunc        func(float64) float64
	FFunc        func(float64) float64
	CosTh, SinTh float64
}

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

func (g *Generator) Eval2(x, y float64) float64 {
	v := x*g.CosTh + y*g.SinTh
	if g.FFunc == nil {
		return g.GFunc(g.VtoT(v))
	}
	return g.FFunc(g.GFunc(g.VtoT(v)))
}

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

func Zero(t float64) float64 {
	return 0
}

func Sin(t float64) float64 {
	t *= 2 * math.Pi
	t -= math.Pi / 2
	return math.Sin(t)
}

func Square(t float64) float64 {
	if t > 0.5 {
		return 1
	}
	return -1
}

func Triangle(t float64) float64 {
	v := 0.0
	if t < 0.5 {
		v = t * 2
	} else {
		v = (1 - t) * 2
	}
	return v*2 - 1
}

func Saw(t float64) float64 {
	return t*2 - 1
}

type NLGF struct {
	NL util.NonLinear
}

func (g *NLGF) NL1(t float64) float64 {
	v := 0.0
	if t < 0.5 {
		v = g.NL.Transform(2 * t)
	} else {
		v = g.NL.Transform(2 * (1 - t))
	}
	return v*2 - 1
}

type NL2GF struct {
	NLU, NLD util.NonLinear
}

func (g *NL2GF) NL2(t float64) float64 {
	v := 0.0
	if t < 0.5 {
		v = g.NLU.Transform(2 * t)
	} else {
		v = g.NLD.Transform(2 * (1 - t))
	}
	return v*2 - 1
}

func closeTo(a, b float64) bool {
	d := a - b
	if d < 0 {
		d = -d
	}

	return d < 0.000001
}
