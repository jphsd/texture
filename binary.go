package texture

import "math/rand"

type Binary struct {
	Name string
	Seed int64
	bits [][]bool
}

func NewBinary(width, height int, seed int64) *Binary {
	lr := rand.New(rand.NewSource(seed))
	ba := make([][]bool, height)
	for i := 0; i < height; i++ {
		ba[i] = make([]bool, width)
		for j := 0; j < width; j++ {
			ba[i][j] = lr.Intn(2) > 0
		}
	}
	return &Binary{"Binary", seed, ba}
}

func (b *Binary) Eval2(x, y float64) float64 {
	w, h := len(b.bits[0]), len(b.bits)

	// Convert x and y into index into bits
	ix, iy := int(x)%w, int(y)%h
	if ix < 0 {
		ix += w
	}
	if iy < 0 {
		iy += h
	}

	if b.bits[iy][ix] {
		return 1
	}
	return -1
}
