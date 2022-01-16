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

func parseGameHistory(b *Buffer) {
	checkVers(b, 2, "GameHistory")

	// TODO: parse recipeUnlocked
	count := b.GetInt32le()
	for i := 0; int32(i) < count; i++ {
		b.GetInt32le() // recipeID
	}

	// TODO: parse tutorialUnlocked
	count = b.GetInt32le()
	for i := 0; int32(i) < count; i++ {
		b.GetInt32le() // id
	}

	// TODO: parse featureKeys
	count = b.GetInt32le()
	for i := 0; int32(i) < count; i++ {
		b.GetInt32le() // id
	}

	// TODO: parse techStates
	count = b.GetInt32le()
	for i := 0; int32(i) < count; i++ {
		b.GetInt32le() // id
		b.GetBoolean() // unlocked
		b.GetInt32le() // curLevel
		b.GetInt32le() // maxLevel
		b.GetInt64le() // hashUnloaded
		b.GetInt64le() // hashNeeded
	}

	b.GetBoolean() // automanageLabItems
	b.GetInt32le() // currentTech

	// techQueue
	count = b.GetInt32le()
	for i := 0; int32(i) < count; i++ {
		b.GetInt32le() // ID
	}

	b.GetInt32le() // universeObserveLevel
	b.GetFloat32() // solarSailLife
	b.GetFloat32() // solarEnergyLossRate
	b.GetBoolean() // useIonLayer
	b.GetInt32le() // inserterStackCount
	b.GetFloat32() // logisticDroneSpeed
	b.GetFloat32() // logisticDroneSpeedScale
	b.GetInt32le() // logisticDroneCarries
	b.GetFloat32() // logisticShipSailSpeed
	b.GetFloat32() // logisticShipWarpSpeed
	b.GetFloat32() // logisticShipSpeedScale
	b.GetBoolean() // logisticShipWarpDrive
	b.GetInt32le() // logisticShipCarries
	b.GetFloat32() // miningCostRate
	b.GetFloat32() // miningSpeedScale
	b.GetInt32le() // storageLevel
	b.GetInt32le() // labLevel
	b.GetInt32le() // techSpeed
	b.GetFloat32() // dysonNodeLatitide
	b.GetInt64le() // universeMatrixPointUploaded
	b.GetBoolean() // missionAccomplished
}
