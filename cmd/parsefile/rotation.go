package main

import "fmt"

type RotationSingle struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
	W float32 `json:"w"`
}

func (rs *RotationSingle) String() string {
	return fmt.Sprintf("RotationSingle{%f,%f,%f,%f}", rs.X, rs.Y, rs.Z, rs.W)
}
