package texture

import (
	"github.com/jphsd/datastruct"
	"image"
	"image/color"
)

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
func NewTextureRGBA(width, height int, src ColorField, ox, oy, dx, dy float64, cache bool) *TextureRGBA {
	var bits datastruct.Bits
	if cache {
		bits = datastruct.NewBits(width * height)
	}
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
	return t.RGBAAt(x, y)
}

func (t *TextureRGBA) RGBAAt(x, y int) color.Color {
	if !(image.Point{x, y}.In(t.Rect)) {
		return color.RGBA{}
	}
	var i int
	if t.bits != nil {
		// Convert x, y to bit index
		i = x + y*t.Stride
		if t.bits.Get(i) {
			return t.Img.RGBAAt(x, y)
		}
	}
	// Pixel not set - evaluate it
	col := t.Src.Eval2(t.Ox+float64(x)*t.Dx, t.Oy+float64(y)*t.Dy)
	rgba, ok := col.(color.RGBA)
	if !ok {
		rgba = color.RGBAModel.Convert(col).(color.RGBA)
	}
	t.Img.Set(x, y, rgba)
	if t.bits != nil {
		t.bits.Set(i)
	}

	return rgba
}

// TextureRGBA64 is a lazily evaluated RGBA64 image. For expensive textures this allows only the requested pixels
// to be calculated, and not the entire image.
type TextureRGBA64 struct {
	Src    ColorField
	Rect   image.Rectangle
	Img    *image.RGBA64 // Evaluated pixels
	Stride int
	Ox, Oy float64
	Dx, Dy float64
	bits   datastruct.Bits // True if pixel has already been evaluated
}

// NewTextureRGBA64 creates a new TextureRGBA64 from the supplied parameters
func NewTextureRGBA64(width, height int, src ColorField, ox, oy, dx, dy float64, cache bool) *TextureRGBA64 {
	var bits datastruct.Bits
	if cache {
		bits = datastruct.NewBits(width * height)
	}
	rect := image.Rectangle{image.Point{}, image.Point{width, height}}
	img := image.NewRGBA64(rect)
	return &TextureRGBA64{src, rect, img, width, ox, oy, dx, dy, bits}
}

// ColorModel implements the ColorModel function in the Image interface.
func (t *TextureRGBA64) ColorModel() color.Model {
	return color.RGBA64Model
}

// Bounds implements the Bounds function in the Image interface.
func (t *TextureRGBA64) Bounds() image.Rectangle {
	return t.Rect
}

// At implements the At function in the Image interface.
func (t *TextureRGBA64) At(x, y int) color.Color {
	return t.RGBA64At(x, y)
}

func (t *TextureRGBA64) RGBA64At(x, y int) color.Color {
	if !(image.Point{x, y}.In(t.Rect)) {
		return color.RGBA64{}
	}
	var i int
	if t.bits != nil {
		// Convert x, y to bit index
		i = x + y*t.Stride
		if t.bits.Get(i) {
			return t.Img.RGBA64At(x, y)
		}
	}
	// Pixel not set - evaluate it
	col := t.Src.Eval2(t.Ox+float64(x)*t.Dx, t.Oy+float64(y)*t.Dy)
	rgba, ok := col.(color.RGBA64)
	if !ok {
		rgba = color.RGBA64Model.Convert(col).(color.RGBA64)
	}
	t.Img.Set(x, y, rgba)
	if t.bits != nil {
		t.bits.Set(i)
	}

	return rgba
}

// TextureGray16 is a lazily evaluated Gray16 image. For expensive textures this allows only the requested pixels
// to be calculated, and not the entire image.
type TextureGray16 struct {
	Src    Field
	Rect   image.Rectangle
	Img    *image.Gray16 // Evaluated pixels
	Stride int
	Ox, Oy float64
	Dx, Dy float64
	bits   datastruct.Bits // True if pixel has already been evaluated
}

// NewTextureGray16 creates a new TextureGray16 from the supplied parameters
func NewTextureGray16(width, height int, src Field, ox, oy, dx, dy float64, cache bool) *TextureGray16 {
	var bits datastruct.Bits
	if cache {
		bits = datastruct.NewBits(width * height)
	}
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
	return t.Gray16At(x, y)
}

func (t *TextureGray16) Gray16At(x, y int) color.Color {
	if !(image.Point{x, y}.In(t.Rect)) {
		return color.Gray16{}
	}
	var i int
	if t.bits != nil {
		// Convert x, y to bit index
		i = x + y*t.Stride
		if t.bits.Get(i) {
			return t.Img.Gray16At(x, y)
		}
	}
	// Pixel not set - evaluate it
	v := (t.Src.Eval2(t.Ox+float64(x)*t.Dx, t.Oy+float64(y)*t.Dy) + 1) / 2
	g16 := color.Gray16{uint16(v * 0xffff)}
	t.Img.Set(x, y, g16)
	if t.bits != nil {
		t.bits.Set(i)
	}

	return g16
}
