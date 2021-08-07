package surface

import "math"

// Quaternion holds a quaternion defined as a vector (unit) and angle.
type Quaternion struct {
	Vec []float64
	Ang float64
	rot [][]float64
}

func NewQuaternion(v []float64, th float64) *Quaternion {
	q := &Quaternion{Unit(v), th, nil}
	q.rot = q.Rotation()
	return q
}

// Rotation returns the rotation matrix that describes a point transformed by the quaternion.
func (q *Quaternion) Rotation() [][]float64 {
	ca, sa := math.Cos(q.Ang), math.Sin(q.Ang)
	omca := 1 - ca
	x, y, z := q.Vec[0], q.Vec[1], q.Vec[2]
	x2, y2, z2 := x*x, y*y, z*z

	return [][]float64{
		{ca + x2*omca, x*y*omca - z*sa, x*z*omca + y*sa},
		{y*x*omca + z*sa, ca + y2*omca, y*z*omca - x*sa},
		{z*x*omca - y*sa, z*y*omca + x*sa, ca + z2*omca},
	}
}

// Apply applies the quaternion to the set of supplied points.
func (q *Quaternion) Apply(pts ...[]float64) [][]float64 {
	npts := make([][]float64, len(pts))
	r1, r2, r3 := q.rot[0], q.rot[1], q.rot[2]
	for i, pt := range pts {
		npt := make([]float64, 3)
		x, y, z := pt[0], pt[1], pt[2]
		npt[0] = r1[0]*x + r1[1]*y + r1[2]*z
		npt[1] = r2[0]*x + r2[1]*y + r2[2]*z
		npt[2] = r3[0]*x + r3[1]*y + r3[2]*z
		npts[i] = npt
	}
	return npts
}
