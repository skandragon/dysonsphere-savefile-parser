package main

import "fmt"

// PositionSingle is a position (x, y, z) with float32 components.
type PositionSingle struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

func (ps *PositionSingle) String() string {
	return fmt.Sprintf("PositionSingle{%f,%f,%f}", ps.X, ps.Y, ps.Z)
}
