//go:build ignore

package main

import (
	"flag"
	"fmt"
	g2d "github.com/jphsd/graphics2d"
	gi "github.com/jphsd/graphics2d/image"
	"github.com/jphsd/texture"
	"math"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	n := flag.Int("n", 6, "width")
	flag.Parse()

	width, height := 800, 800

	r := 400.0
	da := math.Pi / float64(*n)
	cnt := 0
	for true {
		a := 0.0
		cf := texture.MakeColorField(6, 0)
		img := texture.NewRGBA(width, height, cf, 0, 0, 1, 1)
		gi.SaveImage(img, fmt.Sprintf("%06d-0k", cnt))

		pts := make([][]float64, *n*2)
		for i := 0; i < *n; i++ {
			pts[i] = []float64{400 + math.Cos(a)*r, 400 + math.Sin(a)*r}
			cf = texture.NewReflectCF(cf, []float64{400, 400}, pts[i])
			a += da
		}

		// Finish additional pts
		for i := 0; i < *n; i++ {
			pts[*n+i] = []float64{400 + math.Cos(a)*r, 400 + math.Sin(a)*r}
			a += da
		}

		// Place rim
		prev := pts[0]
		for i := 1; i < *n*2; i++ {
			cur := pts[i]
			cf = texture.NewReflectCF(cf, prev, cur)
			prev = cur
		}
		cf = texture.NewReflectCF(cf, prev, pts[0])

		img = texture.NewRGBA(width, height, cf, 0, 0, 1, 1)
		gi.SaveImage(img, fmt.Sprintf("%06d-1k", cnt))

		// Zoom out

		xfm := g2d.NewAff3()
		xfm.ScaleAbout(2, 2, 400, 400)
		cf = texture.NewTransformCF(cf, xfm)

		img = texture.NewRGBA(width, height, cf, 0, 0, 1, 1)
		gi.SaveImage(img, fmt.Sprintf("%06d-2k", cnt))

		cnt++
	}
}
