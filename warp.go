package texture

import (
	"image/color"
	"math"
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

// Eval converts from Euclidean to radial coords.
func (rw *RadialWF) Eval(x, y float64) (float64, float64) {
	lx, ly := x-rw.Center[0], y-rw.Center[1]
	rr := math.Hypot(lx, ly) * rw.RScale
	rx := rr*math.Atan2(ly, lx)*rw.CScale + rw.Center[0]
	ry := rr + rw.Center[1]
	return rx, ry
}
