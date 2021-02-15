package main

import "github.com/skandragon/dysonsphere/types"

// StationStore holds the current status of a station's item storage
type StationStore struct {
	Index       int32                     `json:"index"`
	ItemID      int32                     `json:"item_id"`
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
		ItemID:      b.GetInt32le(),
		Count:       b.GetInt32le(),
		LocalOrder:  b.GetInt32le(),
		RemoteOrder: b.GetInt32le(),
		Max:         b.GetInt32le(),
		LocalLogic:  types.LogisticStorageType(b.GetInt32le()),
		RemoteLogic: types.LogisticStorageType(b.GetInt32le()),
	}
}
