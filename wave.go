package texture

import "math"

type Wave interface {
	Eval(v float64) float64
	Lambda() float64
}

func MapValueToLambda(v, lambda float64) (int, float64) {
	if v < 0 {
		if v > -lambda {
			return 0, v + lambda
		}
		n, f := math.Modf(v / lambda)
		if f < 0 {
			return -int(n), lambda * (1 + f)
		}
		return -int(n) - 1, f
	}

	if v < lambda {
		return 0, v
	}

	n, f := math.Modf(v / lambda)
	return int(n), lambda * f
}

// Non-linear waves

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
		if ov < 0 {
			return -1
		}
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

// DCWave - nlf for up and nlf for down. If only one nlf, use for both up and down. Independent lambdas.

type DCWave struct {
	Name string
	L1   float64
	L2   float64
	Sum  float64
	NL1  *NonLinear
	NL2  *NonLinear
	Once bool
}

func NewDCWave(lambdas []float64, nlfs []*NonLinear, once bool) *DCWave {
	l1 := lambdas[0]
	var l2 float64
	if len(lambdas) > 1 {
		l2 = lambdas[1]
	} else {
		l2 = l1
	}
	nl1 := nlfs[0]
	var nl2 *NonLinear
	if len(nlfs) > 1 {
		nl2 = nlfs[1]
	} else {
		nl2 = nl1
	}
	return &DCWave{"DCWave", l1, l2, l1 + l2, nl1, nl2, once}
}

func (g *DCWave) Eval(v float64) float64 {
	// Map v to n, v [0,sum)
	ov := v
	r, v := MapValueToLambda(v, g.Sum)

	if g.Once && (ov < 0 || r > 0) {
		return -1
	}

	// Find nlf and t for v
	nlf := g.NL1
	var t float64
	if v > g.L1 {
		nlf = g.NL2
		t = 1 - (v-g.L1)/g.L2
	} else {
		t = v / g.L1
	}

	return nlf.Eval(t)
}

func (g *DCWave) Lambda() float64 {
	return g.Sum
}

// ACWave - nlfs for each quadrant. Independent lambdas. Start and end at 0.

type ACWave struct {
	Name      string
	Lambdas   [4]float64
	CumLambda [4]float64
	NLFs      [4]*NonLinear
	Once      bool
}

func NewACWave(lambdas []float64, nlfs []*NonLinear, once bool) *ACWave {
	res := &ACWave{}
	res.Name = "ACWave"
	switch len(lambdas) {
	default:
		fallthrough
	case 1:
		res.Lambdas[0] = lambdas[0]
		res.Lambdas[1] = lambdas[0]
		res.Lambdas[2] = lambdas[0]
		res.Lambdas[3] = lambdas[0]
	case 2:
		res.Lambdas[0] = lambdas[0]
		res.Lambdas[1] = lambdas[1]
		res.Lambdas[2] = lambdas[0]
		res.Lambdas[3] = lambdas[1]
	case 4:
		res.Lambdas[0] = lambdas[0]
		res.Lambdas[1] = lambdas[1]
		res.Lambdas[2] = lambdas[2]
		res.Lambdas[3] = lambdas[3]
	}
	sum := 0.0
	for i := 0; i < 4; i++ {
		sum += res.Lambdas[i]
		res.CumLambda[i] = sum
	}
	switch len(nlfs) {
	default:
		fallthrough
	case 1:
		res.NLFs[0] = nlfs[0]
		res.NLFs[1] = nlfs[0]
		res.NLFs[2] = nlfs[0]
		res.NLFs[3] = nlfs[0]
	case 2:
		res.NLFs[0] = nlfs[0]
		res.NLFs[1] = nlfs[1]
		res.NLFs[2] = nlfs[0]
		res.NLFs[3] = nlfs[1]
	case 4:
		res.NLFs[0] = nlfs[0]
		res.NLFs[1] = nlfs[1]
		res.NLFs[2] = nlfs[2]
		res.NLFs[3] = nlfs[3]
	}
	res.Once = once

	return res
}

func (g *ACWave) Eval(v float64) float64 {
	sum := g.CumLambda[3]

	// Map v to n, v [0,sum)
	ov := v
	r, v := MapValueToLambda(v, sum)

	if g.Once && (ov < 0 || r > 0) {
		return 0
	}

	// Find nlf for v
	q := 0
	for q < 4 {
		if v > g.CumLambda[q] {
			q++
		} else {
			break
		}
	}
	v1 := v
	if q > 0 {
		v1 -= g.CumLambda[q-1]
	}

	// Calculate t
	var t float64
	if q > 3 {
		// Rolled off end due to rounding errors
		t = 1
		q--
	} else {
		t = v1 / g.Lambdas[q]
	}
	// Modify t dep on quadrant
	if q == 1 || q == 3 {
		t = 1 - t
	}

	// Calculate result based on quadrant
	res := g.NLFs[q].Eval0(t)
	if q > 1 {
		res = -res
	}

	return res
}

func (g *ACWave) Lambda() float64 {
	return g.CumLambda[3]
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

// InvertWave - inverts the input wave.

type InvertWave struct {
	Name string
	Src  Wave
}

func NewInvertWave(src Wave) *InvertWave {
	return &InvertWave{"InvertWave", src}
}

func (g *InvertWave) Eval(v float64) float64 {
	return -g.Src.Eval(v)
}

func (g *InvertWave) Lambda() float64 {
	return g.Src.Lambda()
}
