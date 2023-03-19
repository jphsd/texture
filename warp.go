package texture

import "image/color"

// Warp applies a deformation to the values passed into the Eval2 function.
type Warp struct {
	Name string
	Src  Field
	Func func(float64, float64) (float64, float64)
}

func NewWarp(src Field, wf func(float64, float64) (float64, float64)) *Warp {
	return &Warp{"Warp", src, wf}
}

// Eval2 implements the Field interface.
func (w *Warp) Eval2(x, y float64) float64 {
	x, y = w.Func(x, y)
	return w.Src.Eval2(x, y)
}

// WarpVF applies a deformation to the values passed into the Eval2 function.
type WarpVF struct {
	Name string
	Src  VectorField
	Func func(float64, float64) (float64, float64)
}

func NewWarpVF(src VectorField, wf func(float64, float64) (float64, float64)) *WarpVF {
	return &WarpVF{"WarpVF", src, wf}
}

// Eval2 implements the Field interface.
func (w *WarpVF) Eval2(x, y float64) []float64 {
	x, y = w.Func(x, y)
	return w.Src.Eval2(x, y)
}

// WarpCF applies a deformation to the values passed into the Eval2 function.
type WarpCF struct {
	Name string
	Src  ColorField
	Func func(float64, float64) (float64, float64)
}

func NewWarpCF(src ColorField, wf func(float64, float64) (float64, float64)) *WarpCF {
	return &WarpCF{"WarpCF", src, wf}
}

// Eval2 implements the Field interface.
func (w *WarpCF) Eval2(x, y float64) color.Color {
	x, y = w.Func(x, y)
	return w.Src.Eval2(x, y)
}
