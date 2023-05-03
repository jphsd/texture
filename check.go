package texture

// Regular polygon tilings of the plane - triangle, square, hexagon.

type Triangles struct {
	Name  string
	Scale float64
}

func NewTriangles(s float64) *Triangles {
	return &Triangles{"Triangles", s}
}

func (t *Triangles) Eval2(x, y float64) float64 {
	// Even triangles are 1, odd -1
	_, rx := MapValueToLambda(x, t.Scale)
	_, ry := MapValueToLambda(y, t.Scale)
	if (rx + ry) < t.Scale {
		return 1
	}
	return -1
}

type Squares struct {
	Name  string
	Scale float64
}

func NewSquares(s float64) *Squares {
	return &Squares{"Squares", s}
}

func (s *Squares) Eval2(x, y float64) float64 {
	// Even squares are 1, odd -1
	c, _ := MapValueToLambda(x, s.Scale)
	if x < 0 {
		c++
	}
	r, _ := MapValueToLambda(y, s.Scale)
	if y < 0 {
		r++
	}
	v := r + c
	if v&0x1 == 1 {
		return -1
	}
	return 1
}

type Hexagons struct {
	Name  string
	Scale float64
}

func NewHexagons(s float64) *Hexagons {
	return &Hexagons{"Hexagons", s}
}

func (h *Hexagons) Eval2(x, y float64) float64 {
	// Even columnds are 1, 0, -1, odd columns are -1, 1, 0
	c, rx := MapValueToLambda(x, h.Scale)
	r, ry := MapValueToLambda(y, h.Scale)

	r %= 3
	if y < 0 {
		r = 2 - r
	}
	c %= 3
	if x < 0 {
		c = 2 - c
	}

	if r == c {
		// Diagonal requires special handling
		switch r {
		case 0:
			if (rx + ry) < h.Scale {
				return 0
			} else {
				return 1
			}
		case 1:
			if (rx + ry) < h.Scale {
				return 1
			} else {
				return -1
			}
		}
		// case 2
		if (rx + ry) < h.Scale {
			return -1
		}
		return 0
	}
	v := r + c
	switch v {
	case 1:
		return 1
	case 2:
		return 0
	}
	// case 3
	return -1
}
