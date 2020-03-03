package main

import (
	"testing"
)

func TestVec3Equals(t *testing.T) {
	a := Vec3{1, 2, 3}
	b := Vec3{1, 2, 3}

	if !a.Equals(b) {
		t.Error("Equals is incorrect")
	}
}

func TestVec3Add(t *testing.T) {
	a := Vec3{1, 2, 3}
	b := a.Add(Vec3{4, 5, 6})

	if !b.Equals(Vec3{5, 7, 9}) {
		t.Error("Addition is incorrect")
	}
}

func TestVec3Sub(t *testing.T) {
	a := Vec3{6, 1, 3}
	b := a.Sub(Vec3{2, 3, -3})

	if !b.Equals(Vec3{4, -2, 6}) {
		t.Error("Subtract is incorrect")
	}
}

func TestVec3Mult(t *testing.T) {
	a := Vec3{2, 1, -4}
	b := a.Mult(Vec3{4, 8, -2})

	if !b.Equals(Vec3{8, 8, 8}) {
		t.Error("Multiply is incorrect")
	}
}

func TestVec3Div(t *testing.T) {
	a := Vec3{6, 81, 16}
	b := a.Div(Vec3{2, 9, 2})

	if !b.Equals(Vec3{3, 9, 8}) {
		t.Error("Div is incorrect")
	}
}

func TestVec3Dot(t *testing.T) {
	a := Vec3{7, 2, 5}
	b := a.Dot(Vec3{1, 7, 4})

	if b != 41 {
		t.Error("Dot product is incorrect")
	}
}

func TestVec3Cross(t *testing.T) {
	a := Vec3{8, 2, 4}
	b := a.Cross(Vec3{5, 6, 2})

	if !b.Equals(Vec3{-20, 4, 38}) {
		t.Error("Cross product is incorrect")
	}
}

func TestVec3MultF(t *testing.T) {
	a := Vec3{2, 1, 3}
	b := a.MultF(4)

	if !b.Equals(Vec3{8, 4, 12}) {
		t.Error("Multiply by scalar is incorrect")
	}
}

func TestVec3DivF(t *testing.T) {
	a := Vec3{4, 8, 32}
	b := a.DivF(2)

	if !b.Equals(Vec3{2, 4, 16}) {
		t.Error("Divide by scalar is incorrect")
	}
}

func TestVec3Len(t *testing.T) {
	a := Vec3{3, 4, 5}
	b := a.Len()

	if b != 7.0710678118654755 {
		t.Error("Length is incorrect")
	}
}

func TestVec3UnitVec(t *testing.T) {
	a := Vec3{16, 8, 2}
	b := a.UnitVec()

	if !b.Equals(Vec3{0.8888888888888888, 0.4444444444444444, 0.1111111111111111}) {
		t.Error("Unit vector is incorrect")
	}
}
