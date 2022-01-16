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
