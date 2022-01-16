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

package types

// PlanetType indicates the type of the planet's surface.
type PlanetType int32

const (
	// PlanetTypeNone means it doesn't exist I suppose?
	PlanetTypeNone PlanetType = 0

	// PlanetTypeVolcano are hot and likely rocky.
	PlanetTypeVolcano PlanetType = 1

	// PlanetTypeOcean means you'll need lots of landfill to mine.
	PlanetTypeOcean PlanetType = 2

	// PlanetTypeDesert are dry most likely.
	PlanetTypeDesert PlanetType = 3

	// PlanetTypeIce are ice giants.
	PlanetTypeIce PlanetType = 4

	// PlanetTypeGas are gas giants.
	PlanetTypeGas PlanetType = 5
)

func (t PlanetType) String() string {
	return [...]string{
		"None",
		"Volcano",
		"Ocean",
		"Desert",
		"Ice",
		"Gas",
	}[t]
}
