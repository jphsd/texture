package texture

import "math"

type LinearGradient struct {
	Name string
	WF   Wave
}

func NewLinearGradient(wf Wave) *LinearGradient {
	return &LinearGradient{"LinearGradient", wf}
}

func (g *LinearGradient) Eval2(x, y float64) float64 {
	return g.WF.Eval(x)
}

type RadialGradient struct {
	Name string
	WF   Wave
}

func NewRadialGradient(wf Wave) *RadialGradient {
	return &RadialGradient{"RadialGradient", wf}
}

func (g *RadialGradient) Eval2(x, y float64) float64 {
	v := math.Sqrt(x*x + y*y)
	return g.WF.Eval(v)
}

type ConicGradient struct {
	Name string
	WF   Wave
}

func NewConicGradient(wf Wave) *ConicGradient {
	return &ConicGradient{"ConicGradient", wf}
}

func (g *ConicGradient) Eval2(x, y float64) float64 {
	// Convert [-pi,pi] => [0,lambda)
	v := (math.Atan2(y, x)/math.Pi + 1) * g.WF.Lambda() / 2
	return g.WF.Eval(v)
}
