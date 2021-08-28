package texture

import (
	"github.com/jphsd/graphics2d"
	"github.com/jphsd/graphics2d/util"
	"image/color"
	"math"
)

// Reflect contains a line along which a reflection is performed. The line defines where the
// mirror is. Points on the + side of the line remain untransformed, points on the other are
// reflected through the transformation.
type Reflect struct {
	Src   Field
	Start []float64
	End   []float64
	Xfm   *graphics2d.Aff3
}

// NewReflect creates a new Reflection placing the mirror along lp1, lp2.
func NewReflect(src Field, lp1, lp2 []float64) *Reflect {
	xfm := graphics2d.NewAff3()
	xfm.Reflect(lp1[0], lp1[1], lp2[0], lp2[1])
	return &Reflect{src, lp1, lp2, xfm}
}

// Eval2 implements the Field interface.
func (r *Reflect) Eval2(x, y float64) float64 {
	pt := []float64{x, y}
	if util.SideOfLine(r.Start, r.End, pt) < 0 {
		pt = r.Xfm.Apply(pt)[0]
	}
	return r.Src.Eval2(pt[0], pt[1])
}

// ReflectVF contains a line along which a reflection is performed. The line defines where the
// mirror is. Points on the + side of the line remain untransformed, points on the other are
// reflected through the transformation.
type ReflectVF struct {
	Src   VectorField
	Start []float64
	End   []float64
	Xfm   *graphics2d.Aff3
}

// NewReflect creates a new Reflection placing the mirror along lp1, lp2.
func NewReflectVF(src VectorField, lp1, lp2 []float64) *ReflectVF {
	xfm := graphics2d.NewAff3()
	xfm.Reflect(lp1[0], lp1[1], lp2[0], lp2[1])
	return &ReflectVF{src, lp1, lp2, xfm}
}

// Eval2 implements the Field interface.
func (r *ReflectVF) Eval2(x, y float64) []float64 {
	pt := []float64{x, y}
	if util.SideOfLine(r.Start, r.End, pt) > 0 {
		pt = r.Xfm.Apply(pt)[0]
	}
	return r.Src.Eval2(pt[0], pt[1])
}

// ReflectCF contains a line along which a reflection is performed. The line defines where the
// mirror is. Points on the + side of the line remain untransformed, points on the other are
// reflected through the transformation.
type ReflectCF struct {
	Src   ColorField
	Start []float64
	End   []float64
	Xfm   *graphics2d.Aff3
}

// NewReflect creates a new Reflection placing the mirror along lp1, lp2.
func NewReflectCF(src ColorField, lp1, lp2 []float64) *ReflectCF {
	xfm := graphics2d.NewAff3()
	xfm.Reflect(lp1[0], lp1[1], lp2[0], lp2[1])
	return &ReflectCF{src, lp1, lp2, xfm}
}

// Eval2 implements the Field interface.
func (r *ReflectCF) Eval2(x, y float64) color.Color {
	pt := []float64{x, y}
	if util.SideOfLine(r.Start, r.End, pt) > 0 {
		pt = r.Xfm.Apply(pt)[0]
	}
	return r.Src.Eval2(pt[0], pt[1])
}

// Kaleidoscope creates a new field by placing n mirrors, evenly spaced, starting at an angle, offs, and
// meeting at point c.
func Kaleidoscope(src Field, c []float64, n int, offs float64) Field {
	th := math.Pi / float64(n)
	ca := offs
	last := src
	for i := 0; i < n; i++ {
		pt := []float64{c[0] + math.Cos(ca)*10.0, c[1] + math.Sin(ca)*10.0}
		ca += th
		last = NewReflect(last, c, pt)
	}
	return last
}

// KaleidoscopeCF creates a new color field by placing n mirrors, evenly spaced, starting at an angle, offs, and
// meeting at point c.
func KaleidoscopeCF(src ColorField, c []float64, n int, offs float64) ColorField {
	th := math.Pi / float64(n)
	ca := offs
	last := src
	for i := 0; i < n; i++ {
		pt := []float64{c[0] + math.Cos(ca)*10.0, c[1] + math.Sin(ca)*10.0}
		ca += th
		last = NewReflectCF(last, c, pt)
	}
	return last
}

// Kaleidoscope2 creates a new field by placing n mirrors, evenly spaced, starting at an angle, offs, and
// meeting at point c. It then places a second set of mirrors at a distance d from c joing the spokes.
func Kaleidoscope2(src Field, c []float64, d float64, n int, offs float64) Field {
	th := math.Pi / float64(n)
	ca := offs
	last := src
	pts := make([][]float64, n)
	for i := 0; i < n; i++ {
		pt := []float64{c[0] + math.Cos(ca)*d, c[1] + math.Sin(ca)*d}
		pts[i] = pt
		ca += th
		last = NewReflect(last, c, pt)
	}
	prev := pts[0]
	for i := 1; i < n; i++ {
		cur := pts[i]
		last = NewReflect(last, prev, cur)
		prev = cur
	}
	last = NewReflect(last, prev, pts[0])
	return last
}

// Kaleidoscope2CF creates a new color field by placing n mirrors, evenly spaced, starting at an angle, offs, and
// meeting at point c. It then places a second set of mirrors at a distance d from c joing the spokes.
func Kaleidoscope2CF(src ColorField, c []float64, d float64, n int, offs float64) ColorField {
	th := math.Pi / float64(n)
	ca := offs
	last := src
	pts := make([][]float64, n)
	for i := 0; i < n; i++ {
		pt := []float64{c[0] + math.Cos(ca)*d, c[1] + math.Sin(ca)*d}
		pts[i] = pt
		ca += th
		last = NewReflectCF(last, c, pt)
	}
	prev := pts[0]
	for i := 1; i < n; i++ {
		cur := pts[i]
		last = NewReflectCF(last, prev, cur)
		prev = cur
	}
	last = NewReflectCF(last, prev, pts[0])
	return last
}
