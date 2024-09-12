package texture_test

import (
	"fmt"
	"github.com/jphsd/graphics2d/image"
	"github.com/jphsd/texture"
)

func Example_decay() {
	// Original waveform
	nl := texture.NewNLSin()
	w := texture.NewNLWave([]float64{62}, []*texture.NonLinear{nl}, true, false)
	w1 := texture.NewInvertWave(w)
	f := texture.NewLinearGradient(w1)

	// Decay factor
	nl2 := texture.NewNLLinear()
	w2 := texture.NewNLWave([]float64{600}, []*texture.NonLinear{nl2}, false, false)
	w3 := texture.NewInvertWave(w2)
	f2 := texture.NewLinearGradient(w3)

	// Remap [-1,1] => [0,1]
	f4 := texture.NewOffsScaleFilter(f2, 0.5, 1)

	// Combine waves to get decay waveform
	f3 := texture.NewMulCombiner(f, f4)

	img := texture.NewTextureGray16(600, 600, f3, 0, 0, 1, 1, false)
	image.SaveImage(img, "Example_decay")
	fmt.Printf("Generated Example_decay")
	// Output: Generated Example_decay
}
