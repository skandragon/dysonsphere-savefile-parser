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

import "github.com/skandragon/dysonsphere/types"

// StationStore holds the current status of a station's item storage
type StationStore struct {
	Index       int32                     `json:"index"`
	ItemID      types.ItemType            `json:"item_id"`
	Count       int32                     `json:"count"`
	Max         int32                     `json:"max"`
	LocalOrder  int32                     `json:"local_order"`
	RemoteOrder int32                     `json:"remote_order"`
	LocalLogic  types.LogisticStorageType `json:"local_logic"`
	RemoteLogic types.LogisticStorageType `json:"remote_logic"`
}

func parseStationStore(b *Buffer, index int32) *StationStore {
	checkVers(b, 0, "StationStore")

	return &StationStore{
		Index:       index,
		ItemID:      types.ItemType(b.GetInt32le()),
		Count:       b.GetInt32le(),
		LocalOrder:  b.GetInt32le(),
		RemoteOrder: b.GetInt32le(),
		Max:         b.GetInt32le(),
		LocalLogic:  types.LogisticStorageType(b.GetInt32le()),
		RemoteLogic: types.LogisticStorageType(b.GetInt32le()),
	}
}
