package texture

// Box is used to specify an area within a field. The value of In is returned if within the box bounds. The
// value of Out otherwise.
type Box struct {
	Min, Max []float64
	In, Out  float64
}

// Eval2 implements the Field interface.
func (b *Box) Eval2(x, y float64) float64 {
	if x < b.Min[0] || x > b.Max[0] || y < b.Min[1] || y > b.Max[1] {
		return b.Out
	}
	return b.In
}

// Window is used to specify an area for which the source field value is returned. The
// value of Out is returned otherwise.
type Window struct {
	Min, Max []float64
	Src      Field
	Out      float64
}

// Eval2 implements the Field interface.
func (w *Window) Eval2(x, y float64) float64 {
	if x < w.Min[0] || x > w.Max[0] || y < w.Min[1] || y > w.Max[1] {
		return w.Out
	}
	return w.Src.Eval2(x, y)
}

// Window2 is used to specify an area for which the first source field value is returned. The
// second source field value is returned otherwise.
type Window2 struct {
	Min, Max   []float64
	Src1, Src2 Field
}

// Eval2 implements the Field interface.
func (w *Window2) Eval2(x, y float64) float64 {
	if x < w.Min[0] || x > w.Max[0] || y < w.Min[1] || y > w.Max[1] {
		return w.Src2.Eval2(x, y)
	}
	return w.Src1.Eval2(x, y)
}
