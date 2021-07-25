package texture

// Combiner2 takes two source fields and combines them with the specified combiner function. The result is
// passed through an optional filter function.
type Combiner2 struct {
	Src1, Src2 Field
	CFunc      func(float64, float64) float64
	FFunc      func(float64) float64
}

// Eval2 implements the Field interface.
func (c *Combiner2) Eval2(x, y float64) float64 {
	v1, v2 := c.Src1.Eval2(x, y), c.Src2.Eval2(x, y)
	if c.FFunc == nil {
		return c.CFunc(v1, v2)
	}
	return c.FFunc(c.CFunc(v1, v2))
}

// Mul is a multiplying combiner.
func Mul(u, v float64) float64 {
	return u * v
}

// Add is an adding combiner (clamped).
func Add(u, v float64) float64 {
	return clamp(u + v)
}

// Sub is a subtracting combiner (clamped).
func Sub(u, v float64) float64 {
	return clamp(u - v)
}

// Min is a minimizing combiner.
func Min(u, v float64) float64 {
	if u < v {
		return u
	}
	return v
}

// Max is a maximizing combiner.
func Max(u, v float64) float64 {
	if u > v {
		return u
	}
	return v
}

// Avg is an averaging combiner.
func Avg(u, v float64) float64 {
	return (u + v) / 2
}

// Diff combines two values by weighting them in proportion to the difference between them.
func Diff(u, v float64) float64 {
	t := (u - v + 2) / 4
	return (1-t)*u + t*v
}

// CF2 holds values for parameterized combiners.
type CF2 struct {
	V1, V2 float64
}

// Substitute returns either u or v depending on whether u is outside or inside the parameter range.
func (c *CF2) Substitute(u, v float64) float64 {
	if u < c.V1 || u > c.V2 {
		return u
	}
	return v
}

// Sum returns a weighted sum of the inputs.
func (c *CF2) Sum(u, v float64) float64 {
	return c.V1*u + c.V2*v
}

// Combiner3 takes three source fields and combines them with the specified combiner function. The result is
// passed through an optional filter function.
type Combiner3 struct {
	Src1, Src2, Src3 Field
	CFunc            func(...float64) float64
	FFunc            func(float64) float64
}

// Eval2 implements the Field interface.
func (c *Combiner3) Eval2(x, y float64) float64 {
	v1, v2, v3 := c.Src1.Eval2(x, y), c.Src2.Eval2(x, y), c.Src3.Eval2(x, y)
	if c.FFunc == nil {
		return c.CFunc(v1, v2, v3)
	}
	return c.FFunc(c.CFunc(v1, v2, v3))
}

// Blend combines u and v weighted by w.
func Blend(v ...float64) float64 {
	t := (v[2] + 1) / 2
	return (1-t)*v[0] + t*v[1]
}

type CF3 struct {
	Start, End float64
}

// Substitute returns either u or v depending on whether w is outside or inside the parameter range.
func (c *CF3) Substitute(v ...float64) float64 {
	if v[2] < c.Start || v[2] > c.End {
		return v[0]
	}
	return v[1]
}
