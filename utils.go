package texture

import (
	g2d "github.com/jphsd/graphics2d"
	"github.com/jphsd/graphics2d/image"
	"image/color"
	"math"
)

// Utility functions for producing gradient images for use with graphics2d

func NewLinearGray16(w, h int, p1, p2 []float64, wf *NonLinear, mirror, once bool) *TextureGray16 {
	dx, dy := p2[0]-p1[0], p2[1]-p1[1]
	th := math.Atan2(dy, dx)
	lambda := math.Hypot(dx, dy)

	if wf == nil {
		wf = NewNLLinear()
	}
	nl := NewNLWave([]float64{lambda}, []*NonLinear{wf}, mirror, once)
	f1 := NewLinearGradient(nl)
	xfm := g2d.Rotate(-th)
	xfm.Translate(-p1[0], -p1[1])
	f2 := NewTransform(f1, xfm)
	return NewTextureGray16(w, h, f2, 0, 0, 1, 1, false)
}

func NewRadialGray16(w, h int, c []float64, r float64, wf *NonLinear, mirror, once bool) *TextureGray16 {
	if wf == nil {
		wf = NewNLLinear()
	}
	nl := NewNLWave([]float64{r}, []*NonLinear{wf}, mirror, once)
	f1 := NewRadialGradient(nl)
	xfm := g2d.Translate(-c[0], -c[1])
	f2 := NewTransform(f1, xfm)
	return NewTextureGray16(w, h, f2, 0, 0, 1, 1, false)
}

func NewEllipticalGray16(w, h int, c []float64, rx, ry, th float64, wf *NonLinear, mirror, once bool) *TextureGray16 {
	if wf == nil {
		wf = NewNLLinear()
	}
	nl := NewNLWave([]float64{rx}, []*NonLinear{wf}, mirror, once)
	f1 := NewRadialGradient(nl)
	// Reverse order
	xfm := g2d.Translate(c[0], c[1])
	xfm.Rotate(th)
	xfm.Scale(1, ry/rx)
	xfm.Translate(-c[0], -c[1])
	f2 := NewTransform(f1, xfm)
	return NewTextureGray16(w, h, f2, 0, 0, 1, 1, false)
}

func NewConicGray16(w, h int, c []float64, th float64, wf *NonLinear) *TextureGray16 {
	if wf == nil {
		wf = NewNLLinear()
	}
	nl := NewNLWave([]float64{256}, []*NonLinear{wf}, true, false)
	f1 := NewConicGradient(nl)
	xfm := g2d.Rotate(-th)
	xfm.Translate(-c[0], -c[1])
	f2 := NewTransform(f1, xfm)
	return NewTextureGray16(w, h, f2, 0, 0, 1, 1, false)
}

// Colorizer wrappers around the grayscale gradients

func NewLinearRGBA(w, h int, p1, p2 []float64, c1, c2 color.Color, wf *NonLinear, mirror, once bool) *image.Colorizer {
	return image.NewColorizer(NewLinearGray16(w, h, p1, p2, wf, mirror, once), c1, c2, nil, nil, false)
}

func NewRadialRGBA(w, h int, c []float64, r float64, c1, c2 color.Color, wf *NonLinear, mirror, once bool) *image.Colorizer {
	return image.NewColorizer(NewRadialGray16(w, h, c, r, wf, mirror, once), c1, c2, nil, nil, false)
}

func NewEllipticalRGBA(w, h int, c []float64, rx, ry, th float64, c1, c2 color.Color, wf *NonLinear, mirror, once bool) *image.Colorizer {
	return image.NewColorizer(NewEllipticalGray16(w, h, c, rx, ry, th, wf, mirror, once), c1, c2, nil, nil, false)
}

func NewConicRGBA(w, h int, c []float64, th float64, c1, c2 color.Color, wf *NonLinear) *image.Colorizer {
	return image.NewColorizer(NewConicGray16(w, h, c, th, wf), c1, c2, nil, nil, false)
}
