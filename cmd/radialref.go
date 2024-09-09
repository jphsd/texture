//go:build ignore

package main

import (
	"flag"
	"fmt"
	gi "github.com/jphsd/graphics2d/image"
	"github.com/jphsd/texture"
	"github.com/jphsd/texture/random"
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
		cf := random.MakeColorField(6, 0)
		img := texture.NewTextureRGBA(width, height, cf, 0, 0, 1, 1, false)
		gi.SaveImage(img, fmt.Sprintf("%06d-0rr", cnt))

		for i := 0; i < *n; i++ {
			cf = texture.NewReflectCF(cf, []float64{400, 400}, []float64{400 + math.Cos(a)*r, 400 + math.Sin(a)*r})
			a += da
		}
		img = texture.NewTextureRGBA(width, height, cf, 0, 0, 1, 1, false)
		gi.SaveImage(img, fmt.Sprintf("%06d-1rr", cnt))

		cnt++
	}
}
