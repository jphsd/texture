package random

import (
	//"fmt"
	g2d "github.com/jphsd/graphics2d"
	"github.com/jphsd/texture"
	"image"
	"math"
	"math/rand"
)

// Leaf describes a Field that has no predecessors.
type Leaf struct {
	Name string
	Make func() texture.Field
}

var LeafOptions = []Leaf{
	//{"Uniform", MakeUniform},
	{"LinearGradient", MakeLinearGradient},
	{"RadialGradient", MakeRadialGradient},
	{"ConicGradient", MakeConicGradient},
	{"TiledGradient", MakeTiledGradient},
	{"Binary", MakeBinary},
	{"Perlin", MakePerlin},
	{"DistortedPerlin", MakeDistortedPerlin},
	{"Image", MakeImage},
	//{"Shape", MakeShape},
}

// MakeLeaf creates a new leaf.
func MakeLeaf() texture.Field {
	res := LeafOptions[rand.Intn(len(LeafOptions))]
	return res.Make()
}

// The following are Leaves (don't call any fields)

// MakeUniform creates a new Uniform
func MakeUniform() texture.Field {
	return texture.NewUniform(rand.Float64()*2 - 1)
}

func MakeLinearGradient() texture.Field {
	f := texture.NewLinearGradient(MakeWave())

	// Wrap it in a Transform
	xfm := g2d.NewAff3()
	offs := rand.Float64()*200 - 100
	rot := rand.Float64() * math.Pi * 2
	xfm.Rotate(rot)
	xfm.Translate(offs, 0)
	return texture.NewTransform(f, xfm)
}

func MakeRadialGradient() texture.Field {
	w := MakeWave()
	f := texture.NewRadialGradient(w)

	// Wrap it in a Transform
	xfm := g2d.NewAff3()
	offs := -400.0 // Hack alert - assumes 800x800 image
	rot := rand.Float64() * math.Pi * 2
	xfm.Rotate(rot)
	xfm.Translate(offs, offs)
	return texture.NewTransform(f, xfm)
}

func MakeConicGradient() texture.Field {
	w := MakeWave()
	f := texture.NewConicGradient(w)

	// Wrap it in a Transform
	xfm := g2d.NewAff3()
	offs := -400.0 // Hack alert - assumes 800x800 image
	rot := rand.Float64() * math.Pi * 2
	xfm.Rotate(rot)
	xfm.Translate(offs, offs)
	return texture.NewTransform(f, xfm)
}

func MakeTiledGradient() texture.Field {
	var wave texture.Wave

	//Select wave type
	if rand.Intn(2) > 0 {
		w := MakePatternWave().(*texture.PatternWave)
		w.Once = true
		wave = w
	} else {
		w := MakeNLWave().(*texture.NLWave)
		w.Once = true
		wave = w
	}

	var field texture.Field
	// Select gradient type
	if rand.Intn(2) > 0 {
		field = texture.NewRadialGradient(wave)
	} else {
		field = texture.NewConicGradient(wave)
	}

	// Move 0,0 to center
	l := wave.Lambda()
	xfm := g2d.NewAff3()
	xfm.Translate(-l, -l)
	field = texture.NewTransform(field, xfm)

	// Tile the single wave instance
	l *= 2
	field = texture.NewTiler(field, []float64{l, l})

	// Wrap it in a Transform
	xfm = g2d.NewAff3()
	offs := rand.Float64()*200 - 100
	rot := rand.Float64() * math.Pi * 2
	xfm.Rotate(rot)
	xfm.Translate(offs, 0)
	return texture.NewTransform(field, xfm)
}

// MakeBinary creates a new field backed by bit noise.
func MakeBinary() texture.Field {
	xfm := g2d.NewAff3()
	xfm.Scale(0.01, 0.01)
	xfm.Rotate(rand.Float64() * math.Pi * 2)
	f := texture.NewBinary(16, 16, rand.Int63())
	return texture.NewTransform(f, xfm)
}

// MakePerlin creates a new field backed by a perlin noise function.
func MakePerlin() texture.Field {
	xfm := g2d.NewAff3()
	xfm.Scale(0.01, 0.01)
	xfm.Rotate(rand.Float64() * math.Pi * 2)
	f := texture.NewPerlin(rand.Int63())
	return texture.NewTransform(f, xfm)
}

// MakeDistortedPerlin creates a new field backed by a perlin noise function.
func MakeDistortedPerlin() texture.Field {
	return texture.NewDistort(MakePerlin(), 1)
}

var Sample image.Image

func MakeImage() texture.Field {
	if Sample == nil {
		return MakeUniform()
	}

	iw, ih := Sample.Bounds().Dx(), Sample.Bounds().Dy()

	// Make new field from img
	f1 := texture.NewImage(Sample)
	f2 := texture.NewColorToGray(f1)
	f3 := texture.NewTiler(f2, []float64{float64(iw), float64(ih)})
	xfm := g2d.NewAff3()
	xfm.Rotate(rand.Float64() * math.Pi * 2)
	return texture.NewTransform(f3, xfm)
}
