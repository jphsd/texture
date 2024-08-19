package surface

import (
	"github.com/jphsd/texture"
	"github.com/jphsd/texture/color"
	col "image/color"
	"math"
	"math/rand"
)

// Surface collects the ambient light, lights, a material, and normal map required to describe
// an area. If the normal map is nil then the standard normal is use {0, 0, 1}
type Surface struct {
	Ambient Light
	Lights  []Light
	Mat     Material
	Normals texture.VectorField
}

var blinn = false

// Eval2 implements the ColorField interface.
func (s *Surface) Eval2(x, y float64) col.Color {
	// For any point, the color rendered is the sum of the emissive, ambient and the diffuse/specular
	// contributions from all of the lights.

	material := s.Mat
	normals := s.Normals
	if normals == nil {
		normals = texture.DefaultNormal
	}
	ambient := s.Ambient
	view := []float64{0, 0, 1}

	em, amb, diff, spec, shine, rough := material.Eval2(x, y) // Emissive

	// Ambient
	acol, _, _, _ := ambient.Eval2(x, y)
	lamb := amb.Prod(acol) // Ambient
	col := em              // Emissive
	col = col.Add(lamb)
	/*
		if diff == nil {
			return col
		}
	*/

	// Cummulative diffuse and specular for all lights
	normal := normals.Eval2(x, y)
	if rough > 0 {
		normal = Roughen(rough, normal)
	}
	cdiff, cspec := color.FRGBA{}, color.FRGBA{}
	for _, light := range s.Lights {
		lcol, dir, dist, pow := light.Eval2(x, y)
		if lcol.IsBlack() {
			continue
		}
		lambert := Dot(dir, normal)
		if lambert < 0 {
			continue
		}
		if dist > 0 {
			lcol = lcol.Scale(pow / (dist * dist))
		}
		cdiff = cdiff.Add(lcol.Prod(diff.Scale(lambert))) // Diffuse
		//if spec != nil {
		if blinn {
			// Blinn-Phong
			half := Unit([]float64{dir[0] + view[0], dir[1] + view[1], dir[2] + view[2]})
			dp := Dot(half, normal)
			if dp > 0 {
				phong := math.Pow(dp, shine*4)
				cspec = cspec.Add(lcol.Prod(spec.Scale(phong))) // Specular
			}
		} else {
			// Phong
			dp := Dot(Reflect(dir, normal), view)
			if dp > 0 {
				phong := math.Pow(dp, shine)
				cspec = cspec.Add(lcol.Prod(spec.Scale(phong))) // Specular
			}
		}
		//}
	}
	col = col.Add(cdiff)
	col = col.Add(cspec)
	return col
}

// Roughen perturbates a vector by replacing it with a randomly orientented unit vector.
func Roughen(r float64, vec []float64) []float64 {
	// Construct a random unit vector pointing above the XY plane within r * 90 degrees
	theta := rand.Float64() * 2 * math.Pi
	phi := (1 - rand.Float64()*r) * math.Pi / 2
	cp := math.Cos(phi)
	rv := []float64{cp * math.Cos(theta), cp * math.Sin(theta), math.Sin(phi)}

	// Rotate into same plane as vec if necessary
	orig := []float64{0, 0, 1}
	cv := Cross(vec, orig)
	if cv[0]*cv[0]+cv[1]*cv[1]+cv[2]*cv[2] > 0 {
		cv[0], cv[1], cv[2] = -cv[0], -cv[1], -cv[2] // Flipping the normal
		quat := NewQuaternion(cv, math.Acos(Dot(vec, orig)))
		rv = quat.Apply(rv)[0]
	}

	return rv
}
