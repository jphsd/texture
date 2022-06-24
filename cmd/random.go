//go:build ignore
// +build ignore

package main

import (
	"encoding/json"
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
		cf := texture.MakeColorField(6, 0)
		img := texture.NewRGBA(width, height, cf, 0, 0, 1, 1)
		gi.SaveImage(img, fmt.Sprintf("%06d", cnt))

		// Emit JSON description of cf
		b, err := json.Marshal(cf)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		fmt.Printf("%v\n", b)
		cnt++
	}
}
