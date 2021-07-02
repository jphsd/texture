package texture

import (
	g2d "github.com/jphsd/graphics2d"
	g2dcol "github.com/jphsd/graphics2d/color"
	"github.com/jphsd/graphics2d/util"
	"image/color"
	"math"
	"math/rand"
)

type Fields struct {
	Name string
	Make func(int, int) Field
}

var LeafOptions []Fields

// TODO - promote to var?
func GetLeaves() []Fields {
	// Lazy initialization to prevent compiler complaining
	// about initialization loops
	if LeafOptions == nil {
		LeafOptions = []Fields{
			{"Generator", MakeGenerator},
			{"Perlin", MakePerlin},
			{"NonLinear", MakeNonLinear},
		}
	}
	return LeafOptions
}

var NodeOptions []Fields

func GetNodes() []Fields {
	// Lazy initialization to prevent compiler complaining
	// about initialization loops
	if NodeOptions == nil {
		NodeOptions = []Fields{
			{"Fractal", MakeFractal},
			{"Distort", MakeDistort},
			{"Select", MakeSelect},
			{"Direction", MakeDirection},
			{"Magnitude", MakeMagnitude},
			{"Combiner2", MakeCombiner2},
			{"Combiner3", MakeCombiner3},
			{"Displace", MakeDisplace},
			{"VectorCombine", MakeVectorCombine},
		}
	}
	return NodeOptions
}

func MakeField(md, d int) Field {
	l := GetLeaves()
	n := GetNodes()
	if d >= md {
		return l[rand.Intn(len(l))].Make(md, d+1)
	}
	s := len(l) + len(n)
	s = rand.Intn(s)
	if s < len(l) {
		return l[s].Make(md, d+1)
	}
	s -= len(l)
	return n[s].Make(md, d+1)
}

type VectorFields struct {
	Name string
	Make func(int, int) VectorField
}

func MakeVectorField(md, d int) VectorField {
	return MakeNormal(md, d+1)
}

type ColorFields struct {
	Name string
	Make func(int, int) ColorField
}

var ColorFieldOptions []ColorFields

func GetColorFields() []ColorFields {
	// Lazy initialization to prevent compiler complaining
	// about initialization loops
	if ColorFieldOptions == nil {
		ColorFieldOptions = []ColorFields{
			{"Color", MakeColor},
			{"ColorConv", MakeColorConv},
			{"ColorBlend", MakeColorBlend},
			{"ColorSubstitute", MakeColorSubstitute},
		}
	}
	return ColorFieldOptions
}

func MakeColorField(md, d int) ColorField {
	cf := GetColorFields()
	return cf[rand.Intn(len(cf))].Make(md, d+1)
}

// Leaves (don't call any fields)

func MakeGenerator(md, d int) Field {
	f := NewGenerator(PickLambda(), rand.Float64()*math.Pi*2, MakeGeneratorFunc())
	f.Phase = rand.Float64()
	f.FFunc = MakeFilter()
	return f
}

type GenFunc struct {
	Name string
	Make func() func(float64) float64
}

var GFOptions = []GenFunc{
	{"Zero", MakeZero},
	{"Sin", MakeSin},
	{"Square", MakeSquare},
	{"Triangle", MakeTriangle},
	{"Saw", MakeSaw},
	{"NL1", MakeNL1},
	{"NL2", MakeNL2},
}

func MakeGeneratorFunc() func(float64) float64 {
	return GFOptions[rand.Intn(len(GFOptions))].Make()
}

func MakeZero() func(float64) float64 {
	return Zero
}

func MakeSin() func(float64) float64 {
	return Sin
}

func MakeSquare() func(float64) float64 {
	return Square
}

func MakeTriangle() func(float64) float64 {
	return Triangle
}

func MakeSaw() func(float64) float64 {
	return Saw
}

func MakeNL1() func(float64) float64 {
	gf := &NLGF{MakeNL()}
	return gf.NL1
}

func MakeNL2() func(float64) float64 {
	gf := &NL2GF{MakeNL(), MakeNL()}
	return gf.NL2
}

type NLFunc struct {
	Name string
	Make func() util.NonLinear
}

var NLFOptions = []NLFunc{
	{"Linear", MakeNLLinear},
	{"Square", MakeNLSquare},
	{"Cube", MakeNLCube},
	{"Exponential", MakeNLExponential},
	{"Logarithmic", MakeNLLogarithmic},
	{"Sin", MakeNLSin},
	{"Sin1", MakeNLSin1},
	{"Sin2", MakeNLSin2},
	{"Circle1", MakeNLCircle1},
	{"Circle2", MakeNLCircle2},
	{"Catenary", MakeNLCatenary},
	{"Gauss", MakeNLGauss},
	{"Logistic", MakeNLLogistic},
	{"NLP3", MakeNLP3},
	{"NLP5", MakeNLP5},
}

func MakeNL() util.NonLinear {
	return NLFOptions[rand.Intn(len(NLFOptions))].Make()
}

func MakeNLLinear() util.NonLinear {
	return &util.NLLinear{}
}

func MakeNLSquare() util.NonLinear {
	return &util.NLSquare{}
}

func MakeNLCube() util.NonLinear {
	return &util.NLCube{}
}

func MakeNLExponential() util.NonLinear {
	return util.NewNLExponential(10)
}

func MakeNLLogarithmic() util.NonLinear {
	return util.NewNLLogarithmic(10)
}

func MakeNLSin() util.NonLinear {
	return &util.NLSin{}
}

func MakeNLSin1() util.NonLinear {
	return &util.NLSin1{}
}

func MakeNLSin2() util.NonLinear {
	return &util.NLSin2{}
}

func MakeNLCircle1() util.NonLinear {
	return &util.NLCircle1{}
}

func MakeNLCircle2() util.NonLinear {
	return &util.NLCircle2{}
}

func MakeNLCatenary() util.NonLinear {
	return &util.NLCatenary{}
}

func MakeNLGauss() util.NonLinear {
	return util.NewNLGauss(1)
}

func MakeNLLogistic() util.NonLinear {
	return util.NewNLLogistic(10, 0.5)
}

func MakeNLP3() util.NonLinear {
	return &util.NLP3{}
}

func MakeNLP5() util.NonLinear {
	return &util.NLP5{}
}

func MakePerlin(md, d int) Field {
	xfm := g2d.NewAff3()
	xfm.Scale(0.01, 0.01)
	f := NewPerlin()
	f.FFunc = MakeFilter()
	return &Transform{f, xfm}
}

func MakeNonLinear(md, d int) Field {
	f := NewNonLinear(PickLambda(), PickLambda(), rand.Float64()*math.Pi*2, MakeNL(), 10)
	f.PhaseX = rand.Float64()
	f.PhaseY = rand.Float64()
	switch rand.Intn(3) {
	case 0:
	case 1:
		f.OffsetX = 5
	case 2:
		f.OffsetY = 5
	}
	f.FFunc = MakeFilter()
	return f
}

// Processors

func MakeFractal(md, d int) Field {
	lac := 2.0
	hurst := 1.0
	octs := 3
	xfm := g2d.NewAff3()
	xfm.Scale(lac, lac)
	if rand.Intn(2) == 0 {
		fbm := NewFBM(hurst, lac)
		return &Fractal{MakeField(md, d+1), xfm, fbm.Combine, MakeFilter(), octs, 0, 0}
	}
	mf := NewMF(hurst, lac, 0.5)
	return &Fractal{MakeField(md, d+1), xfm, mf.Combine, MakeFilter(), octs, 0, 1}
}

func MakeDistort(md, d int) Field {
	return NewDistort(MakeField(md, d+1), 1)
}

func MakeSelect(md, d int) Field {
	return &Select{MakeVectorField(md, d+1), rand.Intn(3), MakeFilter()}
}

func MakeDirection(md, d int) Field {
	return &Direction{MakeVectorField(md, d+1), MakeFilter()}
}

func MakeMagnitude(md, d int) Field {
	return &Magnitude{MakeVectorField(md, d+1), MakeFilter()}
}

func MakeVectorCombine(md, d int) Field {
	return &VectorCombine{MakeVectorField(md, d+1), MakeCombiner3Func(), MakeFilter()}
}

func MakeNormal(md, d int) VectorField {
	return NewNormal(MakeField(md, d), 10, 10, 2, 2)
}

func MakeColorConv(md, d int) ColorField {
	return &ColorConv{MakeField(md, d+1), MakeColorNL()}
}

func MakeColorNL() *ColorNL {
	c1, c2, c3 := g2dcol.Random(), g2dcol.Random(), g2dcol.Random()
	tval := []float64{rand.Float64()}
	return NewColorNL(c1, c2, []color.Color{c3}, tval, MakeNL(), MakeLerp())
	return nil
}

func MakeLerp() func(float64, color.Color, color.Color) color.Color {
	lerp := ColorRGBALerp
	switch rand.Intn(3) {
	case 0:
	case 1:
		lerp = ColorHSLLerp
	case 2:
		lerp = ColorHSLLerpS
	}
	return lerp
}

// Combiners

func MakeCombiner2(md, d int) Field {
	return &Combiner2{MakeField(md, d+1), MakeField(md, d+1), MakeCombiner2Func(), MakeFilter()}
}

type CF2s struct {
	Name string
	Func func(float64, float64) float64
}

var CF2Options = []CF2s{
	{"Mul", Mul},
	{"Add", Add},
	{"Sub", Sub},
	{"Min", Min},
	{"Max", Max},
	{"Avg", Avg},
	{"Diff", Diff},
}

func MakeCombiner2Func() func(float64, float64) float64 {
	n := len(CF2Options)
	r := rand.Intn(n + 1)
	if r == n {
		s := rand.Float64()
		t := rand.Float64()
		e := t - s*(t+1)
		cf := &CF2{s, e}
		return cf.Substitute
	}
	return CF2Options[r].Func
}

func MakeCombiner3(md, d int) Field {
	return &Combiner3{MakeField(md, d+1), MakeField(md, d+1), MakeField(md, d+1), MakeCombiner3Func(), MakeFilter()}
}

func MakeCombiner3Func() func(float64, float64, float64) float64 {
	if rand.Intn(2) == 0 {
		return Blend
	}
	s := rand.Float64()
	t := rand.Float64()
	e := t - s*(t+1)
	cf := &CF3{s, e}
	return cf.Substitute
}

func MakeDisplace(md, d int) Field {
	return &Displace{MakeField(md, d+1), MakeField(md, d+1), MakeField(md, d+1), 10, 10}
}

func MakeColor(md, d int) ColorField {
	if rand.Intn(2) == 0 {
		return &Color{MakeField(md, d+1), nil, nil, nil, rand.Intn(2) == 0}
	}
	return &Color{MakeField(md, d+1), MakeField(md, d+1), MakeField(md, d+1), nil, rand.Intn(2) == 0}
}

func MakeColorBlend(md, d int) ColorField {
	return &ColorBlend{MakeColorField(md, d+1), MakeColorField(md, d+1), MakeField(md, d+1), MakeLerp()}
}

func MakeColorSubstitute(md, d int) ColorField {
	s := rand.Float64()
	t := rand.Float64()
	e := t - s*(1+t)
	return &ColorSubstitute{MakeColorField(md, d+1), MakeColorField(md, d+1), MakeField(md, d+1), s, e}
}

// Filters

func MakeFilter() func(float64) float64 {
	switch rand.Intn(20) {
	case 0:
		fv := &FilterVals{1, 0, 13}
		return fv.Quantize
	case 1:
		rv := NewRandFV(1, 0, 17)
		return rv.Quantize
	}
	return nil
}

var Lambdas = []float64{
	11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97,
}

func PickLambda() float64 {
	return Lambdas[rand.Intn(len(Lambdas))]
}
