package surface

import "math"

func Dot(a, b []float64) float64 {
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2]
}

func Norm(v []float64) []float64 {
	sum := math.Sqrt(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])
	return []float64{v[0] / sum, v[1] / sum, v[2] / sum}
}

// Return v reflected in n
func Reflect(v, n []float64) []float64 {
	s := Dot(v, n)
	s *= 2
	return []float64{s*n[0] - v[0], s*n[1] - v[1], s*n[2] - v[2]}
}

func Cross(a, b []float64) []float64 {
	return []float64{
		a[1]*b[2] - a[2]*b[1],
		-a[0]*b[2] + a[2]*b[0],
		a[0]*b[1] - a[1]*b[0]}
}
