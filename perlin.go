package texture

import (
	"math"
	"math/rand"
)

// Perlin contains the hash structures for generating a noise value between [-1,1).
// Note noise wraps in 256.
type Perlin struct {
	Name string
	Seed int64
	ph   [512]uint8
}

// NewPerlin initializes a new Perlin hash structure.
func NewPerlin(seed int64) *Perlin {
	res := &Perlin{}
	res.Name = "Perlin"
	res.Seed = seed

	// Initialize hash
	for i := 0; i < 256; i++ {
		res.ph[i] = uint8(i)
	}

	lr := rand.New(rand.NewSource(seed))
	// Scramble hash
	lr.Shuffle(256, func(i, j int) { res.ph[i], res.ph[j] = res.ph[j], res.ph[i] })
	// And replicate
	for i := 0; i < 256; i++ {
		res.ph[i+256] = res.ph[i]
	}

	return res
}

// Eval2 calculates the value at x,y based on the interpolated gradients of the four corners
// of the square which x,y resides in. It implements the Field interface.
func (p *Perlin) Eval2(x, y float64) float64 {
	ix, iy := math.Floor(x), math.Floor(y)
	rx, ry := x-ix, y-iy
	u, v := blend(rx), blend(ry)

	// Select corners from hash of ix and iy % 256
	hx, hy := int(ix)&0xff, int(iy)&0xff
	a := int(p.ph[hx]) + hy
	b := int(p.ph[hx+1]) + hy
	res := lerp(v,
		lerp(u, gradient(p.ph[a], rx, ry), gradient(p.ph[b], rx-1, ry)),
		lerp(u, gradient(p.ph[a+1], rx, ry-1), gradient(p.ph[b+1], rx-1, ry-1)))
	return res
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
