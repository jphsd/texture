/*
Package texture contains functions that can be combined to create textures.

All provide a Field interface that has Eval2(x, y float64) float64 which takes any x and y and
returns a value in [-1,1]. If a type doesn't support the entire 2D plane, then it must return 0
for values of x and y not supported.

1D Generators (have a wavelength, center offset, phase, angle):
  Zero - produces a flat field
  Sin - produces a Sine wave
  Square - produces a square wave
  Triangle - produces a triangular wave
  Saw - produces a saw wave
  NonLinear1 - produces a wave using a NonLinear function (reflected)
  NonLinear2 - produces a wave using two NonLinear functions, one for up and the other for down

1D Random/Multiple wavelength versions of the above

2D Generators:
  Perlin - produces a field using Ken Perlin's improved noise function
  Image - produces a field using an input image (converted to Gray16)

1D Filters map [0,1] to [0,1] (with A, B and C):
  Quantize - quantizes into C buckets
  Clip
  Sine
  Abs
  Pow
  Gaussian
  Saw

2D Combiners mix two or more source fields (rescaled to fit):
  Mul - product of two sources
  Add - sum of two sources
  Sub - difference of two sources (s1 - s2)
  Min - min of two sources
  Max - max of two sources
  Displace - source displaced by X and Y scaled sources

2D Other
  Transform - transforms X and Y with an Aff3 transform

Generators available from other contributors:
  OpenSimplex - github.com/ojrac/opensimplex-go
*/
package texture
