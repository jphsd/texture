package random

import (
	//"fmt"
	"github.com/jphsd/texture"
)

// MakeNLFilter creates a non-linear filter.
func MakeNLFilter(md, d int) texture.Field {
	return texture.NewNLFilter(MakeField(md, d+1), MakeNL(), 1, 0)
}

// MakeQuantizeFilter creates a non-linear filter.
func MakeQuantizeFilter(md, d int) texture.Field {
	return texture.NewQuantizeFilter(MakeField(md, d+1), 1, 0, int(PickLambda()))
}

// MakeRandQuantFilter creates a non-linear filter.
func MakeRandQuantFilter(md, d int) texture.Field {
	return texture.NewRandQuantFilter(MakeField(md, d+1), 1, 0, int(PickLambda()))
}
