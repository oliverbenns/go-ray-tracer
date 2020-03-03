package main

import (
	"testing"
)

func TestPointAtParameter(t *testing.T) {
	origin := Vec3{2, 3, 5}
	direction := Vec3{10, 5, 8}
	ray := Ray{origin, direction}
	point := ray.PointAtParameter(0.5)

	if point.x != 7 || point.y != 5.5 || point.z != 9 {
		t.Error("PointAtParameter is incorrect")
	}
}

func TestPointAtParameterNeg(t *testing.T) {
	origin := Vec3{2, 2, 2}
	direction := Vec3{8, 8, 8}
	ray := Ray{origin, direction}
	point := ray.PointAtParameter(-1)

	if point.x != -6 || point.y != -6 || point.z != -6 {
		t.Error("PointAtParameter is incorrect")
	}
}
