package texture

import (
	"image/color"
	"math"
)

const (
	twoPi = math.Pi * 2
)

type WarpFunc interface {
	Eval(x, y float64) (float64, float64)
}

// Warp applies a deformation to the values passed into the Eval2 function.
type Warp struct {
	Name string
	Src  Field
	Func WarpFunc
}

func NewWarp(src Field, wf WarpFunc) *Warp {
	return &Warp{"Warp", src, wf}
}

// Eval2 implements the Field interface.
func (w *Warp) Eval2(x, y float64) float64 {
	x, y = w.Func.Eval(x, y)
	return w.Src.Eval2(x, y)
}

// WarpVF applies a deformation to the values passed into the Eval2 function.
type WarpVF struct {
	Name string
	Src  VectorField
	Func WarpFunc
}

func NewWarpVF(src VectorField, wf WarpFunc) *WarpVF {
	return &WarpVF{"WarpVF", src, wf}
}

// Eval2 implements the VectorField interface.
func (w *WarpVF) Eval2(x, y float64) []float64 {
	x, y = w.Func.Eval(x, y)
	return w.Src.Eval2(x, y)
}

// WarpCF applies a deformation to the values passed into the Eval2 function.
type WarpCF struct {
	Name string
	Src  ColorField
	Func WarpFunc
}

func NewWarpCF(src ColorField, wf WarpFunc) *WarpCF {
	return &WarpCF{"WarpCF", src, wf}
}

// Eval2 implements the ColorField interface.
func (w *WarpCF) Eval2(x, y float64) color.Color {
	x, y = w.Func.Eval(x, y)
	return w.Src.Eval2(x, y)
}

// RadialWF performs a scaled warp around Center for use in the above warp types.
type RadialWF struct {
	Name   string
	Center []float64 // Center of warp
	RScale float64   // Radial scale
	CScale float64   // Circumference scale
}

func NewRadialWF(c []float64, rs, cs float64) *RadialWF {
	return &RadialWF{"RadialWF", c, rs, cs}
}

// Eval implements the WarpFunc interface
func (wf *RadialWF) Eval(x, y float64) (float64, float64) {
	lx, ly := x-wf.Center[0], y-wf.Center[1]
	rr := math.Hypot(lx, ly) * wf.RScale
	rx := rr*math.Atan2(ly, lx)*wf.CScale + wf.Center[0]
	ry := rr + wf.Center[1]
	return rx, ry
}

// SwirlWF performs a swirl warp around Center in x for use in the above warp types.
// The cordinates are converted to polar (r, th) and th advanced by scale * r.
type SwirlWF struct {
	Name   string
	Center []float64 // Center of warp
	Scale  float64   // Scale factor
}

func NewSwirlWF(c []float64, s float64) *SwirlWF {
	return &SwirlWF{"SwirlWF", c, s / twoPi}
}

// Eval implements the WarpFunc interface
func (wf *SwirlWF) Eval(x, y float64) (float64, float64) {
	dx, dy := x-wf.Center[0], y-wf.Center[1]
	r, th := toPolar(dx, dy)
	th += r * wf.Scale
	dx, dy = toEuclidean(r, th)
	return wf.Center[0] + dx, wf.Center[1] + dy
}

// DrainWF performs a drain warp around Center in x for use in the above warp types.
// The x value is scaled about the central x value by |dy|^alpha*scale
type DrainWF struct {
	Name   string
	Center []float64 // Center of warp
	Scale  float64   // Scale factor
	Effct  float64   // Effect radius
}

func NewDrainWF(c []float64, s, e float64) *DrainWF {
	return &DrainWF{"DrainWF", c, s, e}
}

// Eval implements the WarpFunc interface
func (wf *DrainWF) Eval(x, y float64) (float64, float64) {
	dx, dy := x-wf.Center[0], y-wf.Center[1]
	r, th := toPolar(dx, dy)
	if r > wf.Effct {
		return x, y
	}
	th += (1 - r/wf.Effct) * wf.Scale
	dx, dy = toEuclidean(r, th)
	return wf.Center[0] + dx, wf.Center[1] + dy
}

// RadialNLWF performs a radius warp around Center based on an NL
type RadialNLWF struct {
	Name   string
	Center []float64  // Center of warp
	NL     *NonLinear // [0,1] => [0,1]
	Effct  float64    // Effect radius
}

func NewRadialNLWF(c []float64, nl *NonLinear, e float64) *RadialNLWF {
	return &RadialNLWF{"RadialNLWF", c, nl, e}
}

// Eval implements the WarpFunc interface
func (wf *RadialNLWF) Eval(x, y float64) (float64, float64) {
	dx, dy := x-wf.Center[0], y-wf.Center[1]
	r, th := toPolar(dx, dy)
	if r > wf.Effct {
		return x, y
	}
	t := r / wf.Effct
	tp := (wf.NL.Eval(t) + 1) / 2 // NL returns results in [-1,1]
	r = tp * wf.Effct
	dx, dy = toEuclidean(r, th)
	return wf.Center[0] + dx, wf.Center[1] + dy
}

// PinchXWF performs a pinched warp around Center in x for use in the above warp types.
// The x value is scaled about the central x value by |dy|^alpha*scale
type PinchXWF struct {
	Name   string
	Center []float64 // Center of warp
	Init   float64   // Initial scale
	Scale  float64   // Scale factor
	Alpha  float64   // Power factor
}

func NewPinchXWF(c []float64, i, s, a float64) *PinchXWF {
	return &PinchXWF{"PinchXWF", c, i, s, a}
}

// Eval implements the WarpFunc interface
func (wf *PinchXWF) Eval(x, y float64) (float64, float64) {
	dx, dy := x-wf.Center[0], y-wf.Center[1]
	dy *= wf.Scale
	if dy < 0 {
		dy = -dy
	}
	if !within(wf.Alpha, 1, 0.00001) {
		dy = math.Pow(dy, wf.Alpha)
	}
	dx *= 1 / (dy + wf.Init)
	return wf.Center[0] + dx, y
}

// RippleXWF performs a sin warp using the y value, lambda and offset and applies it to the x value, scaled
// by the amplitude.
type RippleXWF struct {
	Name   string
	Lambda float64
	Amplit float64
	Offset float64
}

func NewRippleXWF(l, a, o float64) *RippleXWF {
	return &RippleXWF{"RippleXWF", l, a, o}
}

// Eval implements the WarpFunc interface
func (wf *RippleXWF) Eval(x, y float64) (float64, float64) {
	_, l := MapValueToLambda(y+wf.Offset, wf.Lambda)
	l = l / wf.Lambda * 2 * math.Pi
	dx := math.Sin(l) * wf.Amplit
	return x + dx, y
}

// RadialRippleWF performs a sin warp using the r value, lambda and offset and applies it to the r value, scaled
// by the amplitude.
type RadialRippleWF struct {
	Name   string
	Center []float64
	Lambda float64
	Amplit float64
	Offset float64
}

func NewRadialRippleWF(c []float64, l, a, o float64) *RadialRippleWF {
	return &RadialRippleWF{"RadialRippleWF", c, l, a, o}
}

// Eval implements the WarpFunc interface
func (wf *RadialRippleWF) Eval(x, y float64) (float64, float64) {
	dx, dy := x-wf.Center[0], y-wf.Center[1]
	r, th := toPolar(dx, dy)
	_, l := MapValueToLambda(r+wf.Offset, wf.Lambda)
	l = l / wf.Lambda * 2 * math.Pi
	dr := math.Sin(l) * wf.Amplit
	return toEuclidean(r+dr, th)
}

// RadialWiggleWF performs a sin warp using the r value, lambda and offset and applies it to the th value, scaled
// by the amplitude.
type RadialWiggleWF struct {
	Name   string
	Center []float64
	Lambda float64
	Amplit float64
	Offset float64
}

func NewRadialWiggleWF(c []float64, l, a, o float64) *RadialWiggleWF {
	return &RadialWiggleWF{"RadialWiggleWF", c, l, a, o}
}

// Eval implements the WarpFunc interface
func (wf *RadialWiggleWF) Eval(x, y float64) (float64, float64) {
	dx, dy := x-wf.Center[0], y-wf.Center[1]
	r, th := toPolar(dx, dy)
	_, l := MapValueToLambda(r+wf.Offset, wf.Lambda)
	l = l / wf.Lambda * 2 * math.Pi
	dth := math.Sin(l) * wf.Amplit
	return toEuclidean(r, th+dth)
}

func toPolar(dx, dy float64) (float64, float64) {
	return math.Hypot(dx, dy), math.Atan2(dy, dx)
}

func toEuclidean(r, th float64) (float64, float64) {
	return r * math.Cos(th), r * math.Sin(th)
}

func within(a, b, eps float64) bool {
	d := a - b
	if d < 0 {
		d = -d
	}
	return d < eps
}
