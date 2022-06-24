package texture

import (
	"github.com/jphsd/graphics2d/util"
	"math"
)

// NonLinear is used to create a field of circles that uses a non-linear function to fill the circles.
type NonLinear struct {
	LambdaX, LambdaY float64 // [1,...)
	PhaseX, PhaseY   float64 // [0,1]
	OffsetX, OffsetY float64 // [0,1]
	FFunc            func(float64) float64
	CosTh, SinTh     float64
	NLFunc           util.NonLinear
	Dist, Inset      float64
}

// NewNonLinear creates a new instance of NonLinear. A circle/elipse is rendered using the supplied
// non-linear function and inset within a box of size lambdaX by lambdaY.
func NewNonLinear(lambdaX, lambdaY, theta float64, nl util.NonLinear, inset float64) *NonLinear {
	if lambdaX < 1 {
		lambdaX = 1
	}
	if lambdaY < 1 {
		lambdaY = 1
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
	if inset < 0 {
		inset = 0
	}
	if lambdaX > lambdaY {
		inset /= lambdaY
	} else {
		inset /= lambdaX
	}
	if inset > 0.5 {
		inset = 0.5
	}
	dist := 1 - 2*inset
	return &NonLinear{lambdaX, lambdaY, 0, 0, 0, 0, nil, ct, st, nl, dist, inset}
}

// Eval2 implements the Field interface.
func (nl *NonLinear) Eval2(x, y float64) float64 {
	u := x*nl.CosTh + y*nl.SinTh
	v := -x*nl.SinTh + y*nl.CosTh
	u, v = nl.XYtoUV(u, v)
	res := 0.0
	if u > nl.Inset && u < nl.Dist+nl.Inset && v > nl.Inset && v < nl.Dist+nl.Inset {
		// Within inset, rescale to [0,1]
		u, v = (u-nl.Inset)/nl.Dist, (v-nl.Inset)/nl.Dist
		dx, dy := 0.5-u, 0.5-v
		d := 2 * math.Sqrt(dx*dx+dy*dy)
		if d <= 1 {
			res = nl.NLFunc.Transform(1 - d)
		}
	}
	if nl.FFunc == nil {
		return res*2 - 1
	}
	return nl.FFunc(res*2 - 1)
}

// XYtoUV converts values in (-inf,inf) to [0,1] based on the generator's orientation, lambdas and phase values.
func (nl *NonLinear) XYtoUV(x, y float64) (float64, float64) {
	nx := 0
	for x < 0 {
		x += nl.LambdaX
		nx--
	}
	for x > nl.LambdaX {
		x -= nl.LambdaX
		nx++
	}
	ny := 0
	for y < 0 {
		y += nl.LambdaY
		ny--
	}
	for y > nl.LambdaY {
		y -= nl.LambdaY
		ny++
	}

	if !util.Equals(0, nl.OffsetX) {
		offs := float64(ny) * nl.OffsetX
		offs -= math.Floor(offs)
		if offs < 0 {
			offs = 1 - offs
		}
		u := x/nl.LambdaX + nl.PhaseX + offs
		for u > 1 {
			u -= 1
		}
		v := y/nl.LambdaY + nl.PhaseY
		if v > 1 {
			v -= 1
		}
		return u, v
	}
	u := x/nl.LambdaX + nl.PhaseX
	for u > 1 {
		u -= 1
	}
	offs := float64(nx) * nl.OffsetY
	offs -= math.Floor(offs)
	if offs < 0 {
		offs = 1 - offs
	}
	v := y/nl.LambdaY + nl.PhaseY + offs
	for v > 1 {
		v -= 1
	}
	return u, v
}

// NLTransform holds the parameters for a non-linear transformation mapping.
type NLTransform struct {
	Src    Field
	NLFunc util.NonLinear
	FFunc  func(float64) float64
}

// Eval2 implements the Field interface.
func (nf *NLTransform) Eval2(x, y float64) float64 {
	v := (nf.Src.Eval2(x, y) + 1) / 2
	v = nf.NLFunc.Transform(v)
	v = v*2 - 1

	if nf.FFunc == nil {
		return v
	}
	return nf.FFunc(v)
}
