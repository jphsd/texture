package texture_test

import (
	"fmt"
	"github.com/jphsd/graphics2d"
	"github.com/jphsd/graphics2d/image"
	"github.com/jphsd/texture"
	"math"
)

func ExampleLinearGradient_Saw() {
	nl := texture.NewNLLinear()
	w := texture.NewNLWave([]float64{125}, []*texture.NonLinear{nl}, false, false)
	f := texture.NewLinearGradient(w)

	img := texture.NewTextureGray16(600, 600, f, 0, 0, 1, 1, false)
	image.SaveImage(img, "ExampleLinearGradient_Saw")
	fmt.Printf("Generated ExampleLinearGradient_Saw")
	// Output: Generated ExampleLinearGradient_Saw
}

func ExampleLinearGradient_Triangle() {
	nl := texture.NewNLLinear()
	w := texture.NewNLWave([]float64{125}, []*texture.NonLinear{nl}, true, false)
	f := texture.NewLinearGradient(w)

	img := texture.NewTextureGray16(600, 600, f, 0, 0, 1, 1, false)
	image.SaveImage(img, "ExampleLinearGradient_Triangle")
	fmt.Printf("Generated ExampleLinearGradient_Triangle")
	// Output: Generated ExampleLinearGradient_Triangle
}

func ExampleLinearGradient_Sine() {
	nl := texture.NewNLSin()
	w := texture.NewNLWave([]float64{125}, []*texture.NonLinear{nl}, true, false)
	f := texture.NewLinearGradient(w)

	img := texture.NewTextureGray16(600, 600, f, 0, 0, 1, 1, false)
	image.SaveImage(img, "ExampleLinearGradient_Sine")
	fmt.Printf("Generated ExampleLinearGradient_Sine")
	// Output: Generated ExampleLinearGradient_Sine
}

func ExampleMulCombiner() {
	nl := texture.NewNLSin()
	w := texture.NewNLWave([]float64{125}, []*texture.NonLinear{nl}, true, false)
	f := texture.NewLinearGradient(w)

	xfm := graphics2d.Rotate(graphics2d.HalfPi)
	f2 := texture.NewTransform(f, xfm)

	f3 := texture.NewMulCombiner(f, f2)

	img := texture.NewTextureGray16(600, 600, f3, 0, 0, 1, 1, false)
	image.SaveImage(img, "ExampleMulCombiner")
	fmt.Printf("Generated ExampleMulCombiner")
	// Output: Generated ExampleMulCombiner
}

func ExampleMaxCombiner() {
	nl := texture.NewNLSin()
	w := texture.NewNLWave([]float64{125}, []*texture.NonLinear{nl}, true, false)
	f := texture.NewLinearGradient(w)

	xfm := graphics2d.Rotate(graphics2d.HalfPi)
	f2 := texture.NewTransform(f, xfm)

	f3 := texture.NewMaxCombiner(f, f2)

	img := texture.NewTextureGray16(600, 600, f3, 0, 0, 1, 1, false)
	image.SaveImage(img, "ExampleMaxCombiner")
	fmt.Printf("Generated ExampleMaxCombiner")
	// Output: Generated ExampleMaxCombiner
}

func ExampleLinearGradient_Circular() {
	nl1 := texture.NewNLCircle1()
	nl2 := texture.NewNLCircle2()
	w := texture.NewDCWave([]float64{125}, []*texture.NonLinear{nl1, nl2}, false)
	f := texture.NewLinearGradient(w)

	img := texture.NewTextureGray16(600, 600, f, 0, 0, 1, 1, false)
	image.SaveImage(img, "ExampleLinearGradient_Circular")
	fmt.Printf("Generated ExampleLinearGradient_Circular")
	// Output: Generated ExampleLinearGradient_Circular
}

func ExampleRadialGradient_Saw() {
	nl := texture.NewNLLinear()
	w := texture.NewNLWave([]float64{125}, []*texture.NonLinear{nl}, false, false)
	f := texture.NewRadialGradient(w)

	img := texture.NewTextureGray16(600, 600, f, 0, 0, 1, 1, false)
	image.SaveImage(img, "ExampleRadialGradient_Saw")
	fmt.Printf("Generated ExampleRadialGradient_Saw")
	// Output: Generated ExampleRadialGradient_Saw
}

func ExampleRadialGradient_Triangle() {
	nl := texture.NewNLLinear()
	w := texture.NewNLWave([]float64{125}, []*texture.NonLinear{nl}, true, false)
	f := texture.NewRadialGradient(w)

	img := texture.NewTextureGray16(600, 600, f, 0, 0, 1, 1, false)
	image.SaveImage(img, "ExampleRadialGradient_Triangle")
	fmt.Printf("Generated ExampleRadialGradient_Triangle")
	// Output: Generated ExampleRadialGradient_Triangle
}

func ExampleRadialGradient_Sine() {
	nl := texture.NewNLSin()
	w := texture.NewNLWave([]float64{125}, []*texture.NonLinear{nl}, true, false)
	f := texture.NewRadialGradient(w)

	img := texture.NewTextureGray16(600, 600, f, 0, 0, 1, 1, false)
	image.SaveImage(img, "ExampleRadialGradient_Sine")
	fmt.Printf("Generated ExampleRadialGradient_Sine")
	// Output: Generated ExampleRadialGradient_Sine
}

func ExampleRadialGradient_Circular() {
	nl1 := texture.NewNLCircle1()
	nl2 := texture.NewNLCircle2()
	w := texture.NewDCWave([]float64{125}, []*texture.NonLinear{nl1, nl2}, false)
	f := texture.NewRadialGradient(w)

	img := texture.NewTextureGray16(600, 600, f, 0, 0, 1, 1, false)
	image.SaveImage(img, "ExampleRadialGradient_Circular")
	fmt.Printf("Generated ExampleRadialGradient_Circular")
	// Output: Generated ExampleRadialGradient_Circular
}

func ExampleConicGradient() {
	nl := texture.NewNLSin()
	w := texture.NewNLWave([]float64{125, 125, 125}, []*texture.NonLinear{nl, nl, nl}, true, false)
	f := texture.NewConicGradient(w)

	img := texture.NewTextureGray16(600, 600, f, -300, -300, 1, 1, false)
	image.SaveImage(img, "ExampleConicGradient")
	fmt.Printf("Generated ExampleConicGradient")
	// Output: Generated ExampleConicGradient
}

func ExamplePerlin() {
	f := texture.NewPerlin(12345)

	img := texture.NewTextureGray16(600, 600, f, 0, 0, .015, .015, false)
	image.SaveImage(img, "ExamplePerlin")
	fmt.Printf("Generated ExamplePerlin")
	// Output: Generated ExamplePerlin
}

func ExampleDisplace() {
	f := texture.NewPerlin(12345)

	fx := texture.NewPerlin(12346)

	fy := texture.NewPerlin(12347)

	f2 := texture.NewDisplace(f, fx, fy, 1)

	img := texture.NewTextureGray16(600, 600, f2, 0, 0, .015, .015, false)
	image.SaveImage(img, "ExampleDisplace")
	fmt.Printf("Generated ExampleDisplace")
	// Output: Generated ExampleDisplace
}

func ExampleTriangles() {
	f := texture.NewTriangles(40)

	img := texture.NewTextureGray16(600, 600, f, 0, 0, 1, 1, false)
	image.SaveImage(img, "ExampleTriangles")
	fmt.Printf("Generated ExampleTriangles")
	// Output: Generated ExampleTriangles
}

func ExampleTransform() {
	f := texture.NewTriangles(40)

	// Apply a transform to convert to equilateral triangles
	xfm := graphics2d.Rotate(-math.Pi / 4)
	xfm.Scale(math.Sqrt2, 1)
	f2 := texture.NewTransform(f, xfm)

	img := texture.NewTextureGray16(600, 600, f2, 0, 0, 1, 1, false)
	image.SaveImage(img, "ExampleTransform")
	fmt.Printf("Generated ExampleTransform")
	// Output: Generated ExampleTransform
}

func ExampleSquares() {
	f := texture.NewSquares(40)

	img := texture.NewTextureGray16(600, 600, f, 0, 0, 1, 1, false)
	image.SaveImage(img, "ExampleSquares")
	fmt.Printf("Generated ExampleSquares")
	// Output: Generated ExampleSquares
}

func ExampleHexagons() {
	f := texture.NewHexagons(20)

	img := texture.NewTextureGray16(600, 600, f, 0, 0, 1, 1, false)
	image.SaveImage(img, "ExampleHexagons")
	fmt.Printf("Generated ExampleHexagons")
	// Output: Generated ExampleHexagons
}

func ExampleBlend() {
	f := texture.NewHexagons(20)

	f2 := texture.NewSquares(40)

	nl := texture.NewNLLinear()
	w := texture.NewNLWave([]float64{600}, []*texture.NonLinear{nl}, false, false)
	f3 := texture.NewLinearGradient(w)

	f4 := texture.NewBlend(f, f2, f3)

	img := texture.NewTextureGray16(600, 600, f4, 0, 0, 1, 1, false)
	image.SaveImage(img, "ExampleBlend")
	fmt.Printf("Generated ExampleBlend")
	// Output: Generated ExampleBlend
}

func ExampleShape() {
	shape := graphics2d.NewShape(graphics2d.ReentrantPolygon([]float64{300, 300}, 300, 5, .65, 0))
	f := texture.NewShape(shape, texture.BinaryStyle)

	img := texture.NewTextureGray16(600, 600, f, 0, 0, 1, 1, false)
	image.SaveImage(img, "ExampleShape")
	fmt.Printf("Generated ExampleShape")
	// Output: Generated ExampleShape
}

func ExampleShapeCombiner() {
	nl := texture.NewNLSin()
	w := texture.NewNLWave([]float64{32}, []*texture.NonLinear{nl}, true, false)
	f := texture.NewRadialGradient(w)

	xfm := graphics2d.Translate(-300, -300)
	f2 := texture.NewTransform(f, xfm)

	f3 := texture.NewSquares(40)

	shape := graphics2d.NewShape(graphics2d.ReentrantPolygon([]float64{300, 300}, 300, 5, .65, 0))
	f4 := texture.NewShapeCombiner(f2, f3, shape)

	img := texture.NewTextureGray16(600, 600, f4, 0, 0, 1, 1, false)
	image.SaveImage(img, "ExampleShapeCombiner")
	fmt.Printf("Generated ExampleShapeCombiner")
	// Output: Generated ExampleShapeCombiner
}
