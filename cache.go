package texture

import "fmt"

// Cache provides a simple caching layer for a field which can be used when the expense of recalcuating a
// source is expensive or for when multiple queries are likely to be made e.g. morphological and convolution
// operations.
type Cache struct {
	Name       string
	Src        Field
	Resolution float64
	Limit      int
	oneovrres  float64
	cache      map[string]float64
}

// NewCache creates a new Cache with the specified resolution and limit. Once the limit is reached, the cache
// will be reset. The resoultion determines the accuracy of the x,y mapping to previous requests.
func NewCache(src Field, resolution float64, limit int) *Cache {
	return &Cache{"Cache", src, resolution, limit, 1 / resolution, make(map[string]float64)}
}

// Eval2 implements the Field interface.
func (c *Cache) Eval2(x, y float64) float64 {
	ind := c.cacheInd(x, y)
	res, ok := c.cache[ind]
	if !ok {
		res = c.Src.Eval2(x, y)
		if len(c.cache) > c.Limit {
			c.cache = make(map[string]float64)
		}
		c.cache[ind] = res
	}
	return res
}

func (c *Cache) cacheInd(x, y float64) string {
	x *= c.oneovrres
	if x < 0 {
		x -= 1
	}
	ix := int(x)
	y *= c.oneovrres
	if y < 0 {
		y -= 1
	}
	iy := int(y)
	return fmt.Sprintf("%d.%d", ix, iy)
}
