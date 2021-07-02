package texture

import "image/color"

type Field interface {
	Eval2(x, y float64) float64
}

type ColorField interface {
	Eval2(x, y float64) color.Color
}

type VectorField interface {
	Eval2(x, y float64) []float64
}
