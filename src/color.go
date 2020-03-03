package main

type Color struct {
	r, g, b float64
}

func (v *Color) Equals(w Color) bool {
	return v.r == w.r && v.g == w.g && v.b == w.b
}

func (v Color) MultF(s float64) Color {
	v.r *= s
	v.g *= s
	v.b *= s

	return v
}
