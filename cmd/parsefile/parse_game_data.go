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

// GameData holds top level game data, like the galaxy, factories, etc.
type GameData struct {
	Name            string           `json:"name"`
	Descriptor      *GameDesc        `json:"descriptor"`
	StarCount       int32            `json:"star_count"`
	PlanetFactories []*PlanetFactory `json:"planet_factories"`
	GameTicks       int64            `json:"game_ticks"`
	Galaxy          *Universe        `json:"galaxy"`
}

func parseGameData(b *Buffer) *GameData {
	const VERSION = 2
	checkVers(b, VERSION, "GameData")

	ret := &GameData{
		Name:       b.GetString(),
		Descriptor: parseGameDesc(b),
	}
	ret.Galaxy = MakeUniverse(ret.Descriptor)

	ret.GameTicks = b.GetInt64le()

	parseGamePreferences(b)
	parseGameHistory(b)

	b.GetBoolean() // hidePlayerModel
	b.GetBoolean() // disableController

	parseGameStatistics(b)

	b.GetInt32le() // planetId

	parsePlayer(b)

	factoryCount := b.GetInt32le()

	parseGalacticTransport(b) // yes, this is out of order in the code too...

	ret.PlanetFactories = make([]*PlanetFactory, int(factoryCount))
	for i := 0; int32(i) < factoryCount; i++ {
		ret.PlanetFactories[i] = parsePlanetFactory(b, i)
	}

	ret.StarCount = b.GetInt32le()

	// dyson sphere mess...

	return ret
}
