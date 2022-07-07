//go:build ignore

package main

import (
	"fmt"
	g2dimg "github.com/jphsd/graphics2d/image"
	"github.com/jphsd/texture"
	"image"
	"math/rand"
	"os"
	"time"

	_ "image/jpeg"
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))

	width, height := 800, 800

	// Read in Sample.jpg for texture.Sample use in random.go
	f, err := os.Open("Sample.jpg")
	if err != nil {
		panic(err)
	}
	texture.Sample, _, err = image.Decode(f)
	if err != nil {
		panic(err)
	}
	_ = f.Close()

	cnt := 0
	for cnt < 100 {
		name := fmt.Sprintf("%06d", cnt)
		cf := texture.MakeColorField(6, 0)
		img := texture.NewRGBA(width, height, cf, 0, 0, 1, 1)
		g2dimg.SaveImage(img, name)
		texture.SaveJSON(cf, name)
		cnt++
	}
}
