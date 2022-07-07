package texture

type MorpOp int

const (
	ErodeOp MorpOp = iota
	DilateOp
)

type Morphological struct {
	Name string
	Src  Field
	Supp [][]float64
	Oper MorpOp
}

func NewMorphological(src Field, supp [][]float64, op MorpOp) *Morphological {
	return &Morphological{"Morphological", src, supp, op}
}

func (m *Morphological) Eval2(x, y float64) float64 {
	support := make([]float64, len(m.Supp)+1)
	support[0] = m.Src.Eval2(x, y)
	for i, s := range m.Supp {
		support[i+1] = m.Src.Eval2(x+s[0], y+s[1])
	}
	switch m.Oper {
	default:
		fallthrough
	case ErodeOp:
		return Erode(support)
	case DilateOp:
		return Dilate(support)
	}
}

// Morphological helper functions

// Erode - min
func Erode(vals []float64) float64 {
	m := 1.0
	for _, v := range vals {
		if m > v {
			m = v
		}
	}
	return m
}

// Dilate - max
func Dilate(vals []float64) float64 {
	m := -1.0
	for _, v := range vals {
		if m < v {
			m = v
		}
	}
	return m
}

// Higher level Morphological operations

// EdgeIn - orig - E
type EdgeIn struct {
	Name string
	Src1 Field
	Src2 Field
}

func NewEdgeIn(src Field, supp [][]float64) *EdgeIn {
	f := NewMorphological(src, supp, ErodeOp)
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
	f := NewMorphological(src, supp, DilateOp)
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
	f1 := NewMorphological(src, supp, DilateOp)
	f2 := NewMorphological(src, supp, ErodeOp)
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
	f1 := NewMorphological(src, supp, DilateOp)
	f2 := NewMorphological(f1, supp, ErodeOp)
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
	f1 := NewMorphological(src, supp, ErodeOp)
	f2 := NewMorphological(f1, supp, DilateOp)
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
