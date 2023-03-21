package texture

import "math/rand"

type BlockNoise struct {
	Name   string
	Domain []float64
	Rows   int
	Cols   int
	Seed   int64
	Samps  int
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
	return &BlockNoise{"BlockNoise", dom, r, c, seed, samps, cw, ch, 3 * c, make(map[int][][]float64)}
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
	hit := false
	bb := bn.cbCache(r, c, bn.Samps, lw, lh)
	for i := 0; i < bn.Samps; i++ {
		if x < bb[i][0] || x > bb[i][2] || y < bb[i][1] || y > bb[i][3] {
			continue
		}
		hit = true
		break
	}
	if hit {
		return 1
	}

	// Need to check the preceeding cell (to the left)
	c -= 1
	if c < 0 {
		c = bn.Cols - 1
	}
	x += bn.CellW

	// Run through #samps to see if we get a hit
	bb = bn.cbCache(r, c, bn.Samps, lw, lh)
	for i := 0; i < bn.Samps; i++ {
		if x < bb[i][0] || x > bb[i][2] || y < bb[i][1] || y > bb[i][3] {
			continue
		}
		hit = true
		break
	}
	if hit {
		return 1
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
	for i := 0; i < bn.Samps; i++ {
		if x < bb[i][0] || x > bb[i][2] || y < bb[i][1] || y > bb[i][3] {
			continue
		}
		hit = true
		break
	}
	if hit {
		return 1
	}

	return -1.0
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
		res[i] = []float64{ox, oy, ox + dx, oy + dy}
	}
	return res
}
