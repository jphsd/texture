package texture

import (
	g2d "github.com/jphsd/graphics2d"
	"math"
)

type Fractal struct {
	Src     Field
	Xfm     *g2d.Aff3
	CFunc   func(float64, float64, int) float64
	FFunc   func(float64) float64
	Octaves int
	Rem     float64
	Initial float64
}

func (f *Fractal) Eval2(x, y float64) float64 {
	v := f.Initial
	for i := 0; i < f.Octaves; i++ {
		nv := f.Src.Eval2(x, y)
		v = f.CFunc(v, nv, i)
		pt := f.Xfm.Apply([]float64{x, y})
		x, y = pt[0][0], pt[0][1]
	}

	if f.Rem > 0 {
		// Note linear and not geometric...
		nv := f.Rem * (f.Src.Eval2(x, y)*2 - 1)
		v = f.CFunc(v, nv, f.Octaves)
	}

	if f.FFunc == nil {
		return clamp((v + 1) / 2)
	}
	return f.FFunc(clamp((v + 1) / 2))
}

const (
	MaxOctaves = 10
)

type FBM struct {
	Weights []float64
}

func NewFBM(hurst, lacunarity float64) *FBM {
	w := make([]float64, MaxOctaves)
	for i := 0; i < MaxOctaves; i++ {
		w[i] = math.Pow(lacunarity, -hurst*float64(i))
	}
	return &FBM{w}
}

func (f *FBM) Combine(cur, noise float64, oct int) float64 {
	return cur + noise*f.Weights[oct]
}

type MF struct {
	Weights []float64
	Offset  float64
}

func NewMF(hurst, lacunarity, offset float64) *MF {
	w := make([]float64, MaxOctaves)
	for i := 0; i < MaxOctaves; i++ {
		w[i] = math.Pow(lacunarity, -hurst*float64(i))
	}
	return &MF{w, offset}
}

func (f *MF) Combine(cur, noise float64, oct int) float64 {
	return cur * (noise + f.Offset) * f.Weights[oct]
}
