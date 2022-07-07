package texture

import (
	"github.com/jphsd/graphics2d"
	"github.com/jphsd/graphics2d/util"
	"image/color"
)

// Reflect contains a line along which a reflection is performed. The line defines where the
// mirror is. Points on the + side of the line remain untransformed, points on the other are
// reflected through the transformation.
type Reflect struct {
	Name  string
	Src   Field
	Start []float64
	End   []float64
	Xfm   *graphics2d.Aff3
}

// NewReflect creates a new Reflection placing the mirror along lp1, lp2.
func NewReflect(src Field, lp1, lp2 []float64) *Reflect {
	xfm := graphics2d.NewAff3()
	xfm.Reflect(lp1[0], lp1[1], lp2[0], lp2[1])
	return &Reflect{"Reflect", src, lp1, lp2, xfm}
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
	Name  string
	Src   VectorField
	Start []float64
	End   []float64
	Xfm   *graphics2d.Aff3
}

// NewReflect creates a new Reflection placing the mirror along lp1, lp2.
func NewReflectVF(src VectorField, lp1, lp2 []float64) *ReflectVF {
	xfm := graphics2d.NewAff3()
	xfm.Reflect(lp1[0], lp1[1], lp2[0], lp2[1])
	return &ReflectVF{"ReflectVF", src, lp1, lp2, xfm}
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
	Name  string
	Src   ColorField
	Start []float64
	End   []float64
	Xfm   *graphics2d.Aff3
}

// NewReflect creates a new Reflection placing the mirror along lp1, lp2.
func NewReflectCF(src ColorField, lp1, lp2 []float64) *ReflectCF {
	xfm := graphics2d.NewAff3()
	xfm.Reflect(lp1[0], lp1[1], lp2[0], lp2[1])
	return &ReflectCF{"ReflectCF", src, lp1, lp2, xfm}
}

// Eval2 implements the Field interface.
func (r *ReflectCF) Eval2(x, y float64) color.Color {
	pt := []float64{x, y}
	if util.SideOfLine(r.Start, r.End, pt) > 0 {
		pt = r.Xfm.Apply(pt)[0]
	}
	return r.Src.Eval2(pt[0], pt[1])
}
