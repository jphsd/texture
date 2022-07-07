//go:build ignore

package main

import (
	"fmt"
	g2d "github.com/jphsd/graphics2d"
	gi "github.com/jphsd/graphics2d/image"
	"github.com/jphsd/texture"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))

	width, height := 600, 600

	//xfm1 := g2d.LineTransform(0, 0, 600, 0, 100, 200, 200, 100)
	//xfm2 := g2d.LineTransform(0, 0, 600, 0, 400, 100, 500, 200)
	//xfm3 := g2d.LineTransform(0, 0, 600, 0, 200, 200, 400, 200)
	xfm1 := g2d.BoxTransform(0, 0, 600, 0, 600, 100, 200, 200, 100, 400)
	xfm2 := g2d.BoxTransform(0, 0, 600, 0, 600, 400, 100, 500, 200, 400)
	xfm3 := g2d.BoxTransform(0, 0, 600, 0, 600, 200, 200, 400, 200, 400)
	for i := 1; i < 7; i++ {
		src := texture.NewIFS([]float64{float64(width), float64(height)}, []*g2d.Aff3{xfm1, xfm2, xfm3}, i)
		cf := texture.NewColorGray(src)
		img := texture.NewRGBA(width, height, cf, 0, 0, 1, 1)
		gi.SaveImage(img, fmt.Sprintf("ifs-%d", i))
	}
}
