//go:build ignore

package main

import (
	"fmt"
	g2d "github.com/jphsd/graphics2d"
	gi "github.com/jphsd/graphics2d/image"
	"github.com/jphsd/texture"
	"github.com/jphsd/texture/random"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))

	width, height := 800, 800

	cnt := 0
	for true {
		cf := random.MakeColorField(6, 0)
		img := texture.NewTextureRGBA(width, height, cf, 0, 0, 1, 1, false)
		gi.SaveImage(img, fmt.Sprintf("%06d-0t", cnt))

		cf = texture.NewReflectCF(cf, []float64{400, 100}, []float64{746, 700})
		cf = texture.NewReflectCF(cf, []float64{746, 700}, []float64{54, 700})
		cf = texture.NewReflectCF(cf, []float64{54, 700}, []float64{400, 100})

		// Repeat to catch other reflections
		cf = texture.NewReflectCF(cf, []float64{400, 100}, []float64{746, 700})
		cf = texture.NewReflectCF(cf, []float64{746, 700}, []float64{54, 700})
		img = texture.NewTextureRGBA(width, height, cf, 0, 0, 1, 1, false)
		gi.SaveImage(img, fmt.Sprintf("%06d-1t", cnt))

		// Zoom out
		xfm := g2d.NewAff3()
		xfm.ScaleAbout(2, 2, 400, 400)
		cf = texture.NewTransformCF(cf, xfm)

		img = texture.NewTextureRGBA(width, height, cf, 0, 0, 1, 1, false)
		gi.SaveImage(img, fmt.Sprintf("%06d-zt", cnt))

		cnt++
	}
}
