/*
Package random contains functions, MakeXXX, that can be used to create a random texture tree.

The following code creates a random tree up to six nodes deep and uses it to render to an 800x800 image,
which is then saved as "example.png"

  package main

  import (
  	"github.com/jphsd/graphics2d/image"
  	"github.com/jphsd/texture"
  	"github.com/jphsd/texture/random"
  )

  func main() {
  	cf := random.MakeColorField(6, 0)
  	img := texture.NewRGBA(800, 800, cf, 0, 0, 1, 1)
  	image.SaveImage(img, "example")
  }
*/
package random
