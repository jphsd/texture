package texture

import (
	"math"
	"math/rand"
)

// [-1,1] -> [-1,1]

func Invert(t float64) float64 {
	return -t
}

type FilterChain struct {
	Filters []func(float64) float64
}

func NewFilterChain(filters ...func(float64) float64) *FilterChain {
	return &FilterChain{filters}
}

func (fc *FilterChain) Eval(t float64) float64 {
	for _, filter := range fc.Filters {
		t = filter(t)
	}
	return t
}

type FilterVals struct {
	A, B float64
	C    int
}

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

func (fv *FilterVals) Clip(t float64) float64 {
	// Clip(aX+b)
	t *= fv.A
	t += fv.B
	return clamp(t)
}

func (fv *FilterVals) Sine(t float64) float64 {
	// Sine(aX)+b
	t = math.Sin(fv.A*t) + fv.B
	return clamp(t)
}

func (fv *FilterVals) Abs(t float64) float64 {
	// Abs(aX+b)
	t *= fv.A
	t += fv.B
	if t < 0 {
		t = -t
	}
	return clamp(t)
}

func (fv *FilterVals) Pow(t float64) float64 {
	// X^a+b
	t = math.Pow(t, fv.A) + fv.B
	return clamp(t)
}

func (fv *FilterVals) Gaussian(t float64) float64 {
	// Gaussian(a(X+b))
	t += fv.B
	t *= fv.A
	t = math.Pow(math.E, -t*t)
	return clamp(t)
}

func (fv *FilterVals) Fold(t float64) float64 {
	// TODO
	// Fold(aX+b)
	t *= fv.A
	t += fv.B
	if t < 0 {
		t = -t
	}
	if t > 1 {
		nt := math.Floor(t)
		// if nt is even, we're going up, else odd
		if int(nt)%2 == 0 {
			return t - nt
		}
		return 1 - (t - nt)
	}
	return t
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

type RandFV struct {
	A, B float64
	C    int
	M    []float64
}

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
