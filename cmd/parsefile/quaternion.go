package main

import "fmt"

// Quaternion holds X, Y, Z, and W components as float32
type Quaternion struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
	W float32 `json:"w"`
}

// ParseQuaternion extracts a RotationSingle from a binary stream
func ParseQuaternion(b *Buffer) *Quaternion {
	return &Quaternion{
		X: b.GetFloat32(),
		Y: b.GetFloat32(),
		Z: b.GetFloat32(),
		W: b.GetFloat32(),
	}
}

func (rs *Quaternion) String() string {
	return fmt.Sprintf("Quaternion{%f,%f,%f,%f}", rs.X, rs.Y, rs.Z, rs.W)
}
