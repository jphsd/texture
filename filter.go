package texture

import (
	"math"
	"math/rand"
)

// Filter functions map [-1,1] -> [-1,1]

// Invert flips the sign of t.
func Invert(t float64) float64 {
	return -t
}

// FilterChain allows multiple filters to be concatenated.
type FilterChain struct {
	Filters []func(float64) float64
}

// NewFilterChain returns a FilterChain instance.
func NewFilterChain(filters ...func(float64) float64) *FilterChain {
	return &FilterChain{filters}
}

// Eval applies the filters in the filter chain to the input.
func (fc *FilterChain) Eval(t float64) float64 {
	for _, filter := range fc.Filters {
		t = filter(t)
	}
	return t
}

// FilterVals provides parameterization for filter functions.
type FilterVals struct {
	A, B float64
	C    int
}

// Quantize maps At+B into one of C buckets.
func (fv *FilterVals) Quantize(t float64) float64 {
	// Quantize(aX+b)
	t *= fv.A
	t += fv.B
	t = clamp(t)
	t = (t + 1) / 2
	c := float64(fv.C)
	if c < 2 {
		c = 2
	}
	t *= c
	return clamp(2*math.Floor(t)/(c-1) - 1)
}

// Clip limits At+B to [-1,1].
func (fv *FilterVals) Clip(t float64) float64 {
	// Clip(aX+b)
	t *= fv.A
	t += fv.B
	return clamp(t)
}

// Sine returns Sin(At) + B (clamped).
func (fv *FilterVals) Sine(t float64) float64 {
	// Sine(aX)+b
	t = math.Sin(fv.A*t) + fv.B
	return clamp(t)
}

// Abs returns Abs(At+B) (clamped).
func (fv *FilterVals) Abs(t float64) float64 {
	// Abs(aX+b)
	t *= fv.A
	t += fv.B
	if t < 0 {
		t = -t
	}
	return clamp(t)
}

// Pow returns t^A + B (clamped).
func (fv *FilterVals) Pow(t float64) float64 {
	// X^a+b
	t = math.Pow(t, fv.A) + fv.B
	return clamp(t)
}

// Gaussian returns Gaussian(A(t+B)) (clamped).
func (fv *FilterVals) Gaussian(t float64) float64 {
	// Gaussian(a(X+b))
	t += fv.B
	t *= fv.A
	t = math.Pow(math.E, -t*t)
	return clamp(t)
}

// Fold wraps values outside of the domain [-1,1] back into it.
func (fv *FilterVals) Fold(t float64) float64 {
	// Fold(aX+b)
	t *= fv.A
	t += fv.B

	// Map [-1,1] -> [0,1]
	t = (t + 1) / 2

	if t < 0 {
		t = -t
	}
	if t > 1 {
		nt := math.Floor(t)
		t -= nt
		// if nt is even, we're going up, else odd
		if int(nt)%2 != 0 {
			t = 1 - t
		}
	}

	// Map [0,1] -> [-1,1]
	return t*2 - 1
}

func clamp(t float64) float64 {
	if t < -1 {
		return -1
	}
	if t > 1 {
		return 1
	}
	return t
}

// RandFV supports a randomized quatization filter.
type RandFV struct {
	A, B float64
	C    int
	M    []float64
}

// NewRandFV returns a new RandFV instance for use in quantization.
func NewRandFV(a, b float64, c int) *RandFV {
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
	return &RandFV{a, b, c, mm}
}

// Quantize maps At+B into one of C buckets where the bucket value has been randomized.
func (rv *RandFV) Quantize(t float64) float64 {
	// Quantize(aX+b)
	t *= rv.A
	t += rv.B
	t = clamp(t)
	t = (t + 1) / 2
	t *= float64(rv.C)
	k := int(math.Floor(t))
	if k == rv.C {
		k--
	}
	return rv.M[k]
}
