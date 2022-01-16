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
