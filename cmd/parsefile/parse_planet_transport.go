package main

import "fmt"

func parsePlanetTransport(b *Buffer) {
	const VERSION = 0
	checkVers(b, VERSION, "PlanetTransport")

	stationCursor := b.GetInt32le()
	b.GetInt32le() // stationCapacity
	stationRecycleCounter := b.GetInt32le()
	//fmt.Printf("station: capacity %d, cursor %d, recycleCursor %d\n",
	//	stationCapacity, stationCursor, stationRecycleCounter)
	for i := int32(1); i < stationCursor; i++ {
		id := b.GetInt32le()
		if id != 0 {
			if id != i {
				panic(fmt.Sprintf("id != i (%d, %d)", id, i))
			}
			parseStationComponent(b)
		}
	}
	for i := 0; int32(i) < stationRecycleCounter; i++ {
		b.GetInt32le() // recycle id?
	}
}
