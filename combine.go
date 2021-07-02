package texture

type Combiner2 struct {
	Src1, Src2 Field
	CFunc      func(float64, float64) float64
	FFunc      func(float64) float64
}

func (c *Combiner2) Eval2(x, y float64) float64 {
	v1, v2 := c.Src1.Eval2(x, y), c.Src2.Eval2(x, y)
	if c.FFunc == nil {
		return c.CFunc(v1, v2)
	}
	return c.FFunc(c.CFunc(v1, v2))
}

func Mul(u, v float64) float64 {
	return u * v
}

func Add(u, v float64) float64 {
	return clamp(u + v)
}

func Sub(u, v float64) float64 {
	return clamp(u - v)
}

func Min(u, v float64) float64 {
	if u < v {
		return u
	}
	return v
}

func Max(u, v float64) float64 {
	if u > v {
		return u
	}
	return v
}

func Avg(u, v float64) float64 {
	return (u + v) / 2
}

func Diff(u, v float64) float64 {
	t := (u - v + 2) / 4
	return (1-t)*u + t*v
}

type CF2 struct {
	Start, End float64
}

func (c *CF2) Substitute(u, v float64) float64 {
	if u < c.Start || u > c.End {
		return u
	}
	return v
}

type Combiner3 struct {
	Src1, Src2, Src3 Field
	CFunc            func(float64, float64, float64) float64
	FFunc            func(float64) float64
}

func (c *Combiner3) Eval2(x, y float64) float64 {
	v1, v2, v3 := c.Src1.Eval2(x, y), c.Src2.Eval2(x, y), c.Src3.Eval2(x, y)
	if c.FFunc == nil {
		return c.CFunc(v1, v2, v3)
	}
	return c.FFunc(c.CFunc(v1, v2, v3))
}

func Blend(u, v, w float64) float64 {
	t := (w + 1) / 2
	return (1-t)*u + t*v
}

type CF3 struct {
	Start, End float64
}

func (c *CF3) Substitute(u, v, w float64) float64 {
	if w < c.Start || w > c.End {
		return u
	}
	return v
}
