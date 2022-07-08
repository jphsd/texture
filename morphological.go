package texture

type Erode struct {
	Name string
	Src  Field
	Supp [][]float64
}

func NewErode(src Field, supp [][]float64) *Erode {
	return &Erode{"Erode", src, supp}
}

func (m *Erode) Eval2(x, y float64) float64 {
	min := m.Src.Eval2(x, y)
	for _, s := range m.Supp {
		v := m.Src.Eval2(x+s[0], y+s[1])
		if v < min {
			min = v
		}
	}
	return min
}

type Dilate struct {
	Name string
	Src  Field
	Supp [][]float64
}

func NewDilate(src Field, supp [][]float64) *Dilate {
	return &Dilate{"Dilate", src, supp}
}

func (m *Dilate) Eval2(x, y float64) float64 {
	max := m.Src.Eval2(x, y)
	for _, s := range m.Supp {
		v := m.Src.Eval2(x+s[0], y+s[1])
		if v > max {
			max = v
		}
	}
	return max
}

// Higher level Morphological operations

// EdgeIn - orig - E
type EdgeIn struct {
	Name string
	Src1 Field
	Src2 Field
}

func NewEdgeIn(src Field, supp [][]float64) *EdgeIn {
	f := NewErode(src, supp)
	return &EdgeIn{"EdgeIn", src, f}
}

func (m *EdgeIn) Eval2(x, y float64) float64 {
	return m.Src1.Eval2(x, y) - m.Src2.Eval2(x, y)
}

// EdgeOut - D - orig
type EdgeOut struct {
	Name string
	Src1 Field
	Src2 Field
}

func NewEdgeOut(src Field, supp [][]float64) *EdgeOut {
	f := NewDilate(src, supp)
	return &EdgeOut{"EdgeOut", f, src}
}

func (m *EdgeOut) Eval2(x, y float64) float64 {
	return m.Src1.Eval2(x, y) - m.Src2.Eval2(x, y)
}

// Edge - D - E
type Edge struct {
	Name string
	Src1 Field
	Src2 Field
}

func NewEdge(src Field, supp [][]float64) *Edge {
	f1 := NewDilate(src, supp)
	f2 := NewErode(src, supp)
	return &Edge{"Edge", f1, f2}
}

func (m *Edge) Eval2(x, y float64) float64 {
	return m.Src1.Eval2(x, y) - m.Src2.Eval2(x, y)
}

// Close - D then E
type Close struct {
	Name string
	Src  Field
}

func NewClose(src Field, supp [][]float64) *Close {
	f1 := NewDilate(src, supp)
	f2 := NewErode(f1, supp)
	return &Close{"Close", f2}
}

func (m *Close) Eval2(x, y float64) float64 {
	return m.Src.Eval2(x, y)
}

// Open - E then D
type Open struct {
	Name string
	Src  Field
}

func NewOpen(src Field, supp [][]float64) *Open {
	f1 := NewErode(src, supp)
	f2 := NewDilate(f1, supp)
	return &Open{"Open", f2}
}

func (m *Open) Eval2(x, y float64) float64 {
	return m.Src.Eval2(x, y)
}

// TopHat - orig - O
type TopHat struct {
	Name string
	Src1 Field
	Src2 Field
}

func NewTopHat(src Field, supp [][]float64) *TopHat {
	f := NewOpen(src, supp)
	return &TopHat{"TopHat", src, f}
}

func (m *TopHat) Eval2(x, y float64) float64 {
	return m.Src1.Eval2(x, y) - m.Src2.Eval2(x, y)
}

// BottomHat - C - orig
type BottomHat struct {
	Name string
	Src1 Field
	Src2 Field
}

func NewBottomHat(src Field, supp [][]float64) *BottomHat {
	f := NewClose(src, supp)
	return &BottomHat{"BottomHat", f, src}
}

func (m *BottomHat) Eval2(x, y float64) float64 {
	return m.Src1.Eval2(x, y) - m.Src2.Eval2(x, y)
}
