package texture_test

import (
	"fmt"
	"github.com/jphsd/graphics2d"
	"github.com/jphsd/graphics2d/image"
	"github.com/jphsd/texture"
	"math"
)

func ExampleLinearGradient_saw() {
	nl := texture.NewNLLinear()
	w := texture.NewNLWave([]float64{125}, []*texture.NonLinear{nl}, false, false)
	f := texture.NewLinearGradient(w)

	img := texture.NewTextureGray16(600, 600, f, 0, 0, 1, 1, false)
	image.SaveImage(img, "ExampleLinearGradient_saw")
	fmt.Printf("Generated ExampleLinearGradient_saw")
	// Output: Generated ExampleLinearGradient_saw
}

func ExampleLinearGradient_triangle() {
	nl := texture.NewNLLinear()
	w := texture.NewNLWave([]float64{125}, []*texture.NonLinear{nl}, true, false)
	f := texture.NewLinearGradient(w)

	img := texture.NewTextureGray16(600, 600, f, 0, 0, 1, 1, false)
	image.SaveImage(img, "ExampleLinearGradient_triangle")
	fmt.Printf("Generated ExampleLinearGradient_triangle")
	// Output: Generated ExampleLinearGradient_triangle
}

func ExampleLinearGradient_sine() {
	nl := texture.NewNLSin()
	w := texture.NewNLWave([]float64{125}, []*texture.NonLinear{nl}, true, false)
	f := texture.NewLinearGradient(w)

	img := texture.NewTextureGray16(600, 600, f, 0, 0, 1, 1, false)
	image.SaveImage(img, "ExampleLinearGradient_sine")
	fmt.Printf("Generated ExampleLinearGradient_sine")
	// Output: Generated ExampleLinearGradient_sine
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

func ExampleLinearGradient_circular() {
	nl1 := texture.NewNLCircle1()
	nl2 := texture.NewNLCircle2()
	w := texture.NewDCWave([]float64{125}, []*texture.NonLinear{nl1, nl2}, false)
	f := texture.NewLinearGradient(w)

	img := texture.NewTextureGray16(600, 600, f, 0, 0, 1, 1, false)
	image.SaveImage(img, "ExampleLinearGradient_circular")
	fmt.Printf("Generated ExampleLinearGradient_circular")
	// Output: Generated ExampleLinearGradient_circular
}

func ExampleRadialGradient_saw() {
	nl := texture.NewNLLinear()
	w := texture.NewNLWave([]float64{125}, []*texture.NonLinear{nl}, false, false)
	f := texture.NewRadialGradient(w)

	img := texture.NewTextureGray16(600, 600, f, 0, 0, 1, 1, false)
	image.SaveImage(img, "ExampleRadialGradient_saw")
	fmt.Printf("Generated ExampleRadialGradient_saw")
	// Output: Generated ExampleRadialGradient_saw
}

func ExampleRadialGradient_triangle() {
	nl := texture.NewNLLinear()
	w := texture.NewNLWave([]float64{125}, []*texture.NonLinear{nl}, true, false)
	f := texture.NewRadialGradient(w)

	img := texture.NewTextureGray16(600, 600, f, 0, 0, 1, 1, false)
	image.SaveImage(img, "ExampleRadialGradient_triangle")
	fmt.Printf("Generated ExampleRadialGradient_triangle")
	// Output: Generated ExampleRadialGradient_triangle
}

func ExampleRadialGradient_sine() {
	nl := texture.NewNLSin()
	w := texture.NewNLWave([]float64{125}, []*texture.NonLinear{nl}, true, false)
	f := texture.NewRadialGradient(w)

	img := texture.NewTextureGray16(600, 600, f, 0, 0, 1, 1, false)
	image.SaveImage(img, "ExampleRadialGradient_sine")
	fmt.Printf("Generated ExampleRadialGradient_sine")
	// Output: Generated ExampleRadialGradient_sine
}

func ExampleRadialGradient_circular() {
	nl1 := texture.NewNLCircle1()
	nl2 := texture.NewNLCircle2()
	w := texture.NewDCWave([]float64{125}, []*texture.NonLinear{nl1, nl2}, false)
	f := texture.NewRadialGradient(w)

	img := texture.NewTextureGray16(600, 600, f, 0, 0, 1, 1, false)
	image.SaveImage(img, "ExampleRadialGradient_circular")
	fmt.Printf("Generated ExampleRadialGradient_circular")
	// Output: Generated ExampleRadialGradient_circular
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

func ExampleWarp_radial() {
	f := texture.NewSquares(40)

	rwf := texture.NewRadialWF([]float64{300, 300}, 1, 1)

	f2 := texture.NewWarp(f, rwf)

	img := texture.NewTextureGray16(600, 600, f2, 0, 0, 1, 1, false)
	image.SaveImage(img, "ExampleWarp_radial")
	fmt.Printf("Generated ExampleWarp_radial")
	// Output: Generated ExampleWarp_radial
}

func ExampleWarp_radialnl1() {
	f := texture.NewSquares(40)

	rwf := texture.NewRadialNLWF([]float64{300, 300}, texture.NewNLExponential(3), 300)

	f2 := texture.NewWarp(f, rwf)

	img := texture.NewTextureGray16(600, 600, f2, 0, 0, 1, 1, false)
	image.SaveImage(img, "ExampleWarp_radialnl1")
	fmt.Printf("Generated ExampleWarp_radialnl1")
	// Output: Generated ExampleWarp_radialnl1
}

func ExampleWarp_radialnl2() {
	f := texture.NewSquares(40)

	rwf := texture.NewRadialNLWF([]float64{300, 300}, texture.NewNLLogarithmic(3), 300)

	f2 := texture.NewWarp(f, rwf)

	img := texture.NewTextureGray16(600, 600, f2, 0, 0, 1, 1, false)
	image.SaveImage(img, "ExampleWarp_radialnl2")
	fmt.Printf("Generated ExampleWarp_radialnl2")
	// Output: Generated ExampleWarp_radialnl2
}

func ExampleWarp_pinch() {
	f := texture.NewSquares(40)

	rwf := texture.NewPinchXWF([]float64{300, 300}, 0.3, 0.002, 2)

	f2 := texture.NewWarp(f, rwf)

	img := texture.NewTextureGray16(600, 600, f2, 0, 0, 1, 1, false)
	image.SaveImage(img, "ExampleWarp_pinch")
	fmt.Printf("Generated ExampleWarp_pinch")
	// Output: Generated ExampleWarp_pinch
}

func ExampleWarp_swirl() {
	f := texture.NewSquares(40)

	rwf := texture.NewSwirlWF([]float64{300, 300}, -0.05)

	f2 := texture.NewWarp(f, rwf)

	img := texture.NewTextureGray16(600, 600, f2, 0, 0, 1, 1, false)
	image.SaveImage(img, "ExampleWarp_swirl")
	fmt.Printf("Generated ExampleWarp_swirl")
	// Output: Generated ExampleWarp_swirl
}

func ExampleWarp_drain() {
	f := texture.NewSquares(40)

	rwf := texture.NewDrainWF([]float64{300, 300}, math.Pi, 250)

	f2 := texture.NewWarp(f, rwf)

	img := texture.NewTextureGray16(600, 600, f2, 0, 0, 1, 1, false)
	image.SaveImage(img, "ExampleWarp_drain")
	fmt.Printf("Generated ExampleWarp_drain")
	// Output: Generated ExampleWarp_drain
}

func ExampleWarp_ripple() {
	f := texture.NewSquares(40)

	rwf := texture.NewRippleXWF(100, 20, 12.5)

	f2 := texture.NewWarp(f, rwf)

	img := texture.NewTextureGray16(600, 600, f2, 0, 0, 1, 1, false)
	image.SaveImage(img, "ExampleWarp_ripple")
	fmt.Printf("Generated ExampleWarp_ripple")
	// Output: Generated ExampleWarp_ripple
}

func ExampleWarp_radialripple() {
	f := texture.NewSquares(40)

	rwf := texture.NewRadialRippleWF([]float64{300, 300}, 100, 10, 0)

	f2 := texture.NewWarp(f, rwf)

	img := texture.NewTextureGray16(600, 600, f2, 0, 0, 1, 1, false)
	image.SaveImage(img, "ExampleWarp_radialripple")
	fmt.Printf("Generated ExampleWarp_radialripple")
	// Output: Generated ExampleWarp_radialripple
}

func ExampleWarp_radialwriggle() {
	f := texture.NewSquares(40)

	rwf := texture.NewRadialWiggleWF([]float64{300, 300}, 100, 0.1, 0)

	f2 := texture.NewWarp(f, rwf)

	img := texture.NewTextureGray16(600, 600, f2, 0, 0, 1, 1, false)
	image.SaveImage(img, "ExampleWarp_radialwriggle")
	fmt.Printf("Generated ExampleWarp_radialwriggle")
	// Output: Generated ExampleWarp_radialwriggle
}

func ExampleTiler() {
	nl := texture.NewNLCircle1()
	w := texture.NewNLWave([]float64{47}, []*texture.NonLinear{nl}, false, true)
	iw := texture.NewInvertWave(w)
	f := texture.NewRadialGradient(iw)

	xfm := graphics2d.Translate(-50, -50)
	f2 := texture.NewTransform(f, xfm)

	f3 := texture.NewTiler(f2, []float64{100, 100})

	img := texture.NewTextureGray16(600, 600, f3, 0, 0, 1, 1, false)
	image.SaveImage(img, "ExampleTiler")
	fmt.Printf("Generated ExampleTiler")
	// Output: Generated ExampleTiler
}

func ExampleReflect() {
	f := texture.NewPerlin(12345)

	f2 := texture.NewReflect(f, []float64{0, 9}, []float64{9, 0})

	img := texture.NewTextureGray16(600, 600, f2, 0, 0, .015, .015, false)
	image.SaveImage(img, "ExampleReflect")
	fmt.Printf("Generated ExampleReflect")
	// Output: Generated ExampleReflect
}

func ExampleDistort() {
	f := texture.NewPerlin(12345)

	f2 := texture.NewDistort(f, 10)

	img := texture.NewTextureGray16(600, 600, f2, 0, 0, .015, .015, false)
	image.SaveImage(img, "ExampleDistort")
	fmt.Printf("Generated ExampleDistort")
	// Output: Generated ExampleDistort
}

func ExampleFractal_fbm() {
	f := texture.NewPerlin(12345)

	lac := 2.0
	hurst := 1.0
	oct := 3.0
	xfm := graphics2d.Scale(lac, lac)
	fbm := texture.NewFBM(hurst, lac, int(oct+1))
	f2 := texture.NewFractal(f, xfm, fbm, 3)

	img := texture.NewTextureGray16(600, 600, f2, 0, 0, .015, .015, false)
	image.SaveImage(img, "ExampleFractal_fbm")
	fmt.Printf("Generated ExampleFractal_fbm")
	// Output: Generated ExampleFractal_fbm
}

func ExampleFractal_mf() {
	f := texture.NewPerlin(12345)

	lac := 2.0
	hurst := 1.0
	oct := 3.0
	xfm := graphics2d.Scale(lac, lac)
	fbm := texture.NewMF(hurst, lac, 0.5, int(oct+1))
	f2 := texture.NewFractal(f, xfm, fbm, 3)

	img := texture.NewTextureGray16(600, 600, f2, 0, 0, .015, .015, false)
	image.SaveImage(img, "ExampleFractal_mf")
	fmt.Printf("Generated ExampleFractal_mf")
	// Output: Generated ExampleFractal_mf
}
