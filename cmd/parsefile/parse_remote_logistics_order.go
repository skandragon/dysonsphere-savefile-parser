package main

func parseRemoteLogisticOrder(b *Buffer) {
	checkVers(b, 0, "RemoteLogisticOrder")

	b.GetInt32le() // otherStationGId = r.ReadInt32()
	b.GetInt32le() // thisIndex = r.ReadInt32()
	b.GetInt32le() // otherIndex = r.ReadInt32()
	b.GetInt32le() // itemId = r.ReadInt32()
	b.GetInt32le() // thisOrdered = r.ReadInt32()
	b.GetInt32le() // otherOrdered = r.ReadInt32()
}
