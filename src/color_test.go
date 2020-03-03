package main

import (
	"testing"
)

func TestColorEquals(t *testing.T) {
	a := Color{1, 2, 3}
	b := Color{1, 2, 3}

	if !a.Equals(b) {
		t.Error("Equals is incorrect")
	}
}

func TestColorMultF(t *testing.T) {
	a := Color{2, 1, 3}
	b := a.MultF(4)

	if !b.Equals(Color{8, 4, 12}) {
		t.Error("Multiply by scalar is incorrect")
	}
}
