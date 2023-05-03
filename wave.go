package texture

// Pattern and NL

type Wave interface {
	Eval(v float64) float64
	Lambda() float64
}

// Given a value in (-inf, inf) and a wavelength, return number of waves [0,inf) and t [0,lambda) for the value.
func MapValueToLambda(v, lambda float64) (int, float64) {
	if v < 0 {
		n, v := MapValueToLambda(-v, lambda)
		if v < 0.000001 {
			return n, 0
		}
		return n, lambda - v
	} else if v < lambda {
		return 0, v
	}

	// v is greater than or equal to lambda
	n := int(v / lambda)
	v -= float64(n) * lambda
	return n, v
}

// Non-linears

type NLWave struct {
	Name      string
	Lambdas   []float64
	CumLambda []float64
	NLFs      []*NonLinear
	Mirrored  bool
	Once      bool
}

func NewNLWave(lambdas []float64, nlfs []*NonLinear, mirror, once bool) *NLWave {
	// Sum and fix lambdas
	sum := 0.0
	nl, nf := len(lambdas), len(nlfs)
	cl := make([]float64, nl)
	for i, l := range lambdas {
		if l < 0 {
			l = -l
		}
		sum += l
		cl[i] = sum
	}

	// Ensure the two slices have the same length
	if nl < nf {
		// Truncate nlfs
		nlfs = nlfs[:nl]
	} else if nl > nf {
		// Extend nlfs with repetition
		for i := 0; i < nl-nf; i++ {
			nlfs = append(nlfs, nlfs[i])
		}
	}

	return &NLWave{"NLWave", lambdas, cl, nlfs, mirror, once}
}

func (g *NLWave) Eval(v float64) float64 {
	nl := len(g.Lambdas)
	sum := g.CumLambda[nl-1]

	// Map v to n, v [0,sum)
	ov := v
	r, v := MapValueToLambda(v, sum)

	if g.Once {
		if g.Mirrored && r > 1 {
			return -1
		}
		if !g.Mirrored && r > 0 {
			return 1
		}
	}

	// Find nlf for v
	i := 0
	for i < nl {
		if v > g.CumLambda[i] {
			i++
		} else {
			break
		}
	}
	v1 := v
	if i > 0 {
		v1 -= g.CumLambda[i-1]
	}

	// Calculate t
	var t float64
	if i > nl-1 {
		// Rolled off end due to rounding errors
		t = 1
		i--
	} else {
		t = v1 / g.Lambdas[i]
	}

	// If mirrored, find direction
	if g.Mirrored {
		if (ov > 0 && (r*nl+i)%2 == 1) ||
			(ov < 0 && (r*nl+nl-i)%2 == 1) {
			t = 1 - t
		}
	}

	return g.NLFs[i].Eval(t)
}

func (g *NLWave) Lambda() float64 {
	l := g.CumLambda[len(g.Lambdas)-1]
	if g.Mirrored {
		l *= 2
	}
	return l
}

// Patterns

type PatternWave struct {
	Name      string
	Lambdas   []float64
	CumLambda []float64
	Patterns  [][]float64
	Mirrored  bool
	Once      bool
}

func NewPatternWave(lambdas []float64, patterns [][]float64, mirror, once bool) *PatternWave {
	// Sum and fix lambdas
	sum := 0.0
	nl, nf := len(lambdas), len(patterns)
	cl := make([]float64, nl)
	for i, l := range lambdas {
		if l < 0 {
			l = -l
		}
		sum += l
		cl[i] = sum
	}

	// Ensure the two slices have the same length
	if nl < nf {
		// Truncate patterns
		patterns = patterns[:nl]
	} else if nl > nf {
		// Extend patterns with repetition
		for i := 0; i < nl-nf; i++ {
			patterns = append(patterns, patterns[i])
		}
	}

	// Validate patterns
	for i, pat := range patterns {
		patterns[i] = patternValidate(pat)
	}

	return &PatternWave{"PatternWave", lambdas, cl, patterns, mirror, once}
}

func patternValidate(pat []float64) []float64 {
	for i, v := range pat {
		pat[i] = clamp(v)
	}
	// Need to pad ends for Cubic()
	pl := len(pat)
	switch pl {
	case 0:
		pat = []float64{0, 0, 0, 0}
	case 1:
		pat = []float64{pat[0], pat[0], pat[0], pat[0]}
	default:
		np := make([]float64, pl+2)
		np[0] = pat[0]
		copy(np[1:], pat)
		np[pl+1] = pat[pl-1]
		pat = np
	}
	return pat
}

// Eval implements the Field interface.
func (g *PatternWave) Eval(v float64) float64 {
	nl := len(g.Lambdas)
	sum := g.CumLambda[nl-1]

	// Map x to r, v [0,sum]
	ov := v
	r, v := MapValueToLambda(v, sum)

	if g.Once {
		if g.Mirrored && r > 1 {
			return -1
		}
		if !g.Mirrored && r > 0 {
			return 1
		}
	}

	// Find pattern for v
	i := 0
	for i < nl {
		if v > g.CumLambda[i] {
			i++
		} else {
			break
		}
	}
	v1 := v
	if i > 0 {
		v1 -= g.CumLambda[i-1]
	}

	// Calculate t
	t := v1 / g.Lambdas[i]

	// If mirrored, find direction
	if g.Mirrored {
		if (ov > 0 && (r*nl+i)%2 == 1) ||
			(ov < 0 && (r*nl+nl-i)%2 == 1) {
			t = 1 - t
		}
	}

	// Use Cubic to smooth between pattern points
	pat := g.Patterns[i]
	pl := len(pat)
	if 1-t < 0.0000001 {
		return pat[pl-1]
	}
	dt := 1 / float64(len(pat)-3)
	j, ft := MapValueToLambda(t, dt)
	p := pat[j : j+4]
	return clamp(Cubic(ft/dt, p))
}

func (g *PatternWave) Lambda() float64 {
	return g.CumLambda[len(g.Lambdas)-1]
}
