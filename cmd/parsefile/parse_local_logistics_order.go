package main

func parseLocalLogisticOrder(b *Buffer) {
	checkVers(b, 0, "LocalLogisticsOrder")

	b.GetInt32le() // otherStationId = r.ReadInt32()
	b.GetInt32le() // thisIndex = r.ReadInt32()
	b.GetInt32le() // otherIndex = r.ReadInt32()
	b.GetInt32le() // itemId = r.ReadInt32()
	b.GetInt32le() // thisOrdered = r.ReadInt32()
	b.GetInt32le() // otherOrdered = r.ReadInt32()
}
