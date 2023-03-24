package texture

import (
	"github.com/jphsd/datastruct"
	"image"
	"image/color"
)

// Texture image convenience functions.

// NewRGBA renders the src ColorField into a new RGBA image.
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

// TextureRGBA is a lazily evaluated RGBA image. For expensive textures this allows only the requested pixels
// to be calculated, and not the entire image.
type TextureRGBA struct {
	Src    ColorField
	Rect   image.Rectangle
	Img    *image.RGBA // Evaluated pixels
	Stride int
	Ox, Oy float64
	Dx, Dy float64
	bits   datastruct.Bits // True if pixel has already been evaluated
}

// NewTextureRGBA creates a new TextureRGBA from the supplied parameters
func NewTextureRGBA(width, height int, src ColorField, ox, oy, dx, dy float64) *TextureRGBA {
	bits := datastruct.NewBits(width * height)
	rect := image.Rectangle{image.Point{}, image.Point{width, height}}
	img := image.NewRGBA(rect)
	return &TextureRGBA{src, rect, img, width, ox, oy, dx, dy, bits}
}

// ColorModel implements the ColorModel function in the Image interface.
func (t *TextureRGBA) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds implements the Bounds function in the Image interface.
func (t *TextureRGBA) Bounds() image.Rectangle {
	return t.Rect
}

// At implements the At function in the Image interface.
func (t *TextureRGBA) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(t.Rect)) {
		return color.RGBA{}
	}
	// Convert x, y to bit index
	i := x + y*t.Stride
	if t.bits.Get(i) {
		return t.Img.RGBAAt(x, y)
	}
	// Pixel not set - evaluate it
	col := t.Src.Eval2(t.Ox+float64(x)*t.Dx, t.Oy+float64(y)*t.Dy)
	t.Img.Set(x, y, col)
	t.bits.Set(i)

	return col
}

// TextureGray16 is a lazily evaluated Gray16 image. For expensive textures this allows only the requested pixels
// to be calculated, and not the entire image.
type TextureGray16 struct {
	Src    ColorField
	Rect   image.Rectangle
	Img    *image.Gray16 // Evaluated pixels
	Stride int
	Ox, Oy float64
	Dx, Dy float64
	bits   datastruct.Bits // True if pixel has already been evaluated
}

// NewTextureGray16 creates a new TextureRGBA from the supplied parameters
func NewTextureGray16(width, height int, src ColorField, ox, oy, dx, dy float64) *TextureGray16 {
	bits := datastruct.NewBits(width * height)
	rect := image.Rectangle{image.Point{}, image.Point{width, height}}
	img := image.NewGray16(rect)
	return &TextureGray16{src, rect, img, width, ox, oy, dx, dy, bits}
}

// ColorModel implements the ColorModel function in the Image interface.
func (t *TextureGray16) ColorModel() color.Model {
	return color.Gray16Model
}

// Bounds implements the Bounds function in the Image interface.
func (t *TextureGray16) Bounds() image.Rectangle {
	return t.Rect
}

// At implements the At function in the Image interface.
func (t *TextureGray16) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(t.Rect)) {
		return color.Gray16{}
	}
	// Convert x, y to bit index
	i := x + y*t.Stride
	if t.bits.Get(i) {
		return t.Img.Gray16At(x, y)
	}
	// Pixel not set - evaluate it
	col := t.Src.Eval2(t.Ox+float64(x)*t.Dx, t.Oy+float64(y)*t.Dy)
	t.Img.Set(x, y, col)
	t.bits.Set(i)

	return col
}
