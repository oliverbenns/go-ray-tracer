package main

import "math"

type Vec3 struct {
	x float64
	y float64
	z float64
}

func (v Vec3) Add(w Vec3) Vec3 {
	v.x += w.x
	v.y += w.y
	v.z += w.z
	return v
}

func (v Vec3) Sub(w Vec3) Vec3 {
	v.x -= w.x
	v.y -= w.y
	v.z -= w.z
	return v
}

func (v Vec3) Mult(w Vec3) Vec3 {
	v.x *= w.x
	v.y *= w.y
	v.z *= w.z
	return v
}

func (v Vec3) Div(w Vec3) Vec3 {
	v.x /= w.x
	v.y /= w.y
	v.z /= w.z
	return v
}

func (v *Vec3) Dot(w Vec3) float64 {
	return v.x*w.x + v.y*w.y + v.z*w.z
}

func (v Vec3) Cross(w Vec3) Vec3 {
	return Vec3{
		v.y*w.z - v.z*w.y,
		v.z*w.x - v.x*w.z,
		v.x*w.y - v.y*w.x,
	}
}

func (v Vec3) MultF(s float64) Vec3 {
	v.x *= s
	v.y *= s
	v.z *= s

	return v
}

func (v Vec3) DivF(s float64) Vec3 {
	v.x /= s
	v.y /= s
	v.z /= s
	return v
}

func (v *Vec3) Len() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

func (v *Vec3) UnitVec() Vec3 {
	return v.DivF(v.Len())
}

// Color getters
// Might be worth sharing an interface and having 2 structs.
func (v *Vec3) R() float64 {
	return v.x
}

func (v *Vec3) G() float64 {
	return v.y
}

func (v *Vec3) B() float64 {
	return v.z
}
