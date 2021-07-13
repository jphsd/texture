package texture

type Box struct {
	Min, Max []float64
	In, Out  float64
}

func (b *Box) Eval2(x, y float64) float64 {
	if x < b.Min[0] || x > b.Max[0] || y < b.Min[1] || y > b.Max[1] {
		return b.Out
	}
	return b.In
}

type Window struct {
	Min, Max []float64
	Src      Field
	Out      float64
}

func (w *Window) Eval2(x, y float64) float64 {
	if x < w.Min[0] || x > w.Max[0] || y < w.Min[1] || y > w.Max[1] {
		return w.Out
	}
	return w.Src.Eval2(x, y)
}

type Window2 struct {
	Min, Max   []float64
	Src1, Src2 Field
}

func (w *Window2) Eval2(x, y float64) float64 {
	if x < w.Min[0] || x > w.Max[0] || y < w.Min[1] || y > w.Max[1] {
		return w.Src2.Eval2(x, y)
	}
	return w.Src1.Eval2(x, y)
}
