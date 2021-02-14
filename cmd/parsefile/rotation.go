package main

import "fmt"

// RotationSingle holds X, Y, Z, and W components as float32
type RotationSingle struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
	W float32 `json:"w"`
}

// ParseRotationSingle extracts a RotationSingle from a binary stream
func ParseRotationSingle(b *Buffer) *RotationSingle {
	return &RotationSingle{
		X: b.GetFloat32(),
		Y: b.GetFloat32(),
		Z: b.GetFloat32(),
		W: b.GetFloat32(),
	}
}

func (rs *RotationSingle) String() string {
	return fmt.Sprintf("RotationSingle{%f,%f,%f,%f}", rs.X, rs.Y, rs.Z, rs.W)
}
