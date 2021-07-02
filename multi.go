package texture

import (
	"math"
	"math/rand"
)

type Multi struct {
	Sums         []float64
	Sum          float64
	Lambdas      []float64 // [1,...)
	Center       float64   // [0,1]
	Phase        float64   // [0,1]
	GFunc        func(float64) float64
	FFunc        func(float64) float64
	CosTh, SinTh float64
}

func NewMulti(lambdas []float64, theta float64, f func(float64) float64) *Multi {
	sums := make([]float64, len(lambdas))
	sum := 0.0
	for i, lambda := range lambdas {
		if lambda < 1 {
			lambdas[i] = 1
		}
		sum += lambdas[i]
		sums[i] = sum
	}

	return &Multi{sums, sum, lambdas, 0.5, 0, f, nil, math.Cos(theta), math.Sin(theta)}
}

func (g *Multi) Eval2(x, y float64) float64 {
	v := x*g.CosTh + y*g.SinTh
	if g.FFunc == nil {
		return g.GFunc(g.VtoT(v))
	}
	return g.FFunc(g.GFunc(g.VtoT(v)))
}

func (g *Multi) VtoT(v float64) float64 {
	// Vs Div and Floor ...
	for v < 0 {
		v += g.Sum
	}
	for v > g.Sum {
		v -= g.Sum
	}
	// Find the local lambda we're in
	i := 0
	for i < len(g.Sums) && v > g.Sums[i] {
		i++
	}
	start := 0.0
	if i > 0 {
		start = g.Sums[i-1]
	}
	t := (v-start)/g.Lambdas[i] + g.Phase
	if t > 1 {
		t -= 1
	}
	if t <= g.Center {
		return t * 0.5 / g.Center
	}
	return 0.5*(t-g.Center)/(1-g.Center) + 0.5
}

func Random(n int, lambda float64, theta float64, f func(float64) float64) *Multi {
	// Generate n lambdas between 50 to 100% of lambda
	lambdas := make([]float64, n)
	lambda /= 2
	for i := 0; i < n; i++ {
		lambdas[i] = (1 + rand.Float64()) * lambda
	}
	return NewMulti(lambdas, theta, f)
}
