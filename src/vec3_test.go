package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	a := Vec3{1, 2, 3}
	b := a.Add(Vec3{4, 5, 6})

	if b.x != 5 || b.y != 7 || b.z != 9 {
		t.Error("Addition is incorrect")
	}
}

func TestSub(t *testing.T) {
	a := Vec3{6, 1, 3}
	b := a.Sub(Vec3{2, 3, -3})

	if b.x != 4 || b.y != -2 || b.z != 6 {
		t.Error("Subtract is incorrect")
	}
}

func TestMult(t *testing.T) {
	a := Vec3{2, 1, -4}
	b := a.Mult(Vec3{4, 8, -2})

	if b.x != 8 || b.y != 8 || b.z != 8 {
		t.Error("Multiply is incorrect")
	}
}

func TestDiv(t *testing.T) {
	a := Vec3{6, 81, 16}
	b := a.Div(Vec3{2, 9, 2})

	if b.x != 3 || b.y != 9 || b.z != 8 {
		t.Error("Div is incorrect")
	}
}

func TestDot(t *testing.T) {
	a := Vec3{7, 2, 5}
	b := a.Dot(Vec3{1, 7, 4})

	if b != 41 {
		t.Error("Dot product is incorrect")
	}
}

func TestCross(t *testing.T) {
	a := Vec3{8, 2, 4}
	b := a.Cross(Vec3{5, 6, 2})

	if b.x != -20 || b.y != 4 || b.z != 38 {
		t.Error("Cross product is incorrect")
	}
}

func TestMultF(t *testing.T) {
	a := Vec3{2, 1, 3}
	b := a.MultF(4)

	if b.x != 8 || b.y != 4 || b.z != 12 {
		t.Error("Multiply by scalar is incorrect")
	}
}

func TestDivF(t *testing.T) {
	a := Vec3{4, 8, 32}
	b := a.DivF(2)

	if b.x != 2 || b.y != 4 || b.z != 16 {
		t.Error("Divide by scalar is incorrect")
	}
}

func TestLen(t *testing.T) {
	a := Vec3{3, 4, 5}
	b := a.Len()

	if b != 7.0710678118654755 {
		t.Error("Length is incorrect")
	}
}

func TestUnitVec(t *testing.T) {
	a := Vec3{16, 8, 2}
	b := a.UnitVec()

	if b.x != 0.8888888888888888 || b.y != 0.4444444444444444 || b.z != 0.1111111111111111 {
		t.Error("Unit vector is incorrect")
	}
}
