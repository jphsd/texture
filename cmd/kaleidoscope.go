//go:build ignore
// +build ignore

package main

import (
	"fmt"
	gi "github.com/jphsd/graphics2d/image"
	"github.com/jphsd/texture"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))

	width, height := 600, 600

	cnt := 0
	for true {
		cf := texture.MakeColorField(5, 0)
		cf = texture.KaleidoscopeCF(cf, []float64{300, 300}, rand.Intn(5)+1, 0)
		img := texture.NewRGBA(width, height, cf, 0, 0, 1, 1)
		gi.SaveImage(img, fmt.Sprintf("%06dk", cnt))
		cnt++
	}
}
