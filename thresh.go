package texture

// ThresholdFilter returns B if src < A, and C if equal or greater.
// See also [ThresholdCombiner].
type ThresholdFilter struct {
	Name    string
	Src     Field
	A, B, C float64
}

func NewThresholdFilter(src Field, a, b, c float64) *ThresholdFilter {
	return &ThresholdFilter{"ThresholdFilter", src, a, b, c}
}

// Eval2 implements the Field interface.
func (f *ThresholdFilter) Eval2(x, y float64) float64 {
	v := f.Src.Eval2(x, y)
	if v < f.A {
		return f.B
	}
	return f.C
}
