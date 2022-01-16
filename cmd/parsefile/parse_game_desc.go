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

// GameDesc describes the initial starting conditions.
type GameDesc struct {
	GalaxyAlgo         int32   `json:"galaxy_algo"`
	GalaxySeed         int32   `json:"galaxy_seed"`
	StarCount          int32   `json:"star_count"`
	PlayerPronto       int32   `json:"player_pronto"`
	ResourceMultiplier float32 `json:"resource_multiplier"`
	ThemeIDs           []int32 `json:"theme_ids"`
}

func parseGameDesc(b *Buffer) *GameDesc {
	checkVers(b, 2, "GameDesc")

	galaxyAlgo := b.GetInt32le()
	galaxySeed := b.GetInt32le()
	starCount := b.GetInt32le()
	playerPronto := b.GetInt32le()
	resourceMultiplier := b.GetFloat32()

	tidlen := b.GetInt32le()
	tids := make([]int32, int(tidlen))
	for i := 0; int32(i) < tidlen; i++ {
		tids[i] = b.GetInt32le()
	}

	gameDesc := &GameDesc{
		GalaxyAlgo:         galaxyAlgo,
		GalaxySeed:         galaxySeed,
		StarCount:          starCount,
		PlayerPronto:       playerPronto,
		ResourceMultiplier: resourceMultiplier,
		ThemeIDs:           tids,
	}
	return gameDesc
}
