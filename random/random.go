package random

import (
	//"fmt"
	g2d "github.com/jphsd/graphics2d"
	g2dcol "github.com/jphsd/graphics2d/color"
	. "github.com/jphsd/texture"
	"image"
	"math"
	"math/rand"
)

// Leaf describes a Field that has no predecessors.
type Leaf struct {
	Name string
	Make func() Field
}

var LeafOptions = []Leaf{
	//{"Uniform", MakeUniform},
	{"LinearGradient", MakeLinearGradient},
	{"RadialGradient", MakeRadialGradient},
	{"ConicGradient", MakeConicGradient},
	{"TiledGradient", MakeTiledGradient},
	{"Binary", MakeBinary},
	{"Perlin", MakePerlin},
	{"DistortedPerlin", MakeDistortedPerlin},
	{"Image", MakeImage},
	//{"Shape", MakeShape},
}

// MakeLeaf creates a new leaf.
func MakeLeaf() Field {
	res := LeafOptions[rand.Intn(len(LeafOptions))]
	return res.Make()
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
			//{"Transform", MakeTransform}, // Normally combined with something else
			{"Strip", MakeStrip},
			{"Distort", MakeDistort},
			{"Displace", MakeDisplace},
			{"Reflect", MakeReflect},
			{"NLFilter", MakeNLFilter},
			{"QuantizeFilter", MakeQuantizeFilter},
			{"RandQuantFilter", MakeRandQuantFilter},
			{"MulCombiner", MakeMulCombiner},
			{"AddCombiner", MakeAddCombiner},
			{"SubCombiner", MakeSubCombiner},
			{"MinCombiner", MakeMinCombiner},
			{"MaxCombiner", MakeMaxCombiner},
			{"DiffCombiner", MakeDiffCombiner},
			{"WindowedCombiner", MakeWindowedCombiner},
			{"WeightedCombiner", MakeWeightedCombiner},
			{"Blend", MakeBlend},
			{"StochasticBlend", MakeStochasticBlend},
			{"JitterBlend", MakeJitterBlend},
			{"SubstituteCombiner", MakeSubstituteCombiner},
			{"Select", MakeSelect},
			{"Direction", MakeDirection},
			{"Magnitude", MakeMagnitude},
			{"Fractal", MakeFractal},
			{"VariableFractal", MakeVariableFractal},
			{"Morphological", MakeMorphological},
		}
	}
	return NodeOptions
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
			{"ColorGray", MakeColorGray},
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

// MakeUniform creates a new Uniform
func MakeUniform() Field {
	return NewUniform(rand.Float64()*2 - 1)
}

func MakeLinearGradient() Field {
	f := NewLinearGradient(MakeWave())

	// Wrap it in a Transform
	xfm := g2d.NewAff3()
	offs := rand.Float64()*200 - 100
	rot := rand.Float64() * math.Pi * 2
	xfm.Rotate(rot)
	xfm.Translate(offs, 0)
	return NewTransform(f, xfm)
}

func MakeRadialGradient() Field {
	w := MakeWave()
	f := NewRadialGradient(w)

	// Wrap it in a Transform
	xfm := g2d.NewAff3()
	offs := -400.0 // Hack alert - assumes 800x800 image
	rot := rand.Float64() * math.Pi * 2
	xfm.Rotate(rot)
	xfm.Translate(offs, offs)
	return NewTransform(f, xfm)
}

func MakeConicGradient() Field {
	w := MakeWave()
	f := NewConicGradient(w)

	// Wrap it in a Transform
	xfm := g2d.NewAff3()
	offs := -400.0 // Hack alert - assumes 800x800 image
	rot := rand.Float64() * math.Pi * 2
	xfm.Rotate(rot)
	xfm.Translate(offs, offs)
	return NewTransform(f, xfm)
}

func MakeTiledGradient() Field {
	var wave Wave

	//Select wave type
	if rand.Intn(2) > 0 {
		w := MakePatternWave().(*PatternWave)
		w.Once = true
		wave = w
	} else {
		w := MakeNLWave().(*NLWave)
		w.Once = true
		wave = w
	}

	var field Field
	// Select gradient type
	if rand.Intn(2) > 0 {
		field = NewRadialGradient(wave)
	} else {
		field = NewConicGradient(wave)
	}

	// Move 0,0 to center
	l := wave.Lambda()
	xfm := g2d.NewAff3()
	xfm.Translate(-l, -l)
	field = NewTransform(field, xfm)

	// Tile the single wave instance
	l *= 2
	field = NewTiler(field, []float64{l, l})

	// Wrap it in a Transform
	xfm = g2d.NewAff3()
	offs := rand.Float64()*200 - 100
	rot := rand.Float64() * math.Pi * 2
	xfm.Rotate(rot)
	xfm.Translate(offs, 0)
	return NewTransform(field, xfm)
}

func MakeWave() Wave {
	if rand.Intn(2) > 0 {
		return MakePatternWave()
	}
	return MakeNLWave()
}

// MakePatternWave
func MakePatternWave() Wave {
	nl := rand.Intn(5) + 1
	lambdas := make([]float64, nl)
	patterns := make([][]float64, nl)
	for i := 0; i < nl; i++ {
		lambdas[i] = PickLambda()
		patterns[i] = MakePattern(5)
	}
	return NewPatternWave(lambdas, patterns, rand.Intn(2) > 0, false)
}

func MakePattern(n int) []float64 {
	pat := make([]float64, n)
	for i := 0; i < n; i++ {
		pat[i] = rand.Float64()*2 - 1
	}
	return pat
}

// MakeNLWave creates a new NLWave
func MakeNLWave() Wave {
	nl := rand.Intn(5) + 1
	lambdas := make([]float64, nl)
	nlfs := make([]*NonLinear, nl)
	for i := 0; i < nl; i++ {
		lambdas[i] = PickLambda()
		nlfs[i] = MakeNL()
	}
	return NewNLWave(lambdas, nlfs, rand.Intn(2) > 0, false)
}

type NLFunc struct {
	Name string
	Make func() *NonLinear
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

func MakeNL() *NonLinear {
	res := NLFOptions[rand.Intn(len(NLFOptions))]
	return res.Make()
}

func MakeNLLinear() *NonLinear {
	return NewNLLinear()
}

func MakeNLSquare() *NonLinear {
	return NewNLSquare()
}

func MakeNLCube() *NonLinear {
	return NewNLCube()
}

func MakeNLExponential() *NonLinear {
	return NewNLExponential(10)
}

func MakeNLLogarithmic() *NonLinear {
	return NewNLLogarithmic(10)
}

func MakeNLSin() *NonLinear {
	return NewNLSin()
}

func MakeNLSin1() *NonLinear {
	return NewNLSin1()
}

func MakeNLSin2() *NonLinear {
	return NewNLSin2()
}

func MakeNLCircle1() *NonLinear {
	return NewNLCircle1()
}

func MakeNLCircle2() *NonLinear {
	return NewNLCircle2()
}

func MakeNLCatenary() *NonLinear {
	return NewNLCatenary()
}

func MakeNLGauss() *NonLinear {
	return NewNLGauss(1)
}

func MakeNLLogistic() *NonLinear {
	return NewNLLogistic(10, 0.5)
}

func MakeNLP3() *NonLinear {
	return NewNLP3()
}

func MakeNLP5() *NonLinear {
	return NewNLP5()
}

func MakeNLRand() *NonLinear {
	return NewNLRand(0.1, 0.05, rand.Intn(2) == 0)
}

// MakeBinary creates a new field backed by bit noise.
func MakeBinary() Field {
	xfm := g2d.NewAff3()
	xfm.Scale(0.01, 0.01)
	xfm.Rotate(rand.Float64() * math.Pi * 2)
	f := NewBinary(16, 16, rand.Int63())
	return NewTransform(f, xfm)
}

// MakePerlin creates a new field backed by a perlin noise function.
func MakePerlin() Field {
	xfm := g2d.NewAff3()
	xfm.Scale(0.01, 0.01)
	xfm.Rotate(rand.Float64() * math.Pi * 2)
	f := NewPerlin(rand.Int63())
	return NewTransform(f, xfm)
}

// MakeDistortedPerlin creates a new field backed by a perlin noise function.
func MakeDistortedPerlin() Field {
	return NewDistort(MakePerlin(), 1)
}

var Sample image.Image

func MakeImage() Field {
	if Sample == nil {
		return MakeUniform()
	}

	iw, ih := Sample.Bounds().Dx(), Sample.Bounds().Dy()

	// Make new field from img
	f1 := NewImage(Sample)
	f2 := NewColorToGray(f1)
	f3 := NewTiler(f2, []float64{float64(iw), float64(ih)})
	xfm := g2d.NewAff3()
	xfm.Rotate(rand.Float64() * math.Pi * 2)
	return NewTransform(f3, xfm)
}

// Processors - these all require a source input of one form or another.

// MakeTransform creates a new transform processor
func MakeTransform(md, d int) Field {
	xfm := g2d.NewAff3()
	offs := rand.Float64()*200 - 100
	sx := rand.Float64()*4 - 2
	sy := rand.Float64()*4 - 2
	rot := rand.Float64() * math.Pi * 2
	xfm.Scale(sx, sy)
	xfm.Rotate(rot)
	xfm.Translate(offs, 0)
	return NewTransform(MakeField(md, d+1), xfm)
}

// MakeStrip creates a new strip processor
func MakeStrip(md, d int) Field {
	xfm := g2d.NewAff3()
	xfm.Rotate(rand.Float64() * math.Pi * 2)
	y := rand.Float64()*2000 - 1000
	return NewTransform(NewStrip(MakeField(md, d+1), y), xfm)
}

// MakeFractal creates a new fractal processor.
func MakeFractal(md, d int) Field {
	lac := 2.0
	hurst := 1.0
	nocts := rand.Intn(3)
	octs := float64(nocts)
	xfm := g2d.NewAff3()
	xfm.Scale(lac, lac)
	xfm.Rotate(rand.Float64() * math.Pi)
	if rand.Intn(2) == 0 {
		fbm := NewFBM(hurst, lac, nocts)
		res := NewFractal(MakeField(md, d+1), xfm, fbm, octs)
		return res
	}
	mf := NewMF(hurst, lac, 0.5, nocts)
	res := NewFractal(MakeField(md, d+1), xfm, mf, octs)
	return res
}

// MakeVariableFractal creates a new fractal processor.
func MakeVariableFractal(md, d int) Field {
	lac := 2.0
	hurst := 1.0
	nocts := rand.Intn(5)
	octs := float64(nocts)
	xfm := g2d.NewAff3()
	xfm.Scale(lac, lac)
	xfm.Rotate(rand.Float64() * math.Pi)
	if rand.Intn(2) == 0 {
		fbm := NewFBM(hurst, lac, nocts)
		res := NewFractal(MakeField(md, d+1), xfm, fbm, octs)
		return res
	}
	mf := NewMF(hurst, lac, 0.5, nocts)
	res := NewVariableFractal(MakeField(md, d+1), xfm, mf, MakeField(md, d+1), octs)
	return res
}

// MakeDistort creates a new distorted processor.
func MakeDistort(md, d int) Field {
	return NewDistort(MakeField(md, d+1), 1)
}

// MakeSelect creates a new field from a vector field.
func MakeSelect(md, d int) Field {
	return NewSelect(MakeVectorField(md, d+1), rand.Intn(3))
}

// MakeDirection creates a new field from a vector field.
func MakeDirection(md, d int) Field {
	return NewDirection(MakeVectorField(md, d+1))
}

// MakeMagnitude creates a new field from a vector field.
func MakeMagnitude(md, d int) Field {
	return NewMagnitude(MakeVectorField(md, d+1))
}

// MakeReflect creates a mirror plane in the field.
func MakeReflect(md, d int) Field {
	a := rand.Float64() * math.Pi * 2
	dx, dy := math.Cos(a)*400, math.Sin(a)*400 // Hack alert - assumes 800x800
	return NewReflect(MakeField(md, d+1), []float64{400, 400}, []float64{400 + dx, 400 + dy})
}

// MakeNLFilter creates a non-linear filter.
func MakeNLFilter(md, d int) Field {
	return NewNLFilter(MakeField(md, d+1), MakeNL(), 1, 0)
}

// MakeQuantizeFilter creates a non-linear filter.
func MakeQuantizeFilter(md, d int) Field {
	return NewQuantizeFilter(MakeField(md, d+1), 1, 0, int(PickLambda()))
}

// MakeRandQuantFilter creates a non-linear filter.
func MakeRandQuantFilter(md, d int) Field {
	return NewRandQuantFilter(MakeField(md, d+1), 1, 0, int(PickLambda()))
}

// MakeNormal creates a new vector field from a field.
func MakeNormal(md, d int) VectorField {
	return NewNormal(MakeField(md, d), 10, 10, 2, 2)
}

// MakeColorConv creates a new color field from a field.
func MakeColorConv(md, d int) ColorField {
	return NewColorConv(MakeField(md, d+1), g2dcol.Random(), g2dcol.Random(), nil, nil, LerpType(rand.Intn(3)))
}

// MakeColorGray creates a new color field from a field.
func MakeColorGray(md, d int) ColorField {
	return NewColorGray(MakeField(md, d+1))
}

// MakeColorSinCos creates a new color field from a field.
func MakeColorSinCos(md, d int) ColorField {
	return NewColorSinCos(MakeField(md, d+1), rand.Intn(6), rand.Intn(2) == 0)
}

// Combiners - these all require source inputs of one form or another.

func MakeMulCombiner(md, d int) Field {
	return NewMulCombiner(MakeField(md, d+1), MakeField(md, d+1))
}

func MakeAddCombiner(md, d int) Field {
	return NewAddCombiner(MakeField(md, d+1), MakeField(md, d+1))
}

func MakeSubCombiner(md, d int) Field {
	return NewSubCombiner(MakeField(md, d+1), MakeField(md, d+1))
}

func MakeMinCombiner(md, d int) Field {
	return NewMinCombiner(MakeField(md, d+1), MakeField(md, d+1))
}

func MakeMaxCombiner(md, d int) Field {
	return NewMaxCombiner(MakeField(md, d+1), MakeField(md, d+1))
}

func MakeDiffCombiner(md, d int) Field {
	return NewDiffCombiner(MakeField(md, d+1), MakeField(md, d+1))
}

func MakeWindowedCombiner(md, d int) Field {
	return NewWindowedCombiner(MakeField(md, d+1), MakeField(md, d+1), -1/3, 1/3)
}

func MakeWeightedCombiner(md, d int) Field {
	return NewWeightedCombiner(MakeField(md, d+1), MakeField(md, d+1), 0.75, 0.25)
}

func MakeBlend(md, d int) Field {
	return NewBlend(MakeField(md, d+1), MakeField(md, d+1), MakeField(md, d+1))
}

func MakeStochasticBlend(md, d int) Field {
	return NewStochasticBlend(MakeField(md, d+1), MakeField(md, d+1), MakeField(md, d+1))
}

func MakeJitterBlend(md, d int) Field {
	return NewJitterBlend(MakeField(md, d+1), MakeField(md, d+1), MakeField(md, d+1), 0.1)
}

func MakeSubstituteCombiner(md, d int) Field {
	return NewSubstituteCombiner(MakeField(md, d+1), MakeField(md, d+1), MakeField(md, d+1), -1/3, 1/3)
}

// MakeDisplace creates a displacement of src1 with src2 and src3.
func MakeDisplace(md, d int) Field {
	return NewDisplace(MakeField(md, d+1), MakeField(md, d+1), MakeField(md, d+1), 10)
}

// MakeColorFields creates a color field from three fields.
func MakeColorFields(md, d int) ColorField {
	return NewColorFields(MakeField(md, d+1), MakeField(md, d+1), MakeField(md, d+1), nil, rand.Intn(2) == 0)
}

// MakeColorBlend creates a color field from two input color fields and a field.
func MakeColorBlend(md, d int) ColorField {
	return NewColorBlend(MakeColorField(md, d+1), MakeColorField(md, d+1), MakeField(md, d+1), LerpType(rand.Intn(3)))
}

// MakeColorSubstitute creates a color field from two input color fields and a field.
func MakeColorSubstitute(md, d int) ColorField {
	s := rand.Float64()
	t := rand.Float64()
	e := t - s*(1+t)
	return NewColorSubstitute(MakeColorField(md, d+1), MakeColorField(md, d+1), MakeField(md, d+1), s, e)
}

// Filters

var Lambdas = []float64{
	11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97,
}

func PickLambda() float64 {
	return Lambdas[rand.Intn(len(Lambdas))]
}

// Component

// MakeComponent creates a new component.
func MakeComponent() *Component {
	// 2 fields feeding displacement
	disp := MakeField(2, 0)
	amt := rand.Float64()*10 + 1
	src := NewDisplace(MakeField(6, 0), disp, disp, amt)

	// Emit color, alpha and bump map
	c1, c2, c3 := g2dcol.Random(), g2dcol.Random(), g2dcol.Random()
	return NewComponent(src, c1, c2, c3, LerpType(rand.Intn(3)), 20)
}
