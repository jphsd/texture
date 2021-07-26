package texture

import (
	g2d "github.com/jphsd/graphics2d"
	"math"
)

// Fractal holds the pieces necessary for fractal generation. Xfm defines the affine transformation
// applied successively to the coordinate space and CFunc, how the multiple resultant values should be combined.
// An optional filter can be specified which will be applied to the final result.
type Fractal struct {
	Src     Field
	Xfm     *g2d.Aff3
	CFunc   func(...float64) float64
	FFunc   func(float64) float64
	Octaves int
	Rem     float64
	N       int
}

// NewFractal returns a new Fractal instance.
func NewFractal(src Field, xfm *g2d.Aff3, comb func(...float64) float64, octaves float64) *Fractal {
	n := int(math.Floor(octaves))
	r := octaves - float64(n)
	vn := n + 1
	if r > 0 {
		vn++
	}
	return &Fractal{src, xfm, comb, nil, n, r, vn}
}

// Eval2 implements the Field interface.
func (f *Fractal) Eval2(x, y float64) float64 {
	nv := make([]float64, f.N)
	for i := 0; i <= f.Octaves; i++ {
		nv[i] = f.Src.Eval2(x, y)
		pt := f.Xfm.Apply([]float64{x, y})
		x, y = pt[0][0], pt[0][1]
	}

	if f.Rem > 0 {
		// Note linear and not geometric...
		nv[f.Octaves] = f.Rem * f.Src.Eval2(x, y)
	}

	v := clamp(f.CFunc(nv...))

	if f.FFunc == nil {
		return v
	}
	return f.FFunc(v)
}

const (
	// MaxOctaves is the maximum number of iterations a Fractal can perform
	MaxOctaves = 10
)

// FBM holds the precomputed weights for an fBM.
type FBM struct {
	Weights []float64
}

// NewFBM returns a new FBM instance based on the Hurst and Lacunarity parameters.
func NewFBM(hurst, lacunarity float64) *FBM {
	w := make([]float64, MaxOctaves)
	for i := 0; i < MaxOctaves; i++ {
		w[i] = math.Pow(lacunarity, -hurst*float64(i+1))
	}
	return &FBM{w}
}

// Combine takes the values from the successive applications of the affine transform and
// combines them using the precomputed weights.
func (f *FBM) Combine(values ...float64) float64 {
	res := 0.0
	for i := 0; i < len(values); i++ {
		res += values[i] * f.Weights[i]
	}
	return res
}

// MF holds the precomputed weights and offset for an multifractal.
type MF struct {
	Weights []float64
	Offset  float64
}

// NewMF returns a new MF instance based on the Hurst and Lacunarity parameters.
func NewMF(hurst, lacunarity, offset float64) *MF {
	w := make([]float64, MaxOctaves)
	for i := 0; i < MaxOctaves; i++ {
		w[i] = math.Pow(lacunarity, -hurst*float64(i+1))
	}
	return &MF{w, offset}
}

// Combine takes the values from the successive applications of the affine transform and
// combines them using the precomputed weights and offset.
func (f *MF) Combine(values ...float64) float64 {
	res := 0.0
	for i := 0; i < len(values); i++ {
		res += (values[i] + f.Offset) * f.Weights[i]
	}
	return res
}
