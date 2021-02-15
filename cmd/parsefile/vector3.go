package main

import "fmt"

// Vector3 is a position (x, y, z) with float32 components.
type Vector3 struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

// ParseVector3 extracts a ParseVector3 from a binary stream.
func ParseVector3(b *Buffer) *Vector3 {
	return &Vector3{
		X: b.GetFloat32(),
		Y: b.GetFloat32(),
		Z: b.GetFloat32(),
	}
}

func (ps *Vector3) String() string {
	return fmt.Sprintf("ParseVector3{%f,%f,%f}", ps.X, ps.Y, ps.Z)
}
