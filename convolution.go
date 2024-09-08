package texture

type Convolution struct {
	Name string
	Src  Field
	Kern [][]float64 // triplet {dx, dy, w}
}

// kernel {dx, dy, w}
func NewConvolution(src Field, kern [][]float64, norm bool) *Convolution {
	if norm {
		var sum, sumpos float64
		for i := range kern {
			sum += kern[i][2]
			if kern[i][2] > 0 {
				sumpos += kern[i][2]
			}
		}
		if sum < 0 || sum > 0 {
			for i := range kern {
				kern[i][2] /= sum
			}
		} else if sumpos > 0 {
			for i := range kern {
				kern[i][2] /= sumpos
			}
		}
	}
	return &Convolution{"Convolution", src, kern}
}

func (c *Convolution) Eval2(x, y float64) float64 {
	var sum float64
	for _, k := range c.Kern {
		v := c.Src.Eval2(x+k[0], y+k[1])
		v *= k[2]
		sum += v
	}
	return sum
}

// TODO kernel generation helper functions
// See https://en.wikipedia.org/wiki/Kernel_(image_processing)
// and http://www.dspguide.com/ch24/1.htm
// and https://docs.gimp.org/2.8/en/plug-in-convmatrix.html
// and https://programmathically.com/understanding-convolutional-filters-and-convolutional-kernels/
// and https://www.taylorpetrick.com/blog/post/convolution-part3
// and https://legacy.imagemagick.org/Usage/convolve/ ***
// Prewitt - edge
// Sobel - edge
// Kirsch - edge
// Roberts - edge
// Ridge
// Sharpen
// Gaussian
// Box
// Laplacian
