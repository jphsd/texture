//go:build ignore

package main

import (
	"fmt"
	gi "github.com/jphsd/graphics2d/image"
	"github.com/jphsd/texture"
	"github.com/jphsd/texture/random"
	"github.com/jphsd/texture/surface"
	"image"
	"image/draw"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))

	width, height := 400, 400

	cnt := 0
	for true {
		res := image.NewRGBA(image.Rect(0, 0, width*3, height))
		cf := random.MakeComponent()
		// Color
		img := texture.NewTextureRGBA(width, height, cf.Color, 0, 0, 1, 1, false)
		draw.Draw(res, image.Rect(0, 0, width, height), img, image.Point{}, draw.Src)
		// Alpha
		gimg := texture.NewTextureGray16(width, height, cf.Value, 0, 0, 1, 1, false)
		draw.Draw(res, image.Rect(width, 0, 2*width, height), gimg, image.Point{}, draw.Src)
		// Bump Map
		bm := &surface.BumpMap{surface.DefaultAmbient, nil, nil, cf.Vector}
		img = texture.NewTextureRGBA(width, height, bm, 0, 0, 1, 1, false)
		draw.Draw(res, image.Rect(2*width, 0, 3*width, height), img, image.Point{}, draw.Src)
		gi.SaveImage(res, fmt.Sprintf("%06d", cnt))
		cnt++
	}
}
