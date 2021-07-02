package texture

import (
	"math"
	"math/rand"
)

// Perlin contains the hash structures for generating a noise value between [-1,1).
// Note noise wraps in 256.
type Perlin struct {
	Ph    [512]uint8
	FFunc func(float64) float64
}

// NewPerlin initializes a new Perlin hash structure.
func NewPerlin() *Perlin {
	res := &Perlin{}
	// Initialize
	for i := 0; i < 256; i++ {
		res.Ph[i] = uint8(i)
	}

	// Scramble
	rand.Shuffle(256, func(i, j int) { res.Ph[i], res.Ph[j] = res.Ph[j], res.Ph[i] })
	// And replicate
	for i := 0; i < 256; i++ {
		res.Ph[i+256] = res.Ph[i]
	}

	return res
}

// Eval2 calculates the value at x,y based on the interpolated gradients of the four corners
// of the square which x,y resides in.
func (p *Perlin) Eval2(x, y float64) float64 {
	ix, iy := math.Floor(x), math.Floor(y)
	rx, ry := x-ix, y-iy
	u, v := blend(rx), blend(ry)

	// Select corners from hash of ix and iy % 256
	hx, hy := int(ix)&0xff, int(iy)&0xff
	a := int(p.Ph[hx]) + hy
	b := int(p.Ph[hx+1]) + hy
	res := lerp(v,
		lerp(u, gradient(p.Ph[a], rx, ry), gradient(p.Ph[b], rx-1, ry)),
		lerp(u, gradient(p.Ph[a+1], rx, ry-1), gradient(p.Ph[b+1], rx-1, ry-1)))
	if p.FFunc == nil {
		return res
	}
	return p.FFunc(res)
}

func lerp(t, s, e float64) float64 {
	return (1-t)*s + t*e
}

func gradient(hash uint8, x, y float64) float64 {
	u, v := 0.0, 0.0
	switch hash % 4 {
	case 0:
		u, v = x, y
	case 1:
		u, v = x, -y
	case 2:
		u, v = -x, y
	case 3:
		u, v = -x, -y
	}
	return u + v
}

// [0, 1] -> [0, 1]
func blend(t float64) float64 {
	// Improved blend poly^5
	// d1: 30t^4-60t^3+30t^2 : 0 at t=0,1
	// d2: 120t^3-180t^2+60t : 0 at t=0,1
	return t * t * t * (t*(t*6-15) + 10)
}
