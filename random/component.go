package random

import (
	//"fmt"
	g2dcol "github.com/jphsd/graphics2d/color"
	"github.com/jphsd/texture"
	"math/rand"
)

// MakeComponent creates a new component.
func MakeComponent() *texture.Component {
	// 2 fields feeding displacement
	disp := MakeField(2, 0)
	amt := rand.Float64()*10 + 1
	src := texture.NewDisplace(MakeField(6, 0), disp, disp, amt)

	// Emit color, alpha and bump map
	c1, c2, c3 := g2dcol.Random(), g2dcol.Random(), g2dcol.Random()
	return texture.NewComponent(src, c1, c2, c3, texture.LerpType(rand.Intn(3)), 20)
}
