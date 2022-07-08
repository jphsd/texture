package random

import (
	//"fmt"
	g2d "github.com/jphsd/graphics2d"
	"github.com/jphsd/texture"
	"math"
	"math/rand"
)

// MakeTransform creates a new transform processor
func MakeTransform(md, d int) texture.Field {
	xfm := g2d.NewAff3()
	offs := rand.Float64()*200 - 100
	sx := rand.Float64()*4 - 2
	sy := rand.Float64()*4 - 2
	rot := rand.Float64() * math.Pi * 2
	xfm.Scale(sx, sy)
	xfm.Rotate(rot)
	xfm.Translate(offs, 0)
	return texture.NewTransform(MakeField(md, d+1), xfm)
}

// MakeStrip creates a new strip processor
func MakeStrip(md, d int) texture.Field {
	xfm := g2d.NewAff3()
	xfm.Rotate(rand.Float64() * math.Pi * 2)
	y := rand.Float64()*2000 - 1000
	return texture.NewTransform(texture.NewStrip(MakeField(md, d+1), y), xfm)
}

// MakeFractal creates a new fractal processor.
func MakeFractal(md, d int) texture.Field {
	lac := 2.0
	hurst := 1.0
	nocts := rand.Intn(3)
	octs := float64(nocts)
	xfm := g2d.NewAff3()
	xfm.Scale(lac, lac)
	xfm.Rotate(rand.Float64() * math.Pi)
	if rand.Intn(2) == 0 {
		fbm := texture.NewFBM(hurst, lac, nocts)
		res := texture.NewFractal(MakeField(md, d+1), xfm, fbm, octs)
		return res
	}
	mf := texture.NewMF(hurst, lac, 0.5, nocts)
	res := texture.NewFractal(MakeField(md, d+1), xfm, mf, octs)
	return res
}

// MakeVariableFractal creates a new fractal processor.
func MakeVariableFractal(md, d int) texture.Field {
	lac := 2.0
	hurst := 1.0
	nocts := rand.Intn(5)
	octs := float64(nocts)
	xfm := g2d.NewAff3()
	xfm.Scale(lac, lac)
	xfm.Rotate(rand.Float64() * math.Pi)
	if rand.Intn(2) == 0 {
		fbm := texture.NewFBM(hurst, lac, nocts)
		res := texture.NewFractal(MakeField(md, d+1), xfm, fbm, octs)
		return res
	}
	mf := texture.NewMF(hurst, lac, 0.5, nocts)
	res := texture.NewVariableFractal(MakeField(md, d+1), xfm, mf, MakeField(md, d+1), octs)
	return res
}

// MakeDistort creates a new distorted processor.
func MakeDistort(md, d int) texture.Field {
	return texture.NewDistort(MakeField(md, d+1), 1)
}

// MakeReflect creates a mirror plane in the field.
func MakeReflect(md, d int) texture.Field {
	a := rand.Float64() * math.Pi * 2
	dx, dy := math.Cos(a)*400, math.Sin(a)*400 // Hack alert - assumes 800x800
	return texture.NewReflect(MakeField(md, d+1), []float64{400, 400}, []float64{400 + dx, 400 + dy})
}

// MakeDisplace creates a displacement of src1 with src2 and src3.
func MakeDisplace(md, d int) texture.Field {
	return texture.NewDisplace(MakeField(md, d+1), MakeField(md, d+1), MakeField(md, d+1), 10)
}
