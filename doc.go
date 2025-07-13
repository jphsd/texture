/*
Package texture contains functions that can be combined to create image textures.

# 0. Getting Started

Here's a simple example that creates a Perlin value field and renders it to an image:

	f := texture.NewPerlin(12345)
	// Create a 600x600 image containing a field rendered from 0,0 in 0.075 steps
	img := texture.NewTextureGray16(600, 600, f, 0, 0, 0.075, 0.075, false)

# 1. Fields

Fields are interfaces that support returning a value for any given point in the field.

Three types of texture field are supported:
 1. [Field] (F) which maintains a field of float64, in the range [-1,1]
 2. [VectorField] (VF) which maintains a field of slice of float64, typically a triplet
 3. [ColorField] (CF) which maintains a field of [image/color.Color]

All fields provide an Eval2(x, y float64) function which takes any x and y and
returns either a value, vector or color.
If a field type doesn't support the entire 2D plane, then it must return 0,
{0, ..., 0}, or [image/color.Black] for values of x and y not supported.

# 2. Texture

A texture is a tree made up of field nodes and leaves.

A field is considered a leaf if it is not dependent on any other fields.
Examples of leaves are the gradient fields [LinearGradient], [RadialGradient] and [ConicGradient].

Any field that relies on another field or fields is considered a node.
Examples of nodes are the filter fields like [AbsFilter], [ClipFilter] and [InvertFilter].
All of these require at least one source field in order to operate.

Textures are built from the leaves upward until a root node has been created.
This node can then be passed to a function that will realize it, such as [NewTextureRGBA] which generates
an [image.RGBA] by repeatedly calling the root's [ColorField] Eval2 function for each image pixel.

# 3. Leaves - 1D

1D leaves vary only in one dimension, typically x.
  - [Uniform]
  - [LinearGradient]

Different rotations, scalings and offsets can be obtained using [Transform].

# 3.1 Uniform (F,VF,CF)

As the name suggests, these return a single value independent of the location within the field.

# 3.2 LinearGradient (F)

This field defines a gradient going from left to right using a [Wave] starting at 0 and repeating as
a function of the wave's wave length (lambda).

# 4. Leaves - 2D

2D leaves vary in x and y.
  - [Chequered]
  - [ConicGradient]
  - [Image]
  - [Perlin]
  - [RadialGradient]
  - [Shape]

Different rotations, scalings and offsets can be obtained using [Transform].

# 4.1 Chequered (F)

A collection of leaves for generating triangles, squares and hexagonal chequer boards.
The size of the cells is determined by the scale value.
  - [NewTriangles]
  - [NewSquares]
  - [NewHexagons]

The triangle and square boards are colored -1 and 1, whereas the hexagonal board is colored -1, 0 and 1.

# 4.2 ConicGradient (F)

This field defines a gradient that rotates around {0,0} using a [Wave] starting at 0 and mapping x
to theta / (2 * Pi).

# 4.3 Image (CF)

The image field uses the underlying image to figure the color value to return for any given location.
Supported interpolations are [NearestInterp], [LinearInterp], [CubicInterp], [P3Interp] and [P5Interp].

# 4.4 Perlin (F)

This field provides Perlin noise (aka gradient noise) using the supplied seed.
The field repeats over [256,256]. See [Perlin93].

# 4.5 RadialGradient (F)

This field defines a gradient that extends from {0,0} using a [Wave] starting at 0 and mapping to the
absolute distance from {0,0}.

# 4.6 Shape (F)

This field is defined by a [graphics2d.Shape].
Locations within the shape return 1 and all others -1.

# 4.7 Experimental (F)

Like the heading says, these are 2D field generator experiments, your mileage may vary.
  - [Binary] grid based binary noise
  - [BlinnField] metaballs. See [Blinn82]
  - [BlockNoise] rectangular block noise
  - [WorleyField] cellular basis functions. See [Worley96]

# 4.8 Third Party Generators (F)
  - [OpenSimplex] open simplex noise

# 5. Nodes - Filters

# 5.1 Value Filters (F)

Filters are nodes that do something with the value supplied by their source.
They map values in [-1,1] to another in [-1,1].
Some filters take A and B parameters, in which case the value filtered is A*value+B
  - [AbsFilter] applies [math.Abs] so the value will be in [0,1]
  - [CeilFilter] limits the value to [-1,C]
  - [ClipFilter] limits the value to [-1,1]
  - [Convolution] calculates value by applying a kernel to the source
  - [FloorFilter] limits the value to [C,1]
  - [FoldFilter] 'folds' a value outside of [-1,1] back in on itself
  - [InvertFilter] applies 0 - value
  - Morphological [Erode], [Dilate], [EdgeIn], [EdgeOut], [Edge], [Close], [Open], [TopHat], [BottomHat]
  - [NLFilter] applies a [NonLinear] to value
  - [OffsScaleFilter] applies A * (B + value)
  - [QuantizeFilter] quantizes the value
  - [RandQuantFilter] like [QuantizeFilter] but the values are scrambled
  - [RemapFilter] maps the value to the new domain [A,B]

# 5.2 Vector Filters (VF)
  - [UnitVector] modifies the magnitude of the vector to 1

# 5.3 Color Filters (CF)

# 6. Nodes - Combiners

The expressive range of the [texture] package is due to the ability to combine multiple source
fields together using some heuristic.

# 6.1 Value Combiners (F)
  - [MulCombiner] src1 * src2
  - [AddCombiner] src1 + src2
  - [SubCombiner] src1 - src2
  - [MinCombiner] min(src1, src2)
  - [MaxCombiner] max(src1, src2)
  - [AvgCombiner] (src1 + src2) / 2
  - [DiffCombiner] (1 - t) * src1 + t * src2, where t = (2 + src1 - src2) / 4
  - [WindowedCombiner] if src1 < A or src1 > B, src1, otherwise src2
  - [WeightedCombiner] src1 * A + src2 * B
  - [Blend] (1 - t) * src1 + t * src2, where t = (1 + src3) / 2
  - [StochasticBlend] src2 if random float < src3, src1 otherwise
  - [JitterBlend] as for [Blend] but with jitter added to t
  - [SubstituteCombiner] src1 if src3 < A or src3 > B, otherwise src2
  - [ShapeCombiner] src1 if location inside of shape, src2 otherwise
  - [ThresholdCombiner] src1 if less than threshold A, src2 otherwise

# 6.2 Vector Combiners (VF)
  - [ShapeCombinerVF] src1 if location inside of shape, src2 otherwise

# 6.3 Color Combiners (CF)
  - [ColorBlend] src1 and src2 are blended according to the value in src3 (F)
  - [ColorSubstitute] src1 if src3 < A or src3 > B, otherwise src2 (src3 is F)
  - [ShapeCombinerCF] src1 if location inside of shape, src2 otherwise

# 7. Nodes - Converters

Convert between F, VF and CF fields
  - [ColorToGray] maps color to [-1,1] via [color.Gray16Model]
  - [ColorSelect] maps one of R, G, B or A channels to [-1,1]
  - [Direction] takes the direction of a vector based on it's first two values and maps it to [-1,1]
  - [Magnitude] takes the magnitude of a vector (scaled and clamped to [-1,1])
  - [Select] takes the selected channel of a vector (scaled and clamped to [-1,1])
  - [Weighted] takes the weighted sum of a vector and clamps it to [-1,1]
  - [VectorFields] takes a slice of fields and creates a vector field
  - [VectorColor] - takes the four channels of a color field and maps them to a vector field
  - [Normal] - converts a field to a vector field of normals using the finite distance method
  - [ColorGray] maps [-1,1] to [Black,White] [image/color.Gray16] values
  - [ColorSinCos] uses one of six modes to convert [-1,1] to color using [math.Sin] and [math.Cos]
  - [ColorConv] uses a color interpolator to map [-1,1] to color
  - [ColorFields] uses 4 sources [-1,1], one each for R, G, B, A or H, S, L, A colors
  - [ColorVector] uses VF triplets to map to either RGB or HSL colors (A is opaque)

# 8. Nodes - Transformers

Transformers affect the value of x and y used when a field's Evals method is called.
They map (x, y) to (x', y').
  - [Displace]
  - [Displace2]
  - [Distort]
  - [Pixelate]
  - [Reflect]
  - [StochasticTiler]
  - [Strip]
  - [Tiler]
  - [Transform]
  - [Warp]

# 8.1 Transform (F,VF,CF)

This transformer applies an affine transform [graphics2d.Aff3] to input coordinates allowing for translations,
rotations, scalings and shearings.
When something other than a gradient in x is required, the affine transform can be used to move, scale and rotate
it to the desired location.

# 8.2 Tiler (F,VF,CF)

Tiler transforms allow finite areas to be replicated across the infinite plane.
Useful for creating repeating patterns and images.

# 8.3 Reflect (F,VF,CF)

The relection transforms take a line defined by two points as the location of a mirror.
Points on the positive side of the line remain unchanged while those on the negative are remapped.

# 8.4 Warp (F,VF,CF)

The warp transforms provide generalized image warping functionality not provided by the preceding.
They rely on a function [WarpFunc] to map points from one domain to the other.
  - [RadialWF]
  - [SwirlWF]
  - [DrainWF]
  - [RadialNLWF]
  - [PinchXWF]
  - [RippleXWF]
  - [RadialRippleWF]
  - [RadialWiggleWF]

# 8.5 Displace (F,VF,CF)

The displace transforms use two fields to perturb the location returned from a third field.
The degree of perturbation is controlled by a scaling factor.
The second version of each transform allows a generalized affine transform to be supplied rather than just
a fixed scaling.

# 8.6 Distort (F)

[Distort] provides a self referential transform that samples the field three times, once each for the x and y
displacements and once with the new x' and y'.

# 8.7 Pixelate (F,VF,CF)

These transforms apply a resolution filter to x and y.
Note that pixelate does not perform true pixelation in terms of averaging values over the desired resolution.

# 8.8 Strip (F,VF,CF)

These transforms replace y with a fixed value when performing Eval2(x, y).
[Strip] also implements the [Wave] interface and can be used in gradient leaves.

# 9. Nodes - Fractal

Three type of fractal nodel are available, [Fractal], [VariableFractal] and [IFS].

# 9.1 Fractal and VariableFractal (F)

Fractal nodes are a combination of both combiner and transformer nodes.
Each location is evaluated for each octave, with an affine transform being applied between evaluations to x and y,
and the resulting values then combined using an [OctaveCombiner].
  - [Fractal] standard fractal
  - [VariableFractal] fractal with octaves derived from a source field

Two [OctaveCombiner] are provided
  - [FBM] from supplied  Hurst and Lacunarity values
  - [MF] from supplied Hurst, Lacunarity and offset valaues

# 9.2 IFS (F)

IFS, or Iterated Fractal Systems [Barnsley88], take a series of contractive affine transformations and apply them
repeatedely to some depth (akin to octaves above).
  - [IFS]
  - [IFSCombiner]

# 10. Wave

Any type that implements [Wave] can be used to drive a gradient field.
This interface defines two methods - Eval(x float64) which returns a value in [-1, 1],
and Lambda() which returns the wave length of the wave.

Three types are defined as starting points and allow a variety of waveforms to be generated:
 1. [NLWave] multiple wave shapes with varying wave lengths
 2. [DCWave] one wave shape for the rising edge and one for the falling edge
 3. [ACWave] a wave shape per quadrant
 4. [InvertWave] takes a wave and inverts it

All of them utilize the non-linear functions provided in [graphics2d].
For convenience these are wrapped in [NonLinear], primarily so that the output is mapped from
[0,1] to [-1,1], and so that the slope name can be captured for JSON marshalling.

# 10.1 NLWave

[NLWave] takes a slice of wave lengths and a slice of slopes, together with flags that indicate
if slopes should be mirrored, and whether only a single cycle shold be generated.

# 10.2 DCWave

[DCWave] takes a slice of one or two wave lengths and a slice of one or two slopes, together with
a flag that indicate if only a single cycle shold be generated.

If only one wave length is provided, then it is used for both the rising and falling halves of the wave.
If only one slope type is provided, then it is used for both the rising and falling halves of the wave.
Hence, providing only one wave length and slope type is equivalent to mirroring.

When once is set, values less than 0 or greater than the wave length are returned as -1.

# 10.3 ACWave

[ACWave] takes a slice of one, two or four wave lengths and a slice of one, two or four slopes, together
with a flag that indicate if only a single cycle shold be generated. Note this waveform starts at 0, unlike
the other two which both start at -1.

If only one wave length is provided, then it is used for all quadrants of the wave.
If only one slope type is provided, then it is used for all quadrants of the wave.
If only two wave lengths are provided, then they are used for both halves of the wave.
If only two slope types are provided, then they are used for both halves of the wave.

When once is set, values less than 0 or greater than the wave length are returned as 0.

# 11. Utilities

# 11.1 Realization

Textures are realized by calling the root node's Eval2(x, y) method.
Image wrappers are provided that perform this step lazily and in a cacheable fashion.
The wrapper is responsible for defining the image bounds, the texture offset and step values.
These images can then be passed to [image/draw.Draw], [graphics2d.RenderShape] or to
[graphics2d.NewFilledPen] as source images.

Caching determines whether a value is evaluated once and cached, or always evaluated.
If a particular texture subgraph is expensive to compute, it may be better to evaluate it's domain
once and cache it in an image that can then be referenced through an [Image] node.
  - [TextureGray16] takes a value field
  - [TextureRGBA] takes a color field
  - [TextureRGBA64] takes a color field

# 11.2 Gradients

The 2D graphics packages in other languages, such as Java and SVG, have a notion of a gradient fill
or paint.
Go, however, doesn't since it's [golang.org/x/image/vector.Draw] takes an image.
To address this [texture] has some utility functions that use simple gradient textures to create the same
effect.

The gradients are all value fields and mapped to either [image/color.Gray16] or [image/color.RGBA].
In the latter case, by using [github.com/jphsd/graphics2d/image.Colorizer].
  - [NewLinearGray16]
  - [NewRadialGray16]
  - [NewConicGray16]
  - [NewLinearRGBA]
  - [NewRadialRGBA]
  - [NewConicRGBA]

# 12. Package Examples

[Chequered]: https://pkg.go.dev/github.com/jphsd/texture#hdr-4_1_Chequered__F_
[Barnsley88]: https://doi.org/10.1016/c2013-0-10335-2
[Blinn82]: https://dl.acm.org/doi/10.1145/357306.357310
[OpenSimplex]: https://pkg.go.dev/github.com/ojrac/opensimplex-go
[Perlin93]: https://dl.acm.org/doi/10.1145/325165.325247
[Worley96]: https://dl.acm.org/doi/10.1145/237170.237267
*/
package texture
