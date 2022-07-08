package random

import (
	"math/rand"
	"github.com/jphsd/texture"
)

func MakeSupport() [][]float64 {
	switch rand.Intn(3) {
	default:
		fallthrough
	case 0:
		return texture.Z4Support(1, 1)
	case 1:
		return texture.X3Support(1, 1)
	case 2:
		return texture.Z8Support(1, 1)
	}
}

func MakeMorphological(md, d int) texture.Field {
	switch rand.Intn(4) {
	default:
		fallthrough
	case 0:
		return texture.NewErode(MakeField(md, d+1), MakeSupport())
	case 1:
		return texture.NewDilate(MakeField(md, d+1), MakeSupport())
	case 2:
		return texture.NewOpen(MakeField(md, d+1), MakeSupport())
	case 3:
		return texture.NewClose(MakeField(md, d+1), MakeSupport())
	}
}
