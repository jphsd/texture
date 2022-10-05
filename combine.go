package texture

import "math/rand"

// MulComnbiner is a multiplying combiner.
type MulCombiner struct {
	Name string
	Src1 Field
	Src2 Field
}

func NewMulCombiner(src1, src2 Field) *MulCombiner {
	return &MulCombiner{"MulCombiner", src1, src2}
}

// Eval2 implements the Field interface.
func (c *MulCombiner) Eval2(x, y float64) float64 {
	v1, v2 := c.Src1.Eval2(x, y), c.Src2.Eval2(x, y)
	return v1 * v2
}

// AddCombiner is an adding combiner (clamped).
type AddCombiner struct {
	Name string
	Src1 Field
	Src2 Field
}

func NewAddCombiner(src1, src2 Field) *AddCombiner {
	return &AddCombiner{"AddCombiner", src1, src2}
}

// Eval2 implements the Field interface.
func (c *AddCombiner) Eval2(x, y float64) float64 {
	v1, v2 := c.Src1.Eval2(x, y), c.Src2.Eval2(x, y)
	return clamp(v1 + v2)
}

// SubCombiner is a subtracting combiner (clamped).
type SubCombiner struct {
	Name string
	Src1 Field
	Src2 Field
}

func NewSubCombiner(src1, src2 Field) *SubCombiner {
	return &SubCombiner{"SubCombiner", src1, src2}
}

// Eval2 implements the Field interface.
func (c *SubCombiner) Eval2(x, y float64) float64 {
	v1, v2 := c.Src1.Eval2(x, y), c.Src2.Eval2(x, y)
	return clamp(v1 - v2)
}

// MinCombiner is a minimizing combiner.
type MinCombiner struct {
	Name string
	Src1 Field
	Src2 Field
}

func NewMinCombiner(src1, src2 Field) *MinCombiner {
	return &MinCombiner{"MinCombiner", src1, src2}
}

// Eval2 implements the Field interface.
func (c *MinCombiner) Eval2(x, y float64) float64 {
	v1, v2 := c.Src1.Eval2(x, y), c.Src2.Eval2(x, y)
	if v1 < v2 {
		return v1
	}
	return v2
}

// MaxCombiner is a maximizing combiner.
type MaxCombiner struct {
	Name string
	Src1 Field
	Src2 Field
}

func NewMaxCombiner(src1, src2 Field) *MaxCombiner {
	return &MaxCombiner{"MaxCombiner", src1, src2}
}

// Eval2 implements the Field interface.
func (c *MaxCombiner) Eval2(x, y float64) float64 {
	v1, v2 := c.Src1.Eval2(x, y), c.Src2.Eval2(x, y)
	if v1 < v2 {
		return v2
	}
	return v1
}

// AvgCombiner is an averaging combiner.
type AvgCombiner struct {
	Name string
	Src1 Field
	Src2 Field
}

func NewAvgCombiner(src1, src2 Field) *AvgCombiner {
	return &AvgCombiner{"AvgCombiner", src1, src2}
}

// Eval2 implements the Field interface.
func (c *AvgCombiner) Eval2(x, y float64) float64 {
	v1, v2 := c.Src1.Eval2(x, y), c.Src2.Eval2(x, y)
	return (v1 + v2) / 2
}

// DiffCombiner combines two values by weighting them in proportion to the difference between them.
type DiffCombiner struct {
	Name string
	Src1 Field
	Src2 Field
}

func NewDiffCombiner(src1, src2 Field) *DiffCombiner {
	return &DiffCombiner{"DiffCombiner", src1, src2}
}

// Eval2 implements the Field interface.
func (c *DiffCombiner) Eval2(x, y float64) float64 {
	v1, v2 := c.Src1.Eval2(x, y), c.Src2.Eval2(x, y)
	t := (v1 - v2 + 2) / 4
	return (1-t)*v1 + t*v2
}

// WindowedfCombiner combines two fields based on the window values.
type WindowedCombiner struct {
	Name string
	Src1 Field
	Src2 Field
	A, B float64
}

func NewWindowedCombiner(src1, src2 Field, a, b float64) *WindowedCombiner {
	return &WindowedCombiner{"WindowedCombiner", src1, src2, a, b}
}

// Eval2 implements the Field interface.
func (c *WindowedCombiner) Eval2(x, y float64) float64 {
	v1 := c.Src1.Eval2(x, y)
	if v1 < c.A || v1 > c.B {
		return v1
	}
	return c.Src2.Eval2(x, y)
}

// WeightedfCombiner combines two fields based on the supplied values.
type WeightedCombiner struct {
	Name string
	Src1 Field
	Src2 Field
	A, B float64
}

func NewWeightedCombiner(src1, src2 Field, a, b float64) *WeightedCombiner {
	return &WeightedCombiner{"WeightedCombiner", src1, src2, a, b}
}

// Eval2 implements the Field interface.
func (c *WeightedCombiner) Eval2(x, y float64) float64 {
	v1, v2 := c.Src1.Eval2(x, y), c.Src2.Eval2(x, y)
	return clamp(c.A*v1 + c.B*v2)
}

type Blend struct {
	Name string
	Src1 Field
	Src2 Field
	Src3 Field
}

func NewBlend(src1, src2, src3 Field) *Blend {
	return &Blend{"Blend", src1, src2, src3}
}

// Eval2 implements the Field interface.
func (b *Blend) Eval2(x, y float64) float64 {
	v1, v2, v3 := b.Src1.Eval2(x, y), b.Src2.Eval2(x, y), b.Src3.Eval2(x, y)
	t := (v3 + 1) / 2
	return (1-t)*v1 + t*v2
}

type StochasticBlend struct {
	Name string
	Src1 Field
	Src2 Field
	Src3 Field
}

func NewStochasticBlend(src1, src2, src3 Field) *StochasticBlend {
	return &StochasticBlend{"StochasticBlend", src1, src2, src3}
}

// Eval2 implements the Field interface.
func (b *StochasticBlend) Eval2(x, y float64) float64 {
	t := b.Src3.Eval2(x, y)
	t = (t + 1) / 2
	if rand.Float64() < t {
		return b.Src2.Eval2(x, y)
	}
	return b.Src1.Eval2(x, y)
}

type JitterBlend struct {
	Name string
	Src1 Field
	Src2 Field
	Src3 Field
	Perc float64
}

func NewJitterBlend(src1, src2, src3 Field, perc float64) *JitterBlend {
	return &JitterBlend{"JitterBlend", src1, src2, src3, perc}
}

// Eval2 implements the Field interface.
func (b *JitterBlend) Eval2(x, y float64) float64 {
	v1, v2, v3 := b.Src1.Eval2(x, y), b.Src2.Eval2(x, y), b.Src3.Eval2(x, y)
	t := (v3 + 1) / 2
	j := (rand.Float64()*2 - 1) * b.Perc
	t += j
	if t < 0 {
		t = 0
	} else if t > 1 {
		t = 1
	}
	return (1-t)*v1 + t*v2
}

// SubstituteCombiner combines two fields based on the supplied values.
type SubstituteCombiner struct {
	Name string
	Src1 Field
	Src2 Field
	Src3 Field
	A, B float64
}

func NewSubstituteCombiner(src1, src2, src3 Field, a, b float64) *SubstituteCombiner {
	return &SubstituteCombiner{"SubstituteCombiner", src1, src2, src3, a, b}
}

// Eval2 implements the Field interface.
func (c *SubstituteCombiner) Eval2(x, y float64) float64 {
	v1, v3 := c.Src1.Eval2(x, y), c.Src3.Eval2(x, y)
	if v3 < c.A || v3 > c.B {
		return v1
	}
	return c.Src2.Eval2(x, y)
}
