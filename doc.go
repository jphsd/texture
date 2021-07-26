/*
Package texture contains functions that can be combined to create textures.

Three types of texture field are supported, value fields which return a value in the range [-1,1]; vector
fields that typically return a triplet of values; and color fields which return a Color from the standard
color package.

All fields provide an Eval2(x, y float64) function which takes any x and y and
returns either a value, vector or color. If a type doesn't support the entire 2D plane, then it must return 0,
{0, ... , 0}, or color.Black for values of x and y not supported.

1D Generators (have a wavelength, center offset, phase, angle):
  Flat - produces a flat field
  NL1 - produces a wave using a NonLinear function (reflected)
  NL2 - produces a wave using two NonLinear functions, one for up and the other for down
  Noise1D - produces a wave derived from the Perlin noise function
  Saw - produces a saw wave
  Sin - produces a Sine wave
  Square - produces a square wave
  Triangle - produces a triangular wave

1D Random/Multiple wavelength versions of the above

2D Generators:
  Box - ove value if inside box or another if not
  NonLinear - produces a field filled with circles/elipses using a non-linear function
  Image - produces a field using an input image (converted to Gray16)
  Perlin - produces a field using Ken Perlin's improved noise function
  Triangles - produces a field filled with triangles

1D Filters map [-1,1] to [-1,1] (with A, B and C):
  Abs
  Clip
  Fold
  Gaussian
  Pow
  Quantize - quantizes into C buckets
  Sine

2D Combiners mix two or more source fields (rescaled to fit):
  ColorBlend - uses the value source to blend between to color sources
  ColorFields - uses four value sources to populate a color field (RGBA or HSLA)
  ColorSubstitute - switches between two color fields based on the values in a value field
  Combiner2
    Add - sum of two sources
    Avg - sum of two sources
    Diff - weighted sum of two sources based on their difference
    Max - max of two sources
    Min - min of two sources
    Mul - product of two sources
    Sub - difference of two sources (s1 - s2)
  Combiner3
    Substitute - substitue one value for another depending on a third value
  Displace - source displaced by X and Y scaled sources
  Window2 - select between two sources based on a region

2D Other:
  Distort - similar to Displace only the field itself is used as the displacement source
  Fractal - combine a source field through an affine transform multiple times
  IFS - combine a source field through an affine transform multiple times
  Transform - transforms a value field's X and Y with an Aff3 affine transform
  TransformCF - transforms a color field's X and Y with an Aff3 affine transform
  TransformVF - transforms a vector field's X and Y with an Aff3 affine transform
  UnitNormal - converts a vector field to a unit length vector field
  Window - window the source field to a region

Converters:
  Color - converts a value field to a gray scale color field
  ColorConv - converts a value field to a color field using a non-linear color map
  ColorSinCos - converts a value field to a clor field using sin and cos
  ColorVector - converts a vector field to a color field (RGB or HSL)
  Normal - converts a value field to a vector field using the finite difference method
  Select - selects one component of a vector field as a value field
  VectorCombine - combines the components of a vector field into a value field

Generators available from other contributors:
  OpenSimplex - github.com/ojrac/opensimplex-go
*/
package texture
