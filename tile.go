package texture

import (
	"image/color"
	"math/rand"
)

type Tiler struct {
	Name   string
	Src    Field
	Domain []float64
}

func NewTiler(src Field, dom []float64) *Tiler {
	return &Tiler{"Tiler", src, dom}
}

func (t *Tiler) Eval2(x, y float64) float64 {
	_, x = MapValueToLambda(x, t.Domain[0])
	_, y = MapValueToLambda(y, t.Domain[1])
	return t.Src.Eval2(x, y)
}

type TilerCF struct {
	Name   string
	Src    ColorField
	Domain []float64
}

func NewTilerCF(src ColorField, dom []float64) *TilerCF {
	return &TilerCF{"TilerCF", src, dom}
}

func (t *TilerCF) Eval2(x, y float64) color.Color {
	_, x = MapValueToLambda(x, t.Domain[0])
	_, y = MapValueToLambda(y, t.Domain[1])
	return t.Src.Eval2(x, y)
}

type TilerVF struct {
	Name   string
	Src    VectorField
	Domain []float64
}

func NewTilerVF(src VectorField, dom []float64) *TilerVF {
	return &TilerVF{"TilerVF", src, dom}
}

func (t *TilerVF) Eval2(x, y float64) []float64 {
	_, x = MapValueToLambda(x, t.Domain[0])
	_, y = MapValueToLambda(y, t.Domain[1])
	return t.Src.Eval2(x, y)
}

type StochasticTiler struct {
	Name   string
	Srcs   []Field
	Domain []float64
	rmap   map[int]int
}

func NewStochasticTiler(srcs []Field, dom []float64) *StochasticTiler {
	return &StochasticTiler{"StochasticTiler", srcs, dom, make(map[int]int)}
}

func (t *StochasticTiler) Eval2(x, y float64) float64 {
	nx, x := MapValueToLambda(x, t.Domain[0])
	ny, y := MapValueToLambda(y, t.Domain[1])

	return t.Srcs[rval(nx, ny, len(t.Srcs), t.rmap)].Eval2(x, y)
}

type StochasticTilerCF struct {
	Name   string
	Srcs   []ColorField
	Domain []float64
	rmap   map[int]int
}

func NewStochasticTilerCF(srcs []ColorField, dom []float64) *StochasticTilerCF {
	return &StochasticTilerCF{"StochasticTilerCF", srcs, dom, make(map[int]int)}
}

func (t *StochasticTilerCF) Eval2(x, y float64) color.Color {
	nx, x := MapValueToLambda(x, t.Domain[0])
	ny, y := MapValueToLambda(y, t.Domain[1])

	return t.Srcs[rval(nx, ny, len(t.Srcs), t.rmap)].Eval2(x, y)
}

type StochasticTilerVF struct {
	Name   string
	Srcs   []VectorField
	Domain []float64
	rmap   map[int]int
}

func NewStochasticTilerVF(srcs []VectorField, dom []float64) *StochasticTilerVF {
	return &StochasticTilerVF{"StochasticTilerVF", srcs, dom, make(map[int]int)}
}

func (t *StochasticTilerVF) Eval2(x, y float64) []float64 {
	nx, x := MapValueToLambda(x, t.Domain[0])
	ny, y := MapValueToLambda(y, t.Domain[1])

	return t.Srcs[rval(nx, ny, len(t.Srcs), t.rmap)].Eval2(x, y)
}

// Simple cache for random mappings
func rval(m, n, l int, rmap map[int]int) int {
	k := m*1024 + n

	if v, ok := rmap[k]; ok {
		return v
	}

	if len(rmap) > 10240 {
		// Start over
		rmap = make(map[int]int)
	}

	lr := rand.New(rand.NewSource(int64(k)))
	v := lr.Intn(l)
	rmap[k] = v
	return v
}
