package texture

import (
	"github.com/jphsd/graphics2d/util"
	"math"
)

// Triangles renders triangles (defined in [0,1]) at the supplied values into the rectangles defined by the
// two lambdas. Note only one of OffsetX, OffsetY can be non-zero. An optional filter can be specified. This
// structure allows any shape to be expressed as triangles and rendered into textures.
type Triangles struct {
	LambdaX, LambdaY float64 // [1,...)
	PhaseX, PhaseY   float64 // [0,1]
	OffsetX, OffsetY float64 // [0,1]
	FFunc            func(float64) float64
	CosTh, SinTh     float64
	Tris             []tri
	Values           []float64
}

// NewTriangles returns a new instance of Triangles based on the supplied slice of triangle points. Each triangle
// is rendered with the matching value from values. The triangles are rendered in order, so if a point is in more
// than one triangle, the point will return the value associated with the last triangle.
func NewTriangles(lambdaX, lambdaY, theta float64, triangles [][][]float64, values []float64) *Triangles {
	if lambdaX < 1 {
		lambdaX = 1
	}
	if lambdaY < 1 {
		lambdaY = 1
	}
	// Snap to quad
	ct := math.Cos(theta)
	if closeTo(0, ct) {
		ct = 0
	} else if closeTo(1, ct) {
		ct = 1
	} else if closeTo(-1, ct) {
		ct = -1
	}
	st := math.Sin(theta)
	if closeTo(0, st) {
		st = 0
	} else if closeTo(1, st) {
		st = 1
	} else if closeTo(-1, st) {
		st = -1
	}
	for i := 0; i < len(values); i++ {
		if values[i] < -1 {
			values[i] = -1
		} else if values[i] > 1 {
			values[i] = 1
		}
	}
	return &Triangles{lambdaX, lambdaY, 0, 0, 0, 0, nil, ct, st, trianglesToTri(triangles), values}
}

type tri struct {
	Points [][]float64
	Bounds [][]float64
}

func trianglesToTri(tris [][][]float64) []tri {
	n := len(tris)
	res := make([]tri, n)
	for i := 0; i < n; i++ {
		res[i].Points = tris[i]
		res[i].Bounds = util.BoundingBox(tris[i]...)
	}
	return res
}

// Eval2 implements the Field interface.
func (t *Triangles) Eval2(x, y float64) float64 {
	u := x*t.CosTh + y*t.SinTh
	v := -x*t.SinTh + y*t.CosTh
	u, v = t.XYToUV(u, v)
	// Run through all triangles until we find one u,v is in
	res := -1.0
	for i, triangle := range t.Tris {
		bb := triangle.Bounds
		if u < bb[0][0] || u > bb[1][0] || v < bb[0][1] || v > bb[1][1] {
			continue
		}
		if util.PointInTriangle([]float64{u, v}, triangle.Points[0], triangle.Points[1], triangle.Points[2]) {
			res = t.Values[i]
			break
		}
	}
	if t.FFunc == nil {
		return res
	}
	return t.FFunc(res)
}

// XYToUV converts values in (-inf,inf) to [0,1] based on the generator's orientation, lambdas and phase values.
func (t *Triangles) XYToUV(x, y float64) (float64, float64) {
	nx := 0
	for x < 0 {
		x += t.LambdaX
		nx--
	}
	for x > t.LambdaX {
		x -= t.LambdaX
		nx++
	}
	ny := 0
	for y < 0 {
		y += t.LambdaY
		ny--
	}
	for y > t.LambdaY {
		y -= t.LambdaY
		ny++
	}

	if !util.Equals(0, t.OffsetX) {
		offs := float64(ny) * t.OffsetX
		offs -= math.Floor(offs)
		if offs < 0 {
			offs = 1 - offs
		}
		u := x/t.LambdaX + t.PhaseX + offs
		for u > 1 {
			u -= 1
		}
		v := y/t.LambdaY + t.PhaseY
		if v > 1 {
			v -= 1
		}
		return u, v
	}
	u := x/t.LambdaX + t.PhaseX
	if u > 1 {
		u -= 1
	}
	offs := float64(nx) * t.OffsetY
	offs -= math.Floor(offs)
	if offs < 0 {
		offs = 1 - offs
	}
	v := y/t.LambdaY + t.PhaseY + offs
	for v > 1 {
		v -= 1
	}

	return u, v
}
