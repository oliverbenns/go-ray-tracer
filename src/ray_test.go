package main

import (
	"testing"
)

func TestPointAtParameter(t *testing.T) {
	origin := Vec3{2, 3, 5}
	direction := Vec3{10, 5, 8}
	ray := Ray{origin, direction}
	point := ray.PointAtParameter(0.5)

	if !point.Equals(Vec3{7, 5.5, 9}) {
		t.Error("PointAtParameter is incorrect")
	}
}

func TestPointAtParameterNeg(t *testing.T) {
	origin := Vec3{2, 2, 2}
	direction := Vec3{8, 8, 8}
	ray := Ray{origin, direction}
	point := ray.PointAtParameter(-1)

	if !point.Equals(Vec3{-6, -6, -6}) {
		t.Error("PointAtParameter is incorrect")
	}
}
