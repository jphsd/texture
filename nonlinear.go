package texture

import (
	"github.com/jphsd/graphics2d/util"
)

type NonLinear struct {
	Name string
	NLF  util.NonLinear
}

func (nl *NonLinear) Eval(t float64) float64 {
	// util.NonLinear is in range [0,1], mpa to [-1,1]
	return nl.NLF.Transform(t)*2 - 1
}

func NewNLLinear() *NonLinear {
	return &NonLinear{"NLLinear", &util.NLLinear{}}
}

func NewNLSquare() *NonLinear {
	return &NonLinear{"NLSquare", &util.NLSquare{}}
}

func NewNLCube() *NonLinear {
	return &NonLinear{"NLCube", &util.NLCube{}}
}

func NewNLExponential(v float64) *NonLinear {
	return &NonLinear{"NLExponential", util.NewNLExponential(v)}
}

func NewNLLogarithmic(v float64) *NonLinear {
	return &NonLinear{"NLLogarithmic", util.NewNLLogarithmic(v)}
}

func NewNLSin() *NonLinear {
	return &NonLinear{"NLSin", &util.NLSin{}}
}

func NewNLSin1() *NonLinear {
	return &NonLinear{"NLSin1", &util.NLSin1{}}
}

func NewNLSin2() *NonLinear {
	return &NonLinear{"NLSin2", &util.NLSin2{}}
}

func NewNLCircle1() *NonLinear {
	return &NonLinear{"NLCircle1", &util.NLCircle1{}}
}

func NewNLCircle2() *NonLinear {
	return &NonLinear{"NLCircle2", &util.NLCircle2{}}
}

func NewNLCatenary() *NonLinear {
	return &NonLinear{"NLCatenary", &util.NLCatenary{}}
}

func NewNLGauss(v float64) *NonLinear {
	return &NonLinear{"NLGauss", util.NewNLGauss(v)}
}

func NewNLLogistic(u, v float64) *NonLinear {
	return &NonLinear{"NLLogistic", util.NewNLLogistic(v, v)}
}

func NewNLP3() *NonLinear {
	return &NonLinear{"NLP3", &util.NLP3{}}
}

func NewNLP5() *NonLinear {
	return &NonLinear{"NLP5", &util.NLP5{}}
}

func NewNLRand(u, v float64, b bool) *NonLinear {
	return &NonLinear{"NLRand", util.NewNLRand(u, v, b)}
}
