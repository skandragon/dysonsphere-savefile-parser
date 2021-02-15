package main

// StationStore holds the current status of a station's item storage
type StationStore struct {
	ItemID      int32
	Count       int32
	Max         int32
	LocalOrder  int32
	RemoteOrder int32
	LocalLogic  LogisticStorageType
	RemoteLogic LogisticStorageType
}

func parseStationStore(b *Buffer) *StationStore {
	checkVers(b, 0, "StationStore")

	return &StationStore{
		ItemID:      b.GetInt32le(),
		Count:       b.GetInt32le(),
		LocalOrder:  b.GetInt32le(),
		RemoteOrder: b.GetInt32le(),
		Max:         b.GetInt32le(),
		LocalLogic:  LogisticStorageType(b.GetInt32le()),
		RemoteLogic: LogisticStorageType(b.GetInt32le()),
	}
}
