package main

import (
	"fmt"
	"math"
)

type VectorLF3 struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

func VectorLF3Zero() *VectorLF3 {
	return &VectorLF3{X: 0, Y: 0, Z: 0}
}

func (v *VectorLF3) String() string {
	return fmt.Sprintf("VectorLF3{%f,%f,%f}", v.X, v.Y, v.Z)
}

func (v *VectorLF3) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v *VectorLF3) Times(s float64) *VectorLF3 {
	return &VectorLF3{
		X: v.X * s,
		Y: v.Y * s,
		Z: v.Z * s,
	}
}
