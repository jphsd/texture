package texture

import "image/color"

// Field defines an evaluation method that given an x and y value, returns a value in the range [-1,1].
type Field interface {
	Eval2(x, y float64) float64
}

// ColorField defines an evaluation method that given an x and y value, returns a Color.
type ColorField interface {
	Eval2(x, y float64) color.Color
}

// VectorField defines an evaluation method that given an x and y value, returns a slice of values in the range [-1,1].
type VectorField interface {
	Eval2(x, y float64) []float64
}
