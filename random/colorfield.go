package random

import (
	//"fmt"
	g2dcol "github.com/jphsd/graphics2d/color"
	"github.com/jphsd/texture"
	"math/rand"
)

// ColorFieldOpts describes the available ColorField functions.
type ColorFieldOpts struct {
	Name string
	Make func(int, int) texture.ColorField
}

var ColorFieldOptions []ColorFieldOpts

// GetColorFields returns the list of ColorField functions.
func GetColorFields() []ColorFieldOpts {
	// Lazy initialization to prevent compiler complaining
	// about initialization loops
	if ColorFieldOptions == nil {
		ColorFieldOptions = []ColorFieldOpts{
			{"ColorGray", MakeColorGray},
			{"ColorSinCos", MakeColorSinCos},
			{"ColorFields", MakeColorFields},
			{"ColorConv", MakeColorConv},
			{"ColorBlend", MakeColorBlend},
			{"ColorSubstitute", MakeColorSubstitute},
		}
	}
	return ColorFieldOptions
}

// MakeColorField creates a new color field.
func MakeColorField(md, d int) texture.ColorField {
	cf := GetColorFields()
	return cf[rand.Intn(len(cf))].Make(md, d+1)
}

// MakeColorConv creates a new color field from a field.
func MakeColorConv(md, d int) texture.ColorField {
	return texture.NewColorConv(MakeField(md, d+1), g2dcol.Random(), g2dcol.Random(), nil, nil, texture.LerpType(rand.Intn(3)))
}

// MakeColorGray creates a new color field from a field.
func MakeColorGray(md, d int) texture.ColorField {
	return texture.NewColorGray(MakeField(md, d+1))
}

// MakeColorSinCos creates a new color field from a field.
func MakeColorSinCos(md, d int) texture.ColorField {
	return texture.NewColorSinCos(MakeField(md, d+1), rand.Intn(6), rand.Intn(2) == 0)
}

// MakeColorFields creates a color field from three fields.
func MakeColorFields(md, d int) texture.ColorField {
	return texture.NewColorFields(MakeField(md, d+1), MakeField(md, d+1), MakeField(md, d+1), nil, rand.Intn(2) == 0)
}

// MakeColorBlend creates a color field from two input color fields and a field.
func MakeColorBlend(md, d int) texture.ColorField {
	return texture.NewColorBlend(MakeColorField(md, d+1), MakeColorField(md, d+1), MakeField(md, d+1), texture.LerpType(rand.Intn(3)))
}

// MakeColorSubstitute creates a color field from two input color fields and a field.
func MakeColorSubstitute(md, d int) texture.ColorField {
	s := rand.Float64()
	t := rand.Float64()
	e := t - s*(1+t)
	return texture.NewColorSubstitute(MakeColorField(md, d+1), MakeColorField(md, d+1), MakeField(md, d+1), s, e)
}
