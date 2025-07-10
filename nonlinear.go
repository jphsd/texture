package texture

import "github.com/jphsd/nonlinear"

// Wrap the NLs into named ones for marshalling

type NonLinear struct {
	Name string
	NLF  nonlinear.NonLinear
}

func (nl *NonLinear) Eval0(t float64) float64 {
	return nl.NLF.Transform(t)
}

func (nl *NonLinear) Eval(t float64) float64 {
	// nonlinear.NonLinear is in range [0,1], map to [-1,1]
	return nl.NLF.Transform(t)*2 - 1
}

func NewNLLinear() *NonLinear {
	return &NonLinear{"NLLinear", &nonlinear.NLLinear{}}
}

func NewNLSquare() *NonLinear {
	return &NonLinear{"NLSquare", &nonlinear.NLSquare{}}
}

func NewNLCube() *NonLinear {
	return &NonLinear{"NLCube", &nonlinear.NLCube{}}
}

func NewNLExponential(v float64) *NonLinear {
	return &NonLinear{"NLExponential", nonlinear.NewNLExponential(v)}
}

func NewNLLogarithmic(v float64) *NonLinear {
	return &NonLinear{"NLLogarithmic", nonlinear.NewNLLogarithmic(v)}
}

func NewNLSin() *NonLinear {
	return &NonLinear{"NLSin", &nonlinear.NLSin{}}
}

func NewNLSin1() *NonLinear {
	return &NonLinear{"NLSin1", &nonlinear.NLSin1{}}
}

func NewNLSin2() *NonLinear {
	return &NonLinear{"NLSin2", &nonlinear.NLSin2{}}
}

func NewNLCircle1() *NonLinear {
	return &NonLinear{"NLCircle1", &nonlinear.NLCircle1{}}
}

func NewNLCircle2() *NonLinear {
	return &NonLinear{"NLCircle2", &nonlinear.NLCircle2{}}
}

func NewNLCatenary() *NonLinear {
	return &NonLinear{"NLCatenary", &nonlinear.NLCatenary{}}
}

func NewNLGauss(v float64) *NonLinear {
	return &NonLinear{"NLGauss", nonlinear.NewNLGauss(v)}
}

func NewNLLogistic(u, v float64) *NonLinear {
	return &NonLinear{"NLLogistic", nonlinear.NewNLLogistic(v, v)}
}

func NewNLP3() *NonLinear {
	return &NonLinear{"NLP3", &nonlinear.NLP3{}}
}

func NewNLP5() *NonLinear {
	return &NonLinear{"NLP5", &nonlinear.NLP5{}}
}
