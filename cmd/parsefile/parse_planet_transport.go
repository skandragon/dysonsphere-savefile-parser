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

// PlanetTransport holds stations and not much else
type PlanetTransport struct {
	StationComponents []*StationComponent `json:"station_components"`
}

func parsePlanetTransport(b *Buffer) *PlanetTransport {
	const VERSION = 0
	checkVers(b, VERSION, "PlanetTransport")

	stationCursor := b.GetInt32le()
	b.GetInt32le() // stationCapacity
	stationRecycleCounter := b.GetInt32le()
	components := make([]*StationComponent, 0)
	for i := int32(1); i < stationCursor; i++ {
		id := b.GetInt32le()
		if id != 0 {
			if id != i {
				panic(fmt.Sprintf("id != i (%d, %d)", id, i))
			}
			components = append(components, parseStationComponent(b))
		}
	}
	for i := 0; int32(i) < stationRecycleCounter; i++ {
		b.GetInt32le() // recycle id?
	}

	return &PlanetTransport{
		StationComponents: components,
	}
}
