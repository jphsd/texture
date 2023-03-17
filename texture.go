package texture

import (
	"github.com/jphsd/datastruct"
	"image"
	"image/color"
)

// Texture image convenience functions.

// NewRGBA renders the texture into a new RGBA image.
func NewRGBA(width, height int, src ColorField, ox, oy, dx, dy float64) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	y := oy
	for r := 0; r < height; r++ {
		x := ox
		for c := 0; c < width; c++ {
			v := src.Eval2(x, y)
			img.Set(c, r, v)
			x += dx
		}
		y += dy
	}

	return img
}

// Texture is a lazily evaluated RGBA image. For expensive textures this allows only the pixels needed
// to be calculated.
type Texture struct {
	Src    ColorField
	Bits   datastruct.Bits // True if pixel has already been evaluated
	Rect   image.Rectangle
	Img    *image.RGBA // Evaulated pixels
	Stride int
	Ox, Oy float64
	Dx, Dy float64
}

// NewTexture creates a new Texture from the supplied parameters
func NewTexture(width, height int, src ColorField, ox, oy, dx, dy float64) *Texture {
	bits := datastruct.NewBits(width * height)
	rect := image.Rectangle{image.Point{}, image.Point{width, height}}
	img := image.NewRGBA(rect)
	return &Texture{src, bits, rect, img, width, ox, oy, dx, dy}
}

// ColorModel implements the ColorModel function in the Image interface.
func (t *Texture) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds implements the Bounds function in the Image interface.
func (t *Texture) Bounds() image.Rectangle {
	return t.Rect
}

// At implements the At function in the Image interface.
func (t *Texture) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(t.Rect)) {
		return color.RGBA{}
	}
	// Convert x, y to bit index
	i := x + y*t.Stride
	if t.Bits.Get(i) {
		return t.Img.RGBAAt(x, y)
	}
	// Pixel not set - evaluate it
	col := t.Src.Eval2(t.Ox+float64(x)*t.Dx, t.Oy+float64(y)*t.Dy)
	t.Img.Set(x, y, col)
	t.Bits.Set(i)

	return col
}
