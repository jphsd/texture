package surface

import (
	"github.com/jphsd/texture"
	"github.com/jphsd/texture/color"
	col "image/color"
	"math"
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

func (s *Surface) Eval2(x, y float64) col.Color {
	// For any point, the color rendered is the sum of the emissive, ambient and the diffuse/specular
	// contributions from all of the lights.

	material := s.Mat
	normals := s.Normals
	if normals == nil {
		normals = &texture.DefaultNormal{}
	}
	ambient := s.Ambient
	view := []float64{0, 0, 1}

	emm, amb, diff, spec, shine := material.Eval2(x, y) // Emissive

	// Emissive
	lemm := &color.FRGBA{}
	if emm != nil {
		lemm = emm
	}

	// Ambient
	acol, _, _, _ := ambient.Eval2(x, y)
	lamb := amb.Prod(acol) // Ambient
	col := lemm
	col = col.Add(lamb)
	if diff == nil {
		return col
	}

	// Cummulative diffuse and specular for all lights
	normal := normals.Eval2(x, y)
	cdiff, cspec := &color.FRGBA{}, &color.FRGBA{}
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
		if spec != nil {
			if blinn {
				// Blinn-Phong
				half := Norm([]float64{dir[0] + view[0], dir[1] + view[1], dir[2] + view[2]})
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
		}
	}
	col = col.Add(cdiff)
	col = col.Add(cspec)
	return col
}
