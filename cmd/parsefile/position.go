package main

import "fmt"

// PositionSingle is a position (x, y, z) with float32 components.
type PositionSingle struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

// ParsePositionSingle extracts a PositionSingle from a binary stream.
func ParsePositionSingle(b *Buffer) *PositionSingle {
	return &PositionSingle{
		X: b.GetFloat32(),
		Y: b.GetFloat32(),
		Z: b.GetFloat32(),
	}
}

func (ps *PositionSingle) String() string {
	return fmt.Sprintf("PositionSingle{%f,%f,%f}", ps.X, ps.Y, ps.Z)
}
