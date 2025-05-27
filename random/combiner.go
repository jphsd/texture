package random

import (
	//"fmt"
	"github.com/jphsd/texture"
)

// Combiners - these all require source inputs of one form or another.

func MakeMulCombiner(md, d int) texture.Field {
	return texture.NewMulCombiner(MakeField(md, d+1), MakeField(md, d+1))
}

func MakeAddCombiner(md, d int) texture.Field {
	return texture.NewAddCombiner(MakeField(md, d+1), MakeField(md, d+1))
}

func MakeSubCombiner(md, d int) texture.Field {
	return texture.NewSubCombiner(MakeField(md, d+1), MakeField(md, d+1))
}

func MakeMinCombiner(md, d int) texture.Field {
	return texture.NewMinCombiner(MakeField(md, d+1), MakeField(md, d+1))
}

func MakeMaxCombiner(md, d int) texture.Field {
	return texture.NewMaxCombiner(MakeField(md, d+1), MakeField(md, d+1))
}

func MakeDiffCombiner(md, d int) texture.Field {
	return texture.NewDiffCombiner(MakeField(md, d+1), MakeField(md, d+1))
}

func MakeWindowedCombiner(md, d int) texture.Field {
	return texture.NewWindowedCombiner(MakeField(md, d+1), MakeField(md, d+1), -1/3, 1/3)
}

func MakeWeightedCombiner(md, d int) texture.Field {
	return texture.NewWeightedCombiner(MakeField(md, d+1), MakeField(md, d+1), 0.75, 0.25)
}

func MakeBlend(md, d int) texture.Field {
	return texture.NewBlend(MakeField(md, d+1), MakeField(md, d+1), MakeField(md, d+1))
}

func MakeStochasticBlend(md, d int) texture.Field {
	return texture.NewStochasticBlend(MakeField(md, d+1), MakeField(md, d+1), MakeField(md, d+1))
}

func MakeJitterBlend(md, d int) texture.Field {
	return texture.NewJitterBlend(MakeField(md, d+1), MakeField(md, d+1), MakeField(md, d+1), 0.1)
}

func MakeSubstituteCombiner(md, d int) texture.Field {
	return texture.NewSubstituteCombiner(MakeField(md, d+1), MakeField(md, d+1), MakeField(md, d+1), -1/3, 1/3)
}

func MakeThresholdCombiner(md, d int) texture.Field {
	return texture.NewThresholdCombiner(MakeField(md, d+1), MakeField(md, d+1), MakeField(md, d+1), 0)
}
