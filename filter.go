package texture

import (
	"math"
	"math/rand"
)

// Filter functions map [-1,1] -> [-1,1]

// NLFilter holds the parameters for a symmetric non-linear filter mapping.
type NLFilter struct {
	Name   string
	Src    Field
	NLFunc *NonLinear
	A, B   float64
}

func NewNLFilter(src Field, nlf *NonLinear, a, b float64) *NLFilter {
	return &NLFilter{"NLFilter", src, nlf, a, b}
}

// Eval2 implements the Field interface.
func (f *NLFilter) Eval2(x, y float64) float64 {
	v := f.Src.Eval2(x, y)
	v *= f.A
	v += f.B
	flag := v < 0
	if flag {
		v = -v
	}
	v = f.NLFunc.NLF.Transform(v)
	if flag {
		v = -v
	}
	return v
}

// InvertFilter flips the sign of v.
type InvertFilter struct {
	Name string
	Src  Field
}

func NewInvertFilter(src Field) *InvertFilter {
	return &InvertFilter{"InvertFilter", src}
}

// Eval2 implements the Field interface.
func (f *InvertFilter) Eval2(x, y float64) float64 {
	return 0 - f.Src.Eval2(x, y)
}

// Quantize maps Av+B into one of C buckets.
type QuantizeFilter struct {
	Name string
	Src  Field
	A, B float64
	C    int
}

func NewQuantizeFilter(src Field, a, b float64, c int) *QuantizeFilter {
	return &QuantizeFilter{"QuantizeFilter", src, a, b, c}
}

// Eval2 implements the Field interface.
func (f *QuantizeFilter) Eval2(x, y float64) float64 {
	v := f.Src.Eval2(x, y)
	v *= f.A
	v += f.B
	v = clamp(v)
	v = (v + 1) / 2
	c := float64(f.C)
	if c < 2 {
		c = 2
	}
	v *= c
	return clamp(2*math.Floor(v)/(c-1) - 1)
}

// Clip limits At+B to [-1,1].
type ClipFilter struct {
	Name string
	Src  Field
	A, B float64
}

func NewClipFilter(src Field, a, b float64) *ClipFilter {
	return &ClipFilter{"ClipFilter", src, a, b}
}

// Eval2 implements the Field interface.
func (f *ClipFilter) Eval2(x, y float64) float64 {
	v := f.Src.Eval2(x, y)
	v *= f.A
	v += f.B
	return clamp(v)
}

// OffsScaleFilter limits A(t+B) to [-1,1].
type OffsScaleFilter struct {
	Name string
	Src  Field
	A, B float64
}

func NewOffsScaleFilter(src Field, a, b float64) *OffsScaleFilter {
	return &OffsScaleFilter{"OffsScaleFilter", src, a, b}
}

// Eval2 implements the Field interface.
func (f *OffsScaleFilter) Eval2(x, y float64) float64 {
	v := f.Src.Eval2(x, y)
	v += f.B
	v *= f.A
	return clamp(v)
}

// Abs returns Abs(At+B) (clamped).
type AbsFilter struct {
	Name string
	Src  Field
	A, B float64
}

func NewAbsFilter(src Field, a, b float64) *AbsFilter {
	return &AbsFilter{"AbsFilter", src, a, b}
}

// Eval2 implements the Field interface.
func (f *AbsFilter) Eval2(x, y float64) float64 {
	v := f.Src.Eval2(x, y)
	v *= f.A
	v += f.B
	if v < 0 {
		v = -v
	}
	return clamp(v)
}

// Fold wraps values (aX+b) outside of the domain [-1,1] back into it.
type FoldFilter struct {
	Name string
	Src  Field
	A, B float64
}

func NewFoldFilter(src Field, a, b float64) *FoldFilter {
	return &FoldFilter{"FoldFilter", src, a, b}
}

// Eval2 implements the Field interface.
func (f *FoldFilter) Eval2(x, y float64) float64 {
	v := f.Src.Eval2(x, y)
	v *= f.A
	v += f.B

	// Map [-1,1] -> [0,1]
	v = (v + 1) / 2

	if v < 0 {
		v = -v
	}
	if v > 1 {
		nv := math.Floor(v)
		v -= nv
		// if nv is even, we're going up, else odd
		if int(nv)%2 != 0 {
			v = 1 - v
		}
	}

	// Map [0,1] -> [-1,1]
	return v*2 - 1
}

// RandQuantFilter supports a randomized quatization filter.
type RandQuantFilter struct {
	Src  Field
	A, B float64
	C    int
	M    []float64
}

// NewRandQuantFilter returns a new RandFilter instance for use in quantization.
func NewRandQuantFilter(src Field, a, b float64, c int) *RandQuantFilter {
	if c < 2 {
		c = 2
	}
	mm := make([]float64, c)
	dx := 2 / float64(c-1)
	mm[0] = -1
	for i := 1; i < c; i++ {
		mm[i] = clamp(mm[i-1] + dx)
	}
	rand.Shuffle(c, func(i, j int) { mm[i], mm[j] = mm[j], mm[i] })
	return &RandQuantFilter{src, a, b, c, mm}
}

// Eval2 implements the Field interface.
func (f *RandQuantFilter) Eval2(x, y float64) float64 {
	v := f.Src.Eval2(x, y)
	v *= f.A
	v += f.B
	v = clamp(v)
	v = (v + 1) / 2
	v *= float64(f.C)
	k := int(math.Floor(v))
	if k == f.C {
		k--
	}
	return f.M[k]
}

func clamp(v float64) float64 {
	if v < -1 {
		return -1
	}
	if v > 1 {
		return 1
	}
	return v
}
