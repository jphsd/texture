package texture

import (
	"github.com/jphsd/graphics2d"
	"image/color"
)

type ShapeStyle int

// Constants for shape styles
const (
	BinaryStyle ShapeStyle = iota
	PathSumStyle
	PathOccStyle
)

type Shape struct {
	Name  string
	Shape *graphics2d.Shape
	Style ShapeStyle
}

func NewShape(shape *graphics2d.Shape, style ShapeStyle) *Shape {
	return &Shape{"Shape", shape, style}
}

func (s *Shape) Eval2(x, y float64) float64 {
	switch s.Style {
	default:
		fallthrough
	case BinaryStyle:
		return s.binary(x, y)
	case PathSumStyle:
		return s.summed(x, y)
	case PathOccStyle:
		return s.occurrence(x, y)
	}
}

// 1 if inside a path, -1 otherwise
func (s *Shape) binary(x, y float64) float64 {
	if s.Shape.PointInShape([]float64{x, y}) {
		return 1
	}
	return -1
}

// -1 if outside of any paths, otherwise path index it first occurs in
func (s *Shape) occurrence(x, y float64) float64 {
	// Treat the order of paths in a shape as a stack
	// and return a value depending on which path in the
	// shape the point hits first.
	paths := s.Shape.Paths()
	np := len(paths)
	dp := 2 / float64(np+1)
	v := 1.0
	for _, path := range paths {
		if path.PointInPath([]float64{x, y}) {
			return v
		}
		v -= dp
	}
	return -1
}

// -1 if outside of any paths, otherwise scaled by # paths it occurs in
func (s *Shape) summed(x, y float64) float64 {
	// Treat the order of paths in a shape as a stack
	// and return a value depending on which path in the
	// shape the point hits first.
	paths := s.Shape.Paths()
	np := len(paths)
	dp := 2 / float64(np+1)
	sum := -1.0
	for _, path := range paths {
		if path.PointInPath([]float64{x, y}) {
			sum += dp
		}
	}
	return sum
}

// Could do these with from primitive operations
type ShapeCombiner struct {
	Name  string
	Src1  Field
	Src2  Field
	Shape *graphics2d.Shape
}

func NewShapeCombiner(src1, src2 Field, shape *graphics2d.Shape) *ShapeCombiner {
	return &ShapeCombiner{"ShapeCombiner", src1, src2, shape}
}

func (s *ShapeCombiner) Eval2(x, y float64) float64 {
	if s.Shape.PointInShape([]float64{x, y}) {
		return s.Src1.Eval2(x, y)
	}
	return s.Src2.Eval2(x, y)
}

type ShapeCombinerCF struct {
	Name  string
	Src1  ColorField
	Src2  ColorField
	Shape *graphics2d.Shape
}

func NewShapeCombinerCF(src1, src2 ColorField, shape *graphics2d.Shape) *ShapeCombinerCF {
	return &ShapeCombinerCF{"ShapeCombinerCF", src1, src2, shape}
}

func (s *ShapeCombinerCF) Eval2(x, y float64) color.Color {
	if s.Shape.PointInShape([]float64{x, y}) {
		return s.Src1.Eval2(x, y)
	}
	return s.Src2.Eval2(x, y)
}

type ShapeCombinerVF struct {
	Name  string
	Src1  VectorField
	Src2  VectorField
	Shape *graphics2d.Shape
}

func NewShapeCombinerVF(src1, src2 VectorField, shape *graphics2d.Shape) *ShapeCombinerVF {
	return &ShapeCombinerVF{"ShapeCombinerVF", src1, src2, shape}
}

func (s *ShapeCombinerVF) Eval2(x, y float64) []float64 {
	if s.Shape.PointInShape([]float64{x, y}) {
		return s.Src1.Eval2(x, y)
	}
	return s.Src2.Eval2(x, y)
}
