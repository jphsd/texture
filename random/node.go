package random

import (
	//"fmt"
	"github.com/jphsd/texture"
	"math/rand"
)

// Node describes a Field that has predecessors.
type Node struct {
	Name string
	Make func(int, int) texture.Field
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
func MakeNode(md, d int) texture.Field {
	n := GetNodes()
	return n[rand.Intn(len(n))].Make(md, d+1)
}

// MakeField creates either a new leaf or node.
func MakeField(md, d int) texture.Field {
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
func MakeVectorField(md, d int) texture.VectorField {
	return MakeNormal(md, d+1)
}

// Processors - these all require a source input of one form or another.

// MakeSelect creates a new field from a vector field.
func MakeSelect(md, d int) texture.Field {
	return texture.NewSelect(MakeVectorField(md, d+1), rand.Intn(3))
}

// MakeDirection creates a new field from a vector field.
func MakeDirection(md, d int) texture.Field {
	return texture.NewDirection(MakeVectorField(md, d+1))
}

// MakeMagnitude creates a new field from a vector field.
func MakeMagnitude(md, d int) texture.Field {
	return texture.NewMagnitude(MakeVectorField(md, d+1))
}

// MakeNormal creates a new vector field from a field.
func MakeNormal(md, d int) texture.VectorField {
	return texture.NewNormal(MakeField(md, d), 10, 10, 2, 2)
}
