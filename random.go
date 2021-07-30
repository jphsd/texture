package texture

import (
	"fmt"
	g2d "github.com/jphsd/graphics2d"
	g2dcol "github.com/jphsd/graphics2d/color"
	"github.com/jphsd/graphics2d/util"
	"image/color"
	"math"
	"math/rand"
)

// Leaf describes a Field that has no predecessors.
type Leaf struct {
	Name string
	Make func() Field
}

var LeafOptions = []Leaf{
	{"Generator", MakeGenerator},
	{"Perlin", MakePerlin},
	{"DistortedPerlin", MakeDistortedPerlin},
	{"NonLinear", MakeNonLinear},
}

// Node describes a Field that has predecessors.
type Node struct {
	Name string
	Make func(int, int) Field
}

var NodeOptions []Node

// GetNodes returns a slice of the available nodes.
func GetNodes() []Node {
	// Lazy initialization to prevent compiler complaining
	// about initialization loops
	if NodeOptions == nil {
		NodeOptions = []Node{
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

// MakeLeaf creates a new leaf.
func MakeLeaf() Field {
	res := LeafOptions[rand.Intn(len(LeafOptions))]
	fmt.Printf("%s\n", res.Name)
	return res.Make()
}

// MakeNode creates a new node.
func MakeNode(md, d int) Field {
	n := GetNodes()
	return n[rand.Intn(len(n))].Make(md, d+1)
}

// MakeField creates either a new leaf or node.
func MakeField(md, d int) Field {
	l := LeafOptions
	if d >= md {
		return l[rand.Intn(len(l))].Make()
	}
	n := GetNodes()
	s := len(l) + len(n)
	s = rand.Intn(s)
	if s < len(l) {
		return l[s].Make()
	}
	s -= len(l)
	return n[s].Make(md, d+1)
}

// MakeVectorField creates a new VectorField.
func MakeVectorField(md, d int) VectorField {
	return MakeNormal(md, d+1)
}

// ColorFieldOpts describes the available ColorField functions.
type ColorFieldOpts struct {
	Name string
	Make func(int, int) ColorField
}

var ColorFieldOptions []ColorFieldOpts

// GetColorFields returns the list of ColorField functions.
func GetColorFields() []ColorFieldOpts {
	// Lazy initialization to prevent compiler complaining
	// about initialization loops
	if ColorFieldOptions == nil {
		ColorFieldOptions = []ColorFieldOpts{
			{"Color", MakeColor},
			{"ColorSinCos", MakeColorSinCos},
			{"ColorFields", MakeColorFields},
			{"ColorConv", MakeColorConv},
			{"ColorBlend", MakeColorBlend},
			{"ColorSubstitute", MakeColorSubstitute},
		}
	}
	return ColorFieldOptions
}

// MakeColorField creates a new color field.
func MakeColorField(md, d int) ColorField {
	cf := GetColorFields()
	return cf[rand.Intn(len(cf))].Make(md, d+1)
}

// The following are Leaves (don't call any fields)

// Make a new field generator.
func MakeGenerator() Field {
	f := NewGenerator(PickLambda(), rand.Float64()*math.Pi*2, MakeGeneratorFunc())
	f.Phase = rand.Float64()
	f.FFunc = MakeFilter()
	return f
}

// GenFunc describes generator functions.
type GenFunc struct {
	Name string
	Make func() func(float64) float64
}

var GFOptions = []GenFunc{
	{"Flat", MakeFlat},
	{"Sin", MakeSin},
	{"Square", MakeSquare},
	{"Triangle", MakeTriangle},
	{"Saw", MakeSaw},
	{"NL1", MakeNL1},
	{"NL2", MakeNL2},
	{"N1D", MakeN1D},
}

// MakeGeneratorFunc creates a generator function.
func MakeGeneratorFunc() func(float64) float64 {
	res := GFOptions[rand.Intn(len(GFOptions))]
	fmt.Printf(" GF: %s\n", res.Name)
	return res.Make()
}

func MakeFlat() func(float64) float64 {
	fgf := &FlatGF{rand.Float64()*2 - 1}
	return fgf.Flat
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
	{"NLRand", MakeNLRand},
}

func MakeNL() util.NonLinear {
	res := NLFOptions[rand.Intn(len(NLFOptions))]
	fmt.Printf(" %s\n", res.Name)
	return res.Make()
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

func MakeNLRand() util.NonLinear {
	return util.NewNLRand(0.1, 0.05, rand.Intn(2) == 0)
}

func MakeN1D() func(float64) float64 {
	gf := NewNoise1DGF(6)
	return gf.Noise1D
}

// MakePerlin creates a new field backed by a perlin noise function.
func MakePerlin() Field {
	xfm := g2d.NewAff3()
	xfm.Scale(0.01, 0.01)
	f := NewPerlin()
	f.FFunc = MakeFilter()
	return &Transform{f, xfm}
}

// MakeDistortedPerlin creates a new field backed by a perlin noise function.
func MakeDistortedPerlin() Field {
	xfm := g2d.NewAff3()
	xfm.Scale(0.01, 0.01)
	f := NewDistort(NewPerlin(), 1)
	f.FFunc = MakeFilter()
	return &Transform{f, xfm}
}

// MakeNonLinear creates a set of circles on the plane using nonlnear functions.
func MakeNonLinear() Field {
	f := NewNonLinear(PickLambda(), PickLambda(), rand.Float64()*math.Pi*2, MakeNL(), 2)
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

// Processors - these all require a source input of one form or another.

// MakeFractal creates a new fractal processor.
func MakeFractal(md, d int) Field {
	lac := 2.0
	hurst := 1.0
	octs := float64(rand.Intn(3))
	xfm := g2d.NewAff3()
	xfm.Scale(lac, lac)
	xfm.Rotate(rand.Float64() * math.Pi)
	if rand.Intn(2) == 0 {
		fbm := NewFBM(hurst, lac)
		res := NewFractal(MakeField(md, d+1), xfm, fbm.Combine, octs)
		res.FFunc = MakeFilter()
		return res
	}
	mf := NewMF(hurst, lac, 0.5)
	res := NewFractal(MakeField(md, d+1), xfm, mf.Combine, octs)
	res.FFunc = MakeFilter()
	return res
}

// MakeDistortion creates a new distorted processor.
func MakeDistort(md, d int) Field {
	return NewDistort(MakeField(md, d+1), 1)
}

// MakeSelect creates a new field from a vector field.
func MakeSelect(md, d int) Field {
	return &Select{MakeVectorField(md, d+1), rand.Intn(3), MakeFilter()}
}

// MakeDirection creates a new field from a vector field.
func MakeDirection(md, d int) Field {
	return &Direction{MakeVectorField(md, d+1), MakeFilter()}
}

// MakeMagnitude creates a new field from a vector field.
func MakeMagnitude(md, d int) Field {
	return &Magnitude{MakeVectorField(md, d+1), MakeFilter()}
}

// MakeVectorCombine creates a new field from a vector field.
func MakeVectorCombine(md, d int) Field {
	return &VectorCombine{MakeVectorField(md, d+1), MakeCombiner3Func(), MakeFilter()}
}

// MakeNormal creates a new vector field from a field.
func MakeNormal(md, d int) VectorField {
	return NewNormal(MakeField(md, d), 10, 10, 2, 2)
}

// MakeColorConv creates a new color field from a field.
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

// MakeColor creates a new color field from a field.
func MakeColor(md, d int) ColorField {
	return &Color{MakeField(md, d+1)}
}

// MakeColorSinCos creates a new color field from a field.
func MakeColorSinCos(md, d int) ColorField {
	return &ColorSinCos{MakeField(md, d+1), rand.Intn(6), rand.Intn(2) == 0}
}

// Combiners - these all require source inputs of one form or another.

// MakeCombiner2 creates a two field combiner.
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

// MakeCombiner3 creates a three field combiner.
func MakeCombiner3(md, d int) Field {
	return &Combiner3{MakeField(md, d+1), MakeField(md, d+1), MakeField(md, d+1), MakeCombiner3Func(), MakeFilter()}
}

func MakeCombiner3Func() func(...float64) float64 {
	if rand.Intn(2) == 0 {
		return Blend
	}
	s := rand.Float64()
	t := rand.Float64()
	e := t - s*(t+1)
	cf := &CF3{s, e}
	return cf.Substitute
}

// MakeDisplace creates a displacement of src1 with src2 and src3.
func MakeDisplace(md, d int) Field {
	return NewDisplace(MakeField(md, d+1), MakeField(md, d+1), MakeField(md, d+1), 10)
}

// MakeColorFields creates a color field from three fields.
func MakeColorFields(md, d int) ColorField {
	return &ColorFields{MakeField(md, d+1), MakeField(md, d+1), MakeField(md, d+1), nil, rand.Intn(2) == 0}
}

// MakeColorBlend creates a color field from two input color fields and a field.
func MakeColorBlend(md, d int) ColorField {
	return &ColorBlend{MakeColorField(md, d+1), MakeColorField(md, d+1), MakeField(md, d+1), MakeLerp()}
}

// MakeColorSubstitute creates a color field from two input color fields and a field.
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
		fv := &FilterVals{1, 0, int(PickLambda())}
		return fv.Quantize
	case 1:
		rv := NewRandFV(1, 0, int(PickLambda()))
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

// Component

// MakeComponent creates a new component.
func MakeComponent() *Component {
	// 2 fractals feeding displacement
	disp := MakeComponentFractal() // Use same for both x and y
	amt := rand.Float64()*10 + 1
	src := NewDisplace(MakeComponentFractal(), disp, disp, amt)

	// Emit color, alpha and bump map
	c1, c2, c3 := g2dcol.Random(), g2dcol.Random(), g2dcol.Random()
	return NewComponent(src, c1, c2, c3, MakeNL(), MakeLerp(), 20)
}

func MakeComponentFractal() Field {
	// 1 or 2 leaves, combined
	var src Field
	if rand.Intn(2) == 0 {
		src = MakeLeaf()
	} else {
		src = &Combiner2{MakeLeaf(), MakeLeaf(), MakeCombiner2Func(), MakeFilter()}
	}

	// fractal
	lac := 2.0
	hurst := 1.0
	octs := float64(rand.Intn(3))
	xfm := g2d.NewAff3()
	xfm.Scale(lac, lac)
	xfm.Rotate(rand.Float64() * math.Pi)
	if rand.Intn(2) == 0 {
		fbm := NewFBM(hurst, lac)
		ff := NewFractal(src, xfm, fbm.Combine, octs)
		ff.FFunc = MakeFilter()
		return ff
	}
	mf := NewMF(hurst, lac, 0.5)
	ff := NewFractal(src, xfm, mf.Combine, octs)
	ff.FFunc = MakeFilter()
	return ff
}
