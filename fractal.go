package texture

import (
	g2d "github.com/jphsd/graphics2d"
	"math"
)

// Fractal holds the pieces necessary for fractal generation. Xfm defines the affine transformation
// applied successively to the coordinate space and Comb, how the multiple resultant values should be combined.
// Bands can also be manipulated through the weight values (typically [0, 1]) which allows certain frequencies
// to be attenuated.
type Fractal struct {
	Name    string
	Src     Field
	Xfm     *g2d.Aff3
	Comb    OctaveCombiner
	Octaves float64
	Weights []float64
}

// NewFractal returns a new Fractal instance.
func NewFractal(src Field, xfm *g2d.Aff3, comb OctaveCombiner, octaves float64) *Fractal {
	oct := int(octaves)
	w := make([]float64, oct)
	for i := 0; i < oct; i++ {
		w[i] = 1
	}
	return &Fractal{"Fractal", src, xfm, comb, octaves, w}
}

// Eval2 implements the Field interface.
func (f *Fractal) Eval2(x, y float64) float64 {
	n := int(f.Octaves)
	r := f.Octaves - float64(n)
	n++
	nv := make([]float64, n)
	for i := 0; i < n; i++ {
		nv[i] = f.Src.Eval2(x, y) * f.Weights[i]
		pt := f.Xfm.Apply([]float64{x, y})
		x, y = pt[0][0], pt[0][1]
	}
	nv[n-1] *= r

	return clamp(f.Comb.Combine(nv...))
}

type VariableFractal struct {
	Name    string
	Src     Field
	Xfm     *g2d.Aff3
	Comb    OctaveCombiner
	OctSrc  Field
	Scale   float64
	Weights []float64
}

// NewFractal returns a new Fractal instance.
func NewVariableFractal(src Field, xfm *g2d.Aff3, comb OctaveCombiner, octsrc Field, scale float64) *VariableFractal {
	n := int(scale)
	w := make([]float64, n)
	for i := 0; i < n; i++ {
		w[i] = 1
	}
	return &VariableFractal{"VariableFractal", src, xfm, comb, octsrc, scale / 2, w}
}

// Eval2 implements the Field interface.
func (f *VariableFractal) Eval2(x, y float64) float64 {
	oct := (f.OctSrc.Eval2(x, y) + 1) * f.Scale
	n := int(oct)
	r := oct - float64(n)
	n++
	nv := make([]float64, n)
	for i := 0; i < n; i++ {
		nv[i] = f.Src.Eval2(x, y)
		pt := f.Xfm.Apply([]float64{x, y})
		x, y = pt[0][0], pt[0][1]
	}
	nv[n-1] *= r

	return clamp(f.Comb.Combine(nv...))
}

type OctaveCombiner interface {
	Combine(...float64) float64
}

// FBM holds the precomputed weights for an fBM.
type FBM struct {
	Name    string
	Weights []float64
}

// NewFBM returns a new FBM instance based on the Hurst and Lacunarity parameters.
func NewFBM(hurst, lacunarity float64, maxoct int) *FBM {
	maxoct++
	w := make([]float64, maxoct)
	for i := 0; i < maxoct; i++ {
		w[i] = math.Pow(lacunarity, -hurst*float64(i+1))
	}
	return &FBM{"FBM", w}
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
	Name    string
	Weights []float64
	Offset  float64
}

// NewMF returns a new MF instance based on the Hurst and Lacunarity parameters.
func NewMF(hurst, lacunarity, offset float64, maxoct int) *MF {
	maxoct++
	w := make([]float64, maxoct)
	for i := 0; i < maxoct; i++ {
		w[i] = math.Pow(lacunarity, -hurst*float64(i+1))
	}
	return &MF{"MF", w, offset}
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
