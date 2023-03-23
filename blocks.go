package texture

import "math/rand"

type BlockNoise struct {
	Name   string
	Domain []float64
	Rows   int
	Cols   int
	Seed   int64
	Samps  int
	UseMax bool
	CellW  float64
	CellH  float64
	CLen   int
	cache  map[int][][]float64
}

func NewBlockNoise(w, h float64, r, c int, d float64) *BlockNoise {
	dom := []float64{w, h}
	seed := rand.Int63()
	cw, ch := w/float64(c), h/float64(r)
	samps := 4 * int(ch*d)
	return &BlockNoise{"BlockNoise", dom, r, c, seed, samps, false, cw, ch, 3 * c, make(map[int][][]float64)}
}

// Eval2 implements the Field interface.
func (bn *BlockNoise) Eval2(x, y float64) float64 {
	// Wrap x, y
	_, x = MapValueToLambda(x, bn.Domain[0])
	_, y = MapValueToLambda(y, bn.Domain[1])
	fr, fc := y/bn.CellH, x/bn.CellW
	r, c := int(fr), int(fc)
	x -= bn.CellW * float64(c)
	y -= bn.CellH * float64(r)
	lw := bn.CellW
	lh := 1.0

	// Run through #samps to see if we get a hit
	var v float64
	hit := false
	bb := bn.cbCache(r, c, bn.Samps, lw, lh)
	if bn.UseMax {
		hit, v = testBBH(x, y, bb)
	} else {
		hit, v = testBB(x, y, bb)
	}
	if hit {
		return v
	}

	// Need to check the preceeding cell (to the left)
	c -= 1
	if c < 0 {
		c = bn.Cols - 1
	}
	x += bn.CellW

	// Run through #samps to see if we get a hit
	bb = bn.cbCache(r, c, bn.Samps, lw, lh)
	if bn.UseMax {
		hit, v = testBBH(x, y, bb)
	} else {
		hit, v = testBB(x, y, bb)
	}
	if hit {
		return v
	}

	// Need to check the cell above
	c += 1
	if c == bn.Cols {
		c = 0
	}
	x -= bn.CellW
	r -= 1
	if r < 0 {
		r = bn.Rows - 1
	}
	y += bn.CellH

	// Run through #samps to see if we get a hit
	bb = bn.cbCache(r, c, bn.Samps, lw, lh)
	if bn.UseMax {
		_, v = testBBH(x, y, bb)
	} else {
		_, v = testBB(x, y, bb)
	}

	return v
}

// first hit wins
func testBB(x, y float64, bbs [][]float64) (bool, float64) {
	v := -1.0
	hit := false
	for _, bb := range bbs {
		if x < bb[0] || x > bb[2] || y < bb[1] || y > bb[3] {
			continue
		}
		hit = true
		v = bb[4]
		break
	}
	return hit, v
}

// largest v wins
func testBBH(x, y float64, bbs [][]float64) (bool, float64) {
	v := -1.0
	hit := false
	for _, bb := range bbs {
		if x < bb[0] || x > bb[2] || y < bb[1] || y > bb[3] {
			continue
		}
		hit = true
		if v < bb[4] {
			v = bb[4]
		}
	}
	return hit, v
}

func (bn *BlockNoise) cbCache(r, c, samps int, w, h float64) [][]float64 {
	cind := r*bn.Cols + c
	res := bn.cache[cind]
	if res != nil {
		return res
	}
	if len(bn.cache) == bn.CLen {
		bn.cache = make(map[int][][]float64)
	}
	res = bn.cellBlocks(r, c, samps, w, h)
	bn.cache[cind] = res
	return res
}

func (bn *BlockNoise) cellBlocks(r, c, samps int, w, h float64) [][]float64 {
	lr := rand.New(rand.NewSource(bn.Seed + int64(r*bn.Cols+c)))
	res := make([][]float64, samps)
	for i := 0; i < samps; i++ {
		ox, oy := lr.Float64()*bn.CellW, lr.Float64()*bn.CellH
		dx, dy := lr.Float64()*w, lr.Float64()*h
		v := lr.Float64()*2 - 1
		res[i] = []float64{ox, oy, ox + dx, oy + dy, v}
	}
	return res
}
