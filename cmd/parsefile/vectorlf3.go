/*
 * Copyright 2021-2022 Michael Graff
 *
 * Licensed under the Apache License, Version 2.0 (the "License")
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"
	"math"
)

// VectorLF3 - a 3d vector with float64 components
type VectorLF3 struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

// VectorLF3Zero - return a new VectorLF3 with all zero components
func VectorLF3Zero() *VectorLF3 {
	return &VectorLF3{X: 0, Y: 0, Z: 0}
}

func (v *VectorLF3) String() string {
	return fmt.Sprintf("VectorLF3{%f,%f,%f}", v.X, v.Y, v.Z)
}

// Magnitude - Return the magnitude of the vector
func (v *VectorLF3) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Times -- multiply each component by s.
func (v *VectorLF3) Times(s float64) *VectorLF3 {
	return &VectorLF3{
		X: v.X * s,
		Y: v.Y * s,
		Z: v.Z * s,
	}
}
