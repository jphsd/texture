package surface

import (
	"github.com/jphsd/texture"
	"image/color"
)

// BumpMap collects the ambient light, a direct light, a material, and normal map required to describe
// an area. If the normal map is nil then the standard normal is use {0, 0, 1}
type BumpMap struct {
	Ambient *Ambient
	Direct  Light
	Mat     Material
	Normals texture.VectorField
}

func (bm *BumpMap) Eval2(x, y float64) color.Color {
	// For any point, the color rendered is the sum of the ambient and the diffuse lights

	normals := bm.Normals
	if normals == nil {
		normals = &texture.DefaultNormal{}
	}

	material := bm.Mat
	if material == nil {
		material = DefaultMaterial
	}
	_, amb, diff, _, _ := material.Eval2(x, y)

	// Ambient
	ambient := bm.Ambient
	if ambient == nil {
		ambient = DefaultAmbient
	}
	col := amb.Prod(ambient.Color)

	// Diffuse
	direct := bm.Direct
	if direct == nil {
		direct = NewDirectional(color.White, []float64{-1, -1, 1})
	}
	lcol, dir, _, _ := direct.Eval2(x, y)
	if lcol.IsBlack() {
		return col
	}
	normal := normals.Eval2(x, y)
	lambert := Dot(dir, normal)
	if lambert < 0 {
		return col
	}

	col = col.Add(lcol.Prod(diff.Scale(lambert)))
	return col
}
