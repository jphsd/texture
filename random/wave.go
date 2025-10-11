package random

import (
	//"fmt"
	"github.com/jphsd/texture"
	"math/rand"
)

func MakeWave() texture.Wave {
	if rand.Intn(2) > 0 {
		return MakePatternWave()
	}
	return MakeNLWave()
}

// MakePatternWave
func MakePatternWave() texture.Wave {
	nl := rand.Intn(5) + 1
	lambdas := make([]float64, nl)
	patterns := make([][]float64, nl)
	for i := 0; i < nl; i++ {
		lambdas[i] = PickLambda()
		patterns[i] = MakePattern(5)
	}
	return texture.NewPatternWave(lambdas, patterns, rand.Intn(2) > 0, false)
}

func MakePattern(n int) []float64 {
	pat := make([]float64, n)
	for i := 0; i < n; i++ {
		pat[i] = rand.Float64()*2 - 1
	}
	return pat
}

// MakeNLWave creates a new NLWave
func MakeNLWave() texture.Wave {
	nl := rand.Intn(5) + 1
	lambdas := make([]float64, nl)
	nlfs := make([]*texture.NonLinear, nl)
	for i := 0; i < nl; i++ {
		lambdas[i] = PickLambda()
		nlfs[i] = MakeNL()
	}
	return texture.NewNLWave(lambdas, nlfs, rand.Intn(2) > 0, false)
}

type NLFunc struct {
	Name string
	Make func() *texture.NonLinear
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

func MakeNL() *texture.NonLinear {
	res := NLFOptions[rand.Intn(len(NLFOptions))]
	return res.Make()
}

func MakeNLLinear() *texture.NonLinear {
	return texture.NewNLLinear()
}

func MakeNLSquare() *texture.NonLinear {
	return texture.NewNLSquare()
}

func MakeNLCube() *texture.NonLinear {
	return texture.NewNLCube()
}

func MakeNLExponential() *texture.NonLinear {
	return texture.NewNLExponential(10)
}

func MakeNLLogarithmic() *texture.NonLinear {
	return texture.NewNLLogarithmic(10)
}

func MakeNLSin() *texture.NonLinear {
	return texture.NewNLSin()
}

func MakeNLSin1() *texture.NonLinear {
	return texture.NewNLSin1()
}

func MakeNLSin2() *texture.NonLinear {
	return texture.NewNLSin2()
}

func MakeNLCircle1() *texture.NonLinear {
	return texture.NewNLCircle1()
}

func MakeNLCircle2() *texture.NonLinear {
	return texture.NewNLCircle2()
}

func MakeNLCatenary() *texture.NonLinear {
	return texture.NewNLCatenary()
}

func MakeNLGauss() *texture.NonLinear {
	return texture.NewNLGauss(1)
}

func MakeNLLogistic() *texture.NonLinear {
	return texture.NewNLLogistic(10, 0.5)
}

func MakeNLP3() *texture.NonLinear {
	return texture.NewNLP3()
}

func MakeNLP5() *texture.NonLinear {
	return texture.NewNLP5()
}

var Lambdas = []float64{
	11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97,
}

func PickLambda() float64 {
	return Lambdas[rand.Intn(len(Lambdas))]
}
