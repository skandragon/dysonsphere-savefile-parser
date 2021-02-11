package main

func parseStationStore(b *Buffer) {
	checkVers(b, 0, "StationStore")

	b.GetInt32le() // itemId = r.ReadInt32();
	b.GetInt32le() // count = r.ReadInt32();
	b.GetInt32le() // localOrder = r.ReadInt32();
	b.GetInt32le() // remoteOrder = r.ReadInt32();
	b.GetInt32le() // max = r.ReadInt32();
	b.GetInt32le() // localLogic = (ELogisticStorage)r.ReadInt32();
	b.GetInt32le() // remoteLogic = (ELogisticStorage)r.ReadInt32();
}
